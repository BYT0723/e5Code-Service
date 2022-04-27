package hub

import (
	"context"
	"e5Code-Service/common/contextx"
	"e5Code-Service/service/ci/rpc/ci"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
	bufSize        = 256
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	ctx  context.Context
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

type BuildReq struct {
	UserID string `json:"user_id"`
	PlanID string `json:"plan_id"`
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(appData string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	req := &BuildReq{}
	for {
		if err := c.conn.ReadJSON(req); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		go func(c *Client) {
			ctx := contextx.SetValueToMetadata(context.Background(), contextx.UserID, req.UserID)
			streamClient, err := c.hub.CIRpc.BuildImage(ctx, &ci.BuildReq{BuildPlanID: req.PlanID})
			if err != nil {
				logx.Error("Fail to BuildImage:", err.Error())
				c.send <- []byte(err.Error())
				return
			}
			for {
				rsp, err := streamClient.Recv()
				if err != nil {
					if err == io.EOF {
						break
					}
					logx.Error("Fail to Recv Log:", err.Error())
					c.send <- []byte(err.Error())
					return
				}
				c.send <- []byte(rsp.LogInfo)
			}
		}(c)
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logx.Error("ServeWs: ", err.Error())
		return
	}
	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte, bufSize),
	}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}
