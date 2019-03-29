package contract

import (
	"blockchainUtil/chain"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"log"
	"strings"
)

// base solidity type
var solidityType map[string]bool

func init() {
	solidityType = make(map[string]bool)
	solidityType["uint256"] = true
	solidityType["address"] = true
	solidityType["bytes"] = true
	solidityType["bool"] = true
}

type Contract interface {
	Address() common.Address
	ABI() *abi.ABI
	Pack(funcName string, args ...interface{}) ([]byte, error)
	EventSigByName(name string) (common.Hash, error)
	Unpack(v interface{}, name string, output []byte) error
	EventHandlers() map[string]eventHandler
}

type eventHandler func([]byte)

type BaseContract struct {
	address common.Address
	abi     *abi.ABI
	ethClient *chain.EthClient
	eventHandlers map[string]eventHandler   // register subscribe event and handlers
}

func NewBaseContract(address common.Address, abi string) *BaseContract {
	c:=&BaseContract{
		address:address,
		eventHandlers:make(map[string]eventHandler),
	}
	c.SetABIFromString(abi)
	return c
}

func (c *BaseContract) RegisterHandler(event string, handler eventHandler) error {
	if c.eventHandlers==nil {
		c.eventHandlers = make(map[string]eventHandler)
	}
	eventSigHash,err:=c.EventSigByName(event)
	if err!=nil {
		return err
	}
	c.eventHandlers[eventSigHash.Hex()] = handler
	return nil
}

func (c *BaseContract) UnRegisterHandler(event string)  {
	eventSigHash,_:= c.EventSigByName(event)
	delete(c.eventHandlers,eventSigHash.Hex())
}

func (c *BaseContract) SetEventHandlers(eventHandlers map[string]eventHandler) {
	c.eventHandlers = eventHandlers
}

func (c *BaseContract) EventHandlers() map[string]eventHandler {
	return c.eventHandlers
}

func (c *BaseContract) Import(file string) error {
	type config struct {
		Address string `json:"address"`
		ABI     string `json:"abi"`
	}
	contractConfig := &config{}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, contractConfig)
	if err != nil {
		return err
	}
	c.address = common.HexToAddress(contractConfig.Address)
	err = c.SetABIFromString(contractConfig.ABI)
	return err
}

func (c *BaseContract) BindClient(client *chain.EthClient) {
	c.ethClient = client
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
	*c.abi, err = abi.JSON(strings.NewReader(a))
	return err
}

func (c *BaseContract) Pack(funcName string, args ...interface{}) ([]byte, error) {
	input, err := c.abi.Pack(funcName, args...)
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (c *BaseContract) EventSigByName(name string) (common.Hash, error) {
	event, ok := c.abi.Events[name]
	if !ok {
		return common.Hash{}, errors.New("on such event")
	}
	var sigString strings.Builder

	eventBody := []rune(event.String())
	sigString.WriteString(event.Name + "(")
	i := 0
	for i < len(eventBody) && eventBody[i] != '(' {
		i++
	}
	if i >= len(eventBody) {
		return common.Hash{}, errors.New("wrong event body")
	}
	i++
	j := i
	for i < len(eventBody) {
		for j < len(eventBody) && eventBody[j] != ' ' && eventBody[j] != ',' && eventBody[j] != ')' {
			j++
		}
		log.Println(i,j)
		tmp := string(eventBody[i:j])
		if _, ok := solidityType[tmp]; ok {
			sigString.WriteString(tmp)
		}
		if j < len(eventBody) {
			if eventBody[j] == ',' {
				sigString.WriteString(",")
			}
		}
		j++
		i = j
	}
	sigString.WriteString(")")
	sigByte:= []byte(sigString.String())
	return crypto.Keccak256Hash(sigByte),nil
}

func (c *BaseContract) Unpack(v interface{}, name string, output []byte) error {
	err:=c.abi.Unpack(v, name, output)
	return err
}

func (c *BaseContract) Subscribe() error {
	topics:= make([][]common.Hash,1)
	topics[0] = make([]common.Hash,0,len(c.eventHandlers))
	for eventSig,_:= range c.eventHandlers {
		topics[0]=append(topics[0],common.HexToHash(eventSig))
	}
	query:= &ethereum.FilterQuery{
		Addresses: [] common.Address{c.address},
		Topics: topics,
	}
	logs,subScribeErr:= c.ethClient.SubscribeEvent(query)
	for {
		select {
		case err:=<-subScribeErr:
			log.Println(err.Error())
		case vLog:= <-logs:
			eventSig:= vLog.Topics[0].Hex()
			if handler,ok:=c.eventHandlers[eventSig];ok {
				go handler(vLog.Data)
			}
		}
	}
}
