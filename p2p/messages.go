package p2p

type MessageKind int

const (
	MessageNewestBlock MessageKind = 1
	MessageAllBlocksRequest MessageKind = 2
	MessageAllBlocksRespinse MessageKind = 3
)

type Message struct {
	Kind MessageKind
	Payload []byte
}

Last