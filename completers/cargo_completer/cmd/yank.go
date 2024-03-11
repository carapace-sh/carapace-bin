package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var yankCmd = &cobra.Command{
	Use:   "yank",
	Short: "Remove a pushed crate from the index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(yankCmd).Standalone()

	yankCmd.Flags().BoolP("help", "h", false, "Print help")
	yankCmd.Flags().String("index", "", "Registry index to yank from")
	yankCmd.Flags().BoolP("quiet", "q", false, "Do not print cargo log messages")
	yankCmd.Flags().String("registry", "", "Registry to use")
	yankCmd.Flags().String("token", "", "API token to use when authenticating")
	yankCmd.Flags().Bool("undo", false, "Undo a yank, putting a version back into the index")
	yankCmd.Flags().String("version", "", "The version to yank or un-yank")
	rootCmd.AddCommand(yankCmd)
}
