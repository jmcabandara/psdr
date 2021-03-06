package tcp

import (
	"fmt"
	"io"
	"net"

	"github.com/joeke80215/psdr/config"
)

type TCP struct {
	addrtcp  *net.TCPAddr
	tPackage io.Reader
}

func (t TCP) Send() error {
	ds, err := net.DialTCP("tcp", nil, t.addrtcp)
	if err != nil {
		return err
	}

	io.Copy(ds, t.tPackage)
	ds.Close()

	return err
}

func New(p io.Reader) *TCP {
	t := new(TCP)
	t.addrtcp, _ = net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", config.Cfg.Host, config.Cfg.Port))
	t.tPackage = p

	return t
}
