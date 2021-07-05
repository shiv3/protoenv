package cmd

import (
	initcmd "github.com/shiv3/protoenv/cmd/commands/init"
	protoc "github.com/shiv3/protoenv/cmd/commands/plugins"
	protoc2 "github.com/shiv3/protoenv/cmd/commands/protoc"
	"github.com/shiv3/protoenv/config"
)

func init() {
	cfg := config.InitConfig()
	initcmd.NewInit(rootCmd, cfg)
	protoc2.NewProtoc(rootCmd, cfg.Protoc)
	protoc.NewPlugins(rootCmd)
}
