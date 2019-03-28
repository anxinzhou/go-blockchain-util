package user

import (
	"Percome19-Crowd-Demo/back-end/contract"
	"blockchainUtil/chain"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
)

type User struct {
	Address common.Address
	privateKey *ecdsa.PrivateKey
	contract *contract.Contract
	ethClient *chain.EthClient
}

func NewUser(address common.Address, privateKey *ecdsa.PrivateKey) *User {
	return &User{
		Address: address,
		privateKey:privateKey,
	}
}

func (u *User) BindEthClient(client *chain.EthClient) {
	u.ethClient = client
}

func (u *User) BindContract(c *contract.Contract) {
	u.contract = c
}

func (u *User) Transfer() {

}
