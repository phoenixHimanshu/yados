package store

import (
	"github.com/davinash/yados/commands/utils"
	"github.com/spf13/cobra"
)

func AddCommands(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "store",
		Short: "Store operations",
		Args:  utils.ExactArgs(0),
		RunE:  utils.ShowHelp(),
	}
	AddCreateStoreCommand(cmd)
	AddDeleteStoreCommand(cmd)
	AddGetStoreCommand(cmd)
	rootCmd.AddCommand(cmd)
}
