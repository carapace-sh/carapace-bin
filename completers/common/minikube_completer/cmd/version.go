package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Print the version of minikube",
	GroupID: "troubleshooting",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()
	versionCmd.Flags().StringP("output", "o", "", "One of 'yaml' or 'json'.")
	versionCmd.Flags().Bool("short", false, "Print just the version number.")
	rootCmd.AddCommand(versionCmd)

	carapace.Gen(versionCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("yaml", "json"),
	})
}
