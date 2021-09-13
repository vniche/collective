FROM debian:buster

RUN apt-get update \
    && apt-get install -y ca-certificates --no-install-recommends \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* \
    && update-ca-certificates 

RUN	addgroup vniche \
    && adduser --disabled-password --gecos "" vniche --uid 1000 --ingroup vniche

COPY --chown=vnihe:vniche collective /usr/local/bin/
RUN chmod +x /usr/local/bin/collective

USER vniche

ENTRYPOINT [ "/usr/local/bin/collective" ]