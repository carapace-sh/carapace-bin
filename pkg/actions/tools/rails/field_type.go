package rails

import "github.com/carapace-sh/carapace"

// ActionFieldTypes completes field type names for generators
//
//	string
//	text
func ActionFieldTypes() carapace.Action {
	return carapace.ActionValues(
		"string",
		"text",
		"integer",
		"bigint",
		"float",
		"decimal",
		"numeric",
		"datetime",
		"timestamp",
		"time",
		"date",
		"binary",
		"boolean",
		"references",
		"belongs_to",
		"digest",
		"token",
		"json",
		"jsonb",
	).Tag("field types").Uid("rails", "field-type")
}
