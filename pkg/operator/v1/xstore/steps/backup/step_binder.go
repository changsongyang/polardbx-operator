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

package backup

import (
	polardbxv1 "github.com/alibaba/polardbx-operator/api/v1"
	"github.com/alibaba/polardbx-operator/pkg/k8s/control"
	xstorev1reconcile "github.com/alibaba/polardbx-operator/pkg/operator/v1/xstore/reconcile"
	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type ConditionFunc func(rc *xstorev1reconcile.BackupContext, log logr.Logger) (bool, error)
type StepFunc func(rc *xstorev1reconcile.BackupContext, flow control.Flow) (reconcile.Result, error)

func NewStepBinder(name string, f StepFunc) control.BindFunc {
	return control.NewStepBinder(
		control.NewStep(
			name, func(rc control.ReconcileContext, flow control.Flow) (reconcile.Result, error) {
				return f(rc.(*xstorev1reconcile.BackupContext), flow)
			},
		),
	)
}

func NewStepIfBinder(conditionName string, condFunc ConditionFunc, binders ...control.BindFunc) control.BindFunc {
	condition := control.NewCachedCondition(
		control.NewCondition(conditionName, func(rc control.ReconcileContext, log logr.Logger) (bool, error) {
			return condFunc(rc.(*xstorev1reconcile.BackupContext), log)
		}),
	)

	ifBinders := make([]control.BindFunc, len(binders))
	for i := range binders {
		ifBinders[i] = control.NewStepIfBinder(condition, control.ExtractStepsFromBindFunc(binders[i])[0])
	}

	return control.CombineBinders(ifBinders...)
}

func WhenDeletedAndNotDeleting(binders ...control.BindFunc) control.BindFunc {
	return NewStepIfBinder("Deleted",
		func(rc *xstorev1reconcile.BackupContext, log logr.Logger) (bool, error) {
			backup := rc.MustGetXStoreBackup()
			if backup.Status.Phase == polardbxv1.XStoreBackupDeleting {
				return false, nil
			}
			return !backup.DeletionTimestamp.IsZero(), nil
		},
		binders...,
	)
}
