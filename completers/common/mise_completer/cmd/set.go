package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:     "set",
	Short:   "Set, show, or list environment variables in mise.toml",
	Aliases: []string{"ev", "env-vars"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setCmd).Standalone()

	setCmd.Flags().Bool("age-encrypt", false, "[experimental] Encrypt the value with age before storing")
	setCmd.Flags().String("age-key-file", "", "[experimental] Age identity file whose public keys to encrypt to")
	setCmd.Flags().StringSlice("age-recipient", nil, "[experimental] Age recipient (x25519 public key) for encryption")
	setCmd.Flags().StringSlice("age-ssh-recipient", nil, "[experimental] SSH recipient (public key or path) for age encryption")
	setCmd.Flags().Bool("complete", false, "Render completions")
	setCmd.Flags().StringP("env", "E", "", "Write to mise.<ENV>.toml in the current directory, or to .mise.<ENV>.toml if that exists")
	setCmd.Flags().String("file", "", "The TOML file to update")
	setCmd.Flags().BoolP("global", "g", false, "Set the environment variable in the global config file")
	setCmd.Flags().BoolP("help", "h", false, "Print help")
	setCmd.Flags().Bool("no-redact", false, "Show raw values instead of redacting secrets")
	setCmd.Flags().Bool("override", false, "Set the environment variable in the global override file, `config.override.toml`")
	setCmd.Flags().String("path", "", "The TOML file to update")
	setCmd.Flags().Bool("prompt", false, "Prompt for environment variable values (input is masked only with --age-encrypt)")
	setCmd.Flags().StringSlice("remove", nil, "Remove the environment variable from config file")
	setCmd.Flags().StringSlice("rm", nil, "Remove the environment variable from config file")
	setCmd.Flags().Bool("stdin", false, "Read the value from stdin, for multiline input")
	setCmd.Flags().StringSlice("unset", nil, "Remove the environment variable from config file")
	setCmd.Flag("complete").Hidden = true
	setCmd.Flag("remove").Hidden = true
	setCmd.Flag("rm").Hidden = true
	setCmd.Flag("unset").Hidden = true
	rootCmd.AddCommand(setCmd)

	carapace.Gen(setCmd).FlagCompletion(carapace.ActionMap{
		"age-key-file": carapace.ActionFiles(),
		"file":         carapace.ActionFiles(),
		"remove":       mise.ActionEnvironmentKeys(),
	})

	carapace.Gen(setCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return mise.ActionEnvironmentKeys().Suffix("=")
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
