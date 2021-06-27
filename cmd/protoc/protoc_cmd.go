package protoc

import (
	"github.com/shiv3/protoenv/config"
	"github.com/spf13/cobra"
	"fmt"
)

type Protoc struct {
	init     Init
	install  Install
	local    Local
	global   Global
	version  Version
	versions Versions
}

func NewProtoc(parentCmd *cobra.Command, cfg config.ProtocConfig) *Protoc {
	cmd := &cobra.Command{
		Use:   "protoc",
		Short: fmt.Sprintf("%s version manager",TargetBinaryFileName),
		Long:  fmt.Sprintf(`%s version manager`,TargetBinaryFileName),
	}
	parentCmd.AddCommand(cmd)
	return &Protoc{
		init:     NewInit(cmd, cfg.InstallPath),
		install:  NewInstall(cmd, cfg.InstallPath),
		local:    NewLocal(cmd, cfg.InstallPath),
		global:   NewGlobal(cmd, cfg.InstallPath),
		version:  NewVersion(cmd, cfg.InstallPath),
		versions: NewVersions(cmd, cfg.InstallPath),
	}
}
