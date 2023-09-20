package env

import "testing"

func TestKnownVariables(t *testing.T) {
	for k, v := range knownVariables {
		for name := range v.Completion {
			if _, ok := v.Names[name]; !ok {
				t.Errorf("variables %#v is unknown in %#v", name, k)
			}
		}
	}
}
