package http

import (
	"net/http"
	"io"
	"../config"
	"fmt"
)

type HTTP struct {
	//ip:port
	addrhttp string

	//package to target server
	hPackage io.Reader
}

func (h HTTP) Send () error {
	req, err := http.NewRequest("POST", h.addrhttp, h.hPackage)

	req.Header.Set("Content-Type", "application/byte")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()

	return nil
}

func New (p io.Reader) *HTTP {
	h := new(HTTP)
	h.addrhttp = fmt.Sprintf("http://%s:%s",config.Cfg.Host,config.Cfg.Port)
	h.hPackage = p

	return h
}
