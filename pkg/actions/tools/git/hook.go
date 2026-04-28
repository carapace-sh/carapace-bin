package git

import (
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
		"applypatch-msg", "invoked before the patch message is applied",
		"pre-applypatch", "invoked after the patch is applied",
		"post-applypatch", "invoked after the patch is applied and a commit is made",
		"pre-commit", "invoked before commit",
		"pre-merge-commit", "invoked before merge commit",
		"prepare-commit-msg", "invoked right after preparing the default log message",
		"commit-msg", "invoked after editing the commit message",
		"post-commit", "invoked after a commit is made",
		"pre-rebase", "invoked before rebase",
		"post-checkout", "invoked after checkout",
		"post-merge", "invoked after merge",
		"pre-push", "invoked before push",
		"pre-receive", "invoked before starting to update refs on the remote repository",
		"update", "invoked once for each ref to be updated",
		"proc-receive", "invoked once for the receive operation",
		"post-receive", "invoked once after all the proposed ref updates are processed",
		"post-update", "invoked once after all the refs have been updated",
		"reference-transaction", "invoked on each reference update",
		"preparing", "invoked before queued reference updates are locked on disk",
		"prepared", "invoked after queued reference updates are locked on disk",
		"committed", "invoked after reference transaction was committed",
		"aborted", "invoked after reference transaction was aborted",
		"push-to-checkout", "invoked before updating reference(s) on push", // TODO
		"pre-auto-gc", "invoked by \"git gc --auto\"",
		"post-rewrite", "invoked by commands that rewrite commits",
		"rebase", "invoked after squash",
		"sendemail-validate", "invoked before sending email",
		"fsmonitor-watchman", "invoked when the configuration option core.fsmonitor is set to watchman",
		"p4-changelist", "invoked after the changelist message has been edited by the user",
		"p4-prepare-changelist", "invoked right after preparing the default changelist message",
		"p4-post-changelist", "invoked after the submit has successfully occurred in P4",
		"p4-pre-submit", "invoked before submit",
		"post-index-change", "invoked when the index is written in read-cache.c do_write_locked_index",
	).Tag("hook events").Uid("git", "hook-event")
}
