package p2p

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
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
	go t.readStream(conn.RemoteAddr().String(), conn)
}

// func Read(key string) {

// }
// func (t *TCPTransport) ReadStream(data []byte) error {
// 	reader := bytes.NewReader(data)
// 	return t.readStream("sample", reader)
// }

func (t *TCPTransport) PathSplitting(s string) (string, error) {
	var WholePath string
	path1 := len(s)
	fmt.Println(path1)
	blocksize := 5
	slicelenth := path1 / blocksize
	paths := make([]string, slicelenth)

	for i := 0; i < slicelenth; i++ {
		from, to := i*blocksize, (i*blocksize)+blocksize
		paths[i] = s[from:to]

	}
	WholePath = strings.Join(paths, "/")
	fmt.Println(WholePath)
	return t.DirCreation(WholePath)

	// for i := 0; i < path1; i++ {
	// 	if i == 2 {
	// 		dir1 := s[0:i]
	// 		dir2 := s[i : len(dir1)+2]
	// 		dir3 := s[len(dir2):]
	// 	}
	// }
	// dir1 := path[0:2]
	// dir2 := path[2:4]
	// dir3 := path[4:]
}
func (t *TCPTransport) DirCreation(s string) (string, error) {
	err := os.MkdirAll(filepath.Dir(s), 0755)
	if err != nil {
		return "", err
	}
	fmt.Println(s)
	return s, nil

}

func (t *TCPTransport) Hashing(line string) (string, error) {
	hash := sha1.Sum([]byte(line))
	hasedString := hex.EncodeToString(hash[:])
	fmt.Println(hasedString)
	return t.PathSplitting(hasedString)

}

func (t *TCPTransport) readStream(key string, conn net.Conn) error {
	reader := bufio.NewReader(conn)
	// buf := new(bytes.Buffer)
	for {
		// buf := make([]byte, 1024)
		// r := strings.NewReader(buf)
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		t.Hashing(line)

		//message := string(buf[:n])
		// n, err := io.Copy(buf, r)
		// if err != nil {
		// 	return nil
		// }
		// if err != nil {
		// 	return err
		// }
		fmt.Println("Your Message is: " + line)
		_, err = conn.Write([]byte("you said: " + line))
		if err != nil {
			return nil
		}
	}
}

// func (t *TCPTransport) ListneAndAccept(conn net.Conn, outbound bool) (net.Conn, error) {

// }
//
//	func CreateHash(key string) string {
//		hash := sha1.Sum([]byte(key))
//		hashStr := hex.EncodeToString(hash[:])
//		return hashStr
//	}
// func (t *TCPTransport) TeeReader(key string, data io.Reader) (int64, error) {
// 	buf := new(bytes.Buffer)
// 	tee := io.TeeReader(data, buf)
// 	os.Create()
// }
