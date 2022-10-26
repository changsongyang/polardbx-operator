/*
Copyright 2021 Alibaba Group Holding Limited.

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

package factory

import (
	corev1 "k8s.io/api/core/v1"

	polardbxv1 "github.com/alibaba/polardbx-operator/api/v1"
	"github.com/alibaba/polardbx-operator/pkg/operator/v1/xstore/convention"
	xstorecommonfactory "github.com/alibaba/polardbx-operator/pkg/operator/v1/xstore/factory"
	xstoreplugincommonfactory "github.com/alibaba/polardbx-operator/pkg/operator/v1/xstore/plugin/common/factory"
	"github.com/alibaba/polardbx-operator/pkg/operator/v1/xstore/reconcile"
)

func NewConfigMap(rc *reconcile.Context, xstore *polardbxv1.XStore, cmType convention.ConfigMapType) (*corev1.ConfigMap, error) {
	switch cmType {
	case convention.ConfigMapTypeConfig:
		return xstorecommonfactory.NewConfigConfigMap(rc, xstore)
	case convention.ConfigMapTypeShared:
		return xstoreplugincommonfactory.NewSharedConfigMap(xstore)
	case convention.ConfigMapTypeTask:
		return xstorecommonfactory.NewTaskConfigMap(xstore), nil
	default:
		panic("unrecognized configmap type: " + cmType)
	}
}
