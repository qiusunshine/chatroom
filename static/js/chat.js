$(function() {
    var conn;
    var msg = $("#msg");
    var log = $("#log");
    function appendLog(msg) {
        var d = log[0];
        var doScroll = d.scrollTop == d.scrollHeight - d.clientHeight;
        msg.appendTo(log);
        if (doScroll) {
            d.scrollTop = d.scrollHeight - d.clientHeight;
        }
    }
    $("#sendButton").click(function () {
        if (!conn) {
            return;
        }
        if (!msg.val()) {
            return ;
        }
        conn.send(msg.val());
        msg.val("");
    });
    $("#form").submit(function() {
        if (!conn) {
            return false;
        }
        if (!msg.val()) {
            return false;
        }
        conn.send(msg.val());
        msg.val("");
        return false
    });
    if (window["WebSocket"]) {
        conn = new WebSocket("ws://localhost:12345/ws");
        conn.onclose = function(evt) {
            appendLog($("<div><b>Connection Closed.</b></div>"));
        };
        conn.onmessage = function(evt) {
            var result = $.parseJSON(evt.data);
            if(result.msgType === "join"){
                appendLog($("<p class=\"am-text-primary\">"+result.content+"</p>"));
            }else if(result.msgType === "msg"){
                appendLog($("<p class=\"am-text-danger\">"+result.time+" <"+result.sender+"> ：</p>"));
                appendLog($("<p class=\"am-text-primary\">"+result.content+"</p>"));
            }else{
                for (var i=0;i<result.msg.length-1;i++){
                    var r = (result.msg)[i];
                    if(r.sender === undefined)continue;
                    appendLog($("<p class=\"am-text-danger\">"+r.time+" <"+r.sender+"> ："+"</p>"));
                    appendLog($("<p class=\"am-text-primary\">"+r.content+"</p>"));
                }
            }
        };
        conn.onopen=function(evt){
            console.log("socket open");
        };
    } else {
        appendLog($("<div><b>WebSockets Not Support.</b></div>"))
    }
});