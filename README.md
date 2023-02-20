# awa-monitoring

## Requirements

- Docker

- Go 1.20

## Set up

- Create `.env` file

```
DOCKER_INFLUXDB_INIT_MODE=setup
DOCKER_INFLUXDB_INIT_USERNAME=changeme
DOCKER_INFLUXDB_INIT_PASSWORD=changeme
DOCKER_INFLUXDB_INIT_ADMIN_TOKEN=changeme
DOCKER_INFLUXDB_INIT_ORG=changeme
DOCKER_INFLUXDB_INIT_BUCKET=changeme
DOCKER_INFLUXDB_INIT_RETENTION=4d
DOCKER_INFLUXDB_INIT_PORT=8086
DOCKER_INFLUXDB_INIT_HOST=influxdb
```

**Note**: Run the code snippet below to create `DOCKER_INFLUXDB_INIT_ADMIN_TOKEN`:

```
$ openssl rand -hex 32
```
