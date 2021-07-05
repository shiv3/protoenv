package protoc

import (
	"fmt"
	"github.com/shiv3/protoenv/cmd/commands/plugins/commands"
	protoc_gen_go "github.com/shiv3/protoenv/cmd/commands/plugins/commands/protoc-gen-go"
	"github.com/spf13/cobra"
)

type Cmd struct {
	PluginsCommands PluginsCommands
	CmdOptions      CmdOptions
}

type PluginsCommands struct {
	list commands.List
}

type CmdOptions struct {
	ShowVersionFormatSimple string
	TargetBinaryFileName    string
}

func NewPlugins(parentCmd *cobra.Command) *Cmd {
	cmd := &cobra.Command{
		Use:   "plugins",
		Short: fmt.Sprintf("plugins version manager"),
		Long:  fmt.Sprintf(`plugins version manager`),
	}

	plugins := map[string]interface{}{
		"protoc-gen-go": protoc_gen_go.NewProtocGenGo("protoc-gen-go", cmd),
	}

	var pluginsAliases []string
	for k, _ := range plugins {
		pluginsAliases = append(pluginsAliases, k)
	}

	parentCmd.AddCommand(cmd)
	return &Cmd{
		PluginsCommands: PluginsCommands{
			list: commands.NewList(cmd, pluginsAliases),
		},
	}
}
