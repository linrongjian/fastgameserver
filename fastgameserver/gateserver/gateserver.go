package gateserver

import (
	"fastgameserver/fastgameserver/gameserver"
	"fastgameserver/fastgameserver/logger"
	"os"
	"os/signal"
)

type GateServer interface {
	gameserver.Server

	Init(...Option) error

	Options() Options
}

type Option func(*Options)

type gateServer struct {
	opts Options
}

func (g *gateServer) Run() error {

	// Ctx = s.Options().Context
	// CreateHTTPServer()
	// ClearOnline()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	logger.Info("Leaf closing down (signal: %v)", sig)

	return nil
}

func (g *gateServer) Stop() error {
	return nil
}

func (g *gateServer) Options() Options {
	return g.opts
}

func (g *gateServer) Init(opts ...Option) error {
	for _, o := range opts {
		o(&g.opts)
	}

	// cmd.AddFlags(defaultFlags)
	// cmd.AddFlags(redis.Flags)
	// cmd.AddFlags(mq.Flags)
	// cmd.AddFlags(mysql.Flags)
	// if err := s.opts.Cmd.Init(); err != nil {
	// logger.Fatal(err)
	// }
	//mq.Startup()
	// mysql.Startup()

	// err := redis.Connect()
	// if err != nil {
	// 	logger.Fatal(err)
	// }
	// s.opts.Rds = redis.S()
	return nil
}

func NewGateServer(opts ...Option) GateServer {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}
	return &gateServer{
		opts: options,
	}
}
