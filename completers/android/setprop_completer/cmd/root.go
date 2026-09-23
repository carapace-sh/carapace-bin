package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "setprop",
	Short: "Set an Android system property",
	Long:  "https://cs.android.com/android/platform/superproject/main/+/main:system/core/toolbox/setprop.cpp",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		actionProperties(),
		carapace.ActionValues(), // VALUE
	)
}

func actionProperties() carapace.Action {
	return carapace.ActionExecCommand("getprop")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if splitted := strings.SplitN(line, ": [", 2); len(splitted) == 2 {
				vals = append(vals, strings.TrimPrefix(splitted[0], "["), strings.TrimSuffix(splitted[1], "]"))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("properties")
}
