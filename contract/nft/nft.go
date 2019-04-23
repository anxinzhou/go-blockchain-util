package nft

import (
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xxRanger/blockchainUtil/contract"
	"math/big"
)

const (
	ABI = "[  {   \"constant\": true,   \"inputs\": [    {     \"name\": \"token_id\",     \"type\": \"uint256\"    }   ],   \"name\": \"infoOfToken\",   \"outputs\": [    {     \"name\": \"nft_type\",     \"type\": \"string\"    },    {     \"name\": \"nft_name\",     \"type\": \"string\"    },    {     \"name\": \"nft_ldef_index\",     \"type\": \"string\"    },    {     \"name\": \"nft_life_index\",     \"type\": \"uint256\"    },    {     \"name\": \"nft_power_index\",     \"type\": \"uint256\"    },    {     \"name\": \"nft_character_id\",     \"type\": \"string\"    },    {     \"name\": \"public_key\",     \"type\": \"bytes32\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"to\",     \"type\": \"address\"    },    {     \"name\": \"token_id\",     \"type\": \"uint256\"    },    {     \"name\": \"nft_type\",     \"type\": \"string\"    },    {     \"name\": \"nft_name\",     \"type\": \"string\"    },    {     \"name\": \"nft_ldef_index\",     \"type\": \"string\"    },    {     \"name\": \"nft_life_index\",     \"type\": \"uint256\"    },    {     \"name\": \"nft_power_index\",     \"type\": \"uint256\"    },    {     \"name\": \"nft_character_id\",     \"type\": \"string\"    },    {     \"name\": \"public_key\",     \"type\": \"bytes32\"    }   ],   \"name\": \"mint\",   \"outputs\": [],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [    {     \"name\": \"tokenId\",     \"type\": \"uint256\"    }   ],   \"name\": \"ownerOf\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"address\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [    {     \"name\": \"owner\",     \"type\": \"address\"    }   ],   \"name\": \"balanceOf\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"uint256\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [    {     \"name\": \"token_id\",     \"type\": \"uint256\"    }   ],   \"name\": \"ldefOfToken\",   \"outputs\": [    {     \"name\": \"nft_ldef_index\",     \"type\": \"string\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"to\",     \"type\": \"address\"    },    {     \"name\": \"token_id\",     \"type\": \"uint256\"    }   ],   \"name\": \"transfer\",   \"outputs\": [],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"inputs\": [],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"constructor\"  } ]"
)

const (
	FuncTransfer = "transfer"
	FuncMint     = "mint"
)

type NFTInfo struct {
	Type        string   `json:"type"`
	Name        string   `json:"name"`
	LdefIndex   string   `json:"ldefIndex"`
	LifeIndex   *big.Int `json:"lifeIndex"`
	PowerIndex  *big.Int `json:"powerIndex"`
	CharacterId string   `json:"characterId"`
	PublicKey   []byte   `json:"publicKey"`
}

type AvatarState struct {
	Gene        *big.Int `json:"gene"`
	AvatarLevel *big.Int `json:"avatarLevel"`
	Weaponed    bool     `json:"weaponed"`
	Armored     bool     `json:"armored"`
}

type NFT struct {
	contract.BaseContract
}

func NewNFT(address common.Address) *NFT {
	nft := &NFT{}
	nft.BaseContract = *contract.NewBaseContract(address, ABI)
	return nft
}

func (c *NFT) InfoOfToken(tokenId *big.Int) (*NFTInfo, error) {
	funcName := "infoOfToken"
	input, err := c.Pack(funcName, tokenId)
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
	if len(data) == 0 {
		return nil, errors.New("transaction rever");
	}

	lenNFTType := new(big.Int).SetBytes(data[:32]).Int64()
	nftType := string(data[32 : 32+lenNFTType])

	nftNameByteStart := 32 + lenNFTType
	nftNameLen := new(big.Int).SetBytes(data[nftNameByteStart:nftNameByteStart+32]).Int64()
	nftName := string(data[nftNameByteStart+32 : nftNameByteStart+32+nftNameLen])

	nftLdefStart := nftNameByteStart + 32 + nftNameLen
	nftLdefLen := new(big.Int).SetBytes(data[nftLdefStart:nftLdefStart+32]).Int64()
	nftLdefIndex := string(data[nftLdefStart+32 : nftLdefStart+32+nftLdefLen])

	nftLifeStart := nftLdefStart + 32 + nftLdefLen
	nftLifeIndex := new(big.Int).SetBytes(data[nftLifeStart : nftLifeStart+32])

	nftPowerStart := nftLifeStart + 32
	nftPowerIndex := new(big.Int).SetBytes(data[nftPowerStart : nftPowerStart+32])

	nftCharacterStart := nftPowerStart + 32
	nftCharacterLen := new(big.Int).SetBytes(data[nftCharacterStart:nftCharacterStart+32]).Int64()
	nftCharacter := string(data[nftCharacterStart+32 : nftCharacterStart+32+nftCharacterLen])

	nftPublicKeyStart := nftCharacterStart + 32 + nftCharacterLen
	nftPublicKeyLen := new(big.Int).SetBytes(data[nftCharacterStart:nftCharacterStart+32]).Int64()
	nftPublickey := data[nftPublicKeyStart+32 : nftPublicKeyStart+32+nftPublicKeyLen]

	nftInfo := &NFTInfo{
		Type:        nftType,
		Name:        nftName,
		LdefIndex:   nftLdefIndex,
		LifeIndex:   nftLifeIndex,
		PowerIndex:  nftPowerIndex,
		CharacterId: nftCharacter,
		PublicKey:   nftPublickey,
	}

	return nftInfo, nil
}

func (c *NFT) LdefOfToken(tokenId *big.Int) (*big.Int, error) {
	funcName := "ownedAvatars"
	input, err := c.Pack(funcName, tokenId)
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
	ldef := new(big.Int).SetBytes(data)
	if ldef.Cmp(big.NewInt(0)) == 0 {
		err := errors.New("user have no token")
		return nil, err
	}
	return ldef, err
}
