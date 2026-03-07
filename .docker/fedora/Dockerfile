FROM fedora

RUN dnf install -y elvish

RUN mkdir -p ~/.config/elvish \
 && echo -e "set paths = [ /carapace-bin/cmd/carapace \$@paths ]\neval (carapace _carapace|slurp)" > ~/.config/elvish/rc.elv
ENV PATH="/carapace-bin/cmd/carapace:$PATH"

ENTRYPOINT [ "elvish" ]
