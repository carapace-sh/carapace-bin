package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/python"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pytest",
	Short: "simple powerful testing with Python",
	Long:  "https://pytest.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("assert", "", "control assertion debugging tools")
	rootCmd.Flags().String("basetemp", "", "base temporary directory for this test run")
	rootCmd.Flags().Bool("cache-clear", false, "remove all cache contents at start of test run")
	rootCmd.Flags().String("cache-show", "", "show cache contents, don't perform collection or tests")
	rootCmd.Flags().String("capture", "", "per-test capturing method: one of fd|sys|no|tee-sys")
	rootCmd.Flags().String("code-highlight", "", "whether code should be highlighted (only if --color is also enabled)")
	rootCmd.Flags().Bool("collect-in-virtualenv", false, "don't ignore tests in a local virtualenv directory")
	rootCmd.Flags().Bool("collect-only", false, "only collect tests, don't execute them")
	rootCmd.Flags().String("color", "", "color terminal output (yes/no/auto)")
	rootCmd.Flags().String("confcutdir", "", "only load conftest.py's relative to specified dir")
	rootCmd.Flags().StringP("config-file", "c", "", "load configuration from `FILE` instead of trying to locate one of the implicit configuration files")
	rootCmd.Flags().Bool("continue-on-collection-errors", false, "force test execution even if collection errors occur")
	rootCmd.Flags().String("debug", "", "store internal tracing debug information in this log file")
	rootCmd.Flags().StringArray("deselect", nil, "deselect item (via node id prefix) during collection (multi-allowed)")
	rootCmd.Flags().Bool("disable-plugin-autoload", false, "disable plugin auto-loading through entry point packaging metadata")
	rootCmd.Flags().Bool("disable-warnings", false, "disable warnings summary")
	rootCmd.Flags().Bool("doctest-continue-on-failure", false, "for a given doctest, continue to run after the first failure")
	rootCmd.Flags().String("doctest-glob", "", "doctests file matching pattern, default: test*.txt")
	rootCmd.Flags().Bool("doctest-ignore-import-errors", false, "ignore doctest collection errors")
	rootCmd.Flags().Bool("doctest-modules", false, "run doctests in all .py modules")
	rootCmd.Flags().String("doctest-report", "", "choose another output format for diffs on doctest failure")
	rootCmd.Flags().String("durations", "", "show N slowest setup/test durations (N=0 for all)")
	rootCmd.Flags().String("durations-min", "", "minimal duration in seconds for inclusion in slowest list")
	rootCmd.Flags().BoolP("exitfirst", "x", false, "exit instantly on first error or failed test")
	rootCmd.Flags().Bool("failed-first", false, "run all tests, but run the last failures first")
	rootCmd.Flags().Bool("fixtures", false, "show available fixtures, sorted by plugin appearance")
	rootCmd.Flags().Bool("fixtures-per-test", false, "show fixtures per test")
	rootCmd.Flags().Bool("force-short-summary", false, "force condensed summary output regardless of verbosity level")
	rootCmd.Flags().Bool("full-trace", false, "don't cut any tracebacks")
	rootCmd.Flags().BoolP("help", "h", false, "show help message and configuration info")
	rootCmd.Flags().String("ignore", "", "ignore path during collection (multi-allowed)")
	rootCmd.Flags().String("ignore-glob", "", "ignore path pattern during collection (multi-allowed)")
	rootCmd.Flags().String("import-mode", "", "prepend/append to sys.path when importing test modules")
	rootCmd.Flags().String("junit-prefix", "", "prepend prefix to classnames in junit-xml output")
	rootCmd.Flags().String("junit-xml", "", "create junit-xml style report file at given path")
	rootCmd.Flags().StringS("k", "k", "", "only run tests which match the given substring expression")
	rootCmd.Flags().Bool("keep-duplicates", false, "keep duplicate tests")
	rootCmd.Flags().Bool("last-failed", false, "rerun only the tests that failed at the last run")
	rootCmd.Flags().String("last-failed-no-failures", "", "determines whether to execute tests when there are no previously known failures")
	rootCmd.Flags().String("log-auto-indent", "", "auto-indent multiline messages passed to the logging module")
	rootCmd.Flags().String("log-cli-date-format", "", "log date format used by the logging module")
	rootCmd.Flags().String("log-cli-format", "", "log format used by the logging module")
	rootCmd.Flags().String("log-cli-level", "", "CLI logging level")
	rootCmd.Flags().String("log-date-format", "", "log date format used by the logging module")
	rootCmd.Flags().StringArray("log-disable", nil, "disable a logger by name (multi-allowed)")
	rootCmd.Flags().String("log-file", "", "path to a file when logging will be written to")
	rootCmd.Flags().String("log-file-date-format", "", "log date format used by the logging module")
	rootCmd.Flags().String("log-file-format", "", "log format used by the logging module")
	rootCmd.Flags().String("log-file-level", "", "log file logging level")
	rootCmd.Flags().String("log-file-mode", "", "log file open mode")
	rootCmd.Flags().String("log-format", "", "log format used by the logging module")
	rootCmd.Flags().String("log-level", "", "level of messages to catch/display")
	rootCmd.Flags().StringS("m", "m", "", "only run tests matching given mark expression")
	rootCmd.Flags().Bool("markers", false, "show markers (builtin, plugin and per-project ones)")
	rootCmd.Flags().String("maxfail", "", "exit after first num failures or errors")
	rootCmd.Flags().Bool("new-first", false, "run tests from new files first, then the rest sorted by file mtime")
	rootCmd.Flags().Bool("no-fold-skipped", false, "do not fold skipped tests in short summary")
	rootCmd.Flags().Bool("no-header", false, "disable header")
	rootCmd.Flags().Bool("no-showlocals", false, "hide locals in tracebacks")
	rootCmd.Flags().Bool("no-summary", false, "disable summary")
	rootCmd.Flags().Bool("noconftest", false, "don't load any conftest.py files")
	rootCmd.Flags().StringArrayP("override-ini", "o", nil, "override configuration option with \"option=value\" style")
	rootCmd.Flags().StringArrayS("p", "p", nil, "early-load given plugin module name or entry point (multi-allowed)")
	rootCmd.Flags().String("pastebin", "", "send failed|all info to bpaste.net pastebin service")
	rootCmd.Flags().Bool("pdb", false, "start the interactive Python debugger on errors or KeyboardInterrupt")
	rootCmd.Flags().String("pdbcls", "", "specify a custom interactive Python debugger for use with --pdb")
	rootCmd.Flags().Bool("pyargs", false, "try to interpret all arguments as Python packages")
	rootCmd.Flags().StringArrayP("pythonwarnings", "W", nil, "set which warnings to report, see -W option of Python itself")
	rootCmd.Flags().BoolP("quiet", "q", false, "decrease verbosity")
	rootCmd.Flags().StringS("r", "r", "", "show extra test summary info as specified by chars")
	rootCmd.Flags().String("rootdir", "", "define root directory for tests")
	rootCmd.Flags().Bool("runxfail", false, "report the results of xfail tests as if they were not marked")
	rootCmd.Flags().BoolS("s", "s", false, "shortcut for --capture=no")
	rootCmd.Flags().Bool("setup-only", false, "only setup fixtures, do not execute tests")
	rootCmd.Flags().Bool("setup-plan", false, "show what fixtures and tests would be executed but don't execute anything")
	rootCmd.Flags().Bool("setup-show", false, "show setup of fixtures while executing tests")
	rootCmd.Flags().String("show-capture", "", "controls how captured stdout/stderr/log is shown on failed tests")
	rootCmd.Flags().BoolP("showlocals", "l", false, "show locals in tracebacks")
	rootCmd.Flags().Bool("stepwise", false, "exit on test failure and continue from last failing test next time")
	rootCmd.Flags().Bool("stepwise-reset", false, "reset stepwise state, restarting the stepwise workflow")
	rootCmd.Flags().Bool("stepwise-skip", false, "ignore the first failing test but stop on the next failing test")
	rootCmd.Flags().Bool("strict", false, "enable the strict option")
	rootCmd.Flags().Bool("strict-config", false, "enable the strict_config option")
	rootCmd.Flags().Bool("strict-markers", false, "enable the strict_markers option")
	rootCmd.Flags().String("tb", "", "traceback print mode (auto/long/short/line/native/no)")
	rootCmd.Flags().Bool("trace", false, "immediately break when running each test")
	rootCmd.Flags().Bool("trace-config", false, "trace considerations of conftest.py files")
	rootCmd.Flags().BoolP("verbose", "V", false, "display pytest version and information about plugins")
	rootCmd.Flags().String("verbosity", "", "set verbosity")
	rootCmd.Flags().Bool("xfail-tb", false, "show tracebacks for xfail")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"assert":                  carapace.ActionValues("plain", "rewrite"),
		"basetemp":                carapace.ActionDirectories(),
		"cache-show":              carapace.ActionFiles(),
		"capture":                 carapace.ActionValues("fd", "sys", "no", "tee-sys"),
		"code-highlight":          carapace.ActionValues("yes", "no"),
		"color":                   carapace.ActionValues("yes", "no", "auto"),
		"confcutdir":              carapace.ActionDirectories(),
		"config-file":             carapace.ActionFiles(),
		"debug":                   carapace.ActionFiles(),
		"doctest-glob":            carapace.ActionFiles(),
		"doctest-report":          carapace.ActionValues("none", "cdiff", "ndiff", "udiff", "only_first_failure"),
		"ignore":                  carapace.ActionFiles(),
		"ignore-glob":             carapace.ActionFiles(),
		"import-mode":             carapace.ActionValues("prepend", "append", "importlib"),
		"junit-xml":               carapace.ActionFiles(),
		"last-failed-no-failures": carapace.ActionValues("all", "none"),
		"log-file":                carapace.ActionFiles(),
		"log-file-mode":           carapace.ActionValues("w", "a"),
		"pastebin":                carapace.ActionValues("failed", "all"),
		"pythonwarnings":          python.ActionWarningControls(),
		"rootdir":                 carapace.ActionDirectories(),
		"show-capture":            carapace.ActionValues("no", "stdout", "stderr", "log", "all"),
		"tb":                      carapace.ActionValues("auto", "long", "short", "line", "native", "no"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
