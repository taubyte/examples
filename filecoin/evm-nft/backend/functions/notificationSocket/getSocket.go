package lib

import (
	"strings"

	utils "github.com/tafseer-khan/tb_utils/err"
	"github.com/taubyte/go-sdk/ethereum/client/bytes"
	"github.com/taubyte/go-sdk/event"
	pubsub "github.com/taubyte/go-sdk/pubsub/node"
)

//export getSocket
func getSocket(e event.Event) uint32 {
	h, err := e.HTTP()
	if err != nil {
		panic(err)
	}

	// domain/path?wallet=<address>
	address, err := h.Query().Get("wallet")
	if err != nil {
		return utils.Write(h, err)
	}

	if _, err = bytes.AddressFromHex(address); err != nil {
		return utils.Write(h, err)
	}

	address = strings.ToLower(address)

	channel, err := pubsub.Channel("notifications/" + address)
	if err != nil {
		return utils.Write(h, err)
	}

	if err := channel.Subscribe(); err != nil {
		return utils.Write(h, err)
	}

	wsUrl, err := channel.WebSocket().Url()
	if err != nil {
		return utils.Write(h, err)
	}

	h.Write([]byte(wsUrl.Path))
	return 0
}
