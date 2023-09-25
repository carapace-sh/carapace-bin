# Variable

> feature in development

Complex environment variable completion is provided with `get-env`, `set-env` and `unset-env`.

In `elvish` the completion is simply overridden.
For other shells custom functions are added.

> Setting `CARAPACE_ENV=0` before sourcing `carapace _carapace` disables this behaviour.

![](./variable.cast)

## Custom variables

Custom variables can be defined in `~/.config/carapace/variables/{group}.yaml`

```yaml
variables:
  CUSTOM_EXAMPLE: example environment variable
  CUSTOM_MACRO: macro example
  HTTPS_PROXY: override existing variable
completion:
  variable:
    CUSTOM_EXAMPLE: ["0\tdisabled\tred", "1\tenabled\tgreen"]
    CUSTOM_MACRO: ["$_tools.gh.Labels({owner: rsteube, name: carapace}) ||| $uniquelist(,)"]
    HTTPS_PROXY: ["https://localhost:8443\tdevelopment", "https://proxy.company:443\tproduction"]
```

![](./variable-custom.cast)

It is also possible to define conditions.

```yaml
condition: ["$Parent([.git])"]
variables:
  CUSTOM_CONDITION: condition example
completion:
  variable:
    CUSTOM_CONDITION: ["within", "git", "repo"]
```

![](./variable-condition.cast)
