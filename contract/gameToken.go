package contract

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

const (
	FUNC_CONUSME="consume"
	FUNC_REWARD="reward"
)

type GameToken struct {
	*ERC20
}

func NewGameToken(address common.Address, abi *abi.ABI) *GameToken {
	gameToken:=&GameToken{}
	gameToken.BaseContract = NewBaseContract(address,abi)
	return gameToken
}
