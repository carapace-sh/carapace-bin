# Environment

Complex environment variable completion is provided with `get-env`, `set-env` and `unset-env`.

In `elvish` the completion is simply overridden.
For other shells custom functions are added.

> Setting the environment variable `CARAPACE_ENV` to `0`
> before sourcing `carapace _carapace` disables this behaviour.

![](./environment.cast)

## Custom variables

Custom variables can be defined in `~/.config/carapace/env.yaml`

```yaml
names:
  CUSTOM_EXAMPLE: example environment variable
  CUSTOM_MACRO: macro example
  HTTPS_PROXY: override existing variable
completion:
  CUSTOM_EXAMPLE: ["0\tdisabled\tred", "1\tenabled\tgreen"]
  CUSTOM_MACRO: ["$_tools.gh.Labels({owner: rsteube, name: carapace}) ||| $uniquelist(,)"]
  HTTPS_PROXY: ["https://localhost:8443\tdevelopment", "https://proxy.company:443\tproduction"]
```

![](./environment-custom.cast)
