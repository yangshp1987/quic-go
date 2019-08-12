package quic

import (
	"fmt"

	"github.com/lucas-clemente/quic-go/internal/qerr"

	"github.com/lucas-clemente/quic-go/internal/wire"

	"github.com/lucas-clemente/quic-go/internal/protocol"
)

type connIDEntry struct {
	ConnID              protocol.ConnectionID
	StatelessResetToken [16]byte
}

type connectionIDManager struct {
	m map[uint64]connIDEntry

	deletedPriorTo uint64
}

func newConnectionIDManager() *connectionIDManager {
	return &connectionIDManager{m: make(map[uint64]connIDEntry)}
}

func (c *connectionIDManager) HandleNewConnectionIDFrame(f *wire.NewConnectionIDFrame) error {
	if f.RetirePriorTo > f.SequenceNumber {
		return qerr.Error(qerr.ProtocolViolation, fmt.Sprintf("Invalid RetirePriorTo value: %d. Sequence number was: %d", f.RetirePriorTo, f.SequenceNumber))
	}
	c.deletePriorTo(f.RetirePriorTo)
	return c.add(f.SequenceNumber, f.ConnectionID, f.StatelessResetToken)
}

func (c *connectionIDManager) add(id uint64, connID protocol.ConnectionID, token [16]byte) error {
	if id < c.deletedPriorTo {
		return nil
	}
	if entry, ok := c.m[id]; ok && (!entry.ConnID.Equal(connID) || entry.StatelessResetToken != token) {
		return qerr.Error(qerr.ProtocolViolation, fmt.Sprintf("Connection ID %d already exists, with different values", id))
	}
	c.m[id] = connIDEntry{ConnID: connID, StatelessResetToken: token}
	return nil
}

func (c *connectionIDManager) deletePriorTo(priorTo uint64) {
	if priorTo < c.deletedPriorTo {
		return
	}
	for id := range c.m {
		if id < priorTo {
			delete(c.m, id)
		}
	}
}
