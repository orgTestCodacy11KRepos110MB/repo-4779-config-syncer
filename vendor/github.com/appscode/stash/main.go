package main

import (
	"os"
	"runtime"

	"github.com/appscode/go/log"
	_ "github.com/appscode/stash/client/clientset/versioned/fake"
	"github.com/appscode/stash/pkg/cmds"
	_ "k8s.io/client-go/kubernetes/fake"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"kmodules.xyz/client-go/logs"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	if err := cmds.NewRootCmd().Execute(); err != nil {
		log.Fatalln("Error in Stash Main:", err)
	}
	log.Infoln("Exiting Stash Main")
}
