package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run tests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().StringS("C", "C", "", "directory to cd into before running")
	testCmd.Flags().Bool("benchmark", false, "Run benchmarks instead of tests.")
	testCmd.Flags().String("exclude", "", "Exclude tests with the given name.")
	testCmd.Flags().Bool("gdb", false, "Run test under gdb.")
	testCmd.Flags().String("gdb-path", "", "Path to the gdb binary (default: gdb).")
	testCmd.Flags().BoolP("interactive", "i", false, "Run tests with interactive input/output.")
	testCmd.Flags().Bool("list", false, "List available tests.")
	testCmd.Flags().String("logbase", "", "Base name for log file.")
	testCmd.Flags().String("max-lines", "", "Maximum number of lines to show from a long test log.")
	testCmd.Flags().StringP("maxfail", "k", "", "Number of failing tests before aborting the test run.")
	testCmd.Flags().Bool("no-rebuild", false, "Do not rebuild before running tests.")
	testCmd.Flags().Bool("no-stdsplit", false, "Do not split stderr and stdout in test logs.")
	testCmd.Flags().String("no-suite", "", "Do not run tests belonging to the given suite.")
	testCmd.Flags().StringP("num-processes", "j", "", "How many parallel processes to use.")
	testCmd.Flags().Bool("print-errorlogs", false, "Whether to print failing tests' logs.")
	testCmd.Flags().BoolP("quiet", "q", false, "Produce less output to the terminal.")
	testCmd.Flags().String("repeat", "", "Number of times to run the tests.")
	testCmd.Flags().String("setup", "", "Which test setup to use.")
	testCmd.Flags().String("slice", "", "Split tests into NUM_SLICES slices and execute slice SLICE.")
	testCmd.Flags().String("suite", "", "Only run tests belonging to the given suite.")
	testCmd.Flags().String("test-args", "", "Arguments to pass to the specified test(s) or all tests")
	testCmd.Flags().StringP("timeout-multiplier", "t", "", "Define a multiplier for test timeout, for example  when running tests in particular conditions they might take more time to execute.")
	testCmd.Flags().BoolP("verbose", "v", false, "Do not redirect stdout and stderr")
	testCmd.Flags().String("wrapper", "", "wrapper to run tests with (e.g. Valgrind)")
	applyCommonFlagCompletion(testCmd)
	rootCmd.AddCommand(testCmd)
}
