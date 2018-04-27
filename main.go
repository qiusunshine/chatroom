package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
	"time"
	"hdy/chat/user"
	"github.com/astaxie/beego"
	_ "hdy/chat/routers"
)

type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

type Client struct {
	id     string
	socket *websocket.Conn
	send   chan []byte
}

type Message struct {
	Sender     string `json:"sender,omitempty"`
	Content    string `json:"content,omitempty"`
	TheTime    string `json:"time,omitempty"`
	ClientsNum int    `json:"clients,omitempty"`
	MsgType    string `json:"msgType,omitempty"`
}

var manager = ClientManager{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
}

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.register:
			manager.clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: (*conn).id + "加入了聊天室", MsgType: "join"})
			manager.send(jsonMessage, conn)
		case conn := <-manager.unregister:
			if _, ok := manager.clients[conn]; ok {
				close(conn.send)
				delete(manager.clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: (*conn).id + "离开了聊天室", MsgType: "join"})
				manager.send(jsonMessage, conn)
			}
		case message := <-manager.broadcast:
			for conn := range manager.clients {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(manager.clients, conn)
				}
			}
		}
	}
}

func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.clients {
		if conn != ignore {
			conn.send <- message
		} else {
			//发送历史记录
			his := user.GetMsgs()
			jsonMessage, _ := json.Marshal(&his)
			conn.send <- jsonMessage
		}
	}
}
//监听获取消息
func (c *Client) read() {
	defer func() {
		manager.unregister <- c
		c.socket.Close()
	}()
	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			manager.unregister <- c
			c.socket.Close()
			break
		}
		m := Message{Sender: c.id, Content: string(message), TheTime: time.Now().Format("2006-01-02 15:04:05"), ClientsNum: len(manager.clients), MsgType: "msg"}
		jsonMessage, _ := json.Marshal(&m)
		manager.broadcast <- jsonMessage
		go user.NewMsg(m.Sender, m.Content, m.TheTime)
	}
}
//发送消息到socket客户
func (c *Client) write() {
	defer c.socket.Close()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
func wsPage(res http.ResponseWriter, req *http.Request) {
	conn, e := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if e != nil {
		http.NotFound(res, req)
		return
	}
	client := &Client{id: user.NewName(), socket: conn, send: make(chan []byte)}
	manager.register <- client
	go client.read()
	go client.write()
}
func main() {
	fmt.Println("Starting application...")
	go manager.start()
	go beego.Run() //开启客户端
	http.HandleFunc("/ws", wsPage)
	http.ListenAndServe(":12345", nil) //监听WebSocket
}
