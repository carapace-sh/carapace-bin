package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_dotfilesCmd = &cobra.Command{
	Use:    "dotfiles",
	Short:  "Deploy, track, and sync dotfiles",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_dotfilesCmd).Standalone()

	bootstrap_dotfilesCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_dotfilesCmd.Flags().Bool("local", false, "Work on this machine's local-only history (`mode = \"track-local\"`), which is never shared")
	bootstrapCmd.AddCommand(bootstrap_dotfilesCmd)
}
