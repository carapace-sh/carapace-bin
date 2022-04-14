package cmd

import (
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "terraform",
	Short: "infrastructure as code software tool",
	Long:  "https://www.terraform.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	carapace.Override(carapace.Opts{
		LongShorthand: true,
	})
	return rootCmd.Execute()
}

func ActionExecute() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO don't change os.Args
		backup := os.Args
		carapace.Override(carapace.Opts{
			LongShorthand: true,
		})
		for index, arg := range c.Args {
			if strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--") {
				c.Args[index] = "-" + arg
			}
		}
		if strings.HasPrefix(c.CallbackValue, "-") && !strings.HasPrefix(c.CallbackValue, "--") {
			c.CallbackValue = "-" + c.CallbackValue
		}
		os.Args = backup
		return carapace.ActionExecute(rootCmd).Invoke(c).ToA()
	})
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("chdir", "", "Switch to a different working directory before executing the given subcommand.")
	rootCmd.Flags().Bool("help", false, "Show this help output, or the help for a specified subcommand.")
	rootCmd.Flags().Bool("version", false, "An alias for the \"version\" subcommand.")

	rootCmd.Flag("chdir").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"chdir": carapace.ActionDirectories(),
	})
}
