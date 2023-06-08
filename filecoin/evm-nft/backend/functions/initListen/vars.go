package lib

import _ "embed"

//go:embed compute_fee_abi.json
var computeFeeAbi []byte

//go:embed mint_abi.json
var mintAbi []byte

const (
	computeFeeAddress = "0xE91dCD50913066Db9E7Ea8D9F6D41138163CE90a"
	mintAddress       = "0x017E5f589795ffc218E97D2B82ECf4c7bd387e63"
)
