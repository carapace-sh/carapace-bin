package jj

import (
	"github.com/rsteube/carapace"
)

type RevsOption struct {
	LocalBranches  bool
	RemoteBranches bool
	Commits        int
	HeadCommits    int
	Tags           bool
}

func (o RevsOption) Default() RevsOption {
	o.LocalBranches = true
	o.RemoteBranches = true
	o.Commits = 100
	o.HeadCommits = 100
	o.Tags = true
	return o

}

// ActionRevs completes refs (commits, branches, tags)
func ActionRevs(revOption RevsOption) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO other refs and ranges
		// if !strings.ContainsAny(c.Value, "~^") {
		batch := carapace.Batch()

		if revOption.LocalBranches {
			batch = append(batch, ActionLocalBranches())
		}

		if revOption.RemoteBranches {
			batch = append(batch, ActionRemoteBranches(""))
		}

		if revOption.Commits > 0 {
			batch = append(batch, ActionRecentCommits(revOption.Commits))
		}

		if revOption.HeadCommits > 0 {
			batch = append(batch, ActionHeadCommits(revOption.HeadCommits))
		}

		if revOption.Tags {
			batch = append(batch, ActionTags())
		}

		return batch.ToA()
		// }

		// index := max(strings.LastIndex(c.Value, "~"), strings.LastIndex(c.Value, "^"))
		// switch c.Value[index] {
		// case '^':
		// 	return ActionRefParents(c.Value[:index]).Prefix(c.Value[:index+1])
		// default: // '~'
		// 	return ActionRefCommits(c.Value[:index]).Prefix(c.Value[:index+1])
		// }
	})
}
