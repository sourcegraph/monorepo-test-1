/*
Copyright 2016 The Kubernetes Authors.

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

package stubs

import "github.com/sourcegraph/monorepo-test-1/kubernetes-14/federation/pkg/dnsprovider/providers/google/clouddns/internal/interfaces"

// Compile time check for interface adherence
var _ interfaces.ChangesService = &ChangesService{}

type ChangesService struct {
	Service *Service
}

func (c *ChangesService) Create(project string, managedZone string, change interfaces.Change) interfaces.ChangesCreateCall {
	return &ChangesCreateCall{c, project, managedZone, change, nil}
}

func (c *ChangesService) NewChange(additions, deletions []interfaces.ResourceRecordSet) interfaces.Change {
	return &Change{c, additions, deletions}
}
