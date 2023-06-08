package lib

import (
	"bytes"
	"time"

	utils "github.com/tafseer-khan/tb_utils/err"
	ethereum "github.com/taubyte/go-sdk/ethereum/client"
	"github.com/taubyte/go-sdk/event"
)

//export initListen
func initListen(e event.Event) uint32 {
	h, err := e.HTTP()
	if err != nil {
		panic(err)
	}

	client, err := ethereum.New("wss://wss.calibration.node.glif.io/apigw/lotus/rpc/v1")
	if err != nil {
		return utils.Write(h, err)
	}

	computeFee, err := client.NewBoundContract(bytes.NewReader(computeFeeAbi), computeFeeAddress)
	if err != nil {
		return utils.Write(h, err)
	}

	computeSubmitted, err := computeFee.Event("ComputeSubmitted")
	if err != nil {
		return utils.Write(h, err)
	}

	if err = computeSubmitted.Subscribe("pubsub://currencyTransfer", 10*time.Minute); err != nil {
		return utils.Write(h, err)
	}

	mint, err := client.NewBoundContract(bytes.NewReader(mintAbi), mintAddress)
	if err != nil {
		return utils.Write(h, err)
	}

	mintEvent, err := mint.Event("NewTauFRC721NFTMinted")
	if err != nil {
		return utils.Write(h, err)
	}

	if err = mintEvent.Subscribe("pubsub://nftMinted", 10*time.Minute); err != nil {
		return utils.Write(h, err)
	}

	return 0
}
