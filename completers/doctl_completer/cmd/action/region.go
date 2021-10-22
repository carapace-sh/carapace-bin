package action

import "github.com/rsteube/carapace"

func ActionRegions() carapace.Action {
	return carapace.ActionValuesDescribed(
		"ams2", "Amsterdam, the Netherlands",
		"ams3", "Amsterdam, the Netherlands",
		"blr1", "Bangalore, India",
		"fra1", "Frankfurt, Germany",
		"lon1", "London, United Kingdom",
		"nyc1", "New York City, United States",
		"nyc2", "New York City, United States",
		"nyc3", "New York City, United States",
		"sfo1", "San Francisco, United States",
		"sfo2", "San Francisco, United States",
		"sfo3", "San Francisco, United States",
		"sgp1", "Singapore",
		"tor1", "Toronto, Canada",
	)
}
