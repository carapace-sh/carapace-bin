package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:     "env",
	Short:   "Print the environment for the current configuration",
	Aliases: []string{"e"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(envCmd).Standalone()

	envCmd.Flags().BoolP("dotenv", "D", false, "Output in dotenv format")
	envCmd.Flags().BoolP("help", "h", false, "Print help")
	envCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	envCmd.Flags().Bool("json-extended", false, "Output in JSON format with additional information (source, tool)")
	envCmd.Flags().Bool("paths", false, "Print the PATH entries mise manages, in PATH order, one per line")
	envCmd.Flags().Bool("redacted", false, "Only include variables marked for redaction; their values are printed unmasked")
	envCmd.Flags().StringP("shell", "s", "", "Shell type to generate environment variables for")
	envCmd.Flags().Bool("values", false, "Only show values of environment variables")
	rootCmd.AddCommand(envCmd)

	carapace.Gen(envCmd).FlagCompletion(carapace.ActionMap{
		"shell": action.ActionShells(),
	})

	carapace.Gen(envCmd).PositionalAnyCompletion(
		mise.ActionToolVersions(),
	)
}
