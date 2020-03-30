package main

import (
	"github.com/ForestEckhardt/cnb-tutorial/node"
	"github.com/cloudfoundry/packit"
)

func main() {
	packit.Detect(node.Detect())
}
