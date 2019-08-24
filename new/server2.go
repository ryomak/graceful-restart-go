package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"syscall"

	"golang.org/x/sys/unix"
)

func init() {
	pidFilePath := "server2.pid"
	if err := os.Remove(pidFilePath); err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
	}
	pidf, err := os.OpenFile(pidFilePath, os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		panic(err)
	}
	if _, err := fmt.Fprint(pidf, syscall.Getpid()); err != nil {
		panic(err)
	}
	pidf.Close()
}

func main() {
	handler := http.NewServeMux()
  handler.HandleFunc("/server2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("This server is 2")
		fmt.Fprintf(w, "This server is 2")
	})
	lc := net.ListenConfig{
		Control: listenCtrl,
	}
	listener, err := lc.Listen(context.Background(), "tcp4", ":8080")
	if err != nil {
		panic(err)
	}
	http.Serve(listener, handler)
}

func listenCtrl(network string, address string, c syscall.RawConn) error {
	var err error
	c.Control(func(s uintptr) {
		err = unix.SetsockoptInt(int(s), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
		if err != nil {
			return
		}
		err = unix.SetsockoptInt(int(s), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
		if err != nil {
			return
		}
	})
	return err
}
