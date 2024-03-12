module github.com/carapace-sh/carapace-bin

go 1.21

require (
	github.com/carapace-sh/carapace v0.50.3-0.20240311152955-f3bd62ff6a23
	github.com/carapace-sh/carapace-bridge v0.2.16-0.20240311173237-760172108463
	github.com/carapace-sh/carapace-shlex v1.0.0
	github.com/carapace-sh/carapace-spec v0.17.1-0.20240311163627-73cd85d39b71
	github.com/pelletier/go-toml v1.9.5
	github.com/spf13/cobra v1.8.0
	github.com/spf13/pflag v1.0.5
	golang.org/x/mod v0.16.0
	gopkg.in/ini.v1 v1.67.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/invopop/jsonschema v0.12.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/wk8/go-ordered-map/v2 v2.1.8 // indirect
)

replace github.com/spf13/pflag => github.com/carapace-sh/carapace-pflag v0.2.0
