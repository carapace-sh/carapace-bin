FROM archlinux

RUN pacman -Sy --noconfirm nix
RUN pacman -Sy --noconfirm elvish

RUN mkdir -p ~/.config/elvish \
&&  echo -e "set paths = [ /carapace-bin/cmd/carapace \$@paths ]\neval (carapace _carapace|slurp)" > ~/.config/elvish/rc.elv
RUN export PATH="/carapace-bin/cmd/carapace:$PATH"

CMD ["elvish"]
