package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
)

var subtreeCmd = &cobra.Command{
	Use:   "subtree",
	Short: "Merge subtrees together and split repository into subtrees",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subtreeCmd).Standalone()

	subtreeCmd.PersistentFlags().BoolP("debug", "d", false, "Produce even more unnecessary output messages on stderr")
	subtreeCmd.PersistentFlags().StringP("gpg-sign", "S", "", "GPG-sign commits")
	subtreeCmd.PersistentFlags().Bool("no-gpg-sign", false, "do not GPG-sign commits")
	subtreeCmd.PersistentFlags().StringP("prefix", "P", "", "Specify the path in the repository to the subtree you want to manipulate")
	subtreeCmd.PersistentFlags().BoolP("quiet", "q", false, "Suppress unnecessary output messages on stderr")
	rootCmd.AddCommand(subtreeCmd)

	subtreeCmd.Flag("gpg-sign").NoOptDefVal = " "

	carapace.Gen(subtreeCmd).FlagCompletion(carapace.ActionMap{
		"gpg-sign": os.ActionGpgKeyIds(),
		"prefix":   carapace.ActionDirectories().ChdirF(traverse.GitWorkTree), // TODO verify/fix
	})
}
