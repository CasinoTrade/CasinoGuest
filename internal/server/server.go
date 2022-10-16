package server

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/CasinoTrade/CasinoGuest/internal/model/log"
)

const pingsBufferSize = 5

type Casino struct {
	log log.Logger

	once   sync.Once
	ctx    context.Context
	cancel context.CancelFunc

	r     *rand.Rand
	pings chan int
}

func New(logger log.Logger) *Casino {
	ctx, cancel := context.WithCancel(context.Background())
	return &Casino{
		r:      rand.New(rand.NewSource(time.Now().Unix())),
		log:    logger,
		pings:  make(chan int, pingsBufferSize),
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *Casino) Start() {
	s.log.Debug("Attempt to start base server")
	s.once.Do(func() {
		s.log.Info("Starting base server")
		go s.run()
	})
}

func (s *Casino) Stop() {
	if s.cancel != nil {
		s.cancel()
	}
}

func (s *Casino) run() {
	defer close(s.pings)
	for {
		select {
		case s.pings <- s.r.Int():
		case <-s.ctx.Done():
			s.log.Info("Stopping base server poller")
			return
		}
	}
}

func (s *Casino) Ping(_ context.Context) int {
	return <-s.pings
}
