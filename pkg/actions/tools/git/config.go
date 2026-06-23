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
//
//	core.autocrlf
//	core.editor
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
	}).UidF(Uid("config"))
}

// ActionDateFormats completes date formats
//
//	relative
//	iso-strict
func ActionDateFormats() carapace.Action {
	return carapace.ActionValuesDescribed(
		"default", "shows the date in the original timezone",
		"relative", "shows the date relative to the current time",
		"local", "shows the date in the local timezone",
		"iso", "shows the date in ISO 8601 format",
		"iso-strict", "shows the date in ISO 8601 strict format",
		"rfc", "shows the date in RFC 2822 format",
		"short", "shows the date as YYYY-MM-DD",
		"raw", "shows the date as seconds since the epoch + timezone",
		"human", "shows the date in a human-readable format",
		"unix", "shows the date as a Unix epoch timestamp (seconds since 1970)",
		"format:", "shows the date with a custom strftime format",
	).Uid("git", "date-format")
}

// ActionColors completes colors
//
//	red
//	green
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
	).Uid("git", "color")
}

// ActionTextAttributes completes text attributes
//
//	bold
//	dim
func ActionTextAttributes() carapace.Action {
	return carapace.ActionStyledValues(
		"bold", style.Bold,
		"dim", style.Dim,
		"ul", style.Underlined,
		"blink", style.Blink,
		"reverse", style.Default,
	).Uid("git", "text-attribute")
}

// ActionColorConfigs completes color configs
//
//	red bold
//	green dim
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
//
//	bool
//	string
func ActionConfigTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"bool", "canonicalize values as either \"true\" or \"false\"",
		"int", "canonicalize values as simple decimal numbers",
		"bool-or-int", "canonicalize according to either bool or int",
		"path", "canonicalize by adding a leading ~ to the value of $HOME",
		"expiry-date", "canonicalize by converting from a fixed or relative date-string to a timestamp",
		"color", "When getting a value, canonicalize by converting to an ANSI color escape sequence",
	).Uid("git", "config-type")
}

