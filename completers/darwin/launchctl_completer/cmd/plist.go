package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/launchctl"
	"github.com/spf13/cobra"
)

var plistCmd = &cobra.Command{
	Use:   "plist",
	Short: "Print the property list for a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plistCmd).Standalone()
	rootCmd.AddCommand(plistCmd)
	carapace.Gen(plistCmd).PositionalAnyCompletion(launchctl.ActionServices())
}
