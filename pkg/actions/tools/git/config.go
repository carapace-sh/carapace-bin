package git

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bridge/pkg/actions/bridge"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionConfigs completes configs
func ActionConfigs() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("git", "help", "--config")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)

			replacer := strings.NewReplacer(
				"*", "<*>",
				"branch.<name>", "branch.<branch>",
				"remote.<name>", "remote.<remote>",
				"submodule.<name>", "submodule.<submodule>",
			)

			for _, line := range lines[:len(lines)-3] {
				vals = append(vals, replacer.Replace(line))
			}

			return carapace.ActionValues(vals...).
				MultiPartsP(".", "<.*>", func(placeholder string, matches map[string]string) carapace.Action {
					// TODO support more placeholders
					switch placeholder {
					case "<branch>":
						return ActionLocalBranches()
					case "<remote>":
						return ActionRemotes()
					case "<submodule>":
						return ActionSubmoduleNames()
					default:
						return carapace.ActionValues()
					}
				})
		})
	})
}

// ActionColors completes colors
func ActionColors() carapace.Action {
	return carapace.ActionStyledValues(
		"normal", style.Default,
		"black", style.Black,
		"red", style.Red,
		"green", style.Green,
		"yellow", style.Yellow,
		"blue", style.Blue,
		"magenta", style.Magenta,
		"cyan", style.Cyan,
		"white", style.White,
	)
}

// ActionTextAttributes completes test attributes
func ActionTextAttributes() carapace.Action {
	return carapace.ActionStyledValues(
		"bold", style.Bold,
		"dim", style.Dim,
		"ul", style.Underlined,
		"blink", style.Blink,
		"reverse", style.Default,
	)
}

// ActionColorConfigs completes color config
func ActionColorConfigs() carapace.Action {
	return carapace.ActionMultiParts(" ", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionColors().NoSpace()
		case 1:
			return ActionColors().NoSpace()
		case 2:
			return ActionTextAttributes()
		default:
			return carapace.ActionValues()
		}
	})
}

// ActionConfigTypes completes config types
func ActionConfigTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"bool", "canonicalize values as either \"true\" or \"false\"",
		"int", "canonicalize values as simple decimal numbers",
		"bool-or-int", "canonicalize according to either bool or int",
		"path", "canonicalize by adding a leading ~ to the value of $HOME",
		"expiry-date", "canonicalize by converting from a fixed or relative date-string to a timestamp",
		"color", "When getting a value, canonicalize by converting to an ANSI color escape sequence",
	)
}

// ActionConfigTypeOptions completes options for a config type
func ActionConfigTypeOptions(t string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch t {
		case "bool":
			return carapace.ActionValues("true", "false")
		case "bool-or-int":
			integers := carapace.ActionValues("1", "2", "3", "4", "5", "6", "7", "8", "9", "0").Invoke(c).Prefix(c.Value)
			return integers.Merge(carapace.ActionValues("true", "false").Invoke(c)).ToA()
		case "int":
			return carapace.ActionValues("1", "2", "3", "4", "5", "6", "7", "8", "9", "0").Invoke(c).Prefix(c.Value).ToA()
		case "path":
			return carapace.ActionFiles()
		case "color":
			return ActionColorConfigs()
		default:
			return carapace.ActionValues()
		}
	})
}

