package p2p

import (
	"encoding/json"
	"fmt"

	"github.com/Hunobas/nomadcoin/blockchain"
	"github.com/Hunobas/nomadcoin/utils"
)

type MessageKind int

const (
	MessageNewestBlock MessageKind = iota
	MessageAllBlocksRequest
	MessageAllBlocksResponse
	MessageNewBlockNotify
	MessageNewTxNotify
)

type Message struct {
	Kind    MessageKind
	Payload []byte
}

func makeMessage(kind MessageKind, payload interface{}) []byte {
	m := Message{
		Kind: kind,
		// 보낼 payload의 정보가 각각 다르기 때문에 Json으로 바꿈. (모든 블록, newest블록, 트랜젝션, 등등...)
		Payload: utils.ToJSON(payload),
	}
	// 다른 port끼리 통신하기 위해서 byte 형태로써 Json으로 바꿈. (ToBytes를 안하는 이유 : 크로스 플랫폼(다른 언어)에서 주고받을 수 있도록 하기 위해)
	return utils.ToJSON(m)
}

func sendNewestBlock(p *peer) {
	fmt.Printf("Sending newest block to %s\n", p.key)
	b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
	utils.HandleErr(err)
	m := makeMessage(MessageNewestBlock, b)
	p.inbox <- m
}

func requestAllBlocks(p *peer) {
	m := makeMessage(MessageAllBlocksRequest, nil)
	p.inbox <- m
}

func sendAllBlocks(p *peer) {
	m := makeMessage(MessageAllBlocksResponse, blockchain.Blocks(blockchain.Blockchain()))
	p.inbox <- m
}

func notifyNewBlock(b *blockchain.Block, p *peer) {
	m := makeMessage(MessageNewBlockNotify, b)
	p.inbox <- m
}

func notifyNewTx(tx *blockchain.Tx, p *peer) {
	m := makeMessage(MessageNewTxNotify, tx)
	p.inbox <- m
}

func handleMsg(m *Message, p *peer) {
	switch m.Kind {
	case MessageNewestBlock:
		fmt.Printf("Received the newest block from %s\n", p.key)
		var newestBlock *blockchain.Block
		utils.HandleErr(json.Unmarshal(m.Payload, &newestBlock))
		b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
		utils.HandleErr(err)
		if newestBlock.Height >= b.Height {
			fmt.Printf("Requesting all blocks from %s\n", p.key)
			requestAllBlocks(p)
		} else {
			fmt.Printf("Sending newest block to %s\n", p.key)
			sendNewestBlock(p)
		}
	case MessageAllBlocksRequest:
		fmt.Printf("%s wants all the blocks.\n", p.key)
		sendAllBlocks(p)
	case MessageAllBlocksResponse:
		fmt.Printf("Received all the blocks from %s\n", p.key)
		var allBlocks []*blockchain.Block
		utils.HandleErr(json.Unmarshal(m.Payload, &allBlocks))
		blockchain.Blockchain().Replace(allBlocks)
	case MessageNewBlockNotify:
		var newBlock *blockchain.Block
		utils.HandleErr(json.Unmarshal(m.Payload, &newBlock))
		blockchain.Blockchain().AddPeerBlock(newBlock)
	case MessageNewTxNotify:
		var newTx *blockchain.Tx
		utils.HandleErr(json.Unmarshal(m.Payload, &newTx))
		blockchain.Mempool().AddPeerTx(newTx)
	}
}