// ActionConfigTypeOptions completes options for a config type
//
//	true
//	false
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
//
//	auto
//	true
func ActionConfigValues(config string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO complete more configurations (see https://git-scm.com/docs/git-config)
		_bool := carapace.ActionValues("true", "false").StyleF(style.ForKeyword)
		if a, ok := (carapace.ActionMap{
			"add.ignoreErrors":                          _bool,
			"add.interactive.useBuiltin":                _bool,
			"advice.addEmbeddedRepo":                    _bool,
			"advice.addEmptyPathspec":                   _bool,
			"advice.addIgnoredFile":                     _bool,
			"advice.amWorkDir":                          _bool,
			"advice.ambiguousFetchRefspec":              _bool,
			"advice.checkoutAmbiguousRemoteBranchName":  _bool,
			"advice.commitBeforeMerge":                  _bool,
			"advice.defaultBranchName":                  _bool,
			"advice.detachedHead":                       _bool,
			"advice.diverging":                          _bool,
			"advice.fetchRemoteHEADWarn":                _bool,
			"advice.fetchShowForcedUpdates":             _bool,
			"advice.forceDeleteBranch":                  _bool,
			"advice.graftFileDeprecated":                _bool,
			"advice.ignoredHook":                        _bool,
			"advice.implicitIdentity":                   _bool,
			"advice.mergeConflict":                      _bool,
			"advice.nestedTag":                          _bool,
			"advice.objectNameWarning":                  _bool,
			"advice.pushAlreadyExists":                  _bool,
			"advice.pushFetchFirst":                     _bool,
			"advice.pushNeedsForce":                     _bool,
			"advice.pushNonFFCurrent":                   _bool,
			"advice.pushNonFFMatching":                  _bool,
			"advice.pushNonFastForward":                 _bool,
			"advice.pushRefNeedsUpdate":                 _bool,
			"advice.pushUnqualifiedRefName":             _bool,
			"advice.pushUpdateRejected":                 _bool,
			"advice.rebaseTodoError":                    _bool,
			"advice.refSyntax":                          _bool,
			"advice.resetNoRefresh":                     _bool,
			"advice.resolveConflict":                    _bool,
			"advice.rmHints":                            _bool,
			"advice.sequencerInUse":                     _bool,
			"advice.setUpstreamFailure":                 _bool,
			"advice.skippedCherryPicks":                 _bool,
			"advice.sparseIndexExpanded":                _bool,
			"advice.statusAheadBehindWarning":           _bool,
			"advice.statusHints":                        _bool,
			"advice.statusUoption":                      _bool,
			"advice.submoduleAlternateErrorStrategyDie": _bool,
			"advice.submoduleMergeConflict":             _bool,
			"advice.submodulesNotUpdated":               _bool,
			"advice.suggestDetachingHead":               _bool,
			"advice.updateSparsePath":                   _bool,
			"advice.waitingForEditor":                   _bool,
			"advice.worktreeAddOrphan":                  _bool,
			"am.keepcr":                                 _bool,
			"am.messageId":                              _bool,
			"am.threeWay":                               _bool,
			"apply.ignoreWhitespace":                    carapace.ActionValues("all", "change", "no", "none", "never", "false", "true").StyleF(style.ForKeyword),
			"apply.whitespace":                          carapace.ActionValues("warn", "error", "error-all", "strip", "fix").StyleF(style.ForKeyword),
			"attr.tree":                                 carapace.ActionValues("ref"),
			"author.email":                              ActionAuthors().NoSpace(),
			"author.name":                               ActionAuthors(),
			"blame.blankBoundary":                       _bool,
			"blame.coloring":                            carapace.ActionValues("repeatedLines", "highlightRecent", "none").StyleF(style.ForKeyword),
			"blame.date":                                ActionDateFormats(),
			"blame.ignoreRevsFile":                      carapace.ActionFiles(),
			"blame.markIgnoredLines":                    _bool,
			"blame.markUnblamableLines":                 _bool,
			"blame.showEmail":                           _bool,
			"blame.showRoot":                            _bool,
			"branch.autoSetupMerge":                     _bool,
			"branch.autoSetupRebase":                    _bool,
			"branch.sort":                               ActionFieldNames(), // TODO verify
			"bundle.version":                            carapace.ActionValues("1", "2"),
			"bundle.mode":                               carapace.ActionValues("all", "any").StyleF(style.ForKeyword),
			"bundle.heuristic":                          carapace.ActionValues("creationToken"),
			"checkout.defaultRemote":                    ActionRemotes(),
			"checkout.guess":                            _bool,
			"checkout.workers":                          carapace.ActionValues().Usage("number of parallel workers"),
			"checkout.thresholdForParallelism":          carapace.ActionValues().Usage("minimum number of files to attempt parallel checkout"),
			"clean.requireForce":                        _bool,
			"clone.defaultRemoteName":                   carapace.ActionValues().Usage("default remote name for new clones (e.g. origin)"),
			"clone.rejectShallow":                       _bool,
			"clone.filterSubmodules":                    _bool,
			"color.advice":                              ActionColorModes(),
			"color.advice.hint":                         ActionColors(),
			"color.blame.highlightRecent":               ActionColors(),
			"color.blame.repeatedLines":                 ActionColors(),
			"color.branch":                              _bool,
			"color.branch.current":                      ActionColors(),
			"color.branch.local":                        ActionColors(),
			"color.branch.plain":                        ActionColors(),
			"color.branch.remote":                       ActionColors(),
			"color.branch.reset":                        ActionColors(),
			"color.branch.upstream":                     ActionColors(),
			"color.branch.worktree":                     ActionColors(),
			"color.decorate.HEAD":                       ActionColors(),
			"color.decorate.branch":                     ActionColors(),
			"color.decorate.grafted":                    ActionColors(),
			"color.decorate.remoteBranch":               ActionColors(),
			"color.decorate.stash":                      ActionColors(),
			"color.decorate.tag":                        ActionColors(),
			"color.diff":                                _bool,
			"color.diff.commit":                         ActionColors(),
			"color.diff.context":                        ActionColors(),
			"color.diff.contextBold":                    ActionColors(),
			"color.diff.contextDimmed":                  ActionColors(),
			"color.diff.frag":                           ActionColors(),
			"color.diff.func":                           ActionColors(),
			"color.diff.meta":                           ActionColors(),
			"color.diff.new":                            ActionColors(),
			"color.diff.newBold":                        ActionColors(),
			"color.diff.newDimmed":                      ActionColors(),
			"color.diff.newMoved":                       ActionColors(),
			"color.diff.newMovedAlternative":            ActionColors(),
			"color.diff.newMovedAlternativeDimmed":      ActionColors(),
			"color.diff.newMovedDimmed":                 ActionColors(),
			"color.diff.old":                            ActionColors(),
			"color.diff.oldBold":                        ActionColors(),
			"color.diff.oldDimmed":                      ActionColors(),
			"color.diff.oldMoved":                       ActionColors(),
			"color.diff.oldMovedAlternative":            ActionColors(),
			"color.diff.oldMovedAlternativeDimmed":      ActionColors(),
			"color.diff.oldMovedDimmed":                 ActionColors(),
			"color.diff.plain":                          ActionColors(),
			"color.diff.whitespace":                     ActionColors(),
			"color.grep":                                _bool,
			"color.grep.column":                         ActionColors(),
			"color.grep.context":                        ActionColors(),
			"color.grep.filename":                       ActionColors(),
			"color.grep.function":                       ActionColors(),
			"color.grep.lineNumber":                     ActionColors(),
			"color.grep.match":                          ActionColors(),
			"color.grep.matchContext":                   ActionColors(),
			"color.grep.matchSelected":                  ActionColors(),
			"color.grep.selected":                       ActionColors(),
			"color.grep.separator":                      ActionColors(),
			"color.interactive":                         _bool,
			"color.interactive.error":                   ActionColors(),
			"color.interactive.header":                  ActionColors(),
			"color.interactive.help":                    ActionColors(),
			"color.interactive.plain":                   ActionColors(),
			"color.interactive.prompt":                  ActionColors(),
			"color.interactive.reset":                   ActionColors(),
			"color.pager":                               _bool,
			"color.push":                                _bool,
			"color.push.error":                          ActionColors(),
			"color.remote":                              _bool,
			"color.remote.error":                        ActionColors(),
			"color.remote.hint":                         ActionColors(),
			"color.remote.success":                      ActionColors(),
			"color.remote.warning":                      ActionColors(),
			"color.showBranch":                          _bool,
			"color.status":                              _bool,
			"color.status.added":                        ActionColors(),
			"color.status.branch":                       ActionColors(),
			"color.status.changed":                      ActionColors(),
			"color.status.header":                       ActionColors(),
			"color.status.localBranch":                  ActionColors(),
			"color.status.noBranch":                     ActionColors(),
			"color.status.remoteBranch":                 ActionColors(),
			"color.status.unmerged":                     ActionColors(),
			"color.status.untracked":                    ActionColors(),
			"color.status.updated":                      ActionColors(),
			"color.transport":                           _bool,
			"color.transport.rejected":                  ActionColors(),
			"color.ui":                                  ActionColorModes(),
			"column.branch":                             carapace.ActionValues("column", "row", "plain", "dense").StyleF(style.ForKeyword),
			"column.clean":                              ActionColumnLayoutModes().UniqueList(","),
			"column.status":                             ActionColumnLayoutModes().UniqueList(","),
			"column.tag":                                ActionColumnLayoutModes().UniqueList(","),
			"column.ui":                                 ActionColumnLayoutModes().UniqueList(","),
			"commit.cleanup":                            ActionCleanupModes(),
			"commit.gpgSign":                            _bool,
			"commit.status":                             _bool,
			"commit.template":                           carapace.ActionFiles(),
			"commit.verbose":                            _bool,
			"commitGraph.changedPaths":                  _bool,
			"commitGraph.changedPathsVersion":           carapace.ActionValues("1", "2"),
			"commitGraph.generationVersion":             carapace.ActionValues("1", "2"),
			"commitGraph.maxNewFilters":                 carapace.ActionValues().Usage("generate at most n new Bloom filters"),
			"commitGraph.readChangedPaths":              _bool,
			"committer.email":                           ActionCommitters().NoSpace(),
			"committer.name":                            ActionCommitters(),
			"completion.commands":                       carapace.ActionValues().Usage("control which commands complete differently when completing with carapace"),
			"core.abbrev":                               carapace.ActionValues("auto", "no").StyleF(style.ForKeyword),
			"core.alternateRefsCommand":                 bridge.ActionCarapaceBin().SplitP(),
			"core.alternateRefsPrefixes":                carapace.ActionValues().Usage("prefixes to use when listing alternate refs"),
			"core.askPass":                              bridge.ActionCarapaceBin().Split(),
			"core.attributesFile":                       carapace.ActionFiles(),
			"core.autocrlf":                             _bool,
			"core.bare":                                 _bool,

			"core.bigFileThreshold":       carapace.ActionValues().Usage("the size of files considered \"big\""),
			"core.checkRoundtripEncoding": text.ActionEncodings().UniqueList(","),
			"core.checkStat":              carapace.ActionValues("default", "minimal"),
			"core.commentChar":            carapace.ActionValues().Usage("consider a line that begins with this character commented"),
			"core.commentString":          carapace.ActionValues().Usage("consider lines that begin with this string commented"),
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
			"core.hideDotFiles":     carapace.ActionValues("true", "false", "dotGitOnly").StyleF(style.ForKeyword),
			"core.hooksPath":        carapace.ActionDirectories(),
			"core.ignoreCase":       _bool,
			"core.ignoreStat":       _bool,
			"core.lockfilePid":      _bool,
			"core.logAllRefUpdates": _bool,
			"core.looseCompression": number.ActionRange(number.RangeOpts{Start: -1, End: 9}),
			"core.maxTreeDepth":     carapace.ActionValues().Usage("maximum tree depth"),
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
			"core.safecrlf":                 carapace.ActionValues("true", "false", "warn").StyleF(style.ForKeyword),
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
			"core.whitespace":         ActionWhitespaceProblems().UniqueList(","),
			"core.worktree":           carapace.ActionDirectories(),
			"credential.helper":       bridge.ActionCarapaceBin().Split(),
			"credential.interactive": carapace.ActionValuesDescribed(
				"auto", "allow automatic prompting",
				"never", "never prompt",
				"always", "always prompt",
			).StyleF(style.ForKeyword),
			"credential.protectProtocol":    _bool,
			"credential.sanitizePrompt":     _bool,
			"credential.useHttpPath":        _bool,
			"credential.username":           os.ActionUsers(),
			"credentialCache.ignoreSIGHUP":  _bool,
			"credentialStore.lockTimeoutMS": carapace.ActionValues().Usage("lock timeout in milliseconds"),
			"diff.algorithm":                ActionDiffAlgorithms(),
			"diff.autoRefreshIndex":         _bool,
			"diff.colorMoved":               ActionColorMovedModes(),
			"diff.colorMovedWS":             ActionColorMovedWsModes(),
			"diff.context":                  carapace.ActionValues().Usage("number of context lines"),
			"diff.dirstat":                  ActionDirstats().UniqueList(","),
			"diff.dstPrefix":                carapace.ActionValues().Usage("destination prefix (default: b/)"),
			"diff.external":                 bridge.ActionCarapaceBin().Split(),
			"diff.guitool":                  ActionDiffTools(),
			"diff.ignoreSubmodules":         carapace.ActionValues("none", "untracked", "dirty", "all"),
			"diff.indentHeuristic":          _bool,
			"diff.interHunkContext":         carapace.ActionValues().Usage("number of extra context lines between hunks"),
			"diff.mnemonicPrefix":           _bool,
			"diff.noPrefix":                 _bool,
			"diff.noprefix":                 _bool,
			"diff.orderFile":                carapace.ActionFiles(),
			"diff.relative":                 _bool,
			"diff.renameLimit":              carapace.ActionValues().Usage("number of files to consider when performing copy/rename detection"),
			"diff.renames":                  carapace.ActionValues("true", "false", "copy").StyleF(style.ForKeyword),
			"diff.srcPrefix":                carapace.ActionValues().Usage("source prefix (default: a/)"),
			"diff.statGraphWidth":           carapace.ActionValues().Usage("width of the graph part in --stat output"),
			"diff.statNameWidth":            carapace.ActionValues().Usage("width of the filename part in --stat output"),
			"diff.submodule": carapace.ActionValuesDescribed(
				"short", "format just shows the names of the commits at the beginning and end of the range",
				"log", "format lists the commits in the range like git-submodule summary does",
				"diff", "format shows an inline diff of the changed contents of the submodule",
			),
			"diff.suppressBlankEmpty":   _bool,
			"diff.tool":                 ActionDiffTools(),
			"diff.trustExitCode":        _bool,
			"diff.wordRegex":            carapace.ActionValues().Usage("regex to split words for word diff"),
			"diff.wsErrorHighlight":     ActionWsErrorHighlightModes().UniqueList(","),
			"difftool.prompt":           _bool,
			"difftool.guiDefault":       carapace.ActionValues().Usage("default guitool to use"),
			"difftool.trustExitCode":    _bool,
			"extensions.objectFormat":   carapace.ActionValues("sha1", "sha256"),
			"extensions.worktreeConfig": _bool,
			"fastimport.unpackLimit":    carapace.ActionValues().Usage("unpack limit for fast-import"),
			"feature.experimental":      _bool,
			"feature.manyFiles":         _bool,
			"fetch.all":                 _bool,
			"fetch.fsckObjects":         _bool,
			"fetch.negotiationAlgorithm": carapace.ActionValuesDescribed(
				"consecutive", "walk over consecutive commits checking each one",
				"skipping", "skip commits in an effort to converge faster",
				"noop", "do not send any information at all",
				"default", "override settings made previously and use the default behaviour",
			),
			"fetch.output":                carapace.ActionValues("full", "compact"),
			"fetch.parallel":              carapace.ActionValues().Usage("number of parallel children to be used for fetching refs"),
			"fetch.prune":                 _bool,
			"fetch.pruneTags":             _bool,
			"fetch.recurseSubmodules":     carapace.ActionValues("true", "false", "ondemand").StyleF(style.ForKeyword),
			"fetch.showForcedUpdates":     _bool,
			"fetch.unpackLimit":           carapace.ActionValues().Usage("unpack limit for fetch"),
			"fetch.writeCommitGraph":      _bool,
			"format.attach":               _bool,
			"format.cc":                   carapace.ActionValues().Usage("email addresses for carbon copy"),
			"format.coverFromDescription": carapace.ActionValues("default", "message", "subject", "auto"),
			"format.coverLetter":          carapace.ActionValues("auto", "none", "email"),
			"format.encodeEmailHeaders":   _bool,
			"format.forceInBodyFrom":      _bool,
			"format.from":                 carapace.ActionValues().Usage("default From: address"),
			"format.mboxrd":               _bool,
			"format.noprefix":             _bool,
			"format.numbered":             carapace.ActionValues("auto", "no"),
			"format.outputDirectory":      carapace.ActionDirectories(),
			"format.pretty":               ActionPrettyFormats(),
			"format.signOff":              _bool,
			"format.signature":            carapace.ActionValues().Usage("email signature"),
			"format.signatureFile":        carapace.ActionFiles(),
			"format.subjectPrefix":        carapace.ActionValues().Usage("subject prefix for patches"),
			"format.suffix":               carapace.ActionValues(".patch", ".diff"),
			"format.thread":               carapace.ActionValues("deep", "shallow"),
			"format.to":                   carapace.ActionValues().Usage("email addresses for primary recipients"),
			"format.useAutoBase":          _bool,
			"fsmonitor.allowRemote":       _bool,
			"fsmonitor.socketDir":         carapace.ActionDirectories(),
			"gc.aggressiveDepth":          carapace.ActionValues().Usage("depth to use in git gc --aggressive"),
			"gc.aggressiveWindow":         carapace.ActionValues().Usage("window to use in git gc --aggressive"),
			"gc.auto":                     carapace.ActionValues().Usage("when there are more than this many loose objects, auto gc"),
			"gc.autoDetach":               _bool,
			"gc.autoPackLimit":            carapace.ActionValues().Usage("when there are more than this many packs, auto gc"),
			"gc.bigPackThreshold":         carapace.ActionValues().Usage("threshold for big packs"),
			"gc.cruftPacks":               _bool,
			"gc.logExpiry":                carapace.ActionValues().Usage("expiry time for gc.log"),
			"gc.maxCruftSize":             carapace.ActionValues().Usage("maximum cruft pack size"),
			"gc.packRefs":                 carapace.ActionValues("true", "false", "notbare"),
			"gc.pruneExpire":              carapace.ActionValues().Usage("expiry time for unreachable loose objects"),
			"gc.reflogExpire":             carapace.ActionValues().Usage("expiry time for reflog entries"),
			"gc.reflogExpireUnreachable":  carapace.ActionValues().Usage("expiry time for unreachable reflog entries"),
			"gc.rerereResolved":           carapace.ActionValues().Usage("retention time for resolved rerere records"),
			"gc.rerereUnresolved":         carapace.ActionValues().Usage("retention time for unresolved rerere records"),
			"gc.worktreePruneExpire":      carapace.ActionValues().Usage("expiry time for pruned worktrees"),
			"gc.writeCommitGraph":         _bool,
			"gpg.format":                  carapace.ActionValues("openpgp", "ssh", "x509"),
			"gpg.minTrustLevel":           carapace.ActionValues("undefined", "never", "marginal", "fully", "ultimate"),
			"gpg.program":                 bridge.ActionCarapaceBin().Split(),
			"gpg.ssh.allowedSignersFile":  carapace.ActionFiles(),
			"gpg.ssh.defaultKeyCommand":   carapace.ActionValues().Usage("command to get default SSH signing key"),
			"gpg.ssh.revocationFile":      carapace.ActionFiles(),
			"grep.column":                 _bool,
			"grep.extendedRegexp":         _bool,
			"grep.fallbackToNoIndex":      _bool,
			"grep.fullName":               _bool,
			"grep.lineNumber":             _bool,
			"grep.patternType":            carapace.ActionValues("basic", "extended", "fixed", "perl"),
			"grep.threads":                carapace.ActionValues().Usage("number of grep worker threads"),
			"gui.GCWarning":               _bool,
			"gui.blamehistoryctx":         carapace.ActionValues().Usage("blame history context radius"),
			"gui.commitMsgWidth":          carapace.ActionValues().Usage("width of the commit message"),
			"gui.diffContext":             carapace.ActionValues().Usage("number of context lines in diff"),
			"gui.displayUntracked":        _bool,
			"gui.encoding":                text.ActionEncodings(),
			"gui.matchTrackingBranch":     carapace.ActionValues().Usage("tracking branch name pattern"),
			"gui.newBranchTemplate":       carapace.ActionValues().Usage("new branch name template"),
			"gui.pruneDuringFetch":        _bool,
			"gui.spellingDictionary":      carapace.ActionValues().Usage("spelling dictionary"),
			"gui.trustmtime":              _bool,
			"help.autoCorrect":            carapace.ActionValues("immediate", "prompt", "never"),
			"help.browser":                bridge.ActionCarapaceBin().Split(),
			"help.format":                 carapace.ActionValues("man", "info", "web", "html"),
			"help.htmlPath":               carapace.ActionDirectories(),
			"http.cookieFile":             carapace.ActionFiles(),
			"http.curloptResolve":         carapace.ActionValues().Usage("host:port:address"),
			"http.delegation":             carapace.ActionValues("none", "always", "ifavailable"),
			"http.emptyAuth":              _bool,
			"http.extraHeader":            carapace.ActionValues().Usage("additional HTTP headers").UniqueList("\n"),
			"http.followRedirects":        _bool,
			"http.lowSpeedLimit":          carapace.ActionValues().Usage("low speed limit in bytes per second"),
			"http.lowSpeedTime":           carapace.ActionValues().Usage("low speed time in seconds"),
			"http.maxRequests":            carapace.ActionValues().Usage("max number of HTTP requests"),
			"http.minSessions":            carapace.ActionValues().Usage("minimum number of curl sessions"),
			"http.noEPSV":                 _bool,
			"http.pinnedPubkey":           carapace.ActionValues().Usage("pinned public key"),
			"http.postBuffer":             carapace.ActionValues().Usage("buffer size for smart HTTP POST"),
			"http.proactiveAuth":          carapace.ActionValues("basic", "digest", "negotiate", "ntlm", "none"),
			"http.proxy":                  carapace.ActionValues().Usage("proxy to use for HTTP/HTTPS"),
			"http.proxyAuthMethod": carapace.ActionValuesDescribed(
				"anyauth", "Automatically pick a suitable authentication method",
				"basic", "HTTP Basic authentication",
				"digest", "HTTP Digest authentication",
				"negotiate", "GSS-Negotiate authentication",
				"ntlm", "NTLM authentication",
			),
			"http.saveCookies":              _bool,
			"http.schannelCheckRevoke":      _bool,
			"http.schannelUseSSLCAInfo":     _bool,
			"http.sslBackend":               carapace.ActionValues("openssl", "schannel", "secure-transport", "mbedtls"),
			"http.sslCAInfo":                carapace.ActionFiles(),
			"http.sslCAPath":                carapace.ActionDirectories(),
			"http.sslCert":                  carapace.ActionFiles(),
			"http.sslCertPasswordProtected": _bool,
			"http.sslCertType":              carapace.ActionValues("pem", "der"),
			"http.sslCipherList":            carapace.ActionValues().Usage("list of ciphers"),
			"http.sslKey":                   carapace.ActionFiles(),
			"http.sslKeyType":               carapace.ActionValues("pem", "der"),
			"http.sslTry":                   _bool,
			"http.sslVerify":                _bool,
			"http.sslVersion":               carapace.ActionValues("sslv2", "sslv3", "tlsv1", "tlsv1.1", "tlsv1.2", "tlsv1.3"),
			"http.userAgent":                carapace.ActionValues().Usage("HTTP user agent"),
			"http.version":                  carapace.ActionValues("HTTP/1.0", "HTTP/1.1", "HTTP/2"),
			"i18n.commitEncoding":           text.ActionEncodings(),
			"i18n.logOutputEncoding":        text.ActionEncodings(),
			"imap.authMethod":               carapace.ActionValues("CRAM-MD5", "LOGIN", "PLAIN"),
			"imap.folder":                   carapace.ActionValues().Usage("IMAP folder to put patches"),
			"imap.host":                     carapace.ActionValues().Usage("IMAP server hostname"),
			"imap.pass":                     carapace.ActionValues().Usage("IMAP server password"),
			"imap.port":                     carapace.ActionValues().Usage("IMAP server port"),
			"imap.preformattedHTML":         _bool,
			"imap.sslverify":                _bool,
			"imap.tunnel":                   carapace.ActionValues().Usage("IMAP tunnel command"),
			"imap.user":                     carapace.ActionValues().Usage("IMAP server username"),
			"include.path":                  carapace.ActionFiles(),
			"index.recordEndOfIndexEntries": _bool,
			"index.recordOffsetTable":       _bool,
			"index.skipHash":                _bool,
			"index.sparse":                  _bool,
			"index.threads":                 carapace.ActionValues().Usage("number of threads for index operations"),
			"index.version":                 carapace.ActionValues("2", "3", "4"),
			"init.defaultBranch":            ActionLocalBranches(),
			"init.defaultObjectFormat":      carapace.ActionValues("sha1", "sha256"),
			"init.defaultRefFormat":         carapace.ActionValues("refs", "reftable"),
			"init.templateDir":              carapace.ActionDirectories(),
			"instaweb.browser":              bridge.ActionCarapaceBin().Split(),
			"instaweb.httpd":                carapace.ActionValues().Usage("HTTP daemon command"),
			"instaweb.local":                _bool,
			"instaweb.modulePath":           carapace.ActionDirectories(),
			"instaweb.port":                 carapace.ActionValues().Usage("port for instaweb"),
			"interactive.diffFilter":        bridge.ActionCarapaceBin().Split(),
			"interactive.singleKey":         _bool,
			"log.abbrevCommit":              _bool,
			"log.date":                      ActionDateFormats(),
			"log.decorate":                  carapace.ActionValues("short", "full", "auto", "no"),
			"log.diffMerges":                carapace.ActionValues("first-parent", "separate", "combined", "dense-combined", "m"),
			"log.excludeDecoration":         carapace.ActionValues().Usage("decoration types to exclude"),
			"log.follow":                    _bool,
			"log.graphColors":               carapace.ActionValues().Usage("colors for graph"),
			"log.initialDecorationSet":      carapace.ActionValues("all", "HEAD", "refs/heads", "refs/tags", "refs/remotes"),
			"log.mailmap":                   _bool,
			"log.showRoot":                  _bool,
			"log.showSignature":             _bool,
			"lsrefs.unborn":                 carapace.ActionValues("advertise", "ignore"),
			"mailinfo.scissors":             _bool,
			"mailmap.blob":                  carapace.ActionValues().Usage("blob containing mailmap"),
			"mailmap.file":                  carapace.ActionFiles(),
			"maintenance.auto":              _bool,
			"maintenance.autoDetach":        _bool,
			"maintenance.strategy":          carapace.ActionValues("incremental", "full"),
			"man.viewer":                    carapace.ActionValues().Usage("man viewer to use"),
			"merge.autoStash":               _bool,
			"merge.branchdesc":              _bool,
			"merge.conflictStyle":           carapace.ActionValues("merge", "diff3", "zdiff3"),
			"merge.defaultToUpstream":       _bool,
			"merge.directoryRenames":        carapace.ActionValues("conflict", "true", "false"),
			"merge.ff":                      carapace.ActionValues("true", "false", "only").StyleF(style.ForKeyword),
			"merge.guitool":                 ActionDiffTools(),
			"merge.log":                     _bool,
			"merge.renameLimit":             carapace.ActionValues().Usage("number of files to consider for rename detection"),
			"merge.renames":                 _bool,
			"merge.renormalize":             _bool,
			"merge.stat":                    _bool,
			"merge.suppressDest":            carapace.ActionValues().Usage("destination to suppress merge"),
			"merge.tool":                    ActionMergeTools(),
			"merge.verbosity":               number.ActionRange(number.RangeOpts{Start: 0, End: 5}),
			"merge.verifySignatures":        _bool,
			"mergetool.guiDefault":          carapace.ActionValues().Usage("default guitool to use"),
			"mergetool.hideResolved":        _bool,
			"mergetool.keepBackup":          _bool,
			"mergetool.keepTemporaries":     _bool,
			"mergetool.prompt":              _bool,
			"mergetool.writeToTemp":         _bool,
			"notes.displayRef":              ActionRefs(RefOption{}.Default()),
			"notes.mergeStrategy":           ActionNotesMergeStrategies(),
			"notes.rewriteMode":             carapace.ActionValues("concatenate", "overwrite", "ignore"),
			"notes.rewriteRef":              ActionRefs(RefOption{}.Default()),
			"pack.allowPackReuse":           _bool,
			"pack.compression":              number.ActionRange(number.RangeOpts{Start: 0, End: 9}),
			"pack.deltaCacheLimit":          carapace.ActionValues().Usage("max bytes for delta cache"),
			"pack.deltaCacheSize":           carapace.ActionValues().Usage("max size for delta cache"),
			"pack.depth":                    carapace.ActionValues().Usage("max delta depth"),
			"pack.indexVersion":             carapace.ActionValues("1", "2"),
			"pack.packSizeLimit":            carapace.ActionValues().Usage("max pack size"),
			"pack.threads":                  carapace.ActionValues().Usage("pack threads"),
			"pack.useBitmaps":               _bool,
			"pack.useSparse":                _bool,
			"pack.window":                   carapace.ActionValues().Usage("pack window size"),
			"pack.windowMemory":             carapace.ActionValues().Usage("pack window memory limit"),
			"pack.writeBitmapHashCache":     _bool,
			"pack.writeReverseIndex":        _bool,
			"promisor.acceptFromServer":     _bool,
			"promisor.advertise":            _bool,
			"promisor.quiet":                _bool,
			"protocol.allow":                carapace.ActionValues("always", "never", "user"),
			"protocol.version":              carapace.ActionValues("0", "1", "2"),
			"pull.autoStash":                _bool,
			"pull.ff":                       _bool,
			"pull.rebase":                   carapace.ActionValues("true", "false", "merges", "interactive").StyleF(style.ForKeyword),
			"push.autoSetupRemote":          _bool,
			"push.default": carapace.ActionValuesDescribed(
				"nothing", "do not push anything",
				"current", "push the current branch to a matching branch",
				"upstream", "push the current branch to its upstream",
				"tracking", "deprecated synonym for upstream",
				"simple", "push the current branch with the same name and upstream",
				"matching", "push all branches that have a matching branch elsewhere",
			),
			"push.followTags":                     _bool,
			"push.gpgSign":                        carapace.ActionValues("if-asked", "never", "true", "false").StyleF(style.ForKeyword),
			"push.negotiate":                      _bool,
			"push.pushOption":                     carapace.ActionValues().Usage("push options").UniqueList("\n"),
			"push.recurseSubmodules":              carapace.ActionValues("check", "on-demand", "only", "no"),
			"push.useBitmaps":                     _bool,
			"push.useForceIfIncludes":             _bool,
			"rebase.abbreviateCommands":           _bool,
			"rebase.autoSquash":                   _bool,
			"rebase.autoStash":                    _bool,
			"rebase.backend":                      carapace.ActionValues("apply", "merge"),
			"rebase.forkPoint":                    _bool,
			"rebase.instructionFormat":            carapace.ActionValues().Usage("format for rebase instruction"),
			"rebase.missingCommitsCheck":          carapace.ActionValues("warn", "error", "ignore"),
			"rebase.rebaseMerges":                 _bool,
			"rebase.rescheduleFailedExec":         _bool,
			"rebase.stat":                         _bool,
			"rebase.updateRefs":                   _bool,
			"receive.advertiseAtomic":             _bool,
			"receive.advertisePushOptions":        _bool,
			"receive.autogc":                      _bool,
			"receive.denyCurrentBranch":           carapace.ActionValues("refuse", "warn", "ignore", "updateInstead"),
			"receive.denyDeleteCurrent":           _bool,
			"receive.denyDeletes":                 _bool,
			"receive.denyNonFastForwards":         _bool,
			"receive.fsckObjects":                 _bool,
			"receive.hideRefs":                    carapace.ActionValues().Usage("refs to hide"),
			"receive.keepAlive":                   carapace.ActionValues().Usage("keep alive time"),
			"receive.maxInputSize":                carapace.ActionValues().Usage("max input size"),
			"receive.shallowUpdate":               _bool,
			"receive.unpackLimit":                 carapace.ActionValues().Usage("unpack limit for receive-pack"),
			"receive.updateServerInfo":            _bool,
			"reftable.blockSize":                  carapace.ActionValues().Usage("reftable block size"),
			"reftable.geometricFactor":            carapace.ActionValues().Usage("reftable geometric factor"),
			"reftable.indexObjects":               _bool,
			"reftable.lockTimeout":                carapace.ActionValues().Usage("reftable lock timeout"),
			"reftable.restartInterval":            carapace.ActionValues().Usage("reftable restart interval"),
			"remote.pushDefault":                  ActionRemotes(),
			"repack.updateServerInfo":             _bool,
			"repack.useDeltaBaseOffset":           _bool,
			"repack.useDeltaIslands":              _bool,
			"repack.writeBitmaps":                 _bool,
			"rerere.autoUpdate":                   _bool,
			"rerere.enabled":                      _bool,
			"revert.reference":                    _bool,
			"safe.bareRepository":                 carapace.ActionValues("all", "explicit"),
			"safe.directory":                      carapace.ActionDirectories(),
			"sendemail.annotate":                  _bool,
			"sendemail.chainReplyTo":              _bool,
			"sendemail.confirm":                   carapace.ActionValues("always", "never", "cc", "compose", "auto"),
			"sendemail.envelopeSender":            carapace.ActionValues().Usage("envelope sender address"),
			"sendemail.from":                      carapace.ActionValues().Usage("From: address"),
			"sendemail.mailmap":                   _bool,
			"sendemail.multiEdit":                 _bool,
			"sendemail.signedOffByCc":             _bool,
			"sendemail.smtpEncryption":            carapace.ActionValues("tls", "ssl"),
			"sendemail.smtpPass":                  carapace.ActionValues().Usage("SMTP password"),
			"sendemail.smtpServer":                carapace.ActionValues().Usage("SMTP server"),
			"sendemail.smtpServerPort":            carapace.ActionValues().Usage("SMTP server port"),
			"sendemail.smtpUser":                  carapace.ActionValues().Usage("SMTP username"),
			"sendemail.suppressCc":                carapace.ActionValues("all", "self", "cc", "bodycc", "sob", "sobcc", "misc", "muttcc"),
			"sendemail.suppressFrom":              _bool,
			"sendemail.thread":                    _bool,
			"sendemail.to":                        carapace.ActionValues().Usage("To: address"),
			"sendemail.transferEncoding":          carapace.ActionValues("7bit", "8bit", "quoted-printable", "base64"),
			"sendemail.validate":                  _bool,
			"sendemail.xmailer":                   _bool,
			"sequence.editor":                     bridge.ActionCarapaceBin().Split(),
			"showBranch.default":                  ActionLocalBranches().UniqueList(" "),
			"sparse.expectFilesOutsideOfPatterns": _bool,
			"splitIndex.maxPercentChange":         carapace.ActionValues().Usage("max percent change for split index"),
			"splitIndex.sharedIndexExpire":        carapace.ActionValues().Usage("shared index expiry"),
			"ssh.variant":                         carapace.ActionValues("auto", "simple", "openssh", "plink", "tortoiseplink", "putty"),
			"stash.showIncludeUntracked":          _bool,
			"stash.showPatch":                     _bool,
			"stash.showStat":                      _bool,
			"status.aheadBehind":                  _bool,
			"status.branch":                       _bool,
			"status.displayCommentPrefix":         _bool,
			"status.relativePaths":                _bool,
			"status.renameLimit":                  carapace.ActionValues().Usage("rename limit"),
			"status.renames":                      _bool,
			"status.short":                        _bool,
			"status.showStash":                    _bool,
			"status.showUntrackedFiles":           carapace.ActionValues("no", "normal", "all"),
			"status.submoduleSummary":             _bool,
			"submodule.active":                    ActionSubmoduleNames().UniqueList(" "),
			"submodule.alternateErrorStrategy":    carapace.ActionValues("die", "info", "ignore"),
			"submodule.alternateLocation":         _bool,
			"submodule.fetchJobs":                 carapace.ActionValues().Usage("number of parallel submodule jobs"),
			"submodule.propagateBranches":         _bool,
			"submodule.recurse":                   _bool,
			"tag.forceSignAnnotated":              _bool,
			"tag.gpgSign":                         _bool,
			"tag.sort":                            ActionFieldNames(),
			"tar.umask": carapace.Batch(
				carapace.ActionValuesDescribed(
					"true", "use the umask of the current process",
					"false", "use the default permission (usually 0002)",
					"umask", "use the umask of the current process",
					"0", "use the default permission (usually 0002)",
				),
				carapace.ActionCallback(func(c carapace.Context) carapace.Action {
					if strings.HasPrefix(c.Value, "0") {
						return fs.ActionFileModesNumeric().Prefix("0")
					}
					return carapace.ActionValues("0").NoSpace('0')
				}),
			).ToA(),
			"trace2.configParams":                 carapace.ActionValues().Usage("config parameters to trace"),
			"trace2.destinationDebug":             _bool,
			"trace2.envVars":                      carapace.ActionValues().Usage("environment variables to trace"),
			"trace2.eventBrief":                   _bool,
			"trace2.eventNesting":                 carapace.ActionValues().Usage("max nesting for trace2 events"),
			"trace2.eventTarget":                  carapace.ActionValues().Usage("trace2 event target"),
			"trace2.maxFiles":                     carapace.ActionValues().Usage("max trace2 log files"),
			"trace2.normalBrief":                  _bool,
			"trace2.normalTarget":                 carapace.ActionValues().Usage("trace2 normal target"),
			"trace2.perfBrief":                    _bool,
			"trace2.perfTarget":                   carapace.ActionValues().Usage("trace2 perf target"),
			"trailer.ifexists":                    carapace.ActionValues("addIfDifferent", "addIfDifferentNeighbor", "add", "replace", "doNothing"),
			"trailer.ifmissing":                   carapace.ActionValues("add", "doNothing"),
			"trailer.separators":                  carapace.ActionValues().Usage("trailer separators"),
			"trailer.where":                       carapace.ActionValues("end", "start", "after", "before"),
			"transfer.fsckObjects":                _bool,
			"transfer.hideRefs":                   carapace.ActionValues().Usage("refs to hide from transfer"),
			"transfer.unpackLimit":                carapace.ActionValues().Usage("unpack limit for transfer"),
			"uploadarchive.allowUnreachable":      _bool,
			"uploadpack.allowAnySHA1InWant":       _bool,
			"uploadpack.allowFilter":              _bool,
			"uploadpack.allowReachableSHA1InWant": _bool,
			"uploadpack.allowRefInWant":           _bool,
			"uploadpack.allowTipSHA1InWant":       _bool,
			"uploadpack.hideRefs":                 carapace.ActionValues().Usage("refs to hide from upload-pack"),
			"uploadpack.keepAlive":                carapace.ActionValues().Usage("keep alive time"),
			"uploadpackfilter.allow":              carapace.ActionValues("always", "never"),
			"uploadpackfilter.tree.maxDepth":      carapace.ActionValues().Usage("max depth for tree filter"),
			"user.email":                          ActionAuthors().NoSpace(),
			"user.name":                           ActionAuthors(),
			"user.signingKey":                     os.ActionGpgKeyIds(),
			"user.useConfigOnly":                  _bool,
			"versionsort.suffix":                  carapace.ActionValues().Usage("suffix for version sort"),
			"web.browser":                         bridge.ActionCarapaceBin().Split(),
			"worktree.guessRemote":                _bool,
			"worktree.useRelativePaths":           _bool,
		}[config]); ok {
			return a
		}

		splitted := strings.Split(config, ".")
		switch splitted[0] {
		case "alias":
			if strings.HasPrefix(c.Value, "!") {
				return bridge.ActionCarapaceBin().Split().Prefix("!")
			}
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
					"cmd":  bridge.ActionCarapaceBin().SplitP(),
					"path": carapace.ActionFiles(),
				}[splitted[2]]
			}
		case "diff":
			switch len(splitted) {
			case 3:
				return carapace.ActionMap{
					"binary":        _bool,
					"cachetextconv": _bool,
					"command":       bridge.ActionCarapaceBin().SplitP(),
					"textconv":      bridge.ActionCarapaceBin().SplitP(),
					"trustExitCode": _bool,
					"wordRegex":     carapace.ActionValues().Usage("word regex"),
					"xfuncname":     carapace.ActionValues().Usage("extended function name regex"),
				}[splitted[2]]
			}
		case "remote":
			switch len(splitted) {
			case 3:
				remoteSub := carapace.ActionMap{
					"fetch":              bridge.ActionCarapaceBin("git", "fetch").Split(),
					"followRemoteHEAD":   carapace.ActionValues("default", "warn"),
					"mirror":             _bool,
					"partialclonefilter": carapace.ActionValues("blob:none", "tree:0"),
					"promisor":           _bool,
					"proxy":              bridge.ActionCarapaceBin().Split(),
					"proxyAuthMethod": carapace.ActionValuesDescribed(
						"anyauth", "Automatically pick a suitable authentication method",
						"basic", "HTTP Basic authentication",
						"digest", "HTTP Digest authentication",
						"negotiate", "GSS-Negotiate authentication",
						"ntlm", "NTLM authentication",
					),
					"prune":             _bool,
					"pruneTags":         _bool,
					"push":              bridge.ActionCarapaceBin("git", "push").Split(),
					"pushurl":           carapace.ActionValues().Usage("URL for pushes"),
					"receivepack":       bridge.ActionCarapaceBin().Split(),
					"serverOption":      carapace.ActionValues().Usage("server options").UniqueList(" "),
					"skipDefaultUpdate": _bool,
					"skipFetchAll":      _bool,
					"tagOpt":            carapace.ActionValues("--tags", "--no-tags"),
					"uploadpack":        bridge.ActionCarapaceBin().Split(),
					"url":               carapace.ActionValues().Usage("remote URL"),
					"vcs":               carapace.ActionValues().Usage("VCS type"),
				}
				return remoteSub[splitted[2]]
			}
		}
		return carapace.ActionValues()
	})
}

