# Selfupdate

With `carapace --selfupdate` specific [nightly]/[stable] releases can be installed.

![](./selfupdate.cast)

Executable is installed to the [GOBIN] directory, essentially shadowing any system installation.


```sh
export PATH="$HOME/.local/bin:$HOME/go/bin:$PATH"
#            │                │            └system installation (e.g. /usr/bin/carapace)
#            │                └selfupdate/go based installation ($GOBIN)
#            └user binaries
```

## Requirements

- [curl] for downloads
- [PATH] containing the [GOBIN] directory

[nightly]:https://github.com/carapace-sh/nightly/releases
[stable]:https://github.com/carapace-sh/carapace-bin/releases

[curl]:https://curl.se
[GOBIN]:https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies
[PATH]:https://en.wikipedia.org/wiki/PATH_(variable)
