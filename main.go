//go:generate go run pkg/codegen/cleanup/main.go
//go:generate rm -rf pkg/generated pkg/crds/yaml/generated
//go:generate go run pkg/codegen/main.go
//go:generate controller-gen crd:generateEmbeddedObjectMeta=true paths=./k3s.cattle.io/... output:crd:dir=./pkg/crds/yaml/generated

package main

func main() {
	// nothing here - this exists to make `go generate` do the right thing.
}
