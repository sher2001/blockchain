package network

import (
	"crypto"
	"fmt"
	"time"

	"github.com/sher2001/blockchain/core"
	"github.com/sirupsen/logrus"
)

type ServerOpts struct {
	Transports []Transport
	BlockTime  time.Duration
	PrivateKey crypto.PrivateKey
}

type Server struct {
	ServerOpts

	BlockTime   time.Duration
	memPool     *TxPool
	isValidator bool
	rpcCh       chan RPC
	quitCh      chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts:  opts,
		BlockTime:   opts.BlockTime,
		memPool:     NewTxPool(),
		isValidator: opts.PrivateKey != nil,
		rpcCh:       make(chan RPC),
		quitCh:      make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.InitTransports()
	ticker := time.NewTicker(s.BlockTime)

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("%+v", rpc)
		case <-s.quitCh:
			break free
		case <-ticker.C:
			if s.isValidator {
				s.CreateNewBlock()
			}
		}
	}

	fmt.Println("Server shutdown")
}

func (s *Server) CreateNewBlock() error {
	fmt.Println("creating a new block")
	return nil
}

func (s *Server) HandleTransaction(tx *core.Transaction) error {
	if err := tx.Verify(); err != nil {
		return err
	}

	hash := tx.Hash(core.TransactionHasher{})
	if s.memPool.Has(hash) {
		logrus.WithFields(logrus.Fields{
			"Hash": hash,
		}).Info("transaction already exists in mempool")
	}

	logrus.WithFields(logrus.Fields{
		"Hash": hash,
	}).Info("adding a new transaction to the mempool")

	return s.memPool.Add(tx)
}

func (s *Server) InitTransports() {
	for _, tr := range s.Transports {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcCh <- rpc
			}
		}(tr)
	}
}
