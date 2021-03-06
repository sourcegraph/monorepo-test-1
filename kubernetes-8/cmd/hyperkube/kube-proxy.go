/*
Copyright 2015 The Kubernetes Authors.

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
	"k8s.io/apiserver/pkg/server/healthz"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-8/cmd/kube-proxy/app"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-8/cmd/kube-proxy/app/options"
)

func init() {
	healthz.DefaultHealthz()
}

// NewKubeProxy creates a new hyperkube Server object that includes the
// description and flags.
func NewKubeProxy() *Server {
	config := options.NewProxyConfig()

	hks := Server{
		name:            "proxy",
		AlternativeName: "kube-proxy",
		SimpleUsage:     "proxy",
		Long: `The Kubernetes proxy server is responsible for taking traffic directed at
		services and forwarding it to the appropriate pods. It generally runs on
		nodes next to the Kubelet and proxies traffic from local pods to remote pods.
		It is also used when handling incoming external traffic.`,
	}

	config.AddFlags(hks.Flags())

	hks.Run = func(_ *Server, _ []string) error {
		s, err := app.NewProxyServerDefault(config)
		if err != nil {
			return err
		}

		return s.Run()
	}

	return &hks
}
