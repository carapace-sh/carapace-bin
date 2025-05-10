package apk

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type PackageSearchOpts struct {
	Arch             string
	From             string
	KeysDir          string
	RepositoriesFile string
	Repository       string
}

func (o PackageSearchOpts) args() []string {
	args := make([]string, 0)
	if o.Arch != "" {
		args = append(args, "--arch", o.Arch)
	}
	if o.From != "" {
		args = append(args, "--from", o.From)
	}
	if o.KeysDir != "" {
		args = append(args, "--keys-dir", o.KeysDir)
	}
	if o.RepositoriesFile != "" {
		args = append(args, "--repositories-file", o.RepositoriesFile)
	}
	if o.Repository != "" {
		args = append(args, "--repository", o.Repository)
	}
	return args
}

// ActionPackageSearch completes installable packages
//
//	2bwm (0.3-r2)
//	2bwm-doc (0.3-r2)
func ActionPackageSearch(opts PackageSearchOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := append([]string{"search"}, opts.args()...)
		return carapace.ActionExecCommand("apk", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			r := regexp.MustCompile(`^(?P<package>.*)-(?P<version>[^-]+-r\d+)$`)

			vals := make([]string, 0)
			for _, line := range lines {
				if matches := r.FindStringSubmatch(line); matches != nil && strings.HasPrefix(matches[1], c.Value) {
					vals = append(vals, matches[1], matches[2])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

type PackageOpts struct {
	InstalledOnly bool

	Arch             string
	From             string
	KeysDir          string
	RepositoriesFile string
	Repository       string
}

func (o PackageOpts) args() []string {
	args := make([]string, 0)
	if o.Arch != "" {
		args = append(args, "--arch", o.Arch)
	}
	if o.From != "" {
		args = append(args, "--from", o.From)
	}
	if o.KeysDir != "" {
		args = append(args, "--keys-dir", o.KeysDir)
	}
	if o.RepositoriesFile != "" {
		args = append(args, "--repositories-file", o.RepositoriesFile)
	}
	if o.Repository != "" {
		args = append(args, "--repository", o.Repository)
	}
	return args
}

// ActionPackages completes packages
// TODO is this actually any different from ActionPackageSearch apart from the `[installed]` suffix in the output?
//
//	2bwm (0.3-r2)
//	2bwm-doc (0.3-r2)
func ActionPackages(opts PackageOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := append([]string{"list"}, opts.args()...)
		args = append(args, c.Value+"*")
		return carapace.ActionExecCommand("apk", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			r := regexp.MustCompile(`^(?P<package>.*)-(?P<version>[^-]+-r\d+) (?P<arch>[^ ]+) {(?P<group>[^ ]+)} (?P<license>\(.+\))( \[(?P<installed>installed)\])?$`)

			vals := make([]string, 0)
			for _, line := range lines {
				if matches := r.FindStringSubmatch(line); matches != nil && strings.HasPrefix(matches[1], c.Value) {
					if opts.InstalledOnly && matches[7] != "installed" {
						continue
					}

					switch matches[7] {
					case "installed":
						vals = append(vals, matches[1], matches[2], style.Blue)
					default:
						vals = append(vals, matches[1], matches[2], style.Default)
					}
				}
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}
