package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "conda",
	Short: "conda is a tool for managing and deploying applications, environments and packages",
	Long:  "https://docs.conda.io/en/latest/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Show this help message and exit.")
	rootCmd.Flags().BoolP("version", "V", false, "Show the conda version number and exit.")

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		addOtherCommands()
	})
}

func addOtherCommands() {
	// hardcoded for now (should check PATH for conda-{COMMAND})
	others := map[string]string{
		"content-trust": "Signing and verification tools for Conda",
		"env":           "Manage conda environments",
	}
	for name, description := range others {
		otherCmd := &cobra.Command{
			Use:                name,
			Short:              description,
			Run:                func(cmd *cobra.Command, args []string) {},
			DisableFlagParsing: true,
		}

		carapace.Gen(otherCmd).PositionalAnyCompletion(
			bridge.ActionCarapaceBin("conda-" + name),
		)

		rootCmd.AddCommand(otherCmd)
	}
}
