package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	addr        NetAddr
	consumeChan chan RPC
	lock        sync.RWMutex
	peers       map[NetAddr]*LocalTransport
}

func NewLocalTransport(addr NetAddr) Transport {
	return &LocalTransport{
		addr:        addr,
		consumeChan: make(chan RPC, 1024),
		peers:       make(map[NetAddr]*LocalTransport),
	}
}

func (t *LocalTransport) Connect(tr Transport) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.peers[tr.Addr()] = tr.(*LocalTransport)

	return nil
}

func (t *LocalTransport) Consume() <-chan RPC {
	return t.consumeChan
}

func (t *LocalTransport) Send_message(to NetAddr, Payload []byte) error {
	t.lock.RLock()
	defer t.lock.RUnlock()

	peer, ok := t.peers[to]
	if !ok {
		return fmt.Errorf("%s: Unable to send message to %s", t.Addr(), to)
	}

	peer.consumeChan <- RPC{
		From:    t.Addr(),
		Payload: Payload,
	}

	return nil
}

func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}
