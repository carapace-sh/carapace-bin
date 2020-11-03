package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().Bool("args", false, "pass the remainder of the command line to the test binary")
	testCmd.Flags().BoolS("c", "c", false, "compile the test binary to pkg.test but do not run it")
	testCmd.Flags().Bool("cpu", false, "cacheable flag")
	testCmd.Flags().String("exec", "", "run the test binary using xprog")
	testCmd.Flags().BoolS("i", "i", false, "install packages that are dependencies of the test")
	testCmd.Flags().Bool("json", false, "convert test output to JSON")
	testCmd.Flags().Bool("list", false, "cacheable flag")
	testCmd.Flags().StringS("o", "o", "", "compile the test binary to the named file")
	testCmd.Flags().Bool("parallel", false, "cacheable flag")
	testCmd.Flags().Bool("run", false, "cacheable flag")
	testCmd.Flags().Bool("short", false, "cacheable flag")
	testCmd.Flags().Bool("v", false, "cacheable flag")
	rootCmd.AddCommand(testCmd)
}
