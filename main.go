package main

import (
	"flag"
	"myhttp/httpclient"
)

const DEFAULT_MAX_PARALLEL_REQUESTS int = 10

func main() {
	//Getting the -parallel flag
	maxParallel := flag.Int("parallel", DEFAULT_MAX_PARALLEL_REQUESTS, "Max no of Parallel Http Requests")
	flag.Parse()

	addresses := flag.Args()
	httpclient.ParallelFetch(*maxParallel, addresses)
}
