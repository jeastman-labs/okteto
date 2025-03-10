// Copyright 2023 The Okteto Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package volumes

import (
	"github.com/okteto/okteto/pkg/constants"
	"github.com/okteto/okteto/pkg/model"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func translate(dev *model.Dev) *apiv1.PersistentVolumeClaim {
	volumeMode := dev.PersistentVolumeMode()
	labels := map[string]string{
		constants.DevLabel: "true",
	}
	for k, v := range dev.PersistentVolumeLabels() {
		labels[k] = v
	}
	pvc := &apiv1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:        dev.GetVolumeName(),
			Labels:      labels,
			Annotations: dev.PersistentVolumeAnnotations(),
		},
		Spec: apiv1.PersistentVolumeClaimSpec{
			AccessModes: []apiv1.PersistentVolumeAccessMode{dev.PersistentVolumeAccessMode()},
			VolumeMode:  &volumeMode,
			Resources: apiv1.VolumeResourceRequirements{
				Requests: apiv1.ResourceList{
					"storage": resource.MustParse(dev.PersistentVolumeSize()),
				},
			},
		},
	}
	if dev.PersistentVolumeStorageClass() != "" {
		storageClass := dev.PersistentVolumeStorageClass()
		pvc.Spec.StorageClassName = &storageClass
	}
	return pvc
}
