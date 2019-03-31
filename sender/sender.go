package sender

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/xxRanger/blockchainUtil/chain"
	"github.com/xxRanger/blockchainUtil/contract"
	"math/big"
	"sync"
)

func init() {
	defaultPrivateChainGasLimit = 3000000
	defaultPrivateChainGasPrice = big.NewInt(0)
	defaultPrivateChainValue = big.NewInt(0)

	defaultPublicChainGasLimit = 3000000
	defaultPublicChainGasPrice = big.NewInt(1000000000)  // 1 Gwei
	defaultPublicChainValue = big.NewInt(0)
}

// chain kind
const (
	CHAIN_KIND_PRIVATE = 0
	CHAIN_KIND_PUBLIC  = 1
)

var (
	defaultPrivateChainGasLimit uint64
	defaultPrivateChainGasPrice *big.Int
	defaultPrivateChainValue    *big.Int

	defaultPublicChainGasLimit uint64
	defaultPublicChainGasPrice *big.Int
	defaultPublicChainValue *big.Int
)

type SendOpts struct {
	To       *common.Address // nil for contract creation
	GasLimit uint64
	GasPrice *big.Int
	Value    *big.Int
	Data     []byte
}

type User struct {
	Address    common.Address
	privateKey *ecdsa.PrivateKey
	*ChainInfo
}

type ChainInfo struct {
	ethClient *chain.EthClient
	nonce     *Nonce
	chainKind int
}

type Nonce struct {
	v     uint64 // only used for compare result get from blockchain to send transaction concurrently
	mutex *sync.Mutex
}

func NewUser(address common.Address, privateKey *ecdsa.PrivateKey) *User {
	return &User{
		Address:    address,
		privateKey: privateKey,
	}
}

// bind and get nonce if non-nil
func (u *User) BindEthClient(client *chain.EthClient, chainKind int) {
	u.ethClient = client
	u.chainKind = chainKind
	u.nonce = &Nonce{
		v:     0,
		mutex: &sync.Mutex{},
	}
}

func (u *User) getNonce() (uint64, error) {
	if u.ethClient == nil {
		return 0, errors.New("use have not binded ethclient")
	}
	u.nonce.mutex.Lock()
	defer u.nonce.mutex.Unlock()
	nonceFromChain, err := u.ethClient.GetNonce(u.Address)
	if err != nil {
		return 0, err
	}
	if nonceFromChain > u.nonce.v {
		u.nonce.v = nonceFromChain // update u.nonce.v
	}
	u.nonce.v += 1
	return u.nonce.v - 1, nil
}

func (u *User) Transfer(to common.Address, value *big.Int) chan error {
	//TODO
	return nil
}

func (u *User) SignTransaction(tx *types.Transaction) (*types.Transaction, error) {
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(u.ethClient.ChainID), u.privateKey)
	return signedTx, err
}

func (u *User) SendTransaction(tx *types.Transaction) chan error {
	return u.ethClient.Send(tx)
}

func (u *User) SendAndSignTransaction(tx *types.Transaction) chan error {
	txError := make(chan error, 1)
	signedTx, err := u.SignTransaction(tx)
	if err != nil {
		txError <- err
		return txError
	}
	return u.SendTransaction(signedTx)
}

func (u *User) CallFunction(c contract.Contract, funcName string, args ...interface{}) ([]byte, error) {
	input, err := c.Pack(funcName, args...)
	if err != nil {
		return nil, err
	}

	contractAddress := c.Address()
	rVal, err := u.ethClient.Client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contractAddress,
		Data: input,
	}, nil)
	return rVal, err
}

func (u *User) SendFunction(c contract.Contract, opt *SendOpts, funcName string, args ... interface{}) chan error {
	txError := make(chan error, 1)
	input, err := c.Pack(funcName, args...)
	if err != nil {
		txError <- err
		return txError
	}
	nonce, err := u.getNonce()
	if err != nil {
		txError <- err
		return txError
	}

	contractAddress := c.Address()
	var tx *types.Transaction
	if opt == nil {
		if u.chainKind == CHAIN_KIND_PRIVATE {
			tx = types.NewTransaction(nonce, contractAddress, defaultPublicChainValue, defaultPublicChainGasLimit, defaultPublicChainGasPrice, input)
		} else if u.chainKind == CHAIN_KIND_PUBLIC {
			tx = types.NewTransaction(nonce, contractAddress, defaultPrivateChainValue, defaultPrivateChainGasLimit, defaultPrivateChainGasPrice, input)
		} else {
			err:= errors.New("unknown chain kind")
			txError <- err
			return txError
		}
	} else {
		tx = types.NewTransaction(nonce, *opt.To, opt.Value, opt.GasLimit, opt.GasPrice, input)
	}
	return u.SendAndSignTransaction(tx)
}
