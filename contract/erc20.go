package contract

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type IERC20 interface {
	Contract
	BalanceOf(address common.Address) (*big.Int,error)
}

type ERC20 struct {
	*BaseContract
}

func NewERC20(address common.Address, abi *abi.ABI) *ERC20 {
	erc20:=& ERC20{}
	erc20.BaseContract = NewBaseContract(address,abi)
	return erc20
}

func (c *ERC20) BalanceOf(address common.Address) (*big.Int,error) {
	funcName:="balanceOf"
	input,err:= c.Pack(funcName,address)
	if err!=nil {
		return nil,err
	}
	data,err:=c.ethClient.Call(&ethereum.CallMsg{
		To: &c.address,
		Data: input,
	})
	if err!=nil {
		return nil,err
	}
	balance:=new(big.Int)
	err= c.Unpack(balance, funcName,data)
	if err!=nil {
		return nil,err
	}
	return balance,err
}
