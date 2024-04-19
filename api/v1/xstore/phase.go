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

package xstore

type Phase string

// Valid phases of xstore.
const (
	PhaseNew        Phase = ""
	PhasePending    Phase = "Pending"
	PhaseCreating   Phase = "Creating"
	PhaseRunning    Phase = "Running"
	PhaseLocked     Phase = "Locked"
	PhaseUpgrading  Phase = "Upgrading"
	PhaseRestoring  Phase = "Restoring"
	PhaseRepairing  Phase = "Repairing"
	PhaseDeleting   Phase = "Deleting"
	PhaseFailed     Phase = "Failed"
	PhaseRestarting Phase = "Restarting"
	PhaseUnknown    Phase = "Unknown"
	PhaseAdapting   Phase = "Adapting"
	PhaseTdeOpening Phase = "PhaseTdeOpening"
)

type Stage string

// Valid stages of xstore.
const (
	StageEmpty   Stage = ""
	StageLocking Stage = "Locking"
	StageClean   Stage = "Clean"
	StageUpdate  Stage = "Update"
)

// valid stage of xstore adapting
const (
	StageAdapting      Stage = "StageAdapting"
	StageFlushMetadata Stage = "StageFlushMetadata"
	StageBeforeSuccess Stage = "StageBeforeSuccess"
)
