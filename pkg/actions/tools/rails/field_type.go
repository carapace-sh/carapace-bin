package rails

import "github.com/carapace-sh/carapace"

// ActionFieldTypes completes field type names for generators
//
//	string
//	text
func ActionFieldTypes() carapace.Action {
	return carapace.ActionValues(
		"belongs_to",
		"bigint",
		"binary",
		"boolean",
		"date",
		"datetime",
		"decimal",
		"digest",
		"float",
		"integer",
		"json",
		"jsonb",
		"numeric",
		"references",
		"string",
		"text",
		"time",
		"timestamp",
		"token",
	).Tag("field types").Uid("rails", "field-type")
}
