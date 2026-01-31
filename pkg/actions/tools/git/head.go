package git

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"strings"
)

// ActionHeads completes heads
//
//	FETCH_HEAD (message)
//	ORIG_HEAD (message)
func ActionHeads() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		refs := []string{
			// https://git-scm.com/docs/gitrevisions#_specifying_revisions
			"HEAD",             // names the commit on which you based the changes in the working tree.
			"FETCH_HEAD",       // records the branch which you fetched from a remote repository with your last git fetch invocation.
			"ORIG_HEAD",        // is created by commands that move your HEAD in a drastic way (git am, git merge, git rebase, git reset), to record the position of the HEAD before their operation, so that you can easily change the tip of the branch back to the state before you ran them.
			"MERGE_HEAD",       // records the commit(s) which you are merging into your branch when you run git merge.
			"REBASE_HEAD",      // during a rebase, records the commit at which the operation is currently stopped, either because of conflicts or an edit command in an interactive rebase.
			"REVERT_HEAD",      // records the commit which you are reverting when you run git revert.
			"CHERRY_PICK_HEAD", // records the commit which you are cherry-picking when you run git cherry-pick.
			"BISECT_HEAD",      // records the current commit to be tested when you run git bisect --no-checkout.
			"AUTO_MERGE",       // records a tree object corresponding to the state the ort merge strategy wrote to the working tree when a merge operation resulted in conflicts.
		}

		if output, err := c.Command("git", "remote").Output(); err == nil {
			lines := strings.Split(string(output), "\n")
			for _, remote := range lines[:len(lines)-1] {
				refs = append(refs, remote+"/HEAD") // head commit of remotes default branch
			}
		}

		batch := carapace.Batch()
		for _, ref := range refs {
			args := []string{"log", "--no-notes", "--pretty=tformat:%<(64,trunc)%s", "--max-count", "1", ref}
			batch = append(batch, carapace.ActionExecCommand("git", args...)(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")

				vals := make([]string, 0)
				for _, line := range lines[:len(lines)-1] {
					vals = append(vals, ref, strings.TrimSpace(line))
				}
				return carapace.ActionValuesDescribed(vals...).Style(styles.Git.Head)
			}).Suppress("unknown revision"))
		}
		return batch.ToA().Tag("heads").UidF(Uid("ref"))
	})
}
