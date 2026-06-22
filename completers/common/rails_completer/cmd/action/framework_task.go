package action

import "github.com/carapace-sh/carapace"

// ActionTimezoneTasks completes time:zones:* tasks
func ActionTimezoneTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"all", "List all time zones",
		"local", "List local time zones",
		"us", "List US time zones",
	).Tag("time zones tasks").UidF(taskUid("time.zones"))
}

// ActionLogTasks completes log:* tasks
func ActionLogTasks() carapace.Action {
	return carapace.ActionValues("clear").Tag("log tasks").UidF(taskUid("log"))
}

// ActionCredentialTasks completes credentials:* tasks
func ActionCredentialTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"edit", "Edit encrypted credentials",
		"show", "Show decrypted credentials",
		"diff", "Enroll/disenroll in credential diffs",
		"fetch", "Fetch a credential value",
	).Tag("credentials tasks").UidF(taskUid("credentials"))
}

// ActionEncryptedTasks completes encrypted:* tasks
func ActionEncryptedTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"edit", "Edit an encrypted file",
		"show", "Show an encrypted file",
	).Tag("encrypted tasks").UidF(taskUid("encrypted"))
}

// ActionDevTasks completes dev:* tasks
func ActionDevTasks() carapace.Action {
	return carapace.ActionValues("cache").Tag("dev tasks").UidF(taskUid("dev"))
}
