package pkg

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

const (
	dialErrMsg      = "error to dial to host "
	connectedMsg    = "...Connected to"
	connectCloseMsg = "...Connection was closed by peer"
	eofMsg          = "...EOF"
)

type TelnetClientInt interface {
	Connect() error
	Close() error
	Receive() error
	Send() error
}

type TelenetClientStr struct {
	add       string
	timeout   time.Duration
	conn      net.Conn
	inputRead *bufio.Reader
	connRead  *bufio.Reader
	out       io.Writer
}

func TelenetClientInit(add string, timeout time.Duration, in io.Reader, out io.Writer) TelnetClientInt {
	inputRead := bufio.NewReader(in)
	res := &TelenetClientStr{
		add:       add,
		timeout:   timeout,
		inputRead: inputRead,
		out:       out,
	}

	return res
}

func (t *TelenetClientStr) Connect() error {
	conn, err := net.DialTimeout("tcp", t.add, t.timeout)
	if err != nil {
		return fmt.Errorf("%s : %w", dialErrMsg, err)
	}
	t.conn = conn
	t.connRead = bufio.NewReader(t.conn)
	fmt.Fprintf(os.Stderr, "%s %s\n", connectedMsg, t.add)
	return nil
}

func (t *TelenetClientStr) Receive() error {
	line, err := t.connRead.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			err = fmt.Errorf(connectCloseMsg)
		}
		return err
	}
	if _, err := fmt.Fprint(t.out, line); err != nil {
		return err
	}
	return nil
}

func (t *TelenetClientStr) Send() error {
	line, err := t.inputRead.ReadString('\n')
	if err != nil {
		return err
	}
	_, err = t.conn.Write([]byte(line))
	if err != nil {
		return fmt.Errorf(connectCloseMsg)
	}
	return nil
}

func (t *TelenetClientStr) Close() error {
	err := t.conn.Close()
	if err != nil {
		return err
	}
	return nil
}
