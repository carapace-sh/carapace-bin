package git

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/time"
	"github.com/carapace-sh/carapace/pkg/style"
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
	Heads          bool
	Tags           bool
	Stashes        bool
	Notes          bool
}

func (o RefOption) Default() RefOption {
	o.LocalBranches = true
	o.RemoteBranches = true
	o.Heads = true
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
		if !strings.ContainsAny(c.Value, "~^@") {
			batch := carapace.Batch()

			if refOption.LocalBranches {
				batch = append(batch, ActionLocalBranches())
			}

			if refOption.RemoteBranches {
				batch = append(batch, ActionRemoteBranches(""))
			}

			if refOption.Heads {
				batch = append(batch, ActionHeads())
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

		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			index := max(strings.LastIndex(c.Value, "~"), strings.LastIndex(c.Value, "^"), strings.LastIndex(c.Value, "@"))
			switch c.Value[index] {
			case '^':
				return ActionRefParents(c.Value[:index])
			case '@':
				return carapace.Batch(
					time.ActionDateTime(time.DateTimeOpts{}),
					carapace.ActionValues("yesterday", "push", "upstream").Style(style.Blue).Suffix("}"),
					carapace.ActionMultiParts(".", func(c carapace.Context) carapace.Action {
						b := carapace.Batch()
						if len(c.Parts)%2 == 1 {
							b = append(b, carapace.ActionValues("year", "month", "week", "day", "hour", "second").Style(style.Blue).Suffix("."))
						}
						if len(c.Parts) > 0 && len(c.Parts)%2 == 0 {
							b = append(b, carapace.ActionValues("ago").Style(style.Blue).Suffix("}"))
						}
						return b.ToA()
					}),
					ActionReflogs(c.Value[:index]).Suffix("}"),
				).ToA().Prefix(c.Value[:index+1] + "{")

			default: // '~'
				return ActionRefCommits(c.Value[:index])
			}
		}).UidF(Uid("ref"))
	})
}

// ActionRefRanges completes refs as range
//
//	HEAD..HEAD~17 (last commit msg)
//	v.0.0.3^2~03...v0.0.4~03 (last commit msg)
func ActionRefRanges(opts RefOption) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if strings.HasPrefix(c.Value, "^") { // negate
			return ActionRefs(opts).NoSpace().Prefix("^")
		}

		delimiter := ".."
		if strings.Contains(c.Value, "...") && strings.Count(c.Value, "..") == 1 {
			delimiter = "..."
		}
		return carapace.ActionMultiPartsN(delimiter, 2, func(c carapace.Context) carapace.Action {
			return ActionRefs(opts).NoSpace()
		})
	})
}
