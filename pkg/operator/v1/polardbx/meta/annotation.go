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

package meta

const (
	AnnotationLock                   = "polardbx/lock"
	AnnotationControllerHints        = "polardbx/controller.hints"
	AnnotationEnableRebalanceOnScale = "polardbx/scale.enable-rebalance"
	AnnotationSchemaCaseInsensitive  = "polardbx/schema.case-insensitive"
)

const (
	HintForbidden = "forbidden"
)

// Guide annotations
const (
	AnnotationConfigGuide       = "polardbx/config-guide"
	AnnotationTopologyModeGuide = "polardbx/topology-mode-guide"
	AnnotationTopologyRuleGuide = "polardbx/topology-rule-guide"
)

// Backup annotations
const (
	AnnotationDummyBackup  = "polardbx/dummy-backup"
	AnnotationBackupBinlog = "polardbx/backupbinlog"
)

// Restore annotations
const (
	// AnnotationImmutableBackupSetPath denotes whether mutate webhook is enabled for RestoreSpec.From.BackupSetPath
	AnnotationImmutableBackupSetPath = "polardbx/immutable-backup-set-path"
)

const (
	AnnotationStorageType = "polardbx/storage-type"
)

const (
	AnnotationPitrConfig = "polardbx/pitr-config"
)
