package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var simplifyParentsCmd = &cobra.Command{
	Use:   "simplify-parents <--source <SOURCE>|--revisions <REVISIONS>>",
	Short: "Simplify parent edges for the specified revision(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(simplifyParentsCmd).Standalone()

	simplifyParentsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	simplifyParentsCmd.Flags().StringSliceP("revisions", "r", nil, "Simplify specified revision(s) (can be repeated)")
	simplifyParentsCmd.Flags().StringSliceP("source", "s", nil, "Simplify specified revision(s) together with their trees of descendants (can be repeated)")
	rootCmd.AddCommand(simplifyParentsCmd)

	carapace.Gen(simplifyParentsCmd).FlagCompletion(carapace.ActionMap{
		"revisions": jj.ActionRevs(jj.RevOption{}.Default()).FilterArgs(),
		"source":    jj.ActionRevs(jj.RevOption{}.Default()).FilterArgs(),
	})
}
