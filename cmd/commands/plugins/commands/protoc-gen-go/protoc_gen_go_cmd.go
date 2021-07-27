package protoc_gen_go

import (
	"fmt"

	commands3 "github.com/shiv3/protoenv/cmd/commands/plugins/commands/protoc-gen-go/commands"
	"github.com/spf13/cobra"
)

type Cmd struct {
	ProtocGenGoCommands ProtocGenGoCommands
	CmdOptions          CmdOptions
}

type ProtocGenGoCommands struct {
	init     commands3.Init
	install  commands3.Install
	local    commands3.Local
	global   commands3.Global
	version  commands3.Version
	versions commands3.Versions
}

type CmdOptions struct {
	ShowVersionFormatSimple string
	TargetBinaryFileName    string
}

func NewProtocGenGo(cmdAlias string, parentCmd *cobra.Command) *Cmd {
	showVersionFormatSimple := "%s\n"
	targetBinaryFileName := "protoc"
	installPath := ""

	cmd := &cobra.Command{
		Use:   cmdAlias,
		Short: fmt.Sprintf("%s version manager", targetBinaryFileName),
		Long:  fmt.Sprintf(`%s version manager`, targetBinaryFileName),
	}

	parentCmd.AddCommand(cmd)
	return &Cmd{
		ProtocGenGoCommands: ProtocGenGoCommands{
			init:     commands3.NewInit(cmd, installPath, targetBinaryFileName),
			install:  commands3.NewInstall(cmd, installPath, showVersionFormatSimple, targetBinaryFileName),
			local:    commands3.NewLocal(cmd, installPath, showVersionFormatSimple, targetBinaryFileName),
			global:   commands3.NewGlobal(cmd, installPath, showVersionFormatSimple, targetBinaryFileName),
			version:  commands3.NewVersion(cmd, installPath, showVersionFormatSimple, targetBinaryFileName),
			versions: commands3.NewVersions(cmd, installPath, showVersionFormatSimple, targetBinaryFileName),
		},
	}
}
