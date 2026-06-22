package rails

import "github.com/carapace-sh/carapace"

// ActionGenerators completes rails generate generator names
//
//	model (Generate a model)
//	scaffold (Generate scaffold (model + controller + views))
func ActionGenerators() carapace.Action {
	return carapace.ActionValuesDescribed(
		"model", "Generate a model",
		"scaffold", "Generate scaffold (model + controller + views)",
		"resource", "Generate resource (model + controller, no views)",
		"controller", "Generate a controller",
		"migration", "Generate a database migration",
		"job", "Generate an Active Job",
		"mailer", "Generate a mailer",
		"channel", "Generate an Action Cable channel",
		"mailbox", "Generate an Action Mailbox",
		"task", "Generate a rake task",
		"helper", "Generate a helper",
		"system_test", "Generate a system test",
		"integration_test", "Generate an integration test",
		"jbuilder", "Generate Jbuilder views",
		"generator", "Generate a custom generator",
		"application_record", "Generate ApplicationRecord",
		"benchmark", "Generate a benchmark",
		"scaffold_controller", "Generate scaffold controller",
		"active_record:migration", "Generate an Active Record migration",
		"active_record:authentication", "Generate authentication model",
		"action_text:install", "Install Action Text",
		"action_mailbox:install", "Install Action Mailbox",
	).Tag("generators").Uid("rails", "generator")
}
