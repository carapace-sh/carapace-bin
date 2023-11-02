package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "List files in a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(filesCmd).Standalone()

	filesCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	filesCmd.Flags().StringP("revision", "r", "", "The revision to list files in")
	rootCmd.AddCommand(filesCmd)
}
