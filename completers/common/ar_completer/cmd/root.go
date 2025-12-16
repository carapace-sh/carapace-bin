package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/ar"
	"github.com/carapace-sh/carapace-bin/pkg/util/embed"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "ar",
	Short: "create, modify, and extract from archives",
	Long:  "https://linux.die.net/man/1/ar",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	// TODO some commands are still missing flags (like -N)
	embed.SubcommandsAsFlagsS(rootCmd,
		deleteCmd,
		moveCmd,
		printCmd,
		quickCmd,
		replaceCmd,
		actCmd,
		displayCmd,
		extractCmd,
	)
}

func addGenericOptions(cmd *cobra.Command) {
	cmd.Flags().BoolS("S", "S", false, "do not build a symbol table")
	cmd.Flags().BoolS("V", "V", false, "display the version number")
	cmd.Flags().BoolS("c", "c", false, "do not warn if the library had to be created")
	cmd.Flags().StringS("l", "l", "", "specify the dependencies of this library")
	cmd.Flags().String("output", "", "specify the output directory for extraction operations")
	cmd.Flags().String("plugin", "", "load the specified plugin")
	cmd.Flags().String("record-libdeps", "", "specify the dependencies of this library")
	cmd.Flags().BoolS("s", "s", false, "create an archive index (cf. ranlib)")
	cmd.Flags().String("target", "", "specify the target object format as BFDNAME")
	cmd.Flags().Bool("thin", false, "make a thin archive")
	cmd.Flags().BoolS("v", "v", false, "be verbose")

	cmd.Flag("target").NoOptDefVal = " "
	cmd.Flag("output").NoOptDefVal = " "
	cmd.Flag("record-libdeps").NoOptDefVal = " "

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionDirectories(),
		"target": ar.ActionTargets(),
	})

	carapace.Gen(cmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if flag == nil && strings.HasPrefix(c.Value, "@") {
				c.Value = strings.TrimPrefix(c.Value, "@")
				return carapace.ActionFiles().Invoke(c).Prefix("@").ToA()
			}

			// TODO somehow skip `@` arg in PreRun (needs support in carapace for it)
			// filtered := []string{}
			// for _, arg := range c.Args {
			// 	if !strings.HasPrefix(arg, "@") {
			// 		filtered = append(filtered, arg)
			// 	}
			// }
			// c.Args = filtered
			// return action.Invoke(c).ToA()
			return action
		})
	})
}
