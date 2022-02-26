package main

import (
	"flag"
	"fmt"
	"myhttp/httpclient"
)

const DEFAULT_MAX_PARALLEL_REQUESTS int = 10

func main() {
	//Getting the -parallel flag
	maxParallel := flag.Int("parallel", DEFAULT_MAX_PARALLEL_REQUESTS, "Max no of Parallel Http Requests")
	flag.Parse()

	addresses := flag.Args()

	if *maxParallel > len(addresses) {
		*maxParallel = len(addresses)
		fmt.Printf("Max Parallel reduced to [%d] no of addresses\n", *maxParallel)
	}

	httpclient.ParallelFetch(*maxParallel, addresses)
}
