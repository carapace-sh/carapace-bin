package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dbCmd = &cobra.Command{
	Use:   "db [--read] [--write] [--] [<connection>]",
	Short: "Start a new database CLI session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dbCmd).Standalone()

	dbCmd.Flags().Bool("read", false, "Connect to the read connection")
	dbCmd.Flags().Bool("write", false, "Connect to the write connection")
	rootCmd.AddCommand(dbCmd)
}
