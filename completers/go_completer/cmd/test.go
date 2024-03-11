package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()
	testCmd.Flags().SetInterspersed(false)

	testCmd.Flags().StringS("C", "C", "", "Change to dir before running the command")
	testCmd.Flags().StringS("bench", "bench", "", "Run only those benchmarks matching a regular expression")
	testCmd.Flags().BoolS("benchmem", "benchmem", false, "Print memory allocation statistics for benchmarks")
	testCmd.Flags().StringS("benchtime", "benchtime", "", "Run enough iterations of each benchmark to take t")
	testCmd.Flags().StringS("blockprofile", "blockprofile", "", "Write a goroutine blocking profile to the specified file when all tests are complete")
	testCmd.Flags().StringS("blockprofilerate", "blockprofilerate", "", "Control the detail provided in goroutine blocking profiles by calling runtime.SetBlockProfileRate with n")
	testCmd.Flags().StringS("count", "count", "", "Run each test, benchmark, and fuzz seed n times")
	testCmd.Flags().BoolS("cover", "cover", false, "Enable coverage analysis")
	testCmd.Flags().StringS("covermode", "covermode", "", "Set the mode for coverage analysis for the package[s] being tested")
	testCmd.Flags().StringS("coverpkg", "coverpkg", "", "Apply coverage analysis in each test to packages matching the patterns")
	testCmd.Flags().StringS("coverprofile", "coverprofile", "", "Write a coverage profile to the file after all tests have passed")
	testCmd.Flags().StringS("cpu", "cpu", "", "Specify a list of GOMAXPROCS values for which the tests, benchmarks or fuzz tests should be executed")
	testCmd.Flags().StringS("cpuprofile", "cpuprofile", "", "Write a CPU profile to the specified file before exiting")
	testCmd.Flags().BoolS("failfast", "failfast", false, "Do not start new tests after the first test failure.")
	testCmd.Flags().BoolS("fullpath", "fullpath", false, "Show full file names in the error messages")
	testCmd.Flags().StringS("fuzz", "fuzz", "", "Run the fuzz test matching the regular expression")
	testCmd.Flags().StringS("fuzzminimizetime", "fuzzminimizetime", "", "Run enough iterations of the fuzz target during each minimization attempt to take t")
	testCmd.Flags().StringS("fuzztime", "fuzztime", "", "Run enough iterations of the fuzz target during fuzzing to take t")
	testCmd.Flags().BoolS("json", "json", false, "Log verbose output and test results in JSON")
	testCmd.Flags().StringS("list", "list", "", "List tests, benchmarks, fuzz tests, or examples matching the regular expression")
	testCmd.Flags().StringS("memprofile", "memprofile", "", "Write an allocation profile to the file after all tests have passed")
	testCmd.Flags().StringS("memprofilerate", "memprofilerate", "", "Enable more precise (and expensive) memory allocation profiles by setting runtime.MemProfileRate")
	testCmd.Flags().StringS("mutexprofile", "mutexprofile", "", "Write a mutex contention profile to the specified file when all tests are complete")
	testCmd.Flags().StringS("mutexprofilefraction", "mutexprofilefraction", "", "Sample 1 in n stack traces of goroutines holding a contended mutex")
	testCmd.Flags().StringS("outputdir", "outputdir", "", "Place output files from profiling in the specified directory")
	testCmd.Flags().StringS("parallel", "parallel", "", "Allow parallel execution of test functions")
	testCmd.Flags().StringS("run", "run", "", "Run only those tests, examples, and fuzz tests matching the regular expression")
	testCmd.Flags().BoolS("short", "short", false, "Tell long-running tests to shorten their run time")
	testCmd.Flags().StringS("shuffle", "shuffle", "", "Randomize the execution order of tests and benchmarks")
	testCmd.Flags().StringS("skip", "skip", "", "Run only those tests, examples, fuzz tests, and benchmarks that do not match the regular expression")
	testCmd.Flags().StringS("timeout", "timeout", "", "If a test binary runs longer than duration d, panic")
	testCmd.Flags().StringS("trace", "trace", "", "Write an execution trace to the specified file before exiting")
	testCmd.Flags().BoolS("v", "v", false, "Verbose output: log all tests as they are run")
	testCmd.Flags().StringS("vet", "vet", "", "Configure the invocation of \"go vet\" during \"go test\" to use the comma-separated list of vet checks")
	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"C":            carapace.ActionDirectories(),
		"bench":        golang.ActionTests(golang.TestOpts{Benchmark: true}).UniqueList("/"),
		"blockprofile": carapace.ActionFiles(),
		"covermode":    carapace.ActionValues("set", "count,atomic"),
		"coverpkg":     golang.ActionPackages().List(","),
		"coverprofile": carapace.ActionFiles(),
		"cpuprofile":   carapace.ActionFiles(),
		"fuzz":         golang.ActionTests(golang.TestOpts{Fuzz: true}),
		"list":         golang.ActionTests(golang.TestOpts{}.Default()),
		"memprofile":   carapace.ActionFiles(),
		"mutexprofile": carapace.ActionFiles(),
		"outputdir":    carapace.ActionDirectories(),
		"run":          golang.ActionTests(golang.TestOpts{}.Default()).UniqueList("/"),
		"shuffle":      carapace.ActionValues("off", "on").StyleF(style.ForKeyword),
		"skip":         golang.ActionTests(golang.TestOpts{}.Default()).UniqueList("/"),
		"trace":        carapace.ActionFiles(),
		"vet": carapace.Batch(
			carapace.ActionValues("off").StyleF(style.ForKeyword),
			golang.ActionAnalyzers(),
		).ToA().UniqueList(","),
	})

	carapace.Gen(testCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.Chdir(cmd.Flag("C").Value.String())
	})

}
