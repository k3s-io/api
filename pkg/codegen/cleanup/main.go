package main

import (
	"github.com/rancher/wrangler/v3/pkg/cleanup"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := cleanup.Cleanup("./k3s.cattle.io"); err != nil {
		logrus.Fatal(err)
	}
}
