package user

import "sync"

type MsgHistory struct {
	Data []Msg `json:"msg"`
	MsgType   string `json:"msgType,omitempty"`
}
type Msg struct {
	Sender    string `json:"sender,omitempty"`
	Content   string `json:"content,omitempty"`
	TheTime   string `json:"time,omitempty"`
}
var msgs =make([]Msg,20,20)
var mutex sync.RWMutex
func NewMsg(sender,msg,time string) {
	mutex.Lock()
	defer mutex.Unlock()
	theMsg:=Msg{Sender:sender,Content:msg,TheTime:time}
	for key, value := range msgs {
		if key==0{
			msgs[19]=theMsg
		}else {
			msgs[key-1]=value
		}
	}
}
func GetMsgs()MsgHistory{
	mutex.RLock()
	defer mutex.RUnlock()
	return MsgHistory{Data:msgs,MsgType:"history"}
}