package cmd

import (
	"github.com/shiv3/protoenv/cmd/initcmd"
	"github.com/shiv3/protoenv/cmd/protoc"
	"github.com/shiv3/protoenv/config"
)

func init() {
	cfg := config.InitConfig()
	initcmd.NewInit(rootCmd, cfg)
	protoc.NewProtoc(rootCmd, cfg.Protoc)
}
