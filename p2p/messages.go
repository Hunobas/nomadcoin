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
	b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
	utils.HandleErr(err)
	m := makeMessage(MessageNewestBlock, b)
	p.inbox <- m
}

func handleMsg(m *Message, p *peer) {
	switch m.Kind {
	case MessageNewestBlock:
		var newestBlock blockchain.Block
		utils.HandleErr(json.Unmarshal(m.Payload, &newestBlock))
		fmt.Println(newestBlock)
	case MessageAllBlocksRequest:
	case MessageAllBlocksResponse:
	}
}
