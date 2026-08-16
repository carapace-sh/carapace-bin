// package devenv contains devenv related actions
package devenv

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/cache/key"
	"github.com/carapace-sh/carapace/pkg/util"
)

// GlobalOpts contains global devenv flags that affect how devenv.nix is evaluated.
type GlobalOpts struct {
	// From is the source for devenv.nix
	From string

	// Impure relaxes the hermeticity of the environment
	Impure bool

	// Options are `<attribute>:<type> <value>` pairs overriding configuration options
	Options []string

	// OverrideInputs are `<name> <uri>` pairs overriding inputs of devenv.yaml
	OverrideInputs []string

	// Profiles are the profiles to activate
	Profiles []string

	// System overrides the target system
	System string
}

func (o GlobalOpts) args() []string {
	args := make([]string, 0)
	if o.From != "" {
		args = append(args, "--from", o.From)
	}
	if o.Impure {
		args = append(args, "--impure")
	}
	for index := 0; index+1 < len(o.Options); index += 2 {
		args = append(args, "--option", o.Options[index], o.Options[index+1])
	}
	for index := 0; index+1 < len(o.OverrideInputs); index += 2 {
		args = append(args, "--override-input", o.OverrideInputs[index], o.OverrideInputs[index+1])
	}
	for _, profile := range o.Profiles {
		args = append(args, "--profile", profile)
	}
	if o.System != "" {
		args = append(args, "--system", o.System)
	}
	return args
}

func (o GlobalOpts) cacheKey() key.Key {
	return key.String(o.args()...)
}

// command prepends the global flags to given devenv arguments
func (o GlobalOpts) command(arg ...string) []string {
	return append(o.args(), arg...)
}

// configFile returns the devenv.nix the current directory belongs to
func configFile(c carapace.Context) (string, error) {
	return util.FindReverse(c.Dir, "devenv.nix")
}
