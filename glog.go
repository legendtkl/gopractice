package main

import (
	"flag"
	//	"fmt"
	"github.com/golang/glog"
	//	"os"
)

func main() {
	flag.Parse()

	glog.Info("This is a Info log")
	glog.Warning("This is a Warning log")
	glog.Error("This is a Error log")

	glog.V(1).Infoln("level 1")
	glog.V(2).Infoln("level 2")

	glog.Flush()
}
