# Provider DigitalOcean

`provider-digitalocean` is a [Crossplane](https://crossplane.io/) provider
that is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the DigitalOcean
API.

## Getting Started

Install using declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-minio
spec:
  package: ghcr.io/eaglesemanation/provider-digitalocean:v0.1.0
EOF
```

You can see the API reference here: https://doc.crds.dev/github.com/eaglesemanation/provider-digitalocean@v0.1.0

## Developing

Run code-generation pipeline:
```console
make generate
```

Test an example against a Kind k8s cluster (provider config is included through setup.sh):
```console
make e2e UPTEST_EXAMPLE_LIST="examples/namespaced/droplet/droplet.yaml,examples/namespaced/reservedip/ipassignment.yaml,examples/namespaced/reservedip/ip.yaml,examples/namespaced/reservedipv6/ipv6assignment.yaml,examples/namespaced/reservedipv6/ipv6.yaml"
kind delete cluster -n local-dev
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/eaglesemanation/provider-digitalocean/issues).
