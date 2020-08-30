package cmd

import (
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "printenv",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("null", "0", false, "end each output line with NUL, not newline")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).PositionalAnyCompletion(ActionEnvironmentVariables())
}

func ActionEnvironmentVariables() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		env := os.Environ()
		vars := make([]string, len(env))
		for index, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			vars[index] = pair[0]
		}
		return carapace.ActionValues(vars...)
	})
}
