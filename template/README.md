A custom metric provider for Drone that exposes additional system metrics. _Please note this project requires Drone server version 1.4 or higher._

## Installation

Create a bearer token:

```console
$ openssl rand -hex 16
bea26a2221fd8090ea38720fc445eca6
```

Download and run the provider:

```console
$ docker run -d \
  --publish=3000:3000 \
  --env=DRONE_DEBUG=true \
  --env=DRONE_TOKEN=bea26a2221fd8090ea38720fc445eca6 \
  --env=DRONE_DATABASE_DRIVER=postgres \
  --env=DRONE_DATABASE_DATASOURCE=postgres://... \
  --restart=always \
  --name=metrics {{DockerRepository}}
```

Configure Prometheus:

```text
global:
    scrape_interval: 60m

    scrape_configs:
    - job_name: 'custom_drone_metrics'
        bearer_token: bea26a2221fd8090ea38720fc445eca6

        static_configs:
        - targets: ['domain.com:3000']
```