package gh

import "github.com/carapace-sh/carapace"

// ActionCodespaceViewFields completes codespace fields for view
//
//	displayName
//	repository
func ActionCodespaceViewFields() carapace.Action {
	return carapace.ActionValues(
		"billableOwner",
		"createdAt",
		"devcontainerPath",
		"displayName",
		"environmentId",
		"gitStatus",
		"idleTimeoutMinutes",
		"lastUsedAt",
		"location",
		"machineDisplayName",
		"machineName",
		"name",
		"owner",
		"prebuild",
		"recentFolders",
		"repository",
		"retentionExpiresAt",
		"retentionPeriodDays",
		"state",
		"vscsTarget",
	)
}
