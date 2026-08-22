package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var action_setPaneFrameStyleCmd = &cobra.Command{
	Use:   "set-pane-frame-style",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_setPaneFrameStyleCmd).Standalone()

	action_setPaneFrameStyleCmd.Flags().BoolP("help", "h", false, "Print help")
	actionCmd.AddCommand(action_setPaneFrameStyleCmd)

	carapace.Gen(action_setPaneFrameStyleCmd).PositionalAnyCompletion(
		carapace.ActionValues("full", "titles", "none").StyleF(style.ForKeyword),
	)
}
