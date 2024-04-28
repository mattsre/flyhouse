# flyhouse

Flyhouse is a CLI for deploying Clickhouse clusters on Fly.io. Cluster configuration is written in [Pkl](https://pkl-lang.org/)

## Getting Started

- Download Flyhouse
- Run `flyhouse login` to authorize to the Fly.io API
- Run `flyhouse deploy -m cluster/` to deploy manifests in the `cluster/` directory
