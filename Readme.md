# Kubernetes Codelab

A simple dummy API written in Golang, deployed to GKE and managed by Kubernetes.

__Note__: `kubernetes_codelab/kubernetes/app.secret.yaml` should not be committed to source control -- it's simply added for demonstration purposes.

## Usage

```go
$ go run main.go
```

## Build

```
$ make build <push>
```

## Deploy

Once the initial image is pushed to GCR, you can update the image via a deploy:

```
# Bump version (TAG)
$ make deploy
```
