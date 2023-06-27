package embed

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// SubcommandsAsFlags embeds subcommands as flags (e.g. `pacman -Syu` where `S` is actually a subcommand).
// Supports a single alias that is used as shorthand.
func SubcommandsAsFlags(cmd *cobra.Command, subcommands ...*cobra.Command) {
	carapace.Gen(cmd).PreRun(func(cmd *cobra.Command, args []string) {
		flags := pflag.NewFlagSet("subcommands", pflag.ContinueOnError)
		for _, s := range subcommands {
			if len(s.Aliases) > 0 {
				flags.BoolP(s.Name(), s.Aliases[0], false, s.Short) // suppport one alias as shorthand
			} else {
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
