package pbcosmos

import (
	"encoding/hex"
	"time"

	"github.com/streamingfast/bstream"
)

// previousID := ""
// if block.Header.LastBlockId != nil {
// 	previousID = hex2string(block.Header.LastBlockId.Hash)
// }
// blk := &bstream.Block{
// 	Id:             hex2string(block.Header.Hash),
// 	PreviousId:     previousID,
// 	Number:         uint64(block.Header.Height),
// 	LibNum:         uint64(block.Header.Height - 1),
// 	Timestamp:      parseTimestamp(block.Header.Time),
// 	PayloadKind:    pbbstream.Protocol_COSMOS,
// 	PayloadVersion: 1,
// }

// if block.Header.Height == bstream.GetProtocolFirstStreamableBlock {
// 	blk.LibNum = bstream.GetProtocolFirstStreamableBlock
// 	blk.PreviousId = ""
// }

func (b *Block) GetFirehoseBlockID() string {
	return hex.EncodeToString(b.Header.Hash)
}

func (b *Block) GetFirehoseBlockNumber() uint64 {
	return b.Header.Height
}

func (b *Block) GetFirehoseBlockParentID() string {
	if b.Header.Height == bstream.GetProtocolFirstStreamableBlock {
		return ""
	}

	if b.Header.LastBlockId != nil {
		return hex.EncodeToString(b.Header.LastBlockId.Hash)
	}

	return ""
}

func (b *Block) GetFirehoseBlockParentNumber() uint64 {
	if b.Header.Height == bstream.GetProtocolFirstStreamableBlock {
		return bstream.GetProtocolFirstStreamableBlock
	}

	return b.Header.Height - 1
}

func (b *Block) GetFirehoseBlockLIBNum() uint64 {
	return b.GetFirehoseBlockParentNumber()
}

func (b *Block) GetFirehoseBlockTime() time.Time {
	ts := b.Header.Time
	return time.Unix(ts.Seconds, int64(ts.Nanos)).UTC()
}

func (b *Block) GetFirehoseBlockVersion() int32 {
	return 1
}
