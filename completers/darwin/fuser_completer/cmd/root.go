package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fuser",
	Short: "list process IDs of all processes that have one or more files open",
	Long:  "https://man.freebsd.org/cgi/man.cgi?fuser",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("c", "c", false, "File is treated as a mount point")
	rootCmd.Flags().BoolS("f", "f", false, "Report only for the named files")
	rootCmd.Flags().BoolS("u", "u", false, "User login name")
}
