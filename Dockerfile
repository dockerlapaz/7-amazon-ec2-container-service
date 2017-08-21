FROM scratch
WORKDIR /votacion
COPY votacion.html /votacion/
COPY static /votacion/static/
COPY app /votacion/

ENTRYPOINT ["/votacion/app"]