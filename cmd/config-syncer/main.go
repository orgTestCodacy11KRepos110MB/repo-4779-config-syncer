/*
Copyright The Config Syncer Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"os"
	"runtime"

	"kubeops.dev/config-syncer/pkg/cmds"

	"gomodules.xyz/logs"
	_ "k8s.io/client-go/kubernetes/fake"
	"k8s.io/klog/v2"
)

func main() {
	rootCmd := cmds.NewCmdConfigSyncer(Version)
	logs.Init(rootCmd, true)
	defer logs.FlushLogs()

	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	if err := rootCmd.Execute(); err != nil {
		klog.Fatal(err)
	}
}
