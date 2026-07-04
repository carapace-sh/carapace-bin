package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/launchctl"
	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print information about a domain or service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(printCmd).Standalone()
	rootCmd.AddCommand(printCmd)
	carapace.Gen(printCmd).PositionalAnyCompletion(launchctl.ActionServices())
}
