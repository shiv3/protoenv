package protoc

import (
	"github.com/shiv3/protoenv/config"
	"github.com/spf13/cobra"
)

type Protoc struct {
	install  Install
	global   Global
	version  Version
	versions Versions
}

func NewProtoc(parentCmd *cobra.Command, cfg config.ProtocConfig) *Protoc {
	cmd := &cobra.Command{
		Use:   "protoc",
		Short: "protoc version manager",
		Long:  `protoc version manager`,
	}
	parentCmd.AddCommand(cmd)
	return &Protoc{
		install:  NewInstall(cmd, cfg.InstallPath),
		global:   NewGlobal(cmd, cfg.InstallPath),
		version:  NewVersion(cmd, cfg.InstallPath),
		versions: NewVersions(cmd, cfg.InstallPath),
	}
}
