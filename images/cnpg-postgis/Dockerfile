# Keep this exact CNPG operand release as the base. The provider catalog pins the
# resulting image digest, so the extension layer is reproducible and immutable.
FROM ghcr.io/cloudnative-pg/postgresql:17.10-standard-trixie@sha256:78c4fdf165e8ffb1b5b7a7fc6b22b3cf37890a338b4c1c8c4d913896129a86da

USER root

RUN apt-get update \
    && apt-get install --yes --no-install-recommends \
        postgresql-17-postgis-3 \
        postgresql-17-postgis-3-scripts \
    && rm -rf /var/lib/apt/lists/*

# Match the uid used by the upstream CloudNativePG operand image.
USER 26
