package blockchain

import (
	"time"

	"github.com/Hunobas/nomadcoin/utils"
)

const (
	minerReward int = 50
)

type mempool struct {
	Txs []*Tx
}

var Mempool *mempool = &mempool{}

type Tx struct {
	Id        string   `json:"id"`
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"txIns"`
	TxOuts    []*TxOut `json:"txOuts"`
}

type TxIn struct {
	TxID  string `json:"txID"`
	Index int    `json:"index"`
	Owner string `json:"owner"`
}

type TxOut struct {
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

type UTxOut struct {
	TxID   string `json:"txID"`
	Index  int    `json:"index"`
	Amount int    `json:"amount"`
}

func (t *Tx) getId() {
	t.Id = utils.Hash(t)
}

func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"", -1, "SATOSHI"},
	}
	txOuts := []*TxOut{
		{address, minerReward},
	}
	tx := Tx{
		Id:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId()
	return &tx
}

func makeTx(from, to string, amount int) (*Tx, error) {

}

func (m *mempool) AddTx(to string, amount int) error {
	tx, err := makeTx("huno", to, amount)
	if err != nil {
		return err
	}
	m.Txs = append(m.Txs, tx)
	return nil
}

func (m *mempool) txToConfirm() []*Tx {
	coinbase := makeCoinbaseTx("huno")
	txs := m.Txs
	txs = append(txs, coinbase)
	m.Txs = nil
	return txs
}
