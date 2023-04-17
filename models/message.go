package models

import (
	"encoding/json"
	"fmt"

	// "ginchat/models"
	"net"
	"net/http"
	"strconv"
	"sync"

	// "github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	FromID   int64
	TargetID int64
	Type     int    //發送群聊、私聊
	Media    int    //消息類型
	Content  string //消息內容
	Picture  string //
	Url      string
	Desc     string
	Amount   int //有的沒的統計
}

func (table *Message) TableName() string {
	return "message"
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

var clientMap map[int64]*Node = make(map[int64]*Node, 0)

var rwLocker sync.RWMutex

func Chat(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	Id := query.Get("userId")
	// token := query.Get("token")
	userId, _ := strconv.ParseInt(Id, 10, 64)
	// msgType := query.Get("type")
	// targetId := query.Get("targetId")
	// context := query.Get("context")
	isValid := true //待處理
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isValid
		},
	}).Upgrade(writer, request, nil)

	if err != nil {
		fmt.Println("err", err)
		return
	}

	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	//userid 跟node綁定 並枷鎖
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()

	go sendProc(node)
	go recvProc(node)
	sendMsg(userId, []byte("歡迎進入聊天室"))

}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()

		if err != nil {
			fmt.Println(err)
			return
		}
		broadMsg(data)
		fmt.Println("[ws] <<<", data)

	}
}

var broadMsgChan chan []byte = make(chan []byte, 1024)

func broadMsg(data []byte) {
	broadMsgChan <- data
}

func init() {
	go udpSendProc()
	go updRecvProc()
}

// 完成udp 數據發送攜程
func udpSendProc() {
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{
		// IP:   net.IPv4(192, 168, 0, 58),
		IP: net.IPv4(192, 168, 68, 106),
		// IP: net.IPv4(192, 168, 1, 143),
		// IP: net.IPv4(192, 168, 0, 208),

		//192.168.68.110
		Port: 3333,
	})

	defer con.Close()
	if err != nil {
		fmt.Println(err)
	}

	for {
		select {
		case data := <-broadMsgChan:
			_, err := con.Write(data)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func updRecvProc() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP: net.IPv4(192, 168, 68, 106),

		Port: 3333,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer con.Close()
	for {
		var buf [512]byte
		n, err := con.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		if n > 0 && buf != [512]byte{} {
			dispatch(buf[0:n])
		}
		// dispatch(buf[0:n])
	}
}

//後端調度邏輯

func dispatch(buf []byte) {
	msg := Message{}
	err := json.Unmarshal(buf, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch msg.Type {
	case 1: //私信
		sendMsg(msg.TargetID, buf)
		fmt.Println("dsadss", msg.TargetID)
		// case 2: // 群發
		// 	sendGroupMsg()
		// case 3: //廣播
		// 	sendAllMsg()
		// case 4:

	}
}

func sendMsg(userId int64, msg []byte) {
	rwLocker.RLock()
	node, ok := clientMap[userId]
	rwLocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}

}
