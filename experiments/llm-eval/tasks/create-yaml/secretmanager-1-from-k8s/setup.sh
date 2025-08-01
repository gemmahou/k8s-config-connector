#!/usr/bin/env bash
# Copyright 2025 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

REPO_ROOT="$(git rev-parse --show-toplevel)"
BUILD_DIR="${REPO_ROOT}/.build/tasks/secretmanager-1-from-k8s"
mkdir -p "${BUILD_DIR}"
rm -f "${BUILD_DIR}"/*

# Create the Kubernetes secret
cat <<EOF > "${BUILD_DIR}/k8s-secret.yaml"
apiVersion: v1
kind: Secret
metadata:
  name: my-k8s-secret
stringData:
  secret-data: "super-secret-value"
EOF

# We would normally apply this with kubectl, but for this test, 
# the file's existence is enough for the prompt.
echo "Created k8s-secret.yaml for reference."

