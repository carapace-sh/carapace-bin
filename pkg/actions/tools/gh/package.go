package gh

import (
	"fmt"

	"github.com/carapace-sh/carapace"
)

type PackageOpts struct {
	Host string
	// Owner is an organization or user name
	Owner string
	// Organization determines which endpoint is used
	Organization bool
	// PackageType filters by package type
	Type string
	// Name of the package for version listings
	Name string
}

func (o PackageOpts) Default() PackageOpts {
	return o
}

func (o PackageOpts) listRoute() string {
	base := fmt.Sprintf("users/%v/packages", o.Owner)
	if o.Organization {
		base = fmt.Sprintf("orgs/%v/packages", o.Owner)
	}
	if o.Type != "" {
		base = fmt.Sprintf("%v?package_type=%v", base, o.Type)
	}
	return base
}

func (o PackageOpts) versionsRoute() string {
	if o.Organization {
		return fmt.Sprintf("orgs/%v/packages/%v/%v/versions", o.Owner, o.Type, o.Name)
	}
	return fmt.Sprintf("users/%v/packages/%v/%v/versions", o.Owner, o.Type, o.Name)
}

// ActionPackages completes packages of an organization or user
//
//	demo-package (container, 3 versions)
func ActionPackages(opts PackageOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult []struct {
			Name         string
			PackageType  string `json:"package_type"`
			VersionCount int    `json:"version_count"`
		}
		return apiV3Action(RepoOpts{Host: opts.Host}, opts.listRoute(), &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, p := range queryResult {
				vals = append(vals, p.Name, fmt.Sprintf("%v (%v versions)", p.PackageType, p.VersionCount))
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	}).Tag("packages")
}

// ActionPackageVersions completes package versions
//
//	1234567 (sha256:0123acf...)
func ActionPackageVersions(opts PackageOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult []struct {
			Id        int
			Name      string
			Version   string
			CreatedAt string `json:"created_at"`
		}
		return apiV3Action(RepoOpts{Host: opts.Host}, opts.versionsRoute(), &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, version := range queryResult {
				label := version.Version
				if label == "" {
					label = version.Name
				}
				vals = append(vals, fmt.Sprintf("%v", version.Id), label)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	}).Tag("package versions")
}
