package main

import (
	"log/slog"

	"github.com/cooldogedev/spectrum-df"
	"github.com/cooldogedev/spectrum-df/transport"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player/chat"
)

func main() {
	log := slog.Default()
	chat.Global.Subscribe(chat.StdoutSubscriber{})
	conf, err := server.DefaultConfig().Config(log)
	if err != nil {
		panic(err)
	}

	conf.Listeners = []func(conf server.Config) (server.Listener, error){func(conf server.Config) (server.Listener, error) {
		return spectrum.NewListener(":19133", transport.NewSpectral())
	}}
	srv := conf.New()
	srv.CloseOnProgramEnd()
	srv.Listen()
	for range srv.Accept() {
	}
}
