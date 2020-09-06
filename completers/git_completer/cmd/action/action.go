package action

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/rsteube/carapace"
)

func rootDir() (string, error) {
	if output, err := exec.Command("git", "rev-parse", "--show-toplevel").Output(); err != nil {
		return "", err
	} else {
		return strings.Split(string(output), "\n")[0], nil
	}
}

func ActionRemotes() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("git", "remote").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return carapace.ActionValues(strings.Split(string(output), "\n")...)
		}
	})
}

// TODO combine with ActionRemoteBranches
func ActionLocalBranches() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("git", "branch", "--format", "%(refname)").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			refs := make([]string, 0)
			for _, line := range strings.Split(string(output), "\n") {
				if len(line) > 12 {
					refs = append(refs, line[11:])
				}
			}
			return carapace.ActionValues(refs...)
		}
	})
}

func ActionRemoteBranches() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("git", "branch", "--remote", "--format", "%(refname)").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			refs := make([]string, 0)
			for _, line := range strings.Split(string(output), "\n") {
				if len(line) > 14 {
					refs = append(refs, line[13:])
				}
			}
			return carapace.ActionValues(refs...)
		}
	})
}

func ActionCommits() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("git", "log", "--pretty=tformat:%H   %<(64,trunc)%s", "--all", "--max-count", "1000").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			refs := make([]string, 0)
			for _, line := range strings.Split(string(output), "\n") {
				if len(line) > 14 {
					refs = append(refs, line[:40], strings.TrimSpace(line[43:]))
				}
			}
			return carapace.ActionValuesDescribed(refs...)
		}
	})
}

// TODO multiparts action to complete step by step
func ActionUnstagedChanges() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("git", "status", "--porcelain").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			if root, err := rootDir(); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				if wd, err := os.Getwd(); err != nil {
					return carapace.ActionMessage(err.Error())
				} else {
					untracked := make([]string, 0)
					for _, line := range strings.Split(string(output), "\n") {
						if len(line) > 3 && line[1] != ' ' { // no blank line and not already staged
							if relativePath, err := filepath.Rel(wd, root+"/"+line[3:]); err != nil {
								return carapace.ActionMessage(err.Error())
							} else {
								untracked = append(untracked, relativePath, line[:2])
							}
						}
					}
					return carapace.ActionValuesDescribed(untracked...)
				}
			}
		}
	})
}

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
