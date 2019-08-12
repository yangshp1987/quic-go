package quic

import (
	"github.com/lucas-clemente/quic-go/internal/protocol"
	"github.com/lucas-clemente/quic-go/internal/wire"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("Conn ID Manager", func() {
	var m *connectionIDManager

	BeforeEach(func() {
		m = newConnectionIDManager()
	})

	It("errors when receiving different connection IDs for the same ID", func() {
		Expect(m.HandleNewConnectionIDFrame(&wire.NewConnectionIDFrame{
			SequenceNumber: 1,
			ConnectionID:   protocol.ConnectionID{0xde, 0xad, 0xbe, 0xef},
		})).To(Succeed())
		Expect(m.HandleNewConnectionIDFrame(&wire.NewConnectionIDFrame{
			SequenceNumber: 1,
			ConnectionID:   protocol.ConnectionID{0xde, 0xca, 0xfb, 0xad},
		})).To(MatchError("PROTOCOL_VIOLATION: Connection ID 1 already exists, with different values"))
	})

	It("errors when receiving different stateless reset tokens for the same ID", func() {
		Expect(m.HandleNewConnectionIDFrame(&wire.NewConnectionIDFrame{
			SequenceNumber:      1,
			ConnectionID:        protocol.ConnectionID{0xde, 0xad, 0xbe, 0xef},
			StatelessResetToken: [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		})).To(Succeed())
		Expect(m.HandleNewConnectionIDFrame(&wire.NewConnectionIDFrame{
			SequenceNumber:      1,
			ConnectionID:        protocol.ConnectionID{0xde, 0xad, 0xbe, 0xef},
			StatelessResetToken: [16]byte{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		})).To(MatchError("PROTOCOL_VIOLATION: Connection ID 1 already exists, with different values"))
	})

	It("errors when receiving a frame with a too high RetirePriorTo value", func() {
		Expect(m.HandleNewConnectionIDFrame(&wire.NewConnectionIDFrame{
			SequenceNumber: 10,
			RetirePriorTo:  11,
		})).To(MatchError("PROTOCOL_VIOLATION: Invalid RetirePriorTo value: 11. Sequence number was: 10"))
	})
})
