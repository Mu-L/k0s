/*
Copyright k0s authors

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta2 "github.com/k0sproject/k0s/pkg/apis/autopilot/v1beta2"
	autopilotv1beta2 "github.com/k0sproject/k0s/pkg/client/clientset/typed/autopilot/v1beta2"
	gentype "k8s.io/client-go/gentype"
)

// fakePlans implements PlanInterface
type fakePlans struct {
	*gentype.FakeClientWithList[*v1beta2.Plan, *v1beta2.PlanList]
	Fake *FakeAutopilotV1beta2
}

func newFakePlans(fake *FakeAutopilotV1beta2) autopilotv1beta2.PlanInterface {
	return &fakePlans{
		gentype.NewFakeClientWithList[*v1beta2.Plan, *v1beta2.PlanList](
			fake.Fake,
			"",
			v1beta2.SchemeGroupVersion.WithResource("plans"),
			v1beta2.SchemeGroupVersion.WithKind("Plan"),
			func() *v1beta2.Plan { return &v1beta2.Plan{} },
			func() *v1beta2.PlanList { return &v1beta2.PlanList{} },
			func(dst, src *v1beta2.PlanList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta2.PlanList) []*v1beta2.Plan { return gentype.ToPointerSlice(list.Items) },
			func(list *v1beta2.PlanList, items []*v1beta2.Plan) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
