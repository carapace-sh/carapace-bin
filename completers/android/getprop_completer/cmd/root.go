package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "getprop",
	Short: "Get an Android system property",
	Long:  "https://cs.android.com/android/platform/superproject/main/+/main:system/core/toolbox/getprop.cpp",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("show-contexts", "Z", false, "show property contexts instead of values")
	rootCmd.Flags().BoolP("show-types", "T", false, "show property types instead of values")

	carapace.Gen(rootCmd).PositionalCompletion(
		actionProperties(),
		carapace.ActionValues(), // DEFAULT
	)
}

func actionProperties() carapace.Action {
	return carapace.ActionExecCommand("getprop")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if splitted := strings.SplitN(line, ": [", 2); len(splitted) == 2 {
				name := strings.TrimPrefix(splitted[0], "[")
				value := strings.TrimSuffix(splitted[1], "]")
				vals = append(vals, name, value)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("properties")
}
