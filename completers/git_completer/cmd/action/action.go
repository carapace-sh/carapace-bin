package action

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
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

func ActionRemoteBranches(remote string) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if branches, err := Branches(RefOption{RemoteBranches: true}); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := make([]string, 0)
			for _, branch := range branches {
				if strings.HasPrefix(branch.Name, remote) {
					vals = append(vals, strings.TrimPrefix(branch.Name, remote+"/"), branch.Message)
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		}
	})
}

func ActionCurrentBranch() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("git", "branch", "--show-current").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return carapace.ActionValues(strings.Split(string(output), "\n")[0])
		}
	})
}

type RefOption struct {
	LocalBranches  bool
	RemoteBranches bool
	Commits        int
	Tags           bool
}

var RefOptionDefault = RefOption{
	LocalBranches:  true,
	RemoteBranches: true,
	Commits:        100,
	Tags:           true,
}

type Branch struct {
	Name    string
	Message string
}

func Branches(refOption RefOption) ([]Branch, error) {
	args := []string{"branch", "--format", "%(refname)\n%(subject)"}
	if refOption.LocalBranches && refOption.RemoteBranches {
		args = append(args, "--all")
	} else if !refOption.LocalBranches && refOption.RemoteBranches {
		args = append(args, "--remote")
	} else if !refOption.LocalBranches && !refOption.RemoteBranches {
		return []Branch{}, nil
	}

	if output, err := exec.Command("git", args...).Output(); err != nil {
		return nil, err
	} else {
		lines := strings.Split(string(output), "\n")
		branches := make([]Branch, len(lines)/2)
		for index, line := range lines[:len(lines)-1] {
			if index%2 == 0 {
				trimmed := strings.TrimPrefix(line, "refs/heads/")
				trimmed = strings.TrimPrefix(trimmed, "refs/remotes/")
				branches[index/2] = Branch{trimmed, lines[index+1]}
			}
		}
		return branches, err
	}
}

type Tag struct {
	Name    string
	Message string
}

func Tags(refOption RefOption) ([]Tag, error) {
	if !refOption.Tags {
		return []Tag{}, nil
	}

	if output, err := exec.Command("git", "tag", "--format", "%(refname)\n%(subject)").Output(); err != nil {
		return nil, err
	} else {
		lines := strings.Split(string(output), "\n")
		tags := make([]Tag, len(lines)/2)
		for index, line := range lines[:len(lines)-1] {
			if index%2 == 0 {
				tags[index/2] = Tag{strings.TrimPrefix(line, "refs/tags/"), lines[index+1]}
			}
		}
		return tags, err
	}
}

type Commit struct {
	Ref     string
	Message string
}

func Commits(refOption RefOption) ([]Commit, error) {
	if refOption.Commits <= 0 {
		return []Commit{}, nil
	}

	if output, err := exec.Command("git", "log", "--pretty=tformat:%h   %<(64,trunc)%s", "--all", "--max-count", strconv.Itoa(refOption.Commits)).Output(); err != nil {
		return nil, err
	} else {
		lines := strings.Split(string(output), "\n")
		commits := make([]Commit, 0)
		for index, line := range lines[:len(lines)-1] {
			if len(line) > 10 { // TODO duh?
				commits = append(commits, Commit{line[:7], strings.TrimSpace(line[10:])})
				if index == 0 {
					commits = append(commits, Commit{"HEAD", strings.TrimSpace(line[10:])}) // TOD fix this
				} else {
					commits = append(commits, Commit{"HEAD~" + fmt.Sprintf("%0"+strconv.Itoa(len(strconv.Itoa(refOption.Commits))-1)+"d", index), strings.TrimSpace(line[10:])}) // TOD fix this
				}
			}
		}
		return commits, nil
	}
}

func ActionRefs(refOption RefOption) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		vals := make([]string, 0)
		if branches, err := Branches(refOption); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			for _, branch := range branches {
				vals = append(vals, branch.Name, branch.Message)
			}
		}
		if commits, err := Commits(refOption); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			for _, commit := range commits {
				vals = append(vals, commit.Ref, commit.Message)
			}
		}
		if tags, err := Tags(refOption); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			for _, tag := range tags {
				vals = append(vals, tag.Name, tag.Message)
			}
		}

		return carapace.ActionValuesDescribed(vals...)

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
