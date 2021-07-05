package protoc

import (
	"fmt"
	commands2 "github.com/shiv3/protoenv/cmd/commands/protoc/commands"
	"github.com/shiv3/protoenv/config"
	"github.com/spf13/cobra"
)

type Cmd struct {
	ProtocCommands ProtocCommands
	CmdOptions     CmdOptions
}

type ProtocCommands struct {
	init     commands2.Init
	install  commands2.Install
	local    commands2.Local
	global   commands2.Global
	version  commands2.Version
	versions commands2.Versions
}

type CmdOptions struct {
	ShowVersionFormatSimple string
	TargetBinaryFileName    string
}

func NewProtoc(parentCmd *cobra.Command, cfg config.ProtocConfig) *Cmd {
	showVersionFormatSimple := "%s\n"
	targetBinaryFileName := "protoc"

	cmd := &cobra.Command{
		Use:   "protoc",
		Short: fmt.Sprintf("%s version manager", targetBinaryFileName),
		Long:  fmt.Sprintf(`%s version manager`, targetBinaryFileName),
	}
	parentCmd.AddCommand(cmd)

	return &Cmd{
		ProtocCommands: ProtocCommands{
			init:     commands2.NewInit(cmd, cfg.InstallPath, targetBinaryFileName),
			install:  commands2.NewInstall(cmd, cfg.InstallPath, showVersionFormatSimple, targetBinaryFileName),
			local:    commands2.NewLocal(cmd, cfg.InstallPath, showVersionFormatSimple, targetBinaryFileName),
			global:   commands2.NewGlobal(cmd, cfg.InstallPath, showVersionFormatSimple, targetBinaryFileName),
			version:  commands2.NewVersion(cmd, cfg.InstallPath, showVersionFormatSimple, targetBinaryFileName),
			versions: commands2.NewVersions(cmd, cfg.InstallPath, showVersionFormatSimple, targetBinaryFileName),
		},
	}
}
