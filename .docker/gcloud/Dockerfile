FROM gcr.io/google.com/cloudsdktool/google-cloud-cli:stable

RUN apt-get update && apt-get install -y curl
RUN curl -so - https://dl.elv.sh/linux-amd64/elvish-HEAD.tar.gz | tar -xzvC /usr/local/bin

RUN mkdir -p ~/.config/elvish \
 && echo "set paths = [ /carapace-bin/cmd/carapace \$@paths ]\neval (carapace _carapace|slurp)" > ~/.config/elvish/rc.elv
ENV PATH="/carapace-bin/cmd/carapace:$PATH"

CMD ["elvish"]
