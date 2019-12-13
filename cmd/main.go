package main

import (
	"flag"

	"github.com/moby/buildkit/frontend/gateway/grpcclient"
	"github.com/moby/buildkit/util/appcontext"
	samplefile "samplefile"
)

var (
	filename string
	graph    bool
)

func main() {
	flag.StringVar(&filename, "f", "Samplefile.yaml", "read")
	flag.Parse()

	if err := grpcclient.RunFromEnvironment(appcontext.Context(), samplefile.Build); err != nil {
		panic(err)
	}
}