// ActionConfigValues completes config values
func ActionConfigValues(config string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO complete more configurations
		_bool := carapace.ActionValues("true", "false").StyleF(style.ForKeyword)
		if a, ok := (carapace.ActionMap{
			"add.ignoreErrors":                     _bool,
			"add.interactive.useBuiltin":           _bool,
			"am.keepcr":                            _bool,
			"am.threeWay":                          _bool,
			"apply.ignoreWhitespace":               carapace.ActionValues("change", "no", "none", "never", "false").StyleF(style.ForKeyword),
			"apply.whitespace":                     _bool,
			"blame.blankBoundary":                  _bool,
			"blame.coloring":                       carapace.ActionValues("repeatedLines", "highlightRecent", "none").StyleF(style.ForKeyword),
			"blame.ignoreRevsFile":                 carapace.ActionFiles(),
			"blame.markIgnoredLines":               _bool,
			"blame.markUnblamableLines":            _bool,
			"blame.showEmail":                      _bool,
			"blame.showRoot":                       _bool,
			"branch.autoSetupMerge":                _bool,
			"branch.autoSetupRebase":               _bool,
			"branch.sort":                          ActionFieldNames(), // TODO verify
			"bundle.version":                       carapace.ActionValues("1"),
			"bundle.mode":                          carapace.ActionValues("all", "any").StyleF(style.ForKeyword),
			"bundle.heuristic":                     carapace.ActionValues("creationToken"),
			"checkout.defaultRemote":               ActionRemotes(),
			"checkout.guess":                       _bool,
			"checkout.workers":                     carapace.ActionValues().Usage("number of parallel workers"),
			"checkout.thresholdForParallelism":     carapace.ActionValues().Usage("minimum number of files to attempt parallel checkout"),
			"clean.requireForce":                   _bool,
			"clone.defaultRemoteName":              ActionRemotes(),
			"clone.rejectShallow":                  _bool,
			"clone.filterSubmodules":               _bool,
			"color.advice":                         ActionColorModes(),
			"color.advice.hint":                    ActionColors(),
			"color.blame.highlightRecent":          ActionColors(),
			"color.blame.repeatedLines":            ActionColors(),
			"color.branch":                         _bool,
			"color.branch.current":                 ActionColors(),
			"color.branch.local":                   ActionColors(),
			"color.branch.plain":                   ActionColors(),
			"color.branch.remote":                  ActionColors(),
			"color.branch.reset":                   ActionColors(),
			"color.branch.upstream":                ActionColors(),
			"color.branch.worktree":                ActionColors(),
			"color.decorate.HEAD":                  ActionColors(),
			"color.decorate.branch":                ActionColors(),
			"color.decorate.grafted":               ActionColors(),
			"color.decorate.remoteBranch":          ActionColors(),
			"color.decorate.stash":                 ActionColors(),
			"color.decorate.tag":                   ActionColors(),
			"color.diff":                           _bool,
			"color.diff.commit":                    ActionColors(),
			"color.diff.context":                   ActionColors(),
			"color.diff.contextBold":               ActionColors(),
			"color.diff.contextDimmed":             ActionColors(),
			"color.diff.frag":                      ActionColors(),
			"color.diff.func":                      ActionColors(),
			"color.diff.meta":                      ActionColors(),
			"color.diff.new":                       ActionColors(),
			"color.diff.newBold":                   ActionColors(),
			"color.diff.newDimmed":                 ActionColors(),
			"color.diff.newMoved":                  ActionColors(),
			"color.diff.newMovedAlternative":       ActionColors(),
			"color.diff.newMovedAlternativeDimmed": ActionColors(),
			"color.diff.newMovedDimmed":            ActionColors(),
			"color.diff.old":                       ActionColors(),
			"color.diff.oldBold":                   ActionColors(),
			"color.diff.oldDimmed":                 ActionColors(),
			"color.diff.oldMoved":                  ActionColors(),
			"color.diff.oldMovedAlternative":       ActionColors(),
			"color.diff.oldMovedAlternativeDimmed": ActionColors(),
			"color.diff.oldMovedDimmed":            ActionColors(),
			"color.diff.plain":                     ActionColors(),
			"color.diff.whitespace":                ActionColors(),
			"color.grep":                           _bool,
			"color.grep.column":                    ActionColors(),
			"color.grep.context":                   ActionColors(),
			"color.grep.filename":                  ActionColors(),
			"color.grep.function":                  ActionColors(),
			"color.grep.lineNumber":                ActionColors(),
			"color.grep.match":                     ActionColors(),
			"color.grep.matchContext":              ActionColors(),
			"color.grep.matchSelected":             ActionColors(),
			"color.grep.selected":                  ActionColors(),
			"color.grep.separator":                 ActionColors(),
			"color.interactive":                    _bool,
			"color.interactive.error":              ActionColors(),
			"color.interactive.header":             ActionColors(),
			"color.interactive.help":               ActionColors(),
			"color.interactive.plain":              ActionColors(),
			"color.interactive.prompt":             ActionColors(),
			"color.interactive.reset":              ActionColors(),
			"color.pager":                          _bool,
			"color.push":                           _bool,
			"color.push.error":                     ActionColors(),
			"color.remote":                         _bool,
			"color.remote.error":                   ActionColors(),
			"color.remote.hint":                    ActionColors(),
			"color.remote.success":                 ActionColors(),
			"color.remote.warning":                 ActionColors(),
			"color.showBranch":                     _bool,
			"color.status":                         _bool,
			"color.status.added":                   ActionColors(),
			"color.status.branch":                  ActionColors(),
			"color.status.changed":                 ActionColors(),
			"color.status.header":                  ActionColors(),
			"color.status.localBranch":             ActionColors(),
			"color.status.noBranch":                ActionColors(),
			"color.status.remoteBranch":            ActionColors(),
			"color.status.unmerged":                ActionColors(),
			"color.status.untracked":               ActionColors(),
			"color.status.updated":                 ActionColors(),
			"color.transport":                      _bool,
			"color.transport.rejected":             ActionColors(),
			"color.ui":                             _bool,
			"column.ui": carapace.ActionValuesDescribed(
				"always", "always show in columns",
				"auto", "show in columns if the output is to the terminal",
				"column", "fill columns before rows",
				"dense", "make unequal size columns to utilize more space",
				"never", "never show in columns",
				"nodense", "make equal size columns",
				"plain", "show in one column",
				"row", "fill rows before columns",
			).StyleF(style.ForKeyword).UniqueList(","),
			"column.branch": _bool,
		}[config]); ok {
			return a
		}

		splitted := strings.Split(config, ".")
		switch splitted[0] {
		case "alias":
			return bridge.ActionCarapaceBin("git").Split()
		case "branch":
			switch len(splitted) {
			case 3:
				return carapace.ActionMap{
					"remote":     ActionRemotes(),
					"pushRemote": ActionRemotes(),
					// TODO merge - use remote configured for given branch
					"mergeOptions": bridge.ActionCarapaceBin("git", "merge").Split(),
					"rebase":       _bool,
				}[splitted[2]]
			}
		case "browser":
			switch len(splitted) {
			case 3:
				return carapace.ActionMap{
					"cmd": bridge.ActionCarapaceBin().SplitP(),
					// TODO path
				}[splitted[2]]
			}
		}
		return carapace.ActionValues()
	})
}
