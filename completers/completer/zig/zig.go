package zig

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zig",
	Short: "Programming language designed for robustness, optimality, and clarity",
	Long:  "https://ziglang.org/",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddCommand(
		newArCmd(),
		newBuildCmd(),
		newBuildExeCmd(),
		newBuildLibCmd(),
		newBuildObjCmd(),
		newCcCmd(),
		newCppCmd(),
		newDlltoolCmd(),
		newEnvCmd(),
		newFetchCmd(),
		newFmtCmd(),
		newHelpCmd(),
		newInitCmd(),
		newLibCmd(),
		newLibcCmd(),
		newObjcopyCmd(),
		newRanlibCmd(),
		newRunCmd(),
		newStripCmd(),
		newTargetsCmd(),
		newTestCmd(),
		newTranslateCCmd(),
		newVersionCmd(),
		newZenCmd(),
	)
}