// ActionPrettyFormats completes pretty print formats
//
//	oneline
//	short
func ActionPrettyFormats() carapace.Action {
	return carapace.ActionValuesDescribed(
		"oneline", "`<sha1> <title line>`",
		"short", "`commit <sha1>\nAuthor: <author>`",
		"medium", "`commit <sha1>\nAuthor: <author>\nDate:   <date>\n\n    <title line>`",
		"full", "`commit <sha1>\nAuthor: <author>\nCommit: <committer>\n\n    <title line>`",
		"fuller", "`commit <sha1>\nAuthor:     <author>\nAuthorDate: <date>\nCommit:     <committer>\nCommitDate: <date>\n\n    <title line>`",
		"reference", "`<abbrev-commit> (<title line>, <short-author-date>)`",
		"email", "`From <sha1> <date>\nFrom: <author>\nDate: <date>\nSubject: [PATCH] <title line>\n\n<body>`",
		"raw", "shows the raw git object",
		"format:", "custom format string (prefix with format:)",
		"tformat:", "custom format string with terminating newline (prefix with tformat:)",
	)
}

// ActionMergeTools completes merge tools
//
//	meld
//	kdiff3
func ActionMergeTools() carapace.Action {
	return carapace.ActionValues(
		"araxis",
		"bc",
		"codecompare",
		"deltawalker",
		"diffmerge",
		"diffuse",
		"ecmerge",
		"emerge",
		"examdiff",
		"guiffy",
		"gvimdiff",
		"kdiff3",
		"kompare",
		"meld",
		"nvimdiff",
		"opendiff",
		"p4merge",
		"smerge",
		"tkdiff",
		"tortoisemerge",
		"vimdiff",
		"winmerge",
		"xxdiff",
	)
}
