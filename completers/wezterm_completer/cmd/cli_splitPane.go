package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wezterm"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
)

var cli_splitPaneCmd = &cobra.Command{
	Use:   "split-pane [OPTIONS] [PROG]...",
	Short: "Move a pane into a new tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cli_splitPaneCmd).Standalone()
	cli_splitPaneCmd.Flags().SetInterspersed(false)

	cli_splitPaneCmd.Flags().Bool("bottom", false, "Split vertically, with the new pane on the bottom")
	cli_splitPaneCmd.Flags().String("cells", "", "The number of cells that the new split should have")
	cli_splitPaneCmd.Flags().String("cwd", "", "Specify the current working directory for the initially spawned program")
	cli_splitPaneCmd.Flags().BoolP("help", "h", false, "Print help")
	cli_splitPaneCmd.Flags().Bool("horizontal", false, "Equivalent to `--right`")
	cli_splitPaneCmd.Flags().Bool("left", false, "Split horizontally, with the new pane on the left")
	cli_splitPaneCmd.Flags().String("move-pane-id", "", "Move the specified pane into the newly created split")
	cli_splitPaneCmd.Flags().String("pane-id", "", "Specify the pane that should be split")
	cli_splitPaneCmd.Flags().String("percent", "", "Specify the number of cells that the new split should have")
	cli_splitPaneCmd.Flags().Bool("right", false, "Split horizontally, with the new pane on the right")
	cli_splitPaneCmd.Flags().Bool("top", false, "Split vertically, with the new pane on the top")
	cli_splitPaneCmd.Flags().Bool("top-level", false, "Rather than splitting the active pane, split the entire window")
	cliCmd.AddCommand(cli_splitPaneCmd)

	carapace.Gen(cli_splitPaneCmd).FlagCompletion(carapace.ActionMap{
		"cwd":          carapace.ActionDirectories(),
		"move-pane-id": wezterm.ActionPanes(),
		"pane-id":      wezterm.ActionPanes(),
	})

	carapace.Gen(cli_splitPaneCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin().ChdirF(traverse.Flag(cli_splitPaneCmd.Flag("cwd"))),
	)
}
