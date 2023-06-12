package lib

import (
	"bytes"
	"image"
	"image/png"

	_ "image/jpeg"
	"io"

	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/transform"
	"github.com/ipfs/go-cid"
	"github.com/taubyte/go-sdk/event"
	ipfs "github.com/taubyte/go-sdk/ipfs/client"
)

//export get
func get(c string) (io.ReadCloser, error) {
	_cid, err := cid.Parse(c)
	if err != nil {
		return nil, err
	}

	ic, err := ipfs.New()
	if err != nil {
		return nil, err
	}

	fr, err := ic.Open(_cid)
	if err != nil {
		return nil, err
	}

	return fr, nil
}

//export bild
func bild(e event.Event) uint32 {
	h, err := e.HTTP()
	if err != nil {
		return 1
	}

	c, err := h.Query().Get("cid")
	if err != nil {
		h.Write([]byte(err.Error()))
		return 1
	}

	ir, err := get(c)
	if err != nil {
		h.Write([]byte(err.Error()))
		return 1
	}
	defer ir.Close()

	img, _, err := image.Decode(ir)
	if err != nil {
		h.Write([]byte("Error Decoding the image. Failed with " + err.Error()))
		return 1
	}

	inverted := effect.Invert(img)
	resized := transform.Resize(inverted, 800, 800, transform.Linear)
	rot := transform.Rotate(resized, 60, nil)

	var b bytes.Buffer
	err = png.Encode(&b, rot)
	if err != nil {
		h.Write([]byte("PNG Encoding failed with " + err.Error()))
		return 1
	}

	h.Headers().Set("Content-Type", "image/png")
	h.Write(b.Bytes())

	return 0
}
