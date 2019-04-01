package gameToken

import (
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xxRanger/blockchainUtil/contract"
	"github.com/xxRanger/blockchainUtil/contract/erc20"
	"log"
	"math/big"
)

const (
	ABI = "[  {   \"constant\": true,   \"inputs\": [],   \"name\": \"name\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"string\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"spender\",     \"type\": \"address\"    },    {     \"name\": \"value\",     \"type\": \"uint256\"    }   ],   \"name\": \"approve\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [    {     \"name\": \"tokenId\",     \"type\": \"uint256\"    }   ],   \"name\": \"avatarState\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"uint256\"    },    {     \"name\": \"\",     \"type\": \"uint256\"    },    {     \"name\": \"\",     \"type\": \"bool\"    },    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [],   \"name\": \"totalSupply\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"uint256\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"to\",     \"type\": \"address\"    },    {     \"name\": \"value\",     \"type\": \"uint256\"    }   ],   \"name\": \"reward\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"by\",     \"type\": \"address\"    },    {     \"name\": \"value\",     \"type\": \"uint256\"    }   ],   \"name\": \"consume\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"from\",     \"type\": \"address\"    },    {     \"name\": \"to\",     \"type\": \"address\"    },    {     \"name\": \"value\",     \"type\": \"uint256\"    }   ],   \"name\": \"transferFrom\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [],   \"name\": \"decimals\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"uint8\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [    {     \"name\": \"owner\",     \"type\": \"address\"    }   ],   \"name\": \"ownedAvatars\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"uint256\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"to\",     \"type\": \"address\"    },    {     \"name\": \"tokenId\",     \"type\": \"uint256\"    }   ],   \"name\": \"mint\",   \"outputs\": [],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"tokenId\",     \"type\": \"uint256\"    }   ],   \"name\": \"upgrade\",   \"outputs\": [],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"tokenId\",     \"type\": \"uint256\"    },    {     \"name\": \"user\",     \"type\": \"address\"    }   ],   \"name\": \"equipArmor\",   \"outputs\": [],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [    {     \"name\": \"\",     \"type\": \"uint256\"    }   ],   \"name\": \"avatar\",   \"outputs\": [    {     \"name\": \"gene\",     \"type\": \"uint256\"    },    {     \"name\": \"avatarLevel\",     \"type\": \"uint256\"    },    {     \"name\": \"weaponed\",     \"type\": \"bool\"    },    {     \"name\": \"armored\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [    {     \"name\": \"tokenId\",     \"type\": \"uint256\"    }   ],   \"name\": \"ownerOf\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"address\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [    {     \"name\": \"owner\",     \"type\": \"address\"    }   ],   \"name\": \"balanceOf\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"uint256\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [],   \"name\": \"symbol\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"string\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"to\",     \"type\": \"address\"    },    {     \"name\": \"value\",     \"type\": \"uint256\"    }   ],   \"name\": \"transfer\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [],   \"name\": \"_owner\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"address\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"tokenId\",     \"type\": \"uint256\"    },    {     \"name\": \"user\",     \"type\": \"address\"    }   ],   \"name\": \"equipWeapon\",   \"outputs\": [],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [    {     \"name\": \"owner\",     \"type\": \"address\"    },    {     \"name\": \"spender\",     \"type\": \"address\"    }   ],   \"name\": \"allowance\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"uint256\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"inputs\": [    {     \"name\": \"totalSupply\",     \"type\": \"uint256\"    },    {     \"name\": \"tokenName\",     \"type\": \"string\"    },    {     \"name\": \"tokenSymbol\",     \"type\": \"string\"    },    {     \"name\": \"decimalUnits\",     \"type\": \"uint8\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"constructor\"  },  {   \"anonymous\": false,   \"inputs\": [    {     \"indexed\": false,     \"name\": \"machine\",     \"type\": \"address\"    },    {     \"indexed\": false,     \"name\": \"player\",     \"type\": \"address\"    },    {     \"indexed\": false,     \"name\": \"value\",     \"type\": \"uint256\"    }   ],   \"name\": \"Reward\",   \"type\": \"event\"  },  {   \"anonymous\": false,   \"inputs\": [    {     \"indexed\": false,     \"name\": \"machine\",     \"type\": \"address\"    },    {     \"indexed\": false,     \"name\": \"player\",     \"type\": \"address\"    },    {     \"indexed\": false,     \"name\": \"value\",     \"type\": \"uint256\"    }   ],   \"name\": \"Consume\",   \"type\": \"event\"  },  {   \"anonymous\": false,   \"inputs\": [    {     \"indexed\": true,     \"name\": \"_from\",     \"type\": \"address\"    },    {     \"indexed\": true,     \"name\": \"_to\",     \"type\": \"address\"    },    {     \"indexed\": false,     \"name\": \"_value\",     \"type\": \"uint256\"    }   ],   \"name\": \"Transfer\",   \"type\": \"event\"  },  {   \"anonymous\": false,   \"inputs\": [    {     \"indexed\": true,     \"name\": \"_owner\",     \"type\": \"address\"    },    {     \"indexed\": true,     \"name\": \"_spender\",     \"type\": \"address\"    },    {     \"indexed\": false,     \"name\": \"_value\",     \"type\": \"uint256\"    }   ],   \"name\": \"Approval\",   \"type\": \"event\"  } ]"
)

const (
	FuncConsume     = "consume"
	FuncReward      = "reward"
	FuncMint        = "mint"
	FuncUpgrade = "upgrade"
	FuncEquipWeapon = "equipWeapon"
	FuncEquipArmor  = "equipArmor"
)

type GameToken struct {
	erc20.ERC20
}

type AvatarState struct {
	Gene        *big.Int `json:"gene"`
	AvatarLevel *big.Int `json:"avatarLevel"`
	Weaponed    bool     `json:"weaponed"`
	Armored     bool     `json:"armored"`
}

func NewGameToken(address common.Address) *GameToken {
	gameToken := &GameToken{}
	gameToken.BaseContract = *contract.NewBaseContract(address, ABI)
	return gameToken
}

func (c *GameToken) AvatarState(tokenId *big.Int) (*AvatarState,error) {
	funcName := "avatarState"
	input, err := c.Pack(funcName, tokenId)
	if err != nil {
		return nil,err
	}
	contractAddress := c.Address()
	data, err := c.Call(&ethereum.CallMsg{
		To:   &contractAddress,
		Data: input,
	})
	if err != nil {
		return nil, err
	}
	if len(data)==0 {
		return nil, errors.New("transaction rever");
	}
	var weaponed bool
	var armored bool
	if data[64]==byte(1) {
		weaponed = true
	}
	if data[65]==byte(1) {
		armored = true
	}

	log.Println(data,"len:", len(data))
	avatarState:=&AvatarState{
		Gene: new(big.Int).SetBytes(data[:32]),
		AvatarLevel:new(big.Int).SetBytes(data[32:64]),
		Weaponed: weaponed,
		Armored: armored,
	}

	return avatarState, nil
}

func (c *GameToken) OwnedAvatar(address common.Address) (*big.Int, error) {
	funcName := "ownedAvatars"
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
	tokenId := new(big.Int).SetBytes(data)
	if tokenId.Cmp(big.NewInt(0))==0 {
		err:=errors.New("user have no token")
		return nil,err
	}
	return tokenId, err
}
