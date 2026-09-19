package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var debugdeltachainCmd = &cobra.Command{
	Use:    "debugdeltachain",
	Short:  "-c|-m|FILE",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugdeltachainCmd).Standalone()

	debugdeltachainCmd.Flags().Bool("all-info", false, "compute all information unless specified otherwise")
	debugdeltachainCmd.Flags().BoolP("changelog", "c", false, "open changelog")
	debugdeltachainCmd.Flags().String("dir", "", "open directory manifest")
	debugdeltachainCmd.Flags().Bool("dist-info", false, "compute information related to base distance")
	debugdeltachainCmd.Flags().BoolP("manifest", "m", false, "open manifest")
	debugdeltachainCmd.Flags().StringArrayP("rev", "r", nil, "restrict processing to these revlog revisions")
	debugdeltachainCmd.Flags().Bool("size-info", false, "compute information related to deltas size")
	debugdeltachainCmd.Flags().Bool("sparse-info", false, "compute information related to sparse read")
	debugdeltachainCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(debugdeltachainCmd)

	carapace.Gen(debugdeltachainCmd).FlagCompletion(carapace.ActionMap{
		"rev": action.ActionRevisions(),
	})
}
