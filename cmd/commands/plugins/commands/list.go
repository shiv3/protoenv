package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

type List struct {
	plugins []string
}

func NewList(parentCmd *cobra.Command, plugins []string) List {
	list := List{
		plugins: plugins,
	}
	cmd := &cobra.Command{
		Use:   "list",
		Short: fmt.Sprintf("list plugins"),
		Long:  fmt.Sprintf(`list plugins`),
		RunE:  list.RunE,
	}
	parentCmd.AddCommand(cmd)
	return list
}

func (i *List) RunE(cmd *cobra.Command, args []string) error {
	for _, plugin := range i.plugins {
		fmt.Printf("%s\n", plugin)
	}
	return nil
}
