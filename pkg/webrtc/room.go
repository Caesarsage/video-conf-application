package webrtc

import (
	"log"
	"sync"
	"videochat/pkg/chat"

	"github.com/gofiber/contrib/websocket"
	"github.com/pion/webrtc/v3"
)

type Room struct {
	Peers *Peers
	Hub *chat.Hub
}

func RoomConn(c *websocket.Conn, p *Peers) {
	var config webrtc.Configuration


	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		log.Print(err)
		return
	}

	newPeer := PeersConnectionState{
		PeerConnection: peerConnection,
		WebSocket: &ThreadSafeWriter(),
		Conn: c,
		Mutex: sync.Mutex()
	}
}
