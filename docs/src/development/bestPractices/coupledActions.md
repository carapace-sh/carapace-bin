# Coupled Actions

Use coupled actions to avoid repetition.

Sometimes an [Action] depends on the same flag values for multiple subcommands.
Since [ActionCallback] is needed to access these the code can become a bit cumbersome and bloated.

```go
carapace.Gen(get_allCmd).PositionalCompletion(
	carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return helm.ActionReleases(helm.ReleasesOpts{
			Namespace:   rootCmd.Flag("namespace").Value.String(),
			KubeContext: rootCmd.Flag("kube-context").Value.String(),
		})
	}),
)
```

An alternative to this is creating a local [Action] that is coupled to the command.
Meaning, passing it the command and expecting specific flags to be present.

```go
// completers/helm_completer/cmd/action/release.go
func ActionReleases(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return helm.ActionReleases(helm.ReleasesOpts{
			Namespace:   cmd.Root().Flag("namespace").Value.String(),
			KubeContext: cmd.Root().Flag("kube-context").Value.String(),
		})
	})
}
```

Thus the call becomes quite compact.

```go
carapace.Gen(get_allCmd).PositionalCompletion(
	action.ActionReleases(get_allCmd),
)
```

[Action]:https://carapace-sh.github.io/carapace/carapace/action.html
[ActionCallback]:https://carapace-sh.github.io/carapace/carapace/defaultActions/actionCallback.html
