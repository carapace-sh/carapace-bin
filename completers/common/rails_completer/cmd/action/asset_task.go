package action

import "github.com/carapace-sh/carapace"

// ActionAssetTasks completes assets:* tasks
func ActionAssetTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"clean", "Remove old compiled assets",
		"clobber", "Remove the entire output path",
		"environment", "Load the asset environment",
		"precompile", "Compile all assets",
	).Tag("asset tasks").UidF(taskUid("assets"))
}

// ActionAssetNamespaces completes assets:* sub-namespaces
func ActionAssetNamespaces() carapace.Action {
	return carapace.ActionValuesDescribed(
		"reveal", "Print all available assets",
	).Tag("asset namespaces").UidF(namespaceUid("assets"))
}

// ActionAssetRevealTasks completes assets:reveal:* tasks
func ActionAssetRevealTasks() carapace.Action {
	return carapace.ActionValues("full").Tag("assets reveal tasks").UidF(taskUid("assets.reveal"))
}
