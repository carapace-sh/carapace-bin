package but

import (
	"net/url"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"github.com/carapace-sh/carapace/pkg/uid"
)

type CliIdsOpts struct {
	Changes  bool // TODO differentiate changes?
	Branches bool
	Commits  bool
	Stacks   bool
}

func (o CliIdsOpts) Default() CliIdsOpts {
	o.Changes = true
	o.Branches = true
	o.Commits = true
	o.Stacks = true
	return o
}

// ActionCliIds completes cli ids
//
//	na
//	c3
func ActionCliIds(opts CliIdsOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return actionStatus(strings.Contains(c.Value, ":"), func(status butStatus) carapace.Action {
			batch := carapace.Batch()

			if opts.Changes {
				for _, unassigned := range status.UnassignedChanges {
					batch = append(batch, carapace.ActionStyledValuesDescribed(unassigned.CliID, unassigned.Description(), styleForChangeType(unassigned.ChangeType)).UidF(
						func(s string, uc uid.Context) (*url.URL, error) {
							return git.Uid("change")(unassigned.FilePath, uc)
						}))
				}
			}

			for _, stack := range status.Stacks {
				if opts.Changes {
					for _, assigned := range stack.AssignedChanges {
						batch = append(batch, carapace.ActionStyledValuesDescribed(assigned.CliID, assigned.Description(), styleForChangeType(assigned.ChangeType)).UidF(
							func(s string, uc uid.Context) (*url.URL, error) {
								return git.Uid("change")(assigned.FilePath, uc)
							}))
					}

				}
				if opts.Stacks && !opts.Branches { // TODO stack has the same id as its branch anyway
					batch = append(batch, carapace.ActionValues(stack.CliID)) // TODO style (details from the corresponding branch)
				}
				for _, branch := range stack.Branches {
					if opts.Branches {
						batch = append(batch, carapace.ActionStyledValuesDescribed(branch.CliID, branch.Name, styles.Git.Branch).UidF(
							func(s string, uc uid.Context) (*url.URL, error) {
								return git.Uid("local-branch")(branch.Name, uc)
							}))
					}
					for _, commit := range branch.Commits {
						batch = append(batch, carapace.ActionValuesDescribed(commit.CliID, commit.Message).UidF(
							func(s string, uc uid.Context) (*url.URL, error) {
								return git.Uid("ref")(commit.CommitID, uc)
							}))

						for _, committed := range commit.Changes {
							batch = append(batch, carapace.ActionStyledValuesDescribed(committed.CliID, committed.Description(), styleForChangeType(committed.ChangeType)).UidF(
								func(s string, uc uid.Context) (*url.URL, error) {
									return git.Uid("change")(committed.FilePath, uc)
								}))
						}
					}
				}
			}
			return batch.ToA().MultiParts(":")
		}).Tag("cli ids")
	})
}
