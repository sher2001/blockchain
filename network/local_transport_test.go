package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_connect(t *testing.T) {
	transportA := NewLocalTransport("A").(*LocalTransport)
	transportB := NewLocalTransport("B").(*LocalTransport)

	transportA.Connect(transportB)
	assert.Equal(t, transportB, transportA.peers[transportB.Addr()])
	transportB.Connect(transportA)
	assert.Equal(t, transportA, transportB.peers[transportA.Addr()])
}

func Test_send_message(t *testing.T) {
	transportA := NewLocalTransport("A")
	transportB := NewLocalTransport("B")

	assert.Error(t, transportA.Send_message(transportB.Addr(), []byte("Something")), "A: Unable to send message to B")

	transportA.Connect(transportB)
	msg := []byte("Hi B")
	transportA.Send_message(transportB.Addr(), msg)

	consume_chan_B := <-transportB.Consume()

	assert.Equal(t, msg, consume_chan_B.Payload)
	assert.Equal(t, transportA.Addr(), consume_chan_B.From)
}
