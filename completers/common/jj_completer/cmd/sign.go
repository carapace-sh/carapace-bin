package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var signCmd = &cobra.Command{
	Use:   "sign [OPTIONS]",
	Short: "Cryptographically sign a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signCmd).Standalone()

	signCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	signCmd.Flags().String("key", "", "The key used for signing")
	signCmd.Flags().StringSliceP("revisions", "r", []string{"@"}, "What revision(s) to sign")
	rootCmd.AddCommand(signCmd)

	carapace.Gen(signCmd).FlagCompletion(carapace.ActionMap{
		"key":       jj.ActionSigningKeys(),
		"revisions": jj.ActionRevSets(jj.RevOption{}.Default()),
	})
}
