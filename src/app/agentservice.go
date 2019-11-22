package main

import (
	"app/client_handler"
	"app/enmu"
	"app/helper/stack"
	"app/misc/packet"
	"app/session"
	"bufio"
	"github.com/golang/glog"
	"github.com/xtaci/kcp-go"
	"net"
	"os"
)

var closed = make(chan struct{}, 1)

func agentRun() {
	lestener, err := kcp.ListenWithOptions(enmu.ServerHost+":"+enmu.ServerPort, nil, 10, 3)
	if err != nil {
		glog.Info("listen error:", err)
		os.Exit(1)
	}
	defer lestener.Close()
	glog.Info("listening on " + enmu.ServerHost + ":" + enmu.ServerPort)
	for {
		conn, err := lestener.Accept()
		if err != nil {
			glog.Info("accept error:", err)
			os.Exit(1)
		}
		glog.Infof("message %s->%s\n", conn.RemoteAddr(), conn.LocalAddr())

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	in := make(chan []byte, 16)
	sess := session.NewSession(in)
	defer func() {
		glog.Info("disconnect:" + conn.RemoteAddr().String())
		sess.OffLine(sess.Id)
		closed <- struct{}{}
		conn.Close()
	}()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	go func() {
		for msg := range in {
			writer.Write(msg)
			writer.Write([]byte("\n"))
			writer.Flush()
		}
	}()
	for {
		msg, _, err := reader.ReadLine()
		if err != nil {
			glog.Info(err)
			return
		}
		reader := packet.Reader(msg)
		c, err := reader.ReadS16()
		if err != nil {
			glog.Info("err=", err)
			return
		}
		bytes := executeHandler(c, sess, reader)
		for _, byt := range bytes {
			in <- byt
		}
	}
}

//执行方法
func executeHandler(code int16, sess *session.Session, reader *packet.Packet) [][]byte {
	defer stack.PrintRecoverFromPanic()
	handle := client_handler.Handlers[code]
	if handle == nil {
		return nil
	}
	retByte := handle(sess, reader)
	return retByte
}
