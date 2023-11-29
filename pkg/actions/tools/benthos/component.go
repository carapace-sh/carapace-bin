package benthos

import (
	"encoding/json"

	"github.com/rsteube/carapace"
)

type components struct {
	BloblangFunctions []string `json:"bloblang-functions"`
	BloblangMethods   []string `json:"bloblang-methods"`
	Buffers           []string `json:"buffers"`
	Caches            []string `json:"caches"`
	Inputs            []string `json:"inputs"`
	Metrics           []string `json:"metrics"`
	Outputs           []string `json:"outputs"`
	Processors        []string `json:"processors"`
	RateLimits        []string `json:"rate-limits"`
	Tracers           []string `json:"tracers"`
}

func actionComponents(f func(c components) carapace.Action) carapace.Action {
	return carapace.ActionExecCommand("benthos", "list", "--format", "json")(func(output []byte) carapace.Action {
		var c components
		if err := json.Unmarshal(output, &c); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return f(c)
	})
}

// ActionBoblangFunctions completes boblang functions
func ActionBoblangFunctions() carapace.Action {
	return actionComponents(func(c components) carapace.Action {
		return carapace.ActionValues(c.BloblangFunctions...)
	}).Tag("boblang functions")
}

// ActionBoblangMethods completes boblang methods
func ActionBoblangMethods() carapace.Action {
	return actionComponents(func(c components) carapace.Action {
		return carapace.ActionValues(c.BloblangMethods...)
	}).Tag("boblang methods")
}

// ActionBuffers completes buffers
func ActionBuffers() carapace.Action {
	return actionComponents(func(c components) carapace.Action {
		return carapace.ActionValues(c.Buffers...)
	}).Tag("buffers")
}

// ActionCaches completes caches
func ActionCaches() carapace.Action {
	return actionComponents(func(c components) carapace.Action {
		return carapace.ActionValues(c.Caches...)
	}).Tag("caches")
}

// ActionInputs completes inputs
func ActionInputs() carapace.Action {
	return actionComponents(func(c components) carapace.Action {
		return carapace.ActionValues(c.Inputs...)
	}).Tag("inputs")
}

// ActionMetrics completes metrics
func ActionMetrics() carapace.Action {
	return actionComponents(func(c components) carapace.Action {
		return carapace.ActionValues(c.Metrics...)
	}).Tag("metrics")
}

// ActionOutputs completes outputs
func ActionOutputs() carapace.Action {
	return actionComponents(func(c components) carapace.Action {
		return carapace.ActionValues(c.Outputs...)
	}).Tag("outputs")
}

// ActionProcessors completes processors
func ActionProcessors() carapace.Action {
	return actionComponents(func(c components) carapace.Action {
		return carapace.ActionValues(c.Processors...)
	}).Tag("processors")
}

// ActionRateLimits completes rate limits
func ActionRateLimits() carapace.Action {
	return actionComponents(func(c components) carapace.Action {
		return carapace.ActionValues(c.RateLimits...)
	}).Tag("rate limits")
}

// ActionTracers completes tracers
func ActionTracers() carapace.Action {
	return actionComponents(func(c components) carapace.Action {
		return carapace.ActionValues(c.Tracers...)
	}).Tag("tracers")
}
