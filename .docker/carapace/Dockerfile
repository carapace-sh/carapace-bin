FROM ghcr.io/carapace-sh/go:1.26.3

RUN curl -L https://github.com/goreleaser/goreleaser/releases/download/v2.20.2/goreleaser_Linux_x86_64.tar.gz \
  | tar -xvz goreleaser \
  && mv goreleaser /usr/local/bin/

RUN curl -fsSL https://github.com/carapace-sh/carapace-man/releases/download/v0.0.7/carapace-man_0.0.7_linux_amd64.tar.gz \
  | tar -xz -C /usr/local/bin
