# Kubernetes Codelab

A simple dummy API written in Golang, deployed to GKE and managed by Kubernetes.

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
# Bump version
$ make deploy
```
