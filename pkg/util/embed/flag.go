package embed

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func subcommandsAsFlags(cmd *cobra.Command, shorthandOnly bool, subcommands ...*cobra.Command) {
	carapace.Gen(cmd).PreRun(func(cmd *cobra.Command, args []string) {
		var nonposix bool
		flags := pflag.NewFlagSet("subcommands", pflag.ContinueOnError)
		for _, s := range subcommands {
			if len(s.Aliases) > 0 {
				nonposix = nonposix || len(s.Aliases[0]) > 1
			}

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

		if arg, ok := findSubcommandFlagArg(args, cmd.Flags(), flags, nonposix); ok {
			flags.Parse([]string{arg})
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
		case subcommand != nil && (!nonposix || len(args) > 1):
			cmd.DisableFlagParsing = true
			carapace.Gen(cmd).PositionalAnyCompletion(
				carapace.ActionExecute(subcommand),
			)
		case len(args) < 2:
			cmd.Flags().AddFlagSet(flags)
		}
	})
}

func findSubcommandFlagArg(args []string, rootFlags, subcommandFlags *pflag.FlagSet, nonposix bool) (string, bool) {
	for index := 0; index < len(args); index++ {
		arg := args[index]

		switch {
		case strings.HasPrefix(arg, "--"):
			name := strings.TrimPrefix(arg, "--")
			if before, _, ok := strings.Cut(name, "="); ok {
				name = before
			}

			if subcommandFlags.Lookup(name) != nil {
				return "--" + name, true
			}
			if flag := rootFlags.Lookup(name); flag != nil && !strings.Contains(arg, "=") && flag.Value.Type() != "bool" {
				index++
			}

		case strings.HasPrefix(arg, "-") && len(arg) > 1:
			shorthand := arg[1:]
			if !nonposix {
				shorthand = arg[1:2]
			}

			if subcommandFlags.ShorthandLookup(shorthand) != nil {
				return "-" + shorthand, true
			}
			if flag := rootFlags.ShorthandLookup(arg[1:2]); flag != nil && len(arg) == 2 && flag.Value.Type() != "bool" {
				index++
			}

		default:
			return "", false
		}
	}

	return "", false
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
