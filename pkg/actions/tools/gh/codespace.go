package gh

import "github.com/carapace-sh/carapace"

// ActionCodespaceViewFields completes codespace fields for view
//
//	displayName
//	repository
func ActionCodespaceViewFields() carapace.Action {
	return carapace.ActionValues(
		"name",
		"displayName",
		"state",
		"owner",
		"billableOwner",
		"location",
		"repository",
		"gitStatus",
		"devcontainerPath",
		"machineName",
		"machineDisplayName",
		"prebuild",
		"createdAt",
		"lastUsedAt",
		"idleTimeoutMinutes",
		"retentionPeriodDays",
		"retentionExpiresAt",
		"recentFolders",
		"vscsTarget",
		"environmentId",
	)
}
