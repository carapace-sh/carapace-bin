package embed

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func subcommandsAsFlags(cmd *cobra.Command, shorthandOnly bool, subcommands ...*cobra.Command) {
	oldRunE := cmd.RunE
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			for _, subcommand := range subcommands {
				switch args[0] {
				case "--" + subcommand.Name(), "-" + subcommand.Short: // TODO name only when shorthandonly
					subcommand.SetArgs(args)
					return subcommand.Execute()
				}
			}
		}

		switch {
		case oldRunE != nil:
			return oldRunE(cmd, args)
		case cmd.Run != nil:
			cmd.Run(cmd, args)
			return nil
		default:
			return nil
		}
	}

	carapace.Gen(cmd).PreRun(func(cmd *cobra.Command, args []string) {
		flags := pflag.NewFlagSet("subcommands", pflag.ContinueOnError)
		for _, s := range subcommands {
			switch {
			case shorthandOnly:
				if len(s.Aliases) > 0 {
					flags.BoolS(s.Name(), s.Aliases[0], false, s.Short)
				}
				// TODO handle misconfiguration

			case len(s.Aliases) > 0:
				flags.BoolP(s.Name(), s.Aliases[0], false, s.Short) // suppport one alias as shorthand

			default:
				flags.Bool(s.Name(), false, s.Short)
			}
		}

		switch {
		case strings.HasPrefix(args[0], "--"):
			flags.Parse(args[:1])
		case strings.HasPrefix(args[0], "-") && len(args[0]) > 1:
			flags.Parse([]string{args[0][:2]})
		}

		var subcommand *cobra.Command
		flags.Visit(func(f *pflag.Flag) {
			for _, s := range subcommands {
				if f.Name == s.Name() {
					subcommand = s
					subcommand.Flags().AddFlag(f)
				}
			}
		})

		switch {
		case subcommand != nil:
			cmd.DisableFlagParsing = true
			carapace.Gen(cmd).PositionalAnyCompletion(
				carapace.ActionExecute(subcommand),
			)
		case len(args) < 2:
			cmd.Flags().AddFlagSet(flags)
		}
	})
}

// SubcommandsAsFlags embeds subcommands as flags (e.g. `pacman -Syu` where `S` is actually a subcommand).
// Supports a single alias that is used as shorthand.
func SubcommandsAsFlags(cmd *cobra.Command, subcommands ...*cobra.Command) {
	subcommandsAsFlags(cmd, false, subcommands...)
}

// SubcommandsAsFlagsS is like SubcommandsAsFlags but only adds the shorthand.
func SubcommandsAsFlagsS(cmd *cobra.Command, subcommands ...*cobra.Command) {
	subcommandsAsFlags(cmd, true, subcommands...)
}
