package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Remove environment variables from mise.toml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unsetCmd).Standalone()

	unsetCmd.Flags().StringP("file", "f", "", "The TOML file to edit instead of the default target")
	unsetCmd.Flags().BoolP("global", "g", false, "Use the global config file")
	unsetCmd.Flags().BoolP("help", "h", false, "Print help")
	unsetCmd.Flags().Bool("override", false, "Use the global override file, `config.override.toml`")
	unsetCmd.Flags().String("path", "", "The TOML file to edit instead of the default target")
	rootCmd.AddCommand(unsetCmd)

	carapace.Gen(unsetCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})

	carapace.Gen(unsetCmd).PositionalAnyCompletion(
		mise.ActionEnvironmentKeys(),
	)
}
