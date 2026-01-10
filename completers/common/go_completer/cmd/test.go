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

	testCmd.Flags().StringS("bench", "bench", "", "Run only those benchmarks matching a regular expression")
	testCmd.Flags().BoolS("benchmem", "benchmem", false, "Print memory allocation statistics for benchmarks")
	testCmd.Flags().StringS("benchtime", "benchtime", "", "Run enough iterations of each benchmark to take t")
	testCmd.Flags().StringS("blockprofile", "blockprofile", "", "Write a goroutine blocking profile to the specified file when all tests are complete")
	testCmd.Flags().StringS("blockprofilerate", "blockprofilerate", "", "Control the detail provided in goroutine blocking profiles by calling runtime.SetBlockProfileRate with n")
	testCmd.Flags().BoolS("c", "c", false, "Compile the test binary to pkg.test in the current directory but do not run it")
	testCmd.Flags().StringS("count", "count", "", "Run each test, benchmark, and fuzz seed n times")
	testCmd.Flags().StringS("covermode", "covermode", "", "Set the mode for coverage analysis for the package[s] being tested")
	testCmd.Flags().StringS("coverprofile", "coverprofile", "", "Write a coverage profile to the file after all tests have passed")
	testCmd.Flags().StringS("cpu", "cpu", "", "Specify a list of GOMAXPROCS values for which the tests, benchmarks or fuzz tests should be executed")
	testCmd.Flags().StringS("cpuprofile", "cpuprofile", "", "Write a CPU profile to the specified file before exiting")
	testCmd.Flags().StringS("exec", "exec", "", "Run the test binary using xprog")
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
	testCmd.Flags().StringS("o", "o", "", "Compile the test binary to the named file")
	testCmd.Flags().StringS("outputdir", "outputdir", "", "Place output files from profiling in the specified directory")
	testCmd.Flags().StringS("parallel", "parallel", "", "Allow parallel execution of test functions")
	testCmd.Flags().StringS("run", "run", "", "Run only those tests, examples, and fuzz tests matching the regular expression")
	testCmd.Flags().BoolS("short", "short", false, "Tell long-running tests to shorten their run time")
	testCmd.Flags().StringS("shuffle", "shuffle", "", "Randomize the execution order of tests and benchmarks")
	testCmd.Flags().StringS("skip", "skip", "", "Run only those tests, examples, fuzz tests, and benchmarks that do not match the regular expression")
	testCmd.Flags().StringS("test.bench", "test.bench", "", "run only benchmarks matching regexp")
	testCmd.Flags().BoolS("test.benchmem", "test.benchmem", false, "print memory allocations for benchmarks")
	testCmd.Flags().StringS("test.benchtime", "test.benchtime", "", "run each benchmark for duration d or N times if `d` is of the form Nx")
	testCmd.Flags().StringS("test.blockprofile", "test.blockprofile", "", "write a goroutine blocking profile to file")
	testCmd.Flags().StringS("test.blockprofilerate", "test.blockprofilerate", "", "set blocking profile rate")
	testCmd.Flags().StringS("test.count", "test.count", "", "run tests and benchmarks n times")
	testCmd.Flags().StringS("test.coverprofile", "test.coverprofile", "", "write a coverage profile to file")
	testCmd.Flags().StringS("test.cpu", "test.cpu", "", "comma-separated list of cpu counts to run each test with")
	testCmd.Flags().StringS("test.cpuprofile", "test.cpuprofile", "", "write a cpu profile to file")
	testCmd.Flags().BoolS("test.failfast", "test.failfast", false, "do not start new tests after the first test failure")
	testCmd.Flags().BoolS("test.fullpath", "test.fullpath", false, "show full file names in error messages")
	testCmd.Flags().StringS("test.fuzz", "test.fuzz", "", "run the fuzz test matching regexp")
	testCmd.Flags().StringS("test.fuzzcachedir", "test.fuzzcachedir", "", "directory where interesting fuzzing inputs are stored")
	testCmd.Flags().StringS("test.fuzzminimizetime", "test.fuzzminimizetime", "", "time to spend minimizing a value after finding a failing input")
	testCmd.Flags().StringS("test.fuzztime", "test.fuzztime", "", "time to spend fuzzing; default is to run indefinitely")
	testCmd.Flags().BoolS("test.fuzzworker", "test.fuzzworker", false, "coordinate with the parent process to fuzz random values")
	testCmd.Flags().StringS("test.gocoverdir", "test.gocoverdir", "", "write coverage intermediate files to this directory")
	testCmd.Flags().StringS("test.list", "test.list", "", "list tests, examples, and benchmarks matching regexp then exit")
	testCmd.Flags().StringS("test.memprofile", "test.memprofile", "", "write an allocation profile to file")
	testCmd.Flags().StringS("test.memprofilerate", "test.memprofilerate", "", "set memory allocation profiling rate")
	testCmd.Flags().StringS("test.mutexprofile", "test.mutexprofile", "", "write a mutex contention profile to the named file after execution")
	testCmd.Flags().StringS("test.mutexprofilefraction", "test.mutexprofilefraction", "", "if >= 0, calls runtime.SetMutexProfileFraction()")
	testCmd.Flags().StringS("test.outputdir", "test.outputdir", "", "write profiles to dir")
	testCmd.Flags().BoolS("test.paniconexit0", "test.paniconexit0", false, "panic on call to os.Exit(0)")
	testCmd.Flags().StringS("test.parallel", "test.parallel", "", "run at most n tests in parallel")
	testCmd.Flags().StringS("test.run", "test.run", "", "run only tests and examples matching regexp")
	testCmd.Flags().BoolS("test.short", "test.short", false, "run smaller test suite to save time")
	testCmd.Flags().StringS("test.shuffle", "test.shuffle", "", "randomize the execution order of tests and benchmarks")
	testCmd.Flags().StringS("test.skip", "test.skip", "", "do not list or run tests matching regexp")
	testCmd.Flags().StringS("test.testlogfile", "test.testlogfile", "", "write test action log to file")
	testCmd.Flags().StringS("test.timeout", "test.timeout", "", "panic test binary after duration d")
	testCmd.Flags().StringS("test.trace", "test.trace", "", "write an execution trace to file")
	testCmd.Flags().BoolS("test.v", "test.v", false, "verbose: print additional output")
	testCmd.Flags().StringS("timeout", "timeout", "", "If a test binary runs longer than duration d, panic")
	testCmd.Flags().StringS("trace", "trace", "", "Write an execution trace to the specified file before exiting")
	testCmd.Flags().StringS("vet", "vet", "", "Configure the invocation of \"go vet\" during \"go test\" to use the comma-separated list of vet checks")
	addBuildFlags(testCmd)
	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"bench":             golang.ActionTests(golang.TestOpts{Benchmark: true}).UniqueList("/"),
		"blockprofile":      carapace.ActionFiles(),
		"covermode":         carapace.ActionValues("set", "count,atomic"),
		"coverpkg":          golang.ActionPackages().List(","),
		"coverprofile":      carapace.ActionFiles(),
		"cpuprofile":        carapace.ActionFiles(),
		"fuzz":              golang.ActionTests(golang.TestOpts{Fuzz: true}),
		"list":              golang.ActionTests(golang.TestOpts{}.Default()),
		"memprofile":        carapace.ActionFiles(),
		"mutexprofile":      carapace.ActionFiles(),
		"o":                 carapace.ActionFiles(),
		"outputdir":         carapace.ActionDirectories(),
		"run":               golang.ActionTests(golang.TestOpts{}.Default()).UniqueList("/"),
		"shuffle":           carapace.ActionValues("off", "on").StyleF(style.ForKeyword),
		"skip":              golang.ActionTests(golang.TestOpts{}.Default()).UniqueList("/"),
		"test.bench":        golang.ActionTests(golang.TestOpts{Benchmark: true}),
		"test.blockprofile": carapace.ActionFiles(),
		"test.coverprofile": carapace.ActionFiles(),
		"test.cpuprofile":   carapace.ActionFiles(),
		"test.fuzz":         golang.ActionTests(golang.TestOpts{Fuzz: true}),
		"test.fuzzcachedir": carapace.ActionDirectories(),
		"test.gocoverdir":   carapace.ActionDirectories(),
		"test.list":         golang.ActionTests(golang.TestOpts{Benchmark: true, Example: true, Test: true}),
		"test.memprofile":   carapace.ActionFiles(),
		"test.mutexprofile": carapace.ActionFiles(),
		"test.outputdir":    carapace.ActionDirectories(),
		"test.run":          golang.ActionTests(golang.TestOpts{Example: true, Test: true}),
		"test.skip":         golang.ActionTests(golang.TestOpts{Test: true}),
		"test.testlogfile":  carapace.ActionFiles(),
		"test.trace":        carapace.ActionFiles(),
		"trace":             carapace.ActionFiles(),
		"vet": carapace.Batch(
			carapace.ActionValues("off").StyleF(style.ForKeyword),
			golang.ActionAnalyzers(),
		).ToA().UniqueList(","),
	})

	carapace.Gen(testCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.Chdir(cmd.Flag("C").Value.String())
	})

}
