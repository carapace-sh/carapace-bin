FROM alpine

RUN apk add gcompat
RUN apk add --no-cache --repository=https://dl-cdn.alpinelinux.org/alpine/edge/community elvish

RUN mkdir -p ~/.config/elvish \
 && echo -e "set paths = [ /carapace-bin/cmd/carapace \$@paths ]\neval (carapace _carapace|slurp)" > ~/.config/elvish/rc.elv
ENV PATH="/carapace-bin/cmd/carapace:$PATH"

CMD ["elvish"]
