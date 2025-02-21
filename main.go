//go:generate go run pkg/codegen/cleanup/main.go
//go:generate rm -rf pkg/generated
//go:generate go run pkg/codegen/main.go

package main

func main() {
	// nothing here - this exists to make `go generate` do the right thing.
}
