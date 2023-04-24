package mvn

import "github.com/rsteube/carapace"

// ActionProfiles completes profiles
func ActionProfiles(file string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if loadedProject, err := loadProject(file); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return carapace.ActionValues(loadedProject.Profiles...)
		}
	})
}
