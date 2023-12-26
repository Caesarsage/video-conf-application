package handlers

import (
	"fmt"
	"time"

	w "videochat/pkg/webrtc"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)


func Stream(c *fiber.Ctx) error {
	suuid := c.Params("suuid")
	if suuid == "" {
		c.Status(400)
		return nil
	}

	web_socket := "web_socket"
	if os.Getenv("ENV") == "PRODUCTION" {
		web_socket = "wss"
	}

	w.RoomsLock.Lock()

	if _, ok := w.Streams[suuid]; ok {
		w.RoomsLock.Unlock()
		return c.Render("stream", fiber.Map{
			"StreamWebSocketAddr": fmt.Sprintf("%s://%s/stream/%s/websocket", web_socket, c.Hostname(), suuid),
			"ChatWebSocketAddr": fmt.Sprintf("%s://%s/stream/%s/chat/websocket", web_socket, c.Hostname(), suuid),
			"ViewWebSocketAddr": fmt.Sprintf("%s://%s/stream/%s/viewer/websocket", web_socket, c.Hostname(), suuid),
			"Type": "stream",
		}, "layout/main")
	}
	w.RoomsLock.Unlock()

	return c.Render("stream", fiber.Map{
		"NoStream": "true",
		"Leave": "true",
	}, "layout/main")
	
}

func StreamWebsocket(c *websocket.Conn) {
	suuid := c.Params("suuid")

	if suuid == ""{
		return
	}
	w.RoomsLock.Lock()
	if stream, ok := w.Streams[uuid]; ok{
		w.RoomsLock.Unlock()
		w.StreamConn(c, stream.Peers)
		return
	}
	w.RoomsLock.Unlock()
}

func StreamViewerWebsocket(c *websocket.Conn) {
	suuid := c.Params("suuid")
	if suuid == "" {
		return
	}
	w.RoomsLock.Lock()
	if stream, ok := w.Streams[suuid]; ok {
		w.RoomsLock.Unlock()
		viewerConn(c, stream.Peers)
		return
	}
	w.RoomsLock.Unlock()
}

func viewerConn( c *websocket.Conn, p *w.Peers) {
	ticker := time.NewTicker(time.Second * 1)
	defer func() {
		ticker.Stop()
		c.Close()
	}()

	for {
		select {
		case <-ticker.C:
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write([]byte(fmt.Sprintf("data: %d\n\n", len(p.Connections))))
		}
	}

}
