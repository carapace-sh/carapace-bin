package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "add dependencies to current module and install them",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCmd).Standalone()

	getCmd.Flags().BoolS("t", "t", false, "consider modules needed to build tests")
	getCmd.Flags().BoolS("u", "u", false, "update modules providing dependencies")
	getCmd.Flags().BoolS("d", "d", false, "only download the source code needed to build")
	getCmd.Flags().BoolS("v", "v", false, "verbose output")
	getCmd.Flags().Bool("insecure", false, "permit using insecure schemes such as HTTP")
	rootCmd.AddCommand(getCmd)
}
