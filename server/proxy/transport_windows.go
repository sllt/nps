//go:build windows
// +build windows

package proxy

import (
	"github.com/sllt/nps/lib/conn"
)

func HandleTrans(c *conn.Conn, s *TunnelModeServer) error {
	return nil
}
