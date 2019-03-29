package main

import (
	"blockchainUtil/contract"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	c:=contract.BaseContract{}
	err:=c.Import("abi.json")
	if err!=nil {
		panic(err)
	}
	abi:=c.ABI()
	log.Println(abi.Methods)
	log.Println(abi.Events)
	log.Println(c.EventSigByName("Aggregate"))
}