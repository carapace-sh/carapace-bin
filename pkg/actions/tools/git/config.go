package git

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/carapace-sh/carapace-bin/pkg/actions/number"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/text"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/style"
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

// ActionTextAttributes completes text attributes
func ActionTextAttributes() carapace.Action {
	return carapace.ActionStyledValues(
		"bold", style.Bold,
		"dim", style.Dim,
		"ul", style.Underlined,
		"blink", style.Blink,
		"reverse", style.Default,
	)
}

// ActionColorConfigs completes color configs
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
		// TODO complete more configurations (see https://git-scm.com/docs/git-config)
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
			"column.ui":                            ActionColumnLayoutModes().UniqueList(","),
			"column.branch":                        _bool,
			"commit.cleanup":                       ActionCleanupModes(),
			"commit.gpgSign":                       _bool,
			"commit.status":                        _bool,
			"commit.template":                      carapace.ActionFiles(),
			"commit.verbose":                       _bool,
			"commitGraph.generationVersion":        carapace.ActionValues("1", "2"),
			"commitGraph.maxNewFilters":            carapace.ActionValues().Usage("generate at most n new Bloom filters"),
			"commitGraph.readChangedPaths":         _bool,
			"core.abbrev":                          carapace.ActionValues("auto", "no").StyleF(style.ForKeyword),
			"core.alternateRefsCommand":            bridge.ActionCarapaceBin().SplitP(),
			"core.askPass":                         _bool,
			"core.attributesFile":                  carapace.ActionFiles(),
			"core.autocrlf":                        _bool,
			"core.bare":                            _bool,

			"core.bigFileThreshold":       carapace.ActionValues().Usage("the size of files considered \"big\""),
			"core.checkRoundtripEncoding": text.ActionEncodings().UniqueList(","),
			"core.checkStat":              carapace.ActionValues("default", "minimal"),
			"core.commentChar":            carapace.ActionValues().Usage("consider a line that begins with this character commented"),
			"core.commitGraph":            _bool,
			"core.compression":            number.ActionRange(number.RangeOpts{Start: 1, End: 9}),
			"core.createObject":           carapace.ActionValues("link", "rename"),
			"core.deltaBaseCacheLimit":    carapace.ActionValues().Usage("Maximum number of bytes per thread to reserve for caching base objects that may be referenced by multiple deltified objects"),
			"core.editor":                 bridge.ActionCarapaceBin().Split(),
			"core.eol":                    carapace.ActionValues("lf", "crlf", "native"),
			"core.excludesFile":           carapace.ActionFiles(),
			"core.fileMode":               _bool,
			"core.filesRefLockTimeout": carapace.ActionValuesDescribed(
				"0", "no retry at all",
				"-1", "try indefinitely",
			).Usage("The length of time, in milliseconds, to retry when trying to lock an individual reference"),
			"core.fsmonitor":            _bool,
			"core.fsmonitorHookVersion": carapace.ActionValues("1", "2"),
			"core.fsync":                carapace.ActionValues("committed", "added", "all"),
			"core.fsyncMethod": carapace.ActionValuesDescribed(
				"fsync", "uses the fsync() system call or platform equivalents",
				"writeout-only", "issues pagecache writeback requests",
				"batch", "enables a mode that uses writeout-only flushes to stage multiple updates in the disk writeback cache",
			),
			"core.fsyncObjectFiles": _bool,
			"core.gitProxy":         bridge.ActionCarapaceBin().Split(),
			"core.hideDotFiles":     _bool,
			"core.hooksPath":        carapace.ActionDirectories(),
			"core.ignoreCase":       _bool,
			"core.ignoreStat":       _bool,
			"core.logAllRefUpdates": _bool,
			"core.looseCompression": number.ActionRange(number.RangeOpts{Start: -1, End: 9}),
			"core.multiPackIndex":   _bool,
			"core.notesRef":         ActionRefs(RefOption{}.Default()),
			// "core.packedGitLimit":
			// "core.packedGitWindowSize":
			// "core.packedRefsTimeout":
			"core.pager":             bridge.ActionCarapaceBin().Split(),
			"core.precomposeUnicode": _bool,
			"core.preferSymlinkRefs": _bool,
			"core.preloadIndex":      _bool,
			"core.protectHFS":        _bool,
			"core.protectNTFS":       _bool,
			"core.quotePath":         _bool,
			// "core.repositoryFormatVersion":
			"core.restrictinheritedhandles": carapace.ActionValues("auto", "true", "false").StyleF(style.ForKeyword),
			"core.safecrlf":                 _bool,
			"core.sharedRepository": carapace.Batch(
				carapace.ActionValuesDescribed(
					"group", "shareable between several users in a group",
					"true", "shareable between several users in a group",
					"all", "readable by all users",
					"world", "readable by all users",
					"everybody", "readable by all users",
					"umask", "use permissions reported by umask",
					"false", "not shared",
				),
				carapace.ActionCallback(func(c carapace.Context) carapace.Action {
					if strings.HasPrefix(c.Value, "0") {
						return fs.ActionFileModesNumeric().Prefix("0")
					}
					return carapace.ActionValues("0").NoSpace('0')
				}),
			).ToA(),
			"core.sparseCheckout":     _bool,
			"core.sparseCheckoutCone": _bool,
			"core.splitIndex":         _bool,
			"core.sshCommand":         bridge.ActionCarapaceBin().Split(),
			"core.symlinks":           _bool,
			"core.trustctime":         _bool,
			"core.unsetenvvars":       os.ActionEnvironmentVariables().UniqueList(","),
			"core.untrackedCache":     _bool,
			"core.useReplaceRefs":     _bool,
			"core.warnAmbiguousRefs":  _bool,
			"core.whitespace":         ActionWhitespaceProblems().UniqueList(","), // TODO this is not correct yet, these can be prefixed with `-` and tabwith accepts a value delimited by `=`
			"core.worktree":           carapace.ActionDirectories(),
			"credential.helper":       bridge.ActionCarapaceBin().Split(),
			"credential.useHttpPath":  _bool,
			"diff.algorithm":          ActionDiffAlgorithms(),
			"diff.autoRefreshIndex":   _bool,
			"diff.colorMoved":         ActionColorMovedModes(),
			"diff.colorMovedWS":       ActionColorMovedWsModes(),
			// "diff.context":
			"diff.dirstat":  ActionDirstats().UniqueList(","),
			"diff.external": bridge.ActionCarapaceBin().Split(),
			// "diff.guitool":
			"diff.ignoreSubmodules": carapace.ActionValues("none", "untracked", "dirty", "all"), // TODO custom action
			"diff.indentHeuristic":  _bool,
			"diff.noprefix":         _bool,
			"diff.orderFile":        carapace.ActionFiles(),
			"diff.relative":         _bool,
			"diff.renames":          carapace.ActionValues("true", "false", "copy").StyleF(style.ForKeyword),
			"diff.submodule": carapace.ActionValuesDescribed(
				"short", "format just shows the names of the commits at the beginning and end of the range",
				"log", "format lists the commits in the range like git-submodule summary does",
				"diff", "format shows an inline diff of the changed contents of the submodule",
			),
			"diff.suppressBlankEmpty":   _bool,
			"diff.tool":                 ActionDiffTools(),
			"diff.wsErrorHighlight":     ActionWsErrorHighlightModes().UniqueList(","),
			"extensions.objectFormat":   carapace.ActionValues("sha1", "sha256"),
			"extensions.worktreeConfig": _bool,
			"feature.experimental":      _bool,
			"feature.manyFiles":         _bool,
			// "fetch.fsck.skipList":
			"fetch.fsckObjects": _bool,
			"fetch.negotiationAlgorithm": carapace.ActionValuesDescribed(
				"consecutive", "walk over consecutive commits checking each one",
				"skipping", "skip commits in an effort to converge faster",
				"noop", "do not send any information at all",
				"default", "override settings made previously and use the default behaviour",
			),
			"fetch.output":            carapace.ActionValues("full", "compact"),
			"fetch.prune":             _bool,
			"fetch.pruneTags":         _bool,
			"fetch.recurseSubmodules": carapace.ActionValues("true", "false", "ondemand").StyleF(style.ForKeyword),
			"fetch.showForcedUpdates": _bool,
			"fetch.writeCommitGraph":  _bool,
			"rerere.enabled":          _bool,
			"remote.pushDefault":      ActionRemotes(),
		}[config]); ok {
			return a
		}

		splitted := strings.Split(config, ".")
		switch splitted[0] {
		case "alias":
			return carapace.Batch(
				bridge.ActionCarapaceBin().Split().Prefix("!").UnlessF(func(c carapace.Context) bool {
					return !strings.HasPrefix(c.Value, "!")
				}),
				bridge.ActionCarapaceBin("git").Split(),
			).ToA()
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
				}[splitted[2]]
			}
		case "diff":
			switch len(splitted) {
			case 3:
				return carapace.ActionMap{
					"command":       bridge.ActionCarapaceBin().SplitP(),
					"binary":        _bool,
					"textconv":      bridge.ActionCarapaceBin().SplitP(),
					"cachetextconv": _bool,
				}[splitted[2]]
			}
		case "remote":
			switch len(splitted) {
			case 3:
				return carapace.ActionMap{
					// "fetch":
					"mirror": _bool,
					// "partialclonefilter":
					"promisor": _bool,
					// "proxy":
					"proxyAuthMethod": carapace.ActionValuesDescribed(
						"anyauth", "Automatically pick a suitable authentication method",
						"basic", "HTTP Basic authentication",
						"digest", "HTTP Digest authentication",
						"negotiate", "GSS-Negotiate authentication",
						"ntlm", "NTLM authentication",
					),
					"prune":     _bool,
					"pruneTags": _bool,
					// "push":
					// "pushurl":
					"receivepack":       bridge.ActionCarapaceBin().Split(),
					"skipDefaultUpdate": _bool,
					"skipFetchAll":      _bool,
					// "tagOpt":
					"uploadpack": bridge.ActionCarapaceBin().Split(),
					// "url":
					// "vcs":
				}[splitted[2]]
			}
		}
		return carapace.ActionValues()
	})
}
