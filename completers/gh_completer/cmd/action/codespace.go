package action

import "github.com/rsteube/carapace"

func ActionCodespaces() carapace.Action {
	return carapace.ActionValues("TODO") // TODO codespace completion
}

func ActionCodespaceFiles(codespace string, expand bool) carapace.Action {
	return carapace.ActionValues("TODO") // TODO codespace completion
}

func ActionCodespaceMachines() carapace.Action {
	return carapace.ActionValues("TODO")
}

func ActionCodespacePorts(codespace string) carapace.Action {
	return carapace.ActionValues("TODO")
}

func ActionCodespaceFields() carapace.Action {
	return carapace.ActionValues(
		"name",
		"owner",
		"repository",
		"state",
		"gitStatus",
		"createdAt",
		"lastUsedAt",
	)
}

func ActionCodespacePortFields() carapace.Action {
	return carapace.ActionValues(
		"sourcePort",
		// "destinationPort", // TODO(mislav): this appears to always be blank?
		"visibility",
		"label",
		"browseUrl",
	)
}
