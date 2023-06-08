package lib

import "math/big"

type computeEvent struct {
	JobId  *big.Int
	Gender string
	Sender string
	Count  *big.Int
}

type avatar struct {
	Cid string
}
