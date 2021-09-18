// package jaeger contains jaeger related actions
package jaeger

import (
	"github.com/rsteube/carapace"
)

// ActionSamplingTypes completes sampling types
//   const (sampler always makes the same decision for all traces)
//   probabilistic (sampler makes a random sampling decision with the probability of sampling equ...)
func ActionSamplingTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"const", "sampler always makes the same decision for all traces",
		"probabilistic", "sampler makes a random sampling decision with the probability of sampling equal to the value of sampler.param property",
		"ratelimiting", "sampler uses a leaky bucket rate limiter to ensure that traces are sampled with a certain constant rate",
		"remote", "sampler consults Jaeger agent for the appropriate sampling strategy to use in the current service",
	)
}
