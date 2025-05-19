package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "molecule [flags] COMMAND",
	Short: "Testing framework to aid in the development of Ansible roles",
	Long:  "https://ansible.readthedocs.io/projects/molecule/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringArrayP("base-config", "c", nil, "Path to a base config (can be used multiple times)")
	rootCmd.Flags().Bool("debug", false, "Enable debug mode")
	rootCmd.Flags().StringP("env-file", "e", ".env.yml", "Path to environment variables file")
	rootCmd.Flags().Bool("help", false, "Show help message and exit")
	rootCmd.Flags().Bool("no-debug", false, "Disable debug mode (default)")
	rootCmd.Flags().BoolSliceP("verbose", "v", nil, "Increase Ansible verbosity level")

	rootCmd.MarkFlagsMutuallyExclusive("debug", "no-debug")
}
