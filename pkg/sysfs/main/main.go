package main

import (

	//	"encoding/json"

	"fmt"

	balloons "github.com/containers/nri-plugins/cmd/plugins/balloons/policy"
	logger "github.com/containers/nri-plugins/pkg/log"
	"github.com/containers/nri-plugins/pkg/sysfs"
)

var log logger.Logger = logger.NewLogger("policy")

func main() {

	log.EnableDebug(true)

	node, err := balloons.NewCpuTreeFromSystem()
	if err != nil {
		panic(err)
	}
	fmt.Println("Node cpu", node.)

	sys, e := sysfs.DiscoverSystem()
	if e != nil {
		panic(e)
	}

	fmt.Println("sys.NodeIDs()", sys.NodeIDs())

	fmt.Println("PackageIDs()", sys.PackageIDs())

	fmt.Println("SocketCount()", sys.SocketCount())

	fmt.Println("CPUIDs", sys.CPUIDs())

	for _, id := range sys.CPUIDs() {
		c := sys.CPU(id)
		fmt.Println(c.Online())
	}

	for _, nodeID := range sys.NodeIDs() {
		node := sys.Node(nodeID)
		memInfo, _ := node.MemoryInfo()
		fmt.Println("Total", memInfo.MemTotal)
		fmt.Println("MemFree", memInfo.MemFree)
		fmt.Println("Total", memInfo.MemUsed)

		fmt.Println(sys.PackageCount())
		fmt.Println(sys.SocketCount())

	}

	// fmt.Println("NodeIDs", sys.NodeIDs())
	// fmt.Println("PackageIDs", sys.PackageIDs())
	// fmt.Println("CPUIDs", sys.CPUIDs())

	// fmt.Println("PackageCount()", sys.PackageCount())
	// fmt.Println("SocketCount()", sys.SocketCount())
	// fmt.Println("CPUCount()", sys.CPUCount())
	// fmt.Println("NUMANodeCount()", sys.NUMANodeCount())
	// fmt.Println("ThreadCount()", sys.ThreadCount())
}
