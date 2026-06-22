package action

import "github.com/carapace-sh/carapace"

// ActionMailboxTasks completes action_mailbox:* tasks
func ActionMailboxTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"install", "Install Action Mailbox",
		"ingress", "Relay inbound email",
	).Tag("action mailbox tasks").UidF(taskUid("action_mailbox"))
}

// ActionMailboxIngressTasks completes action_mailbox:ingress:* tasks
func ActionMailboxIngressTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"exim", "Relay from Exim",
		"postfix", "Relay from Postfix",
		"qmail", "Relay from Qmail",
	).Tag("action mailbox ingress tasks").UidF(taskUid("action_mailbox.ingress"))
}

// ActionMailboxInstallTasks completes action_mailbox:install:* tasks
func ActionMailboxInstallTasks() carapace.Action {
	return carapace.ActionValues("migrations").Tag("action mailbox install tasks").UidF(taskUid("action_mailbox.install"))
}

// ActionTextInstallTasks completes action_text:install:* tasks
func ActionTextInstallTasks() carapace.Action {
	return carapace.ActionValues("migrations").Tag("action text install tasks").UidF(taskUid("action_text.install"))
}

// ActionStorageTasks completes active_storage:* tasks
func ActionStorageTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"install", "Install Active Storage",
		"update", "Update Active Storage migrations",
	).Tag("active storage tasks").UidF(taskUid("active_storage"))
}

// ActionAppTasks completes app:* tasks
func ActionAppTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"template", "Apply an application template",
		"update", "Update config for current Rails version",
	).Tag("app tasks").UidF(taskUid("app"))
}

// ActionAppNamespaces completes app:* sub-namespaces
func ActionAppNamespaces() carapace.Action {
	return carapace.ActionValues("templates").Tag("app namespaces").UidF(namespaceUid("app"))
}

// ActionAppTemplatesTasks completes app:templates:* tasks
func ActionAppTemplatesTasks() carapace.Action {
	return carapace.ActionValues("copy").Tag("app templates tasks").UidF(taskUid("app.templates"))
}

// ActionCacheDigestTasks completes cache_digests:* tasks
func ActionCacheDigestTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"dependencies", "Show template dependencies",
		"nested_dependencies", "Show nested dependencies",
	).Tag("cache digests tasks").UidF(taskUid("cache_digests"))
}
