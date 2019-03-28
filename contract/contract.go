package contract

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"io/ioutil"
	"strings"
)

type Contract interface {
	Address() common.Address
	ABI() *abi.ABI
}

type BaseContract struct {
	address common.Address
	abi *abi.ABI
}

func (c *BaseContract) Import(file string) error {
	type config struct {
		Address string `json:"address"`
		ABI string `json:"abi"`
	}
	contractConfig:=&config{}
	data,err:=ioutil.ReadFile(file)
	if err!=nil {
		return err
	}
	err = json.Unmarshal(data,contractConfig)
	if err!=nil {
		return err
	}
	c.address=common.HexToAddress(contractConfig.Address)
	err=c.SetABIFromString(contractConfig.ABI)
	return err
}

func (c *BaseContract) Address() common.Address {
	return c.address
}

func (c *BaseContract) ABI() *abi.ABI {
	return c.abi
}

func (c *BaseContract) SetABIFromString(a string) error {
	var err error
	c.abi = new(abi.ABI)
	*c.abi,err = abi.JSON(strings.NewReader(a))
	return err
}

func (c *BaseContract) PackFunction(funcName string, args ...interface{}) ([]byte,error) {
	input,err:= c.abi.Pack(funcName, args...)
	if err!=nil {
		return nil,err
	}
	return input,nil
}




