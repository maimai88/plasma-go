package core

import (
	"github.com/ethereum/go-ethereum/common"
)

type Transaction struct {
	// first input
	BlkNum1  uint64 `json:"blkNum1"`
	TxIndex1 uint64	`json:"txIndex1"`
	OIndex1  uint64	`json:"oIndex1"`
	Sig1     []byte	`json:"sig1"`
	Spend1   bool

	// second input
	BlkNum2  uint64 `json:"blkNum2"`
	TxIndex2 uint64	`json:"txIndex2"`
	OIndex2  uint64	`json:"oIndex2"`
	Sig2     []byte	`json:"sig2"`
	Spend2   bool

	// output
	newOwner1 common.Address
	amount1 uint64
	newOwner2 common.Address
	amount2 uint64
}

func MakeTransaction(address common.Address,amount uint64) *Transaction {
	return &Transaction{
		newOwner1:address,amount1:amount,
	}
}

type UTXO struct {
	blkNum uint64
	txIndex uint64
	oIndex uint64
	sig []byte
	spend bool
}

func MakeUTXO(blkNum uint64, txIndex uint64, oIndex uint64,sig []byte, spend bool) UTXO {
	utxo := UTXO{
		blkNum:blkNum,
		txIndex:txIndex,
		oIndex:oIndex,
		sig:sig,
		spend:spend,
	}
	return utxo
}