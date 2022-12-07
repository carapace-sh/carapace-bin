package newrelic

import (
	"fmt"
	"strconv"

	"github.com/rsteube/carapace"
)

// ActionApplicationIds completes application ids
//
//	111111111 (application1)
//	222222222 (application2)
func ActionApplicationIds(profile string) carapace.Action {
	return carapace.ActionCallback(func(carapace.Context) carapace.Action {

		profiles, err := getProfiles()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		// TODO filter by applicationId prefix
		query := fmt.Sprintf(`
{
  actor {
    entitySearch(query: "accountId = %v AND domainType IN ('APM-APPLICATION')") {
      results {
        entities {
          ... on ApmApplicationEntityOutline {
            applicationId
            name
          }
        }
      }
    }
  }
}`, profiles[profile].AccountID)

		type response struct {
			Actor struct {
				EntitySearch struct {
					Results struct {
						Entities []struct {
							ApplicationID int    `json:"applicationId"`
							Name          string `json:"name"`
						} `json:"entities"`
					} `json:"results"`
				} `json:"entitySearch"`
			} `json:"actor"`
		}
		return actionNerdGraph(query, func(r response) carapace.Action {
			vals := make([]string, 0)
			for _, entity := range r.Actor.EntitySearch.Results.Entities {
				vals = append(vals, strconv.Itoa(entity.ApplicationID), entity.Name)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

// ActionApplicationGuids completes application guids
//
//	AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA (application1)
//	BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB (application2)
func ActionApplicationGuids(profile string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		profiles, err := getProfiles()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		query := fmt.Sprintf(`			
{
  actor {
    entitySearch(query: "accountId = %v AND id LIKE '%v%%' AND domainType IN ('APM-APPLICATION')") {
      results {
        entities {
          name
          guid
        }
      }
    }
  }
}`, profiles[profile].AccountID, c.CallbackValue)

		type response struct {
			Actor struct {
				EntitySearch struct {
					Results struct {
						Entities []struct {
							GUID string `json:"guid"`
							Name string `json:"name"`
						} `json:"entities"`
					} `json:"results"`
				} `json:"entitySearch"`
			} `json:"actor"`
		}

		return actionNerdGraph(query, func(r response) carapace.Action {
			vals := make([]string, 0)
			for _, entity := range r.Actor.EntitySearch.Results.Entities {
				vals = append(vals, entity.GUID, entity.Name)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
