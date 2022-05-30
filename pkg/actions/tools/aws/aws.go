// package aws contains amazon web services related actions
package aws

import (
	"strings"

	"github.com/rsteube/carapace"
	"gopkg.in/ini.v1"
)

// ActionRegions completes region names
//   eu-south-1 (Europe Milan)
//   ap-east-1 (Asia Pacific Hong Kong)
func ActionRegions() carapace.Action {
	return carapace.ActionValuesDescribed(
		"af-south-1", "Africa (Cape Town)",
		"ap-east-1", "Asia Pacific (Hong Kong)",
		"ap-northeast-1", "Asia Pacific (Tokyo)",
		"ap-northeast-2", "Asia Pacific (Seoul)",
		"ap-south-1", "Asia Pacific (Mumbai)",
		"ap-southeast-1", "Asia Pacific (Singapore)",
		"ap-southeast-2", "Asia Pacific (Sydney)",
		"ca-central-1", "Canada (Central)",
		"eu-central-1", "Europe (Frankfurt)",
		"eu-north-1", "Europe (Stockholm)",
		"eu-south-1", "Europe (Milan)",
		"eu-west-1", "Europe (Ireland)",
		"eu-west-2", "Europe (London)",
		"eu-west-3", "Europe (Paris)",
		"me-south-1", "Middle East (Bahrain)",
		"sa-east-1", "South America (Sao Paulo)",
		"us-east-1", "US East (N. Virginia)",
		"us-east-2", "US East (Ohio)",
		"us-west-1", "US West (N. California)",
		"us-west-2", "US West (Oregon)",
	)
}

// ActionProfiles completes configuration profile names
//   someprofile (eu-central-1)
//   anotherprofile (us-east-1)
func ActionProfiles() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		profiles := []string{}

		// TODO support windows
		if path, err := c.Abs("~/.aws/config"); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			if cfg, err := ini.Load(path); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				for _, section := range cfg.Sections() {
					if strings.HasPrefix(section.Name(), "profile ") {
						profiles = append(profiles, strings.TrimPrefix(section.Name(), "profile "))
						if key, err := section.GetKey("region"); err != nil {
							profiles = append(profiles, "")
						} else {
							profiles = append(profiles, key.String())
						}
					}
				}
				if len(profiles) == 0 {
					profiles = append(profiles, "default", "")
				}
				return carapace.ActionValuesDescribed(profiles...)
			}
		}
	})
}
