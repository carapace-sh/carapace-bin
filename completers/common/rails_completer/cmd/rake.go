package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rakeCmd = &cobra.Command{
	Use:   "rake",
	Short: "Rake tasks (log:clear, tmp:*, time:zones, etc.)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rakeCmd).Standalone()
	rakeCmd.Flags().BoolP("help", "h", false, "Show help")
}
