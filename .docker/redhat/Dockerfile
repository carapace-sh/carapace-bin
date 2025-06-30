FROM registry.access.redhat.com/ubi10/ubi

RUN curl -so - https://dl.elv.sh/linux-amd64/elvish-HEAD.tar.gz | tar -xzvC /usr/local/bin

RUN mkdir -p ~/.config/elvish \
 && echo -e "set paths = [ /carapace-bin/cmd/carapace \$@paths ]\neval (carapace _carapace|slurp)" > ~/.config/elvish/rc.elv
ENV PATH="/carapace-bin/cmd/carapace:$PATH"

ENTRYPOINT [ "elvish" ]
