package gameToken

import (
	"blockchainUtil/contract"
	"blockchainUtil/contract/erc20"
	"github.com/ethereum/go-ethereum/common"
)

const (
	ABI = "[  {   \"constant\": true,   \"inputs\": [],   \"name\": \"name\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"string\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"spender\",     \"type\": \"address\"    },    {     \"name\": \"value\",     \"type\": \"uint256\"    }   ],   \"name\": \"approve\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [],   \"name\": \"totalSupply\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"uint256\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"to\",     \"type\": \"address\"    },    {     \"name\": \"value\",     \"type\": \"uint256\"    }   ],   \"name\": \"reward\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"by\",     \"type\": \"address\"    },    {     \"name\": \"value\",     \"type\": \"uint256\"    }   ],   \"name\": \"consume\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"from\",     \"type\": \"address\"    },    {     \"name\": \"to\",     \"type\": \"address\"    },    {     \"name\": \"value\",     \"type\": \"uint256\"    }   ],   \"name\": \"transferFrom\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [],   \"name\": \"decimals\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"uint8\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [],   \"name\": \"exchangeRate\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"uint256\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"user\",     \"type\": \"address\"    },    {     \"name\": \"amount\",     \"type\": \"uint256\"    }   ],   \"name\": \"mint\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"user\",     \"type\": \"address\"    },    {     \"name\": \"amount\",     \"type\": \"uint256\"    }   ],   \"name\": \"exchangeForEther\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [    {     \"name\": \"owner\",     \"type\": \"address\"    }   ],   \"name\": \"balanceOf\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"uint256\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"user\",     \"type\": \"address\"    }   ],   \"name\": \"exchangeForToken\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": true,   \"stateMutability\": \"payable\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [],   \"name\": \"exchangeBase\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"uint256\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [],   \"name\": \"symbol\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"string\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"user\",     \"type\": \"address\"    },    {     \"name\": \"amount\",     \"type\": \"uint256\"    }   ],   \"name\": \"burn\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"to\",     \"type\": \"address\"    },    {     \"name\": \"value\",     \"type\": \"uint256\"    }   ],   \"name\": \"transfer\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [    {     \"name\": \"owner\",     \"type\": \"address\"    },    {     \"name\": \"spender\",     \"type\": \"address\"    }   ],   \"name\": \"allowance\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"uint256\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"inputs\": [    {     \"name\": \"totalSupply\",     \"type\": \"uint256\"    },    {     \"name\": \"tokenName\",     \"type\": \"string\"    },    {     \"name\": \"tokenSymbol\",     \"type\": \"string\"    }   ],   \"payable\": true,   \"stateMutability\": \"payable\",   \"type\": \"constructor\"  },  {   \"anonymous\": false,   \"inputs\": [    {     \"indexed\": false,     \"name\": \"user\",     \"type\": \"address\"    },    {     \"indexed\": false,     \"name\": \"amount\",     \"type\": \"uint256\"    }   ],   \"name\": \"Mint\",   \"type\": \"event\"  },  {   \"anonymous\": false,   \"inputs\": [    {     \"indexed\": false,     \"name\": \"user\",     \"type\": \"address\"    },    {     \"indexed\": false,     \"name\": \"amount\",     \"type\": \"uint256\"    }   ],   \"name\": \"Burn\",   \"type\": \"event\"  },  {   \"anonymous\": false,   \"inputs\": [    {     \"indexed\": true,     \"name\": \"_from\",     \"type\": \"address\"    },    {     \"indexed\": true,     \"name\": \"_to\",     \"type\": \"address\"    },    {     \"indexed\": false,     \"name\": \"_value\",     \"type\": \"uint256\"    }   ],   \"name\": \"Transfer\",   \"type\": \"event\"  },  {   \"anonymous\": false,   \"inputs\": [    {     \"indexed\": true,     \"name\": \"_owner\",     \"type\": \"address\"    },    {     \"indexed\": true,     \"name\": \"_spender\",     \"type\": \"address\"    },    {     \"indexed\": false,     \"name\": \"_value\",     \"type\": \"uint256\"    }   ],   \"name\": \"Approval\",   \"type\": \"event\"  } ]"
)

type GameToken struct {
	*erc20.ERC20
}

func NewGameToken(address common.Address) *GameToken {
	gameToken := &GameToken{}
	gameToken.BaseContract = contract.NewBaseContract(address, ABI)
	return gameToken
}