package config

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path"
)

var configFileName = "config.yaml"
var CfgFilePath string

type Config struct {
	ApplicationName   string `default:"protoenv"`
	Debug             bool   `default:"false"`
	RootDirectoryPath string `default:""`

	Protoc ProtocConfig
}

type ProtocConfig struct {
	InstallPath string
}

func getDefaultConfigPath() string {
	h, err := homedir.Dir()
	cobra.CheckErr(err)
	return fmt.Sprintf("%s/.protoenv", h)
}

// InitConfig reads in config file and ENV variables if set.
func InitConfig() Config {
	if CfgFilePath != "" {
		// Use config file from the flag.
		viper.SetConfigFile(CfgFilePath)
	} else {
		// Search config in home directory with name ".protoenv" (without extension).
		viper.AddConfigPath(getDefaultConfigPath())
		viper.SetConfigName(configFileName)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		var config Config
		err := viper.Unmarshal(&config)
		cobra.CheckErr(err)
		return config
	}
	return Config{
		ApplicationName:   "",
		Debug:             false,
		RootDirectoryPath: getDefaultConfigPath(),
		Protoc: ProtocConfig{
			InstallPath: path.Join(getDefaultConfigPath(), "protoc"),
		},
	}
}
