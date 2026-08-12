package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/lore"
	"github.com/spf13/cobra"
)

func globalOpts(cmd *cobra.Command) lore.GlobalOpts {
	opts := lore.GlobalOpts{}
	if flag := cmd.Flag("repository"); flag != nil && flag.Changed {
		opts.Repository = flag.Value.String()
	}
	if flag := cmd.Flag("identity"); flag != nil && flag.Changed {
		opts.Identity = flag.Value.String()
	}
	if flag := cmd.Flag("offline"); flag != nil && flag.Changed {
		opts.Offline = true
	}
	if flag := cmd.Flag("remote"); flag != nil && flag.Changed {
		opts.Remote = true
	}
	if flag := cmd.Flag("local"); flag != nil && flag.Changed {
		opts.Local = true
	}
	return opts
}

// ActionBranches completes branch names from `lore branch list --json`.
func ActionBranches(cmd *cobra.Command) carapace.Action {
	return lore.ActionBranches(globalOpts(cmd))
}

// ActionRevisions completes revision hashes from `lore revision history --json`.
func ActionRevisions(cmd *cobra.Command) carapace.Action {
	opts := lore.RevisionOpts{
		GlobalOpts: globalOpts(cmd),
	}
	if flag := cmd.Flag("branch"); flag != nil && flag.Changed {
		opts.Branch = flag.Value.String()
	}
	if flag := cmd.Flag("revision"); flag != nil && flag.Changed {
		opts.Revision = flag.Value.String()
	}
	return lore.ActionRevisions(opts)
}

// ActionIdentities completes user IDs from `lore auth list --json`.
func ActionIdentities(cmd *cobra.Command) carapace.Action {
	return lore.ActionIdentities(globalOpts(cmd))
}
