module github.com/rsteube/carapace-bin

go 1.19

require (
	github.com/pelletier/go-toml v1.9.5
	github.com/rsteube/carapace v0.42.1
	github.com/rsteube/carapace-bridge v0.1.4
	github.com/rsteube/carapace-shlex v0.0.4
	github.com/rsteube/carapace-spec v0.9.1
	github.com/spf13/cobra v1.7.0
	github.com/spf13/pflag v1.0.5
	golang.org/x/mod v0.12.0
	gopkg.in/ini.v1 v1.67.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/iancoleman/orderedmap v0.0.0-20190318233801-ac98e3ecb4b0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/invopop/jsonschema v0.7.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
)

replace github.com/spf13/pflag => github.com/rsteube/carapace-pflag v0.2.0
