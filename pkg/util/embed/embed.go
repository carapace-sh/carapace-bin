package embed

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/bridge"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type stringValue string

func (s *stringValue) Set(val string) error {
	return nil
}
func (s *stringValue) Type() string {
	return "string"
}

func (s *stringValue) String() string { return string(*s) }

// EmbedCarapaceBin **EXPERIMENTAL** adds a two-pass positional completion for commands like `sudo [CMD] [ARGS]...`
func EmbedCarapaceBin(cmd *cobra.Command) {
	carapace.Gen(cmd).PositionalCompletion(
		carapace.Batch(
			os.ActionPathExecutables(),
			carapace.ActionFiles(),
		).ToA(),
	)

	carapace.Gen(cmd).PreRun(func(cmd *cobra.Command, args []string) {
		flags := pflag.NewFlagSet("copy", pflag.ContinueOnError)
		flags.ParseErrorsWhitelist = pflag.ParseErrorsWhitelist{
			UnknownFlags: true,
		}
		var value stringValue
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			flags.AddFlag(&pflag.Flag{
				Name:                f.Name,
				Shorthand:           f.Shorthand,
				Style:               f.Style,
				Usage:               f.Usage,
				Value:               &value,
				DefValue:            f.DefValue,
				Changed:             f.Changed,
				NoOptDefVal:         f.NoOptDefVal,
				Deprecated:          f.Deprecated,
				Hidden:              f.Hidden,
				ShorthandDeprecated: f.ShorthandDeprecated,
				Annotations:         f.Annotations,
			})
		})

		if err := flags.Parse(args); err == nil && len(flags.Args()) > 1 {
			subCmd := &cobra.Command{
				Use:                flags.Args()[0],
				Run:                func(cmd *cobra.Command, args []string) {},
				DisableFlagParsing: true,
			}
			cmd.AddCommand(subCmd)

			carapace.Gen(subCmd).PositionalAnyCompletion(
				bridge.ActionCarapaceBin(flags.Args()[0]),
			)
		}
	})
}
