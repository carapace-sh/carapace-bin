package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "zpaq",
	Short:              "journaling archiver for incremental backups",
	Long:               "https://mattmahoney.net/dc/zpaq.html",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"add", "append changes in files to archive",
			"a", "append changes in files to archive",
			"extract", "extract files from archive",
			"x", "extract files from archive",
			"list", "list the archive contents",
			"l", "list the archive contents",
		),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return carapace.ActionValues()
			}

			switch c.Args[0] {
			case "add", "a":
				return actionArgs(c.Args[1:], addFlags())
			case "extract", "x":
				return actionArgs(c.Args[1:], extractFlags())
			case "list", "l":
				return actionArgs(c.Args[1:], listFlags())
			default:
				return carapace.ActionValues()
			}
		}),
	)
}

type flagSpec struct {
	description string
	value       carapace.Action // completion for flag value (ActionValues() for bools)
	takesValue  bool            // false for bools, true if flag consumes next arg(s)
	embedded    bool            // true if value is embedded in the flag itself (e.g. -m1, -s5, -t4)
	equivalents []string        // other flag names that are equivalent (e.g. -m and -method)
}

type flagMap map[string]flagSpec

// usedFlags returns a set of flag names already present in args,
// including embedded forms (e.g. -m1 matches -m) and their equivalents.
func usedFlags(args []string, flags flagMap) map[string]bool {
	used := make(map[string]bool)
	for _, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			continue
		}
		name := arg
		if _, ok := flags[arg]; !ok {
			// Check embedded form: -m1, -s5, -t4
			for n, spec := range flags {
				if spec.embedded && strings.HasPrefix(arg, n) && len(arg) > len(n) {
					name = n
					break
				}
			}
		}
		if spec, ok := flags[name]; ok {
			used[name] = true
			for _, eq := range spec.equivalents {
				used[eq] = true
			}
		}
	}
	return used
}

func addFlags() flagMap {
	return flagMap{
		"-all":          {"list all saved versions", carapace.ActionValues("2", "3", "4"), true, false, nil},
		"-f":            {"add files even if date unchanged", carapace.ActionValues(), false, false, []string{"-force"}},
		"-force":        {"add files even if date unchanged", carapace.ActionValues(), false, false, []string{"-f"}},
		"-fragment":     {"set dedupe fragment size range", carapace.ActionValues("0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"), true, false, nil},
		"-index":        {"create archive suffix for remote backup", carapace.ActionFiles(), true, false, nil},
		"-key":          {"encrypt with password", carapace.ActionValues(), true, false, nil},
		"-m":            {"compression method (e.g. -m1)", actionMethods(), false, true, []string{"-method"}},
		"-method":       {"select compression method", actionMethods(), true, false, []string{"-m"}},
		"-noattributes": {"do not save attributes or permissions", carapace.ActionValues(), false, false, nil},
		"-not":          {"do not add files matching pattern", carapace.ActionFiles(), true, false, nil},
		"-only":         {"only add files matching pattern", carapace.ActionFiles(), true, false, nil},
		"-s":            {"show only percent completed (e.g. -s5)", carapace.ActionValues(), false, true, []string{"-summary"}},
		"-summary":      {"show only percent completed", carapace.ActionValues(), true, false, []string{"-s"}},
		"-t":            {"add at most N blocks in parallel (e.g. -t4)", carapace.ActionValues("0", "1", "2", "3", "4", "5", "6", "7", "8"), false, true, []string{"-threads"}},
		"-threads":      {"add at most N blocks in parallel", carapace.ActionValues("0", "1", "2", "3", "4", "5", "6", "7", "8"), true, false, []string{"-t"}},
		"-to":           {"rename external files to internal names", carapace.ActionFiles(), true, false, nil},
		"-until":        {"ignore part of archive after date or version", carapace.ActionValues(), true, false, nil},
	}
}

