package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/go_completer/cmd/action"
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
	testCmd.Flags().String("bench", "", "run only those benchmarks matching a regular expression")
	testCmd.Flags().String("benchtime", "", "Run enough iterations of each benchmark to take given duration")
	testCmd.Flags().BoolS("c", "c", false, "compile the test binary to pkg.test but do not run it")
	testCmd.Flags().String("count", "", "run each test and benchmark n times")
	testCmd.Flags().Bool("cover", false, "enable coverage analysis")
	testCmd.Flags().String("covermode", "", "set the mode for coverage analysis for the package")
	testCmd.Flags().String("coverpkg", "", "apply coverage analysis in each test to packages matching the patterns")
	testCmd.Flags().Bool("cpu", false, "specify a list of GOMAXPROCS values for which the tests or benchmarks should be executed")
	testCmd.Flags().String("exec", "", "run the test binary using xprog")
    testCmd.Flags().Bool("failfast", false, "Do not start new tests after the first test failure")
	testCmd.Flags().BoolS("i", "i", false, "install packages that are dependencies of the test")
	testCmd.Flags().Bool("json", false, "convert test output to JSON")
	testCmd.Flags().String("list", "", "list tests, benchmarks, or examples matching the regular expression")
	testCmd.Flags().StringS("o", "o", "", "compile the test binary to the named file")
	testCmd.Flags().String("parallel", "", "Allow parallel execution of test functions that call t.Paralle")
	testCmd.Flags().String("run", "", "run only those tests and examples matching the regular expression")
	testCmd.Flags().Bool("short", false, "tell long-running tests to shorten their run time")
	testCmd.Flags().String("timeout", "", "if a test binary runs longer than duration d, panic")
	testCmd.Flags().Bool("v", false, "verbose output")
	testCmd.Flags().String("vet", "", "configure the invocation of \"go vet\" during \"go test\" to use the comma-separated list of vet check")
	rootCmd.AddCommand(testCmd)

	testCmd.Flag("bench").NoOptDefVal = " "
	testCmd.Flag("run").NoOptDefVal = " "

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"bench": carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
			return action.ActionTests(c.Args, action.TestOpts{Benchmark: true}).Invoke(c).Filter(c.Parts).ToA()
		}),
		"covermode": carapace.ActionValues("set", "count,atomic"),
		"run": carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
			return action.ActionTests(c.Args, action.TestOpts{Example: true, Test: true}).Invoke(c).Filter(c.Parts).ToA()
		}),
	})

}
