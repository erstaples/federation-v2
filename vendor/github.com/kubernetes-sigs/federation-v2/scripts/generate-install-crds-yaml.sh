#!/usr/bin/env bash

# Copyright 2018 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

INSTALL_CRDS_YAML="${INSTALL_CRDS_YAML:-hack/install-crds-latest.yaml}"

kubebuilder create config --crds --output "${INSTALL_CRDS_YAML}"

# Add crd-install to make sure the CRDs can be installed first in a helm chart.
sed -i 's/^metadata:/metadata:\n  annotations:\n    "helm.sh\/hook": crd-install/g' "${INSTALL_CRDS_YAML}"
