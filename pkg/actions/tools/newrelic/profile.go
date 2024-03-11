package newrelic

import (
	"fmt"
	"strconv"

	"github.com/carapace-sh/carapace"
)

// ActionAccountIds completes account ids
//
//	123456 (eu - profile1)
//	234567 (us - profile2)
func ActionAccountIds() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		profiles, err := getProfiles()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for name, profile := range profiles {
			if name != "" {
				vals = append(vals, strconv.Itoa(profile.AccountID), fmt.Sprintf("%v - %v", profile.Region, name))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionProfiles completes profiles
//
//	profile1 (eu - 123456)
//	profile2 (us - 234567)
func ActionProfiles() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		profiles, err := getProfiles()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for name, profile := range profiles {
			if name != "" {
				vals = append(vals, name, fmt.Sprintf("%v - %v", profile.Region, profile.AccountID))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
