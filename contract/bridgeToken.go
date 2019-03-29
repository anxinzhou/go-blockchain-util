package contract

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

var (
	BridgeTokenFunc *BridgeTokenFuncT
	BridgeTokenEvent *BridgeTokenEventT
)

// function name
type BridgeTokenFuncT struct {
	FuncPayToken string
	FuncExchangeToken string
	FuncPayNFT string
	FuncExchangeNFT string
}

// event name
type BridgeTokenEventT struct {
	EventPayToken string
	EventPayNFT string
	EventExchangeToken string
	EventExchangeNFT string
}

type BridgeTokenEventExchangeToken struct {
	User common.Address `json:"user"`
	Amount common.Address `json:"amount"`
}

//event ExchangeNFT(uint256 tokenID, address owner, uint256 gene, uint256 avatarLevel, bool weaponed, bool armored);
type BridgeTokenEventExchangeNFT struct {
	TokenID *big.Int `json:"tokenID"`
	Owner common.Address `json:"owner"`
	Gene *big.Int `json:"gene"`
	AvatarLevel *big.Int `json:"avatarLevel"`
	Weaponed bool `json:"weaponed"`
	Armored bool `json:"armored"`
}

func init() {
	BridgeTokenFunc=&BridgeTokenFuncT{
		FuncPayToken:"pay",
		FuncExchangeToken:"exchange",
		FuncPayNFT: "payNFT",
		FuncExchangeNFT: "exchangeNFT",
	}
	BridgeTokenEvent=&BridgeTokenEventT{
		EventPayToken:"pay",
		EventPayNFT:"payNFT",
		EventExchangeToken:"exchange",
		EventExchangeNFT:"exchangeNFT",
	}
}

type BridgeToken struct {
	*GameToken
	FuncPayToken string
	FuncExchangeToken string
	FuncPayNFT string
}

func NewBridgeToken(address common.Address,abi *abi.ABI) *BridgeToken {
	bridgeToken:=&BridgeToken{}
	bridgeToken.BaseContract = NewBaseContract(address,abi)
	return bridgeToken
}




