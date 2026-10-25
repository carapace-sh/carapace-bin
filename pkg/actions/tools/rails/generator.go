package rails

import "github.com/carapace-sh/carapace"

// ActionGenerators completes rails generate generator names
//
//	model (Generate a model)
//	scaffold (Generate scaffold (model + controller + views))
func ActionGenerators() carapace.Action {
	return carapace.ActionValuesDescribed(
		"action_mailbox:install", "Install Action Mailbox",
		"action_text:install", "Install Action Text",
		"active_record:authentication", "Generate authentication model",
		"active_record:migration", "Generate an Active Record migration",
		"application_record", "Generate ApplicationRecord",
		"benchmark", "Generate a benchmark",
		"channel", "Generate an Action Cable channel",
		"controller", "Generate a controller",
		"generator", "Generate a custom generator",
		"helper", "Generate a helper",
		"integration_test", "Generate an integration test",
		"jbuilder", "Generate Jbuilder views",
		"job", "Generate an Active Job",
		"mailbox", "Generate an Action Mailbox",
		"mailer", "Generate a mailer",
		"migration", "Generate a database migration",
		"model", "Generate a model",
		"resource", "Generate resource (model + controller, no views)",
		"scaffold", "Generate scaffold (model + controller + views)",
		"scaffold_controller", "Generate scaffold controller",
		"system_test", "Generate a system test",
		"task", "Generate a rake task",
	).Tag("generators").Uid("rails", "generator")
}
