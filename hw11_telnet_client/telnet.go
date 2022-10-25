package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type Telnet struct {
	conn    net.Conn
	addr    string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
}

func (t *Telnet) Connect() error {
	var err error
	t.conn, err = net.DialTimeout("tcp", t.addr, t.timeout)
	if err != nil {
		return fmt.Errorf("net.DialTimeout: %w", err)
	}
	return nil
}

func (t Telnet) Close() error {
	if err := t.conn.Close(); err != nil {
		return fmt.Errorf("t.conn.Close: %w", err)
	}
	return nil
}

func (t Telnet) Send() error {
	if _, err := io.Copy(t.conn, t.in); err != nil {
		return fmt.Errorf("io.Copy: %w", err)
	}
	return nil
}

func (t Telnet) Receive() error {
	if _, err := io.Copy(t.out, t.conn); err != nil {
		return fmt.Errorf("io.Copy: %w", err)
	}
	return nil
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &Telnet{
		addr:    address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}
