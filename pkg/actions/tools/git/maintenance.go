package git

import "github.com/carapace-sh/carapace"

// ActionMaintenanceTasks completes maintenance tasks
//
//	commit-graph (Update the commit-graph files incrementally)
//	prefetch (Updates the object directory with the latest objects from all registered remotes)
func ActionMaintenanceTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"commit-graph", "Update the commit-graph files incrementally",
		"prefetch", "Updates the object directory with the latest objects from all registered remotes",
		"gc", "Clean up unnecessary files and optimize the local repository",
		"geometric", "Avoid all-into-one repacks when possible",
		"loose-objects", "Clean up loose objects and places them into pack-files",
		"incremental-repack", "Repack the object directory using the multi-pack-index feature",
		"pack-refs", "Collect the loose reference files and collects them into a single file",
		"reflog-expire", "Delete entries in the reflog older than the expiry threshold",
		"rerere-gc", "Garbage collection for stale entries in the rerere cache",
		"worktree-prune", "Delete stale or broken worktrees",
	).Uid("git", "maintenance-task")
}
