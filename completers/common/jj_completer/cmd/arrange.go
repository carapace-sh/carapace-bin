package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jjlex/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var arrangeCmd = &cobra.Command{
	Use:   "arrange",
	Short: "Interactively arrange the commit graph",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arrangeCmd).Standalone()

	arrangeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	arrangeCmd.Flags().StringSliceS("r", "r", nil, "")
	arrangeCmd.Flag("r").Hidden = true
	rootCmd.AddCommand(arrangeCmd)

	carapace.Gen(arrangeCmd).PositionalAnyCompletion(
		jj.ActionRevsets(jj.RevOpts{}.Default()),
	)
}
