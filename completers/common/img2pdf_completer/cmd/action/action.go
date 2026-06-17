package action

import "github.com/carapace-sh/carapace"

func ActionRotations() carapace.Action {
	return carapace.ActionValues(
		"0",
		"180",
		"270",
		"90",
		"auto",
		"ifvalid",
		"none",
	).Tag("rotations")
}

func ActionPaperSizes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"A0", "841mmx1189mm",
		"A1", "594mmx841mm",
		"A2", "420mmx594mm",
		"A3", "297mmx420mm",
		"A4", "210mmx297mm",
		"A5", "148mmx210mm",
		"A6", "105mmx148mm",
		"B0", "1000mmx1414mm",
		"B1", "707mmx1000mm",
		"B2", "500mmx707mm",
		"B3", "353mmx500mm",
		"B4", "250mmx353mm",
		"B5", "176mmx250mm",
		"B6", "125mmx176mm",
		"JB0", "1030mmx1456mm",
		"JB1", "728mmx1030mm",
		"JB2", "515mmx728mm",
		"JB3", "364mmx515mm",
		"JB4", "257mmx364mm",
		"JB5", "182mmx257mm",
		"JB6", "128mmx182mm",
		"Legal", "8.5inx14in",
		"Letter", "8.5inx11in",
		"Tabloid", "11inx17in",
	).Usage("append ^T for landscape (case insensitive)").
		Tag("paper sizes")
}
