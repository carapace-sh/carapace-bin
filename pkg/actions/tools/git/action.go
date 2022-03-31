package git

import (
	"strings"

	exec "golang.org/x/sys/execabs"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

func rootDir() (string, error) {
	if output, err := exec.Command("git", "rev-parse", "--show-toplevel").Output(); err != nil {
		return "", err
	} else {
		return strings.Split(string(output), "\n")[0], nil
	}
}

type RefOption struct {
	LocalBranches  bool
	RemoteBranches bool
	Commits        int
	Tags           bool
	Stashes        bool
}

var RefOptionDefault = RefOption{
	LocalBranches:  true,
	RemoteBranches: true,
	Commits:        100,
	Tags:           true,
	Stashes:        true,
}

// ActionRefs completes git references (commits, branches, tags)
//   HEAD~1 (last commit msg)
//   v0.0.1 (last commit msg)
func ActionRefs(refOption RefOption) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := make([]string, 0)
		if branches, err := branches(refOption); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			for _, branch := range branches {
				vals = append(vals, branch.Name, branch.Message, style.Blue)
			}
		}
		if commits, err := commits(refOption); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			for _, commit := range commits {
				vals = append(vals, commit.Ref, commit.Message, style.Default)
			}
		}
		if tags, err := tags(refOption); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			for _, tag := range tags {
				vals = append(vals, tag.Name, tag.Message, style.Yellow)
			}
		}

		if refOption.Stashes {
			return carapace.Batch(
				carapace.ActionStyledValuesDescribed(vals...),
				ActionStashes().Style(style.Green),
			).ToA()
		}
		return carapace.ActionStyledValuesDescribed(vals...)

	})
}

// ActionFieldNames completes field names
//   author (the author header-field)
//   body (the body of the message)
func ActionFieldNames() carapace.Action {
	return carapace.ActionValuesDescribed(
		"authordate", "the date component of the author header-field",
		"authoremail", "the email component of the author header-field",
		"authorname", "the name component of the author header-field",
		"author", "the author header-field",
		"body", "the body of the message",
		"committerdate", "the date component of the committer header-field",
		"committeremail", "the email component of the committer header-field",
		"committername", "the name component of the committer header-field",
		"committer", "the committer header-field",
		"contents", "complete message",
		"creatordate", "the date component of the creator header-field",
		"creator", "the creator header-field",
		"deltabase", "object name of the delta base of the object",
		"HEAD", "* if HEAD matches ref or space otherwise",
		"numparent", "number of parent objects",
		"objectname", "object name (SHA-1)",
		"objectsize", "the size of the object",
		"object", "the object header-field",
		"objecttype", "the type of the object",
		"parent", "the parent header-field",
		"push", "name of a local ref which represents the @{push} location for the displayed ref",
		"refname", "name of the ref",
		"subject", "the subject of the message",
		"symref", "the ref which the given symbolic ref refers to",
		"taggerdate", "the date component of the tagger header-field",
		"taggeremail", "the email component of the tagger header-field",
		"taggername", "the name component of the tagger header-field",
		"tagger", "the tagger header-field",
		"tag", "the tag header-field",
		"trailers", "structured information in commit messages",
		"tree", "the tree header-field",
		"type", "the type header-field",
		"upstream", "name of a local ref which can be considered “upstream” from the displayed ref",
		"version:refname", "sort by versions",
	)
}

// ActionCleanupMode completes cleanup modes
//   strip (strip empty lines and trailing whitespace)
//   whitespace (same as strip except #commentary is not removed)
func ActionCleanupMode() carapace.Action {
	return carapace.ActionValuesDescribed(
		"strip", "strip empty lines and trailing whitespace",
		"whitespace", "same as strip except #commentary is not removed",
		"verbatim", "do not change the message at all",
		"scissors", "same as whitespace except that everything from (and including) the line found below is truncated",
		"default", " Same as strip if the message is to be edited. Otherwise whitespace",
	)
}

// ActionMergeStrategy completes merge strategies
//   octopus (resolve cases with more than two heads)
//   ours (auto-resolve cleanly by favoring our version)
func ActionMergeStrategy() carapace.Action {
	return carapace.ActionValuesDescribed(
		"octopus", "resolve cases with more than two heads",
		"ours", "auto-resolve cleanly by favoring our version",
		"recursive", "recursively resolve two heads using a 3-way merge algorithm",
		"resolve", "resolve two heads using a 3-way merge algorithm",
		"subtree", "modified recursive straty with tree adjustment",
	)
}

// ActionMergeStrategyOptions completes merge strategy options
//   ours (auto-resolve favoring ours)
//   theirs (auto-resolve favoring theirs)
func ActionMergeStrategyOptions(strategy string) carapace.Action {
	switch strategy {
	case "recursive":
		return carapace.ActionValuesDescribed(
			"ours", "auto-resolve favoring ours",
			"theirs", "auto-resolve favoring theirs",
			"patience", "spend extra time to avoid mismerges",
			"diff-algorithm=", "set dif allgorithm",
			"ignore-space-change", "ignore space change",
			"ignore-all-space", "ignore all space",
			"ignore-space-at-eol", "ignore <space> at end of line",
			"ignore-cr-at-eol", "ignore <cr> at end of line",
			"renormalize", "enable renormalize",
			"no-renormalize", "disable renormalize",
			"no-renames", "turn off rename detection",
			"find-renames", "turn on rename detection",
			"subtree", "advance subtree stratebgy",
		)
	default:
		return carapace.ActionValues()
	}
}
