package p2p

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net"
)

type TCPTransport struct {
	listenAddr  string
	listener    net.Listener
	NewPeerConn NewPeerConn
}

type NewPeerConn struct {
	conn     net.Conn
	outbound bool
}

func NewPeer(conn net.Conn, outbound bool) *NewPeerConn {
	return &NewPeerConn{
		conn:     conn,
		outbound: outbound,
	}
}

// type ConnHolder struct {

// }

func NewTCPConn(listenaddr string) *TCPTransport {
	return &TCPTransport{
		listenAddr: listenaddr,
	}
}

func (t *TCPTransport) GetTCPConn() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddr)
	if err != nil {
		return err
	}
	go t.StartAccpetLoop()
	return nil
}

func (t *TCPTransport) StartAccpetLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("tcp accept error: %v", err)
		}
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	fmt.Println("New Incoming Connection", conn)
	NewPeer(conn, true)
}

// func (t *TCPTransport) ListneAndAccept(conn net.Conn, outbound bool) (net.Conn, error) {

// }
func CreateHash(key string) string {
	hash := sha1.Sum([]byte(key))
	hashStr := hex.EncodeToString(hash[:])
	return hashStr
}
