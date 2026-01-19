FROM gcr.io/distroless/static-debian12:nonroot

# Copie avec ownership vers nonroot (UID/GID 65532)
COPY --chown=65532:65532 bin/hugo-mx-gateway LICENSE /app/
COPY --chown=65532:65532 templates /app/templates

WORKDIR /app

# User nonroot prédéfini dans l'image distroless
USER 65532:65532

# Exécution directe du binaire
ENTRYPOINT ["/app/hugo-mx-gateway"]
