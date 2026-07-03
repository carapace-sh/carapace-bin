package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gdu",
	Short: "Pretty fast disk usage analyzer written in Go",
	Long:  "https://github.com/dundee/gdu",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().Bool("archive-browsing", false, "Enable browsing of zip/jar/tar archives")
	rootCmd.Flags().Bool("collapse-path", false, "Collapse single-child directory chains")
	rootCmd.Flags().String("config-file", "", "Read config from file (default is $HOME/.gdu.yaml)")
	rootCmd.Flags().BoolP("const-gc", "g", false, "Enable memory garbage collection during analysis with constant level set by GOGC")
	rootCmd.Flags().StringP("db", "D", "", "Store analysis in database (*.sqlite for SQLite, *.badger for BadgerDB)")
	rootCmd.Flags().Int("depth", 0, "Show directory structure up to specified depth in non-interactive mode")
	rootCmd.Flags().Bool("enable-profiling", false, "Enable collection of profiling data and provide it on http://localhost:6060/debug/pprof/")
	rootCmd.Flags().BoolP("help", "h", false, "help for gdu")
	rootCmd.Flags().StringSliceP("ignore-dirs", "i", []string{"/proc", "/dev", "/sys", "/run"}, "Absolute paths to ignore (separated by comma)")
	rootCmd.Flags().StringSliceP("ignore-dirs-pattern", "I", nil, "Absolute path patterns to ignore (separated by comma)")
	rootCmd.Flags().StringP("ignore-from", "X", "", "Read absolute path patterns to ignore from file")
	rootCmd.Flags().StringP("input-file", "f", "", "Import analysis from JSON file")
	rootCmd.Flags().Bool("interactive", false, "Force interactive mode even when output is not a TTY")
	rootCmd.Flags().StringP("log-file", "l", "/dev/null", "Path to a logfile")
	rootCmd.Flags().String("max-age", "", "Include files with mtime no older than DURATION (e.g., 7d, 2h30m, 1y2mo)")
	rootCmd.Flags().IntP("max-cores", "m", 4, "Set max cores that GDU will use. 4 cores available")
	rootCmd.Flags().String("min-age", "", "Include files with mtime at least DURATION old (e.g., 30d, 1w)")
	rootCmd.Flags().Bool("mouse", false, "Use mouse")
	rootCmd.Flags().BoolP("no-color", "c", false, "Do not use colorized output")
	rootCmd.Flags().BoolP("no-cross", "x", false, "Do not cross filesystem boundaries")
	rootCmd.Flags().Bool("no-delete", false, "Do not allow deletions")
	rootCmd.Flags().BoolP("no-hidden", "H", false, "Ignore hidden directories (beginning with dot)")
	rootCmd.Flags().Bool("no-prefix", false, "Show sizes as raw numbers without any prefixes in non-interactive mode")
	rootCmd.Flags().BoolP("no-progress", "p", false, "Do not show progress in non-interactive mode")
	rootCmd.Flags().Bool("no-spawn-shell", false, "Do not allow spawning shell")
	rootCmd.Flags().BoolP("no-unicode", "u", false, "Do not use Unicode symbols (for size bar)")
	rootCmd.Flags().Bool("no-view-file", false, "Do not allow viewing file contents")
	rootCmd.Flags().BoolP("non-interactive", "n", false, "Do not run in interactive mode")
	rootCmd.Flags().StringP("output-file", "o", "", "Export all info into file as JSON")
	rootCmd.Flags().Bool("reverse-sort", false, "Reverse sorting order (smallest to largest) in non-interactive mode")
	rootCmd.Flags().Bool("sequential", false, "Use sequential scanning (intended for rotating HDDs)")
	rootCmd.Flags().BoolP("show-annexed-size", "A", false, "Use apparent size of git-annex'ed files")
	rootCmd.Flags().BoolP("show-apparent-size", "a", false, "Show apparent size")
	rootCmd.Flags().BoolP("show-disks", "d", false, "Show all mounted disks")
	rootCmd.Flags().BoolP("show-relative-size", "B", false, "Show relative size")
	rootCmd.Flags().Bool("si", false, "Show sizes with decimal SI prefixes (kB, MB, GB) instead of binary prefixes (KiB, MiB, GiB)")
	rootCmd.Flags().String("since", "", "Include files with mtime >= WHEN (RFC3339 timestamp or YYYY-MM-DD)")
	rootCmd.Flags().BoolP("summarize", "s", false, "Show only a total in non-interactive mode")
	rootCmd.Flags().String("until", "", "Include files with mtime <= WHEN (RFC3339 timestamp or YYYY-MM-DD)")
	rootCmd.Flags().BoolP("version", "v", false, "Print version")
	rootCmd.Flags().Bool("write-config", false, "Write current configuration to file (default is $HOME/.gdu.yaml)")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
		"db":          carapace.ActionFiles(),
		"ignore-dirs": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionDirectories().NoSpace()
		}),
		"ignore-from": carapace.ActionFiles(),
		"input-file":  carapace.ActionFiles(),
		"log-file":    carapace.ActionFiles(),
		"output-file": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
