package action

import "github.com/carapace-sh/carapace"

// ActionEventSubscriptionFields completes event subscription fields
func ActionEventSubscriptionFields() carapace.Action {
	return carapace.ActionValues(
		"createdAt",
		"eventTypes",
		"eventsUrl",
		"expiresAt",
		"id",
		"resourceUrl",
		"url",
	)
}
