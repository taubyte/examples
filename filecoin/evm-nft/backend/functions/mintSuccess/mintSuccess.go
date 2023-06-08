package lib

import (
	"math/big"
	"strings"

	"github.com/mailru/easyjson"
	"github.com/taubyte/go-sdk/ethereum/client/logs"
	"github.com/taubyte/go-sdk/event"
	pubsub "github.com/taubyte/go-sdk/pubsub/node"
)

//export mintSuccess
func mintSuccess(e event.Event) uint32 {
	p, err := e.PubSub()
	if err != nil {
		panic(err)
	}

	data, err := p.Data()
	if err != nil {
		panic(err)
	}

	_logs := logs.EventLog{}
	if err = easyjson.Unmarshal(data, &_logs); err != nil {
		panic(err)
	}

	var nft mintedNFT
	if err = easyjson.Unmarshal(_logs.Log.Data, &nft); err != nil {
		panic(err)
	}

	channel, err := pubsub.Channel("notifications/" + strings.ToLower(nft.Sender))
	if err != nil {
		panic(err)
	}

	channel.Publish(_logs.Log.Data)

	return 0
}

type mintedNFT struct {
	Sender   string
	TokenId  *big.Int
	TokenUri string
}