func extractFlags() flagMap {
	return flagMap{
		"-all":          {"extract all saved versions in numbered directories", carapace.ActionValues("2", "3", "4"), true, false, nil},
		"-f":            {"overwrite existing output files", carapace.ActionValues(), false, false, []string{"-force"}},
		"-force":        {"overwrite existing output files", carapace.ActionValues(), false, false, []string{"-f"}},
		"-index":        {"create index for archive without extracting", carapace.ActionFiles(), true, false, nil},
		"-key":          {"decrypt with password", carapace.ActionValues(), true, false, nil},
		"-noattributes": {"ignore saved attributes and use defaults", carapace.ActionValues(), false, false, nil},
		"-not":          {"do not extract files matching pattern", carapace.ActionFiles(), true, false, nil},
		"-only":         {"only extract files matching pattern", carapace.ActionFiles(), true, false, nil},
		"-repack":       {"store extracted files in new archive", carapace.ActionFiles(), true, false, nil},
		"-s":            {"show only percent completed (e.g. -s5)", carapace.ActionValues(), false, true, []string{"-summary"}},
		"-summary":      {"show only percent completed", carapace.ActionValues(), true, false, []string{"-s"}},
		"-t":            {"extract at most N blocks in parallel (e.g. -t4)", carapace.ActionValues("0", "1", "2", "3", "4", "5", "6", "7", "8"), false, true, []string{"-threads"}},
		"-threads":      {"extract at most N blocks in parallel", carapace.ActionValues("0", "1", "2", "3", "4", "5", "6", "7", "8"), true, false, []string{"-t"}},
		"-test":         {"do not write to disk", carapace.ActionValues(), false, false, nil},
		"-to":           {"rename internal files to external names", carapace.ActionFiles(), true, false, nil},
		"-until":        {"ignore part of archive after date or version", carapace.ActionValues(), true, false, nil},
	}
}

func listFlags() flagMap {
	return flagMap{
		"-all":          {"list all saved versions", carapace.ActionValues("2", "3", "4"), true, false, nil},
		"-f":            {"compare files by computing SHA-1 hashes", carapace.ActionValues(), false, false, []string{"-force"}},
		"-force":        {"compare files by computing SHA-1 hashes", carapace.ActionValues(), false, false, []string{"-f"}},
		"-key":          {"decrypt with password", carapace.ActionValues(), true, false, nil},
		"-noattributes": {"do not list or compare attributes", carapace.ActionValues(), false, false, nil},
		"-not":          {"do not list files matching pattern", actionNotList(), true, false, nil},
		"-only":         {"only list files matching pattern", carapace.ActionFiles(), true, false, nil},
		"-s":            {"sort by decreasing size (e.g. -s5)", carapace.ActionValues(), false, true, []string{"-summary"}},
		"-summary":      {"sort by decreasing size and show N largest", carapace.ActionValues(), true, false, []string{"-s"}},
		"-t":            {"use N threads (e.g. -t4)", carapace.ActionValues("0", "1", "2", "3", "4", "5", "6", "7", "8"), false, true, []string{"-threads"}},
		"-threads":      {"use N threads", carapace.ActionValues("0", "1", "2", "3", "4", "5", "6", "7", "8"), true, false, []string{"-t"}},
		"-to":           {"rename external files to respective internal names", carapace.ActionFiles(), true, false, nil},
		"-until":        {"ignore part of archive after date or version", carapace.ActionValues(), true, false, nil},
	}
}

func actionMethods() carapace.Action {
	return carapace.ActionValuesDescribed(
		"0", "store with deduplication but no compression",
		"1", "default, recommended for backups",
		"2", "slower compression, fast decompression",
		"3", "higher compression",
		"4", "higher compression",
		"5", "highest compression",
		"i", "create index (metadata only)",
		"x", "experimental journaling mode",
		"s", "experimental streaming mode",
	)
}

func actionNotList() carapace.Action {
	return carapace.Batch(
		carapace.ActionFiles(),
		carapace.ActionValuesDescribed(
			"=+", "exclude missing internal files",
			"=-", "exclude missing external files",
			"=#", "exclude different files",
			"=?", "exclude identical files",
			"=^", "exclude duplicates",
		),
	).ToA()
}

func actionArgs(args []string, flags flagMap) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(args) > 0 {
			prev := args[len(args)-1]

			if spec, ok := flags[prev]; ok && spec.takesValue {
				return spec.value
			}
		}

		if strings.HasPrefix(c.Value, "-") {
			used := usedFlags(args, flags)

			actions := []carapace.Action{actionFlagNames(flags, used)}

			// Only offer embedded values when prefix matches an embedded flag (e.g. -m, -s, -t)
			// and that flag hasn't been used yet
			for name, spec := range flags {
				if spec.embedded && !used[name] && strings.HasPrefix(c.Value, name) {
					actions = append(actions, spec.value.Prefix(name))
				}
			}

			return carapace.Batch(actions...).Invoke(c).Merge().ToA()
		}

		return carapace.ActionFiles()
	})
}

func actionFlagNames(flags flagMap, used map[string]bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := make([]string, 0, len(flags)*2)
		for name, spec := range flags {
			if used[name] {
				continue
			}
			vals = append(vals, name, spec.description)
		}
		return carapace.ActionValuesDescribed(vals...).Style(style.Blue).Tag("flags")
	})
}
