package erc20

import (
	"github.com/xxRanger/blockchainUtil/contract"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type IERC20 interface {
	contract.Contract
	BalanceOf(address common.Address) (*big.Int, error)
}

const (
	ABI = "[  {   \"constant\": true,   \"inputs\": [],   \"name\": \"name\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"string\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"spender\",     \"type\": \"address\"    },    {     \"name\": \"value\",     \"type\": \"uint256\"    }   ],   \"name\": \"approve\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [],   \"name\": \"totalSupply\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"uint256\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"from\",     \"type\": \"address\"    },    {     \"name\": \"to\",     \"type\": \"address\"    },    {     \"name\": \"value\",     \"type\": \"uint256\"    }   ],   \"name\": \"transferFrom\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [],   \"name\": \"decimals\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"uint8\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [    {     \"name\": \"owner\",     \"type\": \"address\"    }   ],   \"name\": \"balanceOf\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"uint256\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [],   \"name\": \"symbol\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"string\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"to\",     \"type\": \"address\"    },    {     \"name\": \"value\",     \"type\": \"uint256\"    }   ],   \"name\": \"transfer\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [    {     \"name\": \"owner\",     \"type\": \"address\"    },    {     \"name\": \"spender\",     \"type\": \"address\"    }   ],   \"name\": \"allowance\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"uint256\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"inputs\": [    {     \"name\": \"totalSupply\",     \"type\": \"uint256\"    },    {     \"name\": \"tokenName\",     \"type\": \"string\"    },    {     \"name\": \"tokenSymbol\",     \"type\": \"string\"    },    {     \"name\": \"decimalUnits\",     \"type\": \"uint8\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"constructor\"  },  {   \"anonymous\": false,   \"inputs\": [    {     \"indexed\": true,     \"name\": \"_from\",     \"type\": \"address\"    },    {     \"indexed\": true,     \"name\": \"_to\",     \"type\": \"address\"    },    {     \"indexed\": false,     \"name\": \"_value\",     \"type\": \"uint256\"    }   ],   \"name\": \"Transfer\",   \"type\": \"event\"  },  {   \"anonymous\": false,   \"inputs\": [    {     \"indexed\": true,     \"name\": \"_owner\",     \"type\": \"address\"    },    {     \"indexed\": true,     \"name\": \"_spender\",     \"type\": \"address\"    },    {     \"indexed\": false,     \"name\": \"_value\",     \"type\": \"uint256\"    }   ],   \"name\": \"Approval\",   \"type\": \"event\"  } ]"
)

type ERC20 struct {
	contract.BaseContract
}

func NewERC20(address common.Address) *ERC20 {
	erc20 := &ERC20{}
	erc20.BaseContract = *contract.NewBaseContract(address, ABI)
	return erc20
}

func (c *ERC20) BalanceOf(address common.Address) (*big.Int, error) {
	funcName := "balanceOf"
	input, err := c.Pack(funcName, address)
	if err != nil {
		return nil, err
	}
	contractAddress := c.Address()
	data, err := c.Call(&ethereum.CallMsg{
		To:   &contractAddress,
		Data: input,
	})
	if err != nil {
		return nil, err
	}
	balance := new(big.Int)
	err = c.Unpack(balance, funcName, data)
	if err != nil {
		return nil, err
	}
	return balance, err
}
