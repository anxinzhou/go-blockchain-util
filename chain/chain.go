package chain

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

const (
	MAX_WAITING_BLOCK = 50 // come from web3js, if wait more than 50 blocks, transaction time out
	LOG_BUFFER        = 128
)

type EthClient struct {
	Client  *ethclient.Client
	ChainID *big.Int // chain Id of current blockchain
}

func NewEthClient(socket string) (*EthClient, error) {
	c := &EthClient{}
	var err error
	c.Client, err = ethclient.Dial(socket)
	if err != nil {
		return nil, err
	}
	chainId, err := c.GetChainId()
	if err != nil {
		return nil, err
	}
	c.ChainID = chainId
	return c, nil
}

func (c *EthClient) Close() {
	if c.Client == nil {
		panic("close nil eth client")
	}
	c.Client.Close()
}

func (c *EthClient) GetChainId() (*big.Int, error) {
	chainID, err := c.Client.NetworkID(context.Background())
	return chainID, err
}

func (c *EthClient) GetNonce(address common.Address) (uint64, error) {
	nonce, err := c.Client.PendingNonceAt(context.Background(), address)
	return nonce, err
}

func (c *EthClient) GetEther(address common.Address) (*big.Int, error) {
	balance, err := c.Client.BalanceAt(context.Background(), address, nil)
	return balance, err
}

func (c *EthClient) GetTransactionReceipt(txHash common.Hash) (chan *types.Receipt, chan error) {
	count := 0
	ch := make(chan *types.Header)
	receipt := make(chan *types.Receipt, 1)
	receiptError := make(chan error, 1)
	go func() {
		sub, err := c.Client.SubscribeNewHead(context.Background(), ch)
		if err != nil {
			receiptError <- err
			return
		}
		for {
			select {
			case err := <-sub.Err():
				log.Println(err.Error())
				receiptError <- err
				return
			case <-ch:
				count += 1
				if count >= MAX_WAITING_BLOCK {
					receiptError <- errors.New("transaction time out")
				} else {
					r, err := c.Client.TransactionReceipt(context.Background(), txHash)
					if err == nil {
						log.Println("found receipt")
						receipt <- r
						return
					}
				}
			}
		}
	}()
	return receipt, receiptError
}

func (c *EthClient) Send(tx *types.Transaction) chan error {
	txError := make(chan error, 1)
	go func() {
		err := c.Client.SendTransaction(context.Background(), tx)
		if err != nil {
			log.Println(err.Error())
			txError <- err
			return
		}
		receipt, receiptError := c.GetTransactionReceipt(tx.Hash())
		select {
		case err := <-receiptError:
			log.Println(err.Error())
			txError <- err
		case r := <-receipt:
			if r.Status == 0 {
				err:= errors.New("transaction revert")
				log.Println(err.Error())
				txError <- err
			} else {
				log.Println("transaction success")
				txError <- nil
			}
		}
	}()
	return txError
}

func (c *EthClient) Call(msg *ethereum.CallMsg) ([]byte, error) {
	rVal, err := c.Client.CallContract(context.Background(), *msg, nil)
	return rVal, err
}

func (c *EthClient) SubscribeEvent(query *ethereum.FilterQuery) (chan types.Log, <-chan error) {
	logs := make(chan types.Log, LOG_BUFFER)
	sub, err := c.Client.SubscribeFilterLogs(context.Background(), *query, logs)
	if err != nil {
		subScribeError := make(chan error, 1)
		subScribeError <- err
		return logs, subScribeError
	}
	return logs, sub.Err()
}
