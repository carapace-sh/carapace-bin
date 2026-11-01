package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dockerCliPluginHooksCmd = &cobra.Command{
	Use:    "docker-cli-plugin-hooks",
	Short:  "",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dockerCliPluginHooksCmd).Standalone()

	rootCmd.AddCommand(dockerCliPluginHooksCmd)
}
