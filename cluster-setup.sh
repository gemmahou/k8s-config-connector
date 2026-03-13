#!/bin/bash
# Copyright 2022 Google LLC
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

# Script to set up GCP
set -o errexit

# Default values
CLUSTER_1="test-cluster-1"
LOCATION_1="us-west1"
CLUSTER_2="test-cluster-2"
LOCATION_2="us-central1"

# Parse arguments
while [[ $# -gt 0 ]]; do
  case $1 in
    --cluster1)
      CLUSTER_1="$2"
      shift 2
      ;;
    --location1)
      LOCATION_1="$2"
      shift 2
      ;;
    --cluster2)
      CLUSTER_2="$2"
      shift 2
      ;;
    --location2)
      LOCATION_2="$2"
      shift 2
      ;;
    *)
      echo "Unknown option $1"
      exit 1
      ;;
  esac
done

# Configure gcloud with your login credentials.
gcloud auth login
gcloud auth application-default login

# Set PROJECT_ID to your current project
export PROJECT_ID=$(gcloud config get-value project)

# Enable required services
gcloud services enable containerregistry.googleapis.com \
    container.googleapis.com \
    stackdriver.googleapis.com

# Configure gcloud to allow docker to authorize and recognize the gcr.io
# registry.
gcloud auth configure-docker

function setup_cluster() {
    local cluster_name=$1
    local location=$2

    echo "Setting up cluster: ${cluster_name} in ${location}"

    if [[ ! $(gcloud beta container clusters list --filter="name:${cluster_name} AND location:${location}" --format="value(name)") ]]; then
        # Create a GKE cluster with Workload Identity enabled.
        gcloud beta container clusters create "${cluster_name}" \
            --location="${location}" \
            --workload-pool="${PROJECT_ID}.svc.id.goog"
    fi

    # Configure kubectl to communicate with the cluster.
    gcloud container clusters get-credentials "${cluster_name}" --location="${location}"
}

setup_cluster "${CLUSTER_1}" "${LOCATION_1}"
setup_cluster "${CLUSTER_2}" "${LOCATION_2}"

sudo apt-get update && sudo apt-get install google-cloud-cli

sudo apt-get install google-cloud-sdk-gke-gcloud-auth-plugin

# Add an annotation to your default K8s namespace to bind it to your GCP project.
kubectl annotate namespace default \
    "cnrm.cloud.google.com/project-id=${PROJECT_ID}" \
    --overwrite

GREEN='\033[0;32m'
NC='\033[0m'
echo -e "${GREEN}CLUSTER SETUP SUCCESSFUL${NC}"
