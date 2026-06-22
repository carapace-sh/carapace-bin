package action

import "github.com/carapace-sh/carapace"

// ActionTmpTasks completes tmp:* tasks
func ActionTmpTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"clear", "Clear all temp files",
		"create", "Create tmp directories",
	).Tag("tmp tasks").UidF(taskUid("tmp"))
}

// ActionTmpNamespaces completes tmp:* sub-namespaces
func ActionTmpNamespaces() carapace.Action {
	return carapace.ActionValues(
		"cache", "sockets", "pids", "screenshots", "storage",
	).Tag("tmp namespaces").UidF(namespaceUid("tmp"))
}

// ActionTmpSubdirTasks completes tmp:<subdir>:* tasks
func ActionTmpSubdirTasks() carapace.Action {
	return carapace.ActionValues("clear").Tag("tmp subdir tasks").UidF(taskUid("tmp.subdir"))
}
