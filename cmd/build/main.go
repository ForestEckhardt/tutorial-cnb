package main

import (
	"github.com/ForestEckhardt/cnb-tutorial/node"
	"github.com/cloudfoundry/packit"
)

func main() {
	packit.Build(node.Build())
}
