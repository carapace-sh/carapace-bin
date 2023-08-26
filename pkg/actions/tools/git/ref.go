package git

import (
	"strings"

	"github.com/rsteube/carapace"
)

func rootDir(c carapace.Context) (string, error) {
	if output, err := c.Command("git", "rev-parse", "--show-toplevel").Output(); err != nil {
		return "", err
	} else {
		return strings.Split(string(output), "\n")[0], nil
	}
}

type RefOption struct {
	LocalBranches  bool
	RemoteBranches bool
	Commits        int
	HeadCommits    int
	Tags           bool
	Stashes        bool
	Notes          bool
}

func (o RefOption) Default() RefOption {
	o.LocalBranches = true
	o.RemoteBranches = true
	o.Commits = 100
	o.HeadCommits = 100
	o.Tags = true
	o.Stashes = true
	o.Notes = false
	return o

}

// ActionRefs completes refs (commits, branches, tags)
//
//	HEAD~1 (last commit msg)
//	v0.0.1 (last commit msg)
func ActionRefs(refOption RefOption) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if !strings.ContainsAny(c.Value, "~^") {
			batch := carapace.Batch()

			if refOption.LocalBranches {
				batch = append(batch, ActionLocalBranches())
			}

			if refOption.RemoteBranches {
				batch = append(batch, ActionRemoteBranches(""))
			}

			if refOption.Commits > 0 {
				batch = append(batch, ActionRecentCommits(refOption.Commits))
			}

			if refOption.HeadCommits > 0 {
				batch = append(batch, ActionHeadCommits(refOption.HeadCommits))
			}

			if refOption.Tags {
				batch = append(batch, ActionTags())
			}

			if refOption.Stashes {
				batch = append(batch, ActionStashes())
			}

			if refOption.Notes {
				batch = append(batch, ActionNotes())
			}

			return batch.ToA()
		}

		index := max(strings.LastIndex(c.Value, "~"), strings.LastIndex(c.Value, "^"))
		switch c.Value[index] {
		case '^':
			return ActionRefParents(c.Value[:index]).Prefix(c.Value[:index+1])
		default: // '~'
			return ActionRefCommits(c.Value[:index]).Prefix(c.Value[:index+1])
		}
	})
}
