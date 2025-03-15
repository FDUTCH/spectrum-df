package internal

import (
	"bytes"
	packet2 "github.com/cooldogedev/spectrum/server/packet"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"sync"
)

var BufferPool = sync.Pool{
	New: func() any {
		return bytes.NewBuffer(make([]byte, 0, 256))
	},
}

func addPackets(pool packet.Pool) {
	pool[packet2.IDConnectionRequest] = func() packet.Packet { return &packet2.ConnectionRequest{} }
	pool[packet2.IDLatency] = func() packet.Packet { return &packet2.Latency{} }

	pool[packet2.IDConnectionResponse] = func() packet.Packet { return &packet2.ConnectionResponse{} }
	pool[packet2.IDLatency] = func() packet.Packet { return &packet2.Latency{} }
	pool[packet2.IDTransfer] = func() packet.Packet { return &packet2.Transfer{} }
}
