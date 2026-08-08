FROM golang
ARG VERSION

WORKDIR /
RUN git clone https://github.com/balajz/pgxcli --branch "${VERSION}" --depth 1

ADD pgxcli.patch /

WORKDIR /pgxcli
RUN git apply ../pgxcli.patch
RUN go get github.com/carapace-sh/carapace
RUN GOBIN=/usr/bin/ go install ./cmd/pgxcli

ENTRYPOINT ["pgxcli", "_carapace", "spec"]