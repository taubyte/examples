package lib

import (
	"io"

	"github.com/ipfs/go-cid"
	"github.com/taubyte/go-sdk/event"
	ipfs "github.com/taubyte/go-sdk/ipfs/client"
)

//export get
func get(e event.Event) uint32 {
	h, err := e.HTTP()
	if err != nil {
		return 1
	}

	c, err := h.Query().Get("cid")
	if err != nil {
		h.Write([]byte(err.Error()))
		return 1
	}

	_cid, err := cid.Parse(c)
	if err != nil {
		h.Write([]byte(err.Error()))
		return 1
	}

	ic, err := ipfs.New()
	if err != nil {
		h.Write([]byte(err.Error()))
		return 1
	}

	fr, err := ic.Open(_cid)
	if err != nil {
		h.Write([]byte(err.Error()))
		return 1
	}

	defer fr.Close()

	_, err = io.Copy(h, fr)
	if err != nil {
		h.Write([]byte(err.Error()))
		return 1
	}

	return 0
}
