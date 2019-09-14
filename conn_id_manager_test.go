package quic

import (
	"github.com/lucas-clemente/quic-go/internal/protocol"
	"github.com/lucas-clemente/quic-go/internal/wire"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Connection ID Manager", func() {
	var (
		m          *connIDManager
		frameQueue []wire.Frame
	)

	BeforeEach(func() {
		frameQueue = nil
		m = newConnIDManager(func(f wire.Frame) {
			frameQueue = append(frameQueue, f)
		})
	})

	It("returns nil if empty", func() {
		c, rt := m.Get()
		Expect(c).To(BeNil())
		Expect(rt).To(BeNil())
	})

	It("adds and gets connection IDs", func() {
		Expect(m.Add(&wire.NewConnectionIDFrame{
			SequenceNumber:      10,
			ConnectionID:        protocol.ConnectionID{2, 3, 4, 5},
			StatelessResetToken: [16]byte{0xe, 0xd, 0xc, 0xb, 0xa, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		})).To(Succeed())
		Expect(m.Add(&wire.NewConnectionIDFrame{
			SequenceNumber:      4,
			ConnectionID:        protocol.ConnectionID{1, 2, 3, 4},
			StatelessResetToken: [16]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0xa, 0xb, 0xc, 0xd, 0xe},
		})).To(Succeed())
		c1, rt1 := m.Get()
		Expect(c1).To(Equal(protocol.ConnectionID{1, 2, 3, 4}))
		Expect(*rt1).To(Equal([16]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0xa, 0xb, 0xc, 0xd, 0xe}))
		c2, rt2 := m.Get()
		Expect(c2).To(Equal(protocol.ConnectionID{2, 3, 4, 5}))
		Expect(*rt2).To(Equal([16]byte{0xe, 0xd, 0xc, 0xb, 0xa, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}))
		c3, rt3 := m.Get()
		Expect(c3).To(BeNil())
		Expect(rt3).To(BeNil())
	})

	It("accepts duplicates", func() {
		f := &wire.NewConnectionIDFrame{
			SequenceNumber:      1,
			ConnectionID:        protocol.ConnectionID{1, 2, 3, 4},
			StatelessResetToken: [16]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0xa, 0xb, 0xc, 0xd, 0xe},
		}
		Expect(m.Add(f)).To(Succeed())
		Expect(m.Add(f)).To(Succeed())
		c1, rt1 := m.Get()
		Expect(c1).To(Equal(protocol.ConnectionID{1, 2, 3, 4}))
		Expect(*rt1).To(Equal([16]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0xa, 0xb, 0xc, 0xd, 0xe}))
		c2, rt2 := m.Get()
		Expect(c2).To(BeNil())
		Expect(rt2).To(BeNil())
	})

	It("rejects duplicates with different connection IDs", func() {
		Expect(m.Add(&wire.NewConnectionIDFrame{
			SequenceNumber: 42,
			ConnectionID:   protocol.ConnectionID{1, 2, 3, 4},
		})).To(Succeed())
		Expect(m.Add(&wire.NewConnectionIDFrame{
			SequenceNumber: 42,
			ConnectionID:   protocol.ConnectionID{2, 3, 4, 5},
		})).To(MatchError("received conflicting connection IDs for sequence number 42"))
	})

	It("rejects duplicates with different connection IDs", func() {
		Expect(m.Add(&wire.NewConnectionIDFrame{
			SequenceNumber:      42,
			ConnectionID:        protocol.ConnectionID{1, 2, 3, 4},
			StatelessResetToken: [16]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0xa, 0xb, 0xc, 0xd, 0xe},
		})).To(Succeed())
		Expect(m.Add(&wire.NewConnectionIDFrame{
			SequenceNumber:      42,
			ConnectionID:        protocol.ConnectionID{1, 2, 3, 4},
			StatelessResetToken: [16]byte{0xe, 0xd, 0xc, 0xb, 0xa, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		})).To(MatchError("received conflicting stateless reset tokens for sequence number 42"))
	})

	It("retires connection IDs", func() {
		Expect(m.Add(&wire.NewConnectionIDFrame{
			SequenceNumber: 10,
			ConnectionID:   protocol.ConnectionID{1, 2, 3, 4},
		})).To(Succeed())
		Expect(m.Add(&wire.NewConnectionIDFrame{
			SequenceNumber: 13,
			ConnectionID:   protocol.ConnectionID{2, 3, 4, 5},
		})).To(Succeed())
		Expect(frameQueue).To(BeEmpty())
		Expect(m.Add(&wire.NewConnectionIDFrame{
			RetirePriorTo:  11,
			SequenceNumber: 17,
			ConnectionID:   protocol.ConnectionID{3, 4, 5, 6},
		})).To(Succeed())
		Expect(frameQueue).To(HaveLen(1))
		Expect(frameQueue[0].(*wire.RetireConnectionIDFrame).SequenceNumber).To(BeEquivalentTo(10))
		c, _ := m.Get()
		Expect(c).To(Equal(protocol.ConnectionID{2, 3, 4, 5}))
		c, _ = m.Get()
		Expect(c).To(Equal(protocol.ConnectionID{3, 4, 5, 6}))
	})
})
