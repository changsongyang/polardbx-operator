/*
Copyright 2022 Alibaba Group Holding Limited.

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

package polardbx

type TDE struct {
	// Enable represents whether tde is enabled
	// +kubebuilder:default=false
	// +optional
	Enable bool `json:"enable,omitempty"`

	// KeyringPath represents the keyring path
	// +kubebuilder:default=/data/mysql/mysql-keyring/keyring
	// +optional
	KeyringPath string `json:"keyringPath,omitempty"`
}
