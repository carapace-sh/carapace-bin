FROM ubuntu

RUN apt-get update \
 && apt-get install -y bash curl elvish git scala

RUN curl -fL "https://github.com/coursier/launchers/raw/master/cs-x86_64-pc-linux.gz" \
  | gzip -d > /usr/bin/cs \
 && chmod +x /usr/bin/cs

RUN cs setup -y
RUN cs install bloop --only-prebuilt=true

ENV PATH="/root/.local/share/coursier/bin:$PATH"
RUN bloop

RUN mkdir -p ~/.config/elvish \
 && echo "set paths = [ /carapace-bin/cmd/carapace \$@paths ]\neval (carapace _carapace|slurp)" > ~/.config/elvish/rc.elv
ENV PATH="/carapace-bin/cmd/carapace:$PATH"

ENTRYPOINT [ "elvish" ]
