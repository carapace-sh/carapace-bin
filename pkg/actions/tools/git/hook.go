package git

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/traverse"
)

// ActionHooks completes hooks
//
//	pre-commit.sample
//	pre-merge-commit.sample
func ActionHooks() carapace.Action {
	return carapace.ActionExecutables("hooks").
		ChdirF(traverse.GitDir).
		Tag("hooks").
		UidF(Uid("hook"))
}

// ActionHookEvents completes hook events
//
//	applypatch-msg (invoked before the patch message is applied)
//	pre-applypatch (invoked after the patch is applied)
func ActionHookEvents() carapace.Action {
	return carapace.ActionValuesDescribed(
		"aborted", "invoked after reference transaction was aborted",
		"applypatch-msg", "invoked before the patch message is applied",
		"commit-msg", "invoked after editing the commit message",
		"committed", "invoked after reference transaction was committed",
		"fsmonitor-watchman", "invoked when the configuration option core.fsmonitor is set to watchman",
		"p4-changelist", "invoked after the changelist message has been edited by the user",
		"p4-post-changelist", "invoked after the submit has successfully occurred in P4",
		"p4-pre-submit", "invoked before submit",
		"p4-prepare-changelist", "invoked right after preparing the default changelist message",
		"post-applypatch", "invoked after the patch is applied and a commit is made",
		"post-checkout", "invoked after checkout",
		"post-commit", "invoked after a commit is made",
		"post-index-change", "invoked when the index is written in read-cache.c do_write_locked_index",
		"post-merge", "invoked after merge",
		"post-receive", "invoked once after all the proposed ref updates are processed",
		"post-rewrite", "invoked by commands that rewrite commits",
		"post-update", "invoked once after all the refs have been updated",
		"pre-applypatch", "invoked after the patch is applied",
		"pre-auto-gc", "invoked by \"git gc --auto\"",
		"pre-commit", "invoked before commit",
		"pre-merge-commit", "invoked before merge commit",
		"pre-push", "invoked before push",
		"pre-rebase", "invoked before rebase",
		"pre-receive", "invoked before starting to update refs on the remote repository",
		"prepare-commit-msg", "invoked right after preparing the default log message",
		"prepared", "invoked after queued reference updates are locked on disk",
		"preparing", "invoked before queued reference updates are locked on disk",
		"proc-receive", "invoked once for the receive operation",
		"push-to-checkout", "invoked before updating reference(s) on push", // TODO
		"rebase", "invoked after squash",
		"reference-transaction", "invoked on each reference update",
		"sendemail-validate", "invoked before sending email",
		"update", "invoked once for each ref to be updated",
	).Tag("hook events").Uid("git", "hook-event")
}

// ActionHookNames completes configured hook names
//
//	linter
//	no-leaks
func ActionHookNames() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommandE("git", "config", "--name-only", "--get-regexp", "^hook\\.")(func(output []byte, err error) carapace.Action {
			if err != nil {
				return carapace.ActionValues()
			}
			names := make(map[string]bool)
			for line := range strings.SplitSeq(string(output), "\n") {
				line = strings.TrimSpace(line)
				if line == "" {
					continue
				}
				parts := strings.SplitN(line, ".", 3)
				if len(parts) >= 2 {
					names[parts[1]] = true
				}
			}
			vals := make([]string, 0, len(names))
			for name := range names {
				vals = append(vals, name)
			}
			return carapace.ActionValues(vals...).Tag("hook names")
		})
	}).UidF(Uid("hook-name"))
}
