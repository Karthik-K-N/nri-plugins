package main

import (

	//"fmt"
	topologyaware "github.com/containers/nri-plugins/cmd/plugins/topology-aware/policy"

)

func main() {

	err := topologyaware.BUILD()
	if err != nil {
		panic(err)
	}






}
