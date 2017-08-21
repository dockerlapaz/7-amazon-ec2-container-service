FROM scratch
COPY . /votacion/
WORKDIR /votacion
ENTRYPOINT ["/votacion/app"]