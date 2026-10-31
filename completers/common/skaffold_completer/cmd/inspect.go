package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:    "inspect",
	Short:  "Helper commands for Cloud Code IDEs to interact with and modify skaffold configuration files.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspectCmd).Standalone()

	inspectCmd.PersistentFlags().StringP("filename", "f", "skaffold.yaml", "Path to the local Skaffold config file. Defaults to `skaffold.yaml`")
	inspectCmd.PersistentFlags().StringP("format", "o", "json", "Output format. One of: json(default)")
	inspectCmd.PersistentFlags().String("remote-cache-dir", "", "Specify the location of the remote cache (defaults to $HOME/.skaffold/remote-cache)")
	applyCommonFlagCompletion(inspectCmd)
	rootCmd.AddCommand(inspectCmd)
}
