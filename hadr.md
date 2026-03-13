# Draft Fishfood Guide

The setup is intended for a self-install use case and requires the following to be provisioned:

## Setup environment

The setup involves configuring KCC for multicluster leader election and installing necessary components: KCC, MCL, and Syncer.

#### Install KCC and test CRDs in both clusters
1. **Prepare test clusters**
   1. Clone the Config Connector repository
      ```bash
      git clone git@github.com:GoogleCloudPlatform/k8s-config-connector.git
      make manifests
      ```
   2. Setup local environment
      ```bash
      # Feel free to skip below steps if you have done them already
      # Update apt and install build-essential
      sudo apt-get update
      sudo apt install build-essential
      # Set up sudoless Docker
      cd scripts/environment-setup
      ./docker-setup.sh
      # verify docker by `docker run hello-world`
      # Install Golang
      ./golang-setup.sh
      source ~/.profile
      # Install other build dependencies
      ./repo-setup.sh
      source ~/.profile
      ```
   3. Create test clusters
   
      Use the existing `./gcp-setup.sh` to create and initialize test clusters.
      Ensure you make below changes before executing `gcp-setup.sh`:
      - Line 38: Update zones , i.e. us-west1, us-central1
      - Line 40: Update cluster names, i.e. test-cluster-1, test-cluster-2
      - Line 67: Update Google service account name, default service account name is `cnrm-system`. use can create a new service account, i.e. test-sa or use an existing service account.
      - Line 72 and 78: ensure the service account name in the command is correct if you are using a different service account other than the default `cnrm-system`.
      
      Execute the script twice with different cluster names and zones to create two test clusters.
      ```bash
      ./gcp-setup.sh
      ```
2. **KCC Installation:**
   1. Configure Config Connector
      ```bash
      kubectl apply -f operator/config/crd/bases/core.cnrm.cloud.google.com_configconnectors.yaml
      # apply configconnector.yaml in both clusters
      ```
   2. Deploy the KCC workloads
   
      Because the feature is still WIP, we'll need to work on the feature branch.
      ```bash
      git checkout <remote-branch>
      ```
      ```bash
      make deploy-controller
      # deploy workloads in both clusters
      ```
   3. Verify your installation
      ```bash
      kubectl wait -n cnrm-system --for=condition=Ready pod --all
      # If Config Connector is installed correctly, the output is similar to the following:
      # pod/cnrm-controller-manager-0 condition met
      ```
3. **CRD Installation:** 

   Install CRDs for below test resources
   - Regular resource: `ArtifactRegistryRepository`
   - Resource with service-generated ID: `EssentialContactsContact`
   - Unadoptable resource: `DataflowFlexTemplateJob`
   - You are also welcome to test against other resources, full KCC supported CRDs can be found at `config/crds/resources`.
   ```bash
   kubectl apply -f config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_artifactregistryrepositories.artifactregistry.cnrm.cloud.google.com.yaml
   kubectl apply -f config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_essentialcontactscontacts.essentialcontacts.cnrm.cloud.google.com.yaml
   kubectl apply -f config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_dataflowflextemplatejobs.dataflow.cnrm.cloud.google.com.yaml
   # install CRDs in both clusters
   ```
#### Configure Inter-Cluster Secrets
Assume we have two clusters named `test-cluster-1(us-west1)` and `test-cluster-2(us-central1)`, `test-cluster-1` is the current cluster.
1. Identify the current cluster and find the other cluster context name:
   ```bash
   kubectl config get-contexts
   ```
   ```bash
   CURRENT   NAME                                                  
             gke_test-project_us-central1_test-cluster-2       
   *         gke_test-project_us-west1_test-cluster-1
   ```
2. Export the context to file
   ```bash
   kubectl config view --context=<ANOTHER_CONTEXT_NAME> --minify --flatten > test-cluster-2-kubeconfig.yaml
   # Replace <ANOTHER_CONTEXT_NAME> with the name found above, i.e. gke_test-project_us-central1_test-cluster-2
   # --minify: Only includes the information for that specific context.
   # --flatten: Embeds the certificate data directly into the file so it doesn't rely on external file paths.
   ```
3. Verify the file
   ```bash
   kubectl --kubeconfig=test-cluster-2-kubeconfig.yaml get nodes
   # If this command works, `test-cluster-2-kubeconfig.yaml` is ready to be used.
   ```
4. Create the Kubeconfig Secret
   ```bash
   kubectl create secret generic test-cluster-2 \
   --from-file=kubeconfig=test-cluster-2-kubeconfig.yaml
   # Ensure secret name is the cluster name
   ```    
`test-cluster-1` can now access `test-cluster-2` using its kubeconfig secret. 
Repeat the steps above on `test-cluster-2` to grant it access to `test-cluster-1`.
1. Switch context:
   ```bash
   kubectl config use-context gke_test-project_us-central1_test-cluster-2
   ```
   ```bash
   kubectl config get-contexts
   
   CURRENT   NAME                                                  
   *         gke_test-project_us-central1_test-cluster-2       
             gke_test-project_us-west1_test-cluster-1
   ```
2. Export the context to file
   ```bash
   kubectl config view --context=gke_test-project_us-west1_test-cluster-1 --minify --flatten > test-cluster-1-kubeconfig.yaml
   ```
3. Create the Kubeconfig Secret
   ```bash
   kubectl create secret generic test-cluster-1 \
   --from-file=kubeconfig=test-cluster-1-kubeconfig.yaml
   ```    
4. Switch back to test-cluster-1 for a clean initial test environment
   ```bash
   kubectl config use-context gke_test-project_us-west1_test-cluster-1
   ```
Ensuring both clusters have cross-cluster access is a prerequisite for HA/DR. 
This allows the standby cluster to sync resources from the remote leader during failover.
#### Configure GCP Service Account and GCS Bucket
1. Create a GCS bucket
   ```bash
   gcloud storage buckets create gs://<TEST_BUCKET>
   # GCS bucket must have a unique name.
   # We recommend using a new GCS bucket for testing, this bucket acts as the global lease lock for leader election.
   ```
2. Grant the GSA the roles/storage.admin role on the GCS bucket.
   ```bash
   gcloud storage buckets add-iam-policy-binding gs://<TEST_BUCKET> \
   -member="serviceAccount:<TEST_SA>@<TEST_PROJECT_ID>.iam.gserviceaccount.com" \
   --role="roles/storage.admin"
   ```
#### Install the Multi-Cluster Lease (MCL) controller and its CRD in both clusters
1. Build and push the controller image
   ```bash
   git clone git@github.com:gke-labs/multicluster-leader-election.git
   cd multicluster-leader-election
   export PROJECT_ID=$(gcloud config get-value project)
   export IMG=gcr.io/<TEST_PROJECT_ID>/multiclusterlease/controller:latest
   cd config/manager && kustomize edit set image controller=${IMG}
   cd ../..
   make docker-build && make docker-push
   ```
2. Apply RBAC permissions
   ```bash
   # Ensure the controller has necessary permissions to manage MultiClusterLeases object.
   # Update name and namespace in role_binding.yaml with the KSA and namespace:
   apiVersion: rbac.authorization.k8s.io/v1
   kind: ClusterRoleBinding
   metadata:
     name: manager-rolebinding
   roleRef:
     apiGroup: rbac.authorization.k8s.io
     kind: ClusterRole
     name: manager-role
   subjects:
   - kind: ServiceAccount
     name: cnrm-controller-manager
     namespace: cnrm-system
   ```
   ```bash
   kubectl apply -f config/rbac/role.yaml
   kubectl apply -f config/rbac/role_binding.yaml
   
   kubectl --context=<ANOTHER_CONTEXT_NAME> apply -f config/rbac/role.yaml          
   kubectl --context=<ANOTHER_CONTEXT_NAME> apply -f config/rbac/role_binding.yaml
   ```
3. Deploy the MCL controller
   ```bash
   # mclcontroller.yaml
   apiVersion: apps/v1
   kind: Deployment
   metadata:
     name: multiclusterlease-controller-manager
     namespace: cnrm-system
   spec:
     replicas: 1
     template:
       metadata:
         labels:
           control-plane: controller-manager-mcl
       spec:
         serviceAccountName: cnrm-controller-manager              
    containers:
           - name: manager
             image: gcr.io/<TEST_PROJECT_ID>/multiclusterlease/controller:latest 
             command:
               - /manager
             args:
               - --gcs-bucket=<TEST_BUCKET>
     selector:
       matchLabels:
         control-plane: controller-manager-mcl
   ```
   ```bash
   kubectl apply -f mclcontroller.yaml
   kubectl --context=<ANOTHER_CONTEXT_NAME> apply -f mclcontroller.yaml 
   ```
4. Install MultiClusterLease CRD
   ```bash
   kubectl apply -f config/crd/bases/multicluster.core.cnrm.cloud.google.com_multiclusterleases.yaml
   kubectl --context=<ANOTHER_CONTEXT_NAME> apply -f config/crd/bases/multicluster.core.cnrm.cloud.google.com_multiclusterleases.yaml
   ```
#### Install the Syncer controller and its CRD in both clusters
1. Build and push the controller image
   ```bash
   git clone git@github.com:gke-labs/kube-etl.git
   cd syncer
   # image name is default to gcr.io/<TEST_PROJECT_ID>/krmsyncer/controller:latest
   make docker-build && make docker-push
   ```
2. Apply RBAC permissions
   ```bash
   kubectl apply -f config/rbac/role.yaml
   kubectl --context=<ANOTHER_CONTEXT_NAME> apply -f config/rbac/role.yaml
   ```
3. Deploy the Syncer controller
   ```bash
   # krmsyncercontroller.yaml
   apiVersion: apps/v1
   kind: Deployment
   metadata:
     name: krmsyncer-controller-manager
     namespace: cnrm-system
   spec:
     replicas: 1
     template:
       metadata:
         labels:
           control-plane: controller-manager-krmsyncer
       spec:
         serviceAccountName: cnrm-controller-manager              
    containers:
           - name: manager
             image: gcr.io/<TEST_PROJECT_ID>/krmsyncer/controller:latest 
             command:
               - /manager
     selector:
       matchLabels:
         control-plane: controller-manager-krmsyncer
   ```
   ```bash
   kubectl apply -f krmsyncercontroller.yaml
   kubectl --context=<ANOTHER_CONTEXT_NAME> apply -f krmsyncercontroller.yaml 
   ```
4. Install KRMSyncer CRD
   ```bash
   kubectl apply -f config/crd/syncer.gkelabs.io_krmsyncers.yaml
   kubectl --context=<ANOTHER_CONTEXT_NAME> apply -f config/crd/syncer.gkelabs.io_krmsyncers.yaml
   ```
#### Configure Multi-Cluster Lease and Syncer Resources in both clusters
1. Update the ConfigConnector.yaml file in both clusters to integrate the lease and syncer mechanism.
   ```bash
   # In current cluster test-cluster-1
   apiVersion: core.cnrm.cloud.google.com/v1beta1
   kind: ConfigConnector
   metadata:
     name: configconnector.core.cnrm.cloud.google.com
   spec:
     mode: cluster
     googleServiceAccount: "<TEST_SA>@<TEST_PROJECT_ID>.iam.gserviceaccount.com"
     stateIntoSpec: Absent 
     experiments:
       multiClusterLease:
         leaseName: test-lease
         namespace: cnrm-system
         clusterCandidateIdentity: test-cluster-1 >>>>> use the name for the another cluster, i.e. test-cluster-2
   ```
   ```bash
   kubectl apply -f configconnector.yaml
   ```
   ```bash
   # In another cluster test-cluster-2
   apiVersion: core.cnrm.cloud.google.com/v1beta1
   kind: ConfigConnector
   metadata:
     name: configconnector.core.cnrm.cloud.google.com
   spec:
     mode: cluster
     googleServiceAccount: "<TEST_SA>@<TEST_PROJECT_ID>.iam.gserviceaccount.com"
     stateIntoSpec: Absent 
     experiments:
       multiClusterLease:
         leaseName: test-lease
         namespace: cnrm-system
         clusterCandidateIdentity: test-cluster-2
   ```
   ```bash
   kubectl --context=<ANOTHER_CONTEXT_NAME> apply -f configconnector.yaml 
   ```

## Expected Leader Election Behavior

The KCC manager pod that acquires the global lease from the GCS bucket becomes the **Leader** and performs reconciliation. The other manager pod enters **Standby** mode. If the Leader fails to renew its lease (e.g., due to pod failure or network partition), the Standby pod automatically attempts to take over as the new Leader after the lease expires (typically 15-20 seconds).

1. Verify MultiClusterLease objects in both clusters
   ```bash  
   kubectl describe MultiClusterLease test-lease
   kubectl --context=<ANOTHER_CONTEXT_NAME> describe MultiClusterLease test-lease 
   ``` 
   One candidate will be elected as leader, typically `test-cluster-1`, because it acquires the lock earlier than `test-cluster-2`.
   Both MultiClusterLease objects should have the same `globalHolderIdentity` indicates the current leader. For example:
   ```bash
   # Leader
   status:
     conditions:
     - lastTransitionTime: '2026-03-12T23:48:26Z'
       message: successfully communicated with backend
       observedGeneration: 62007
       reason: BackendHealthy
       status: 'True'
       type: BackendHealthy
     globalHolderIdentity: test-cluster-1
   
   # Standby
   status:
     conditions:
     - lastTransitionTime: '2026-03-12T23:47:38Z'
       message: 'failed to communicate with backend: candidate lease is stale'
       observedGeneration: 40525
       reason: BackendError
       status: 'False'
       type: BackendHealthy
     globalHolderIdentity: test-cluster-1
   ```
2. Verify KRMSyncer objects in both clusters
   ```bash
   kubectl describe KRMSyncer test-syncer
   kubectl --context=<ANOTHER_CONTEXT_NAME> describe KRMSyncer test-syncer
   ```
   Both KRMSyncer objects should be active with the correct suspend configuration.
   Under our `pull` model, the KRMSyncer should remain suspended on the leader and active on the standby cluster.
   For example:
   ```bash
   # Leader
   spec:
     suspend: true 
   status:
     conditions:
     - lastTransitionTime: '2026-03-12T23:48:26Z'
       message: Controller is active
       reason: Active
       status: 'True'
       type: Active
       
   # Standby
   spec:
     suspend: false
   status:
     conditions:
     - lastTransitionTime: '2026-03-12T23:47:38Z'
       message: Controller is active
       reason: Active
       status: 'True'
       type: Active
   ```
3. (optional) Check controller logs

   Please review the controller logs if your objects are in an unexpected or error state.
   Reach out to acpana@ or yuhou@ to confirm if a bug report/ticket is needed.

## Test Cases

The testing is divided into the following scenarios to verify HA/DR functionality and resource reconciliation after failover.
### Enable GCP service API
```bash
gcloud services enable artifactregistry.googleapis.com \
                       dataflow.googleapis.com \
                       essentialcontacts.googleapis.com
```
### Failover with a Regular Resource (ArtifactRegistryRepository)

This test verifies that a standard resource is correctly reconciled by the new leader after a failover.

1. **Apply Resource to Leader Cluster:**
   ```yaml
   # artifactregistry.yaml
   apiVersion: artifactregistry.cnrm.cloud.google.com/v1beta1
   kind: ArtifactRegistryRepository
   metadata:
     name: test-repo
   spec:
     location: us-central1
     format: DOCKER
   ```
   ```bash
   kubectl apply -f artifactregistry.yaml
   ```
2. **Verify Reconciliation in Leader:**

   Ensure the resource becomes `Ready`.
   ```bash
   kubectl wait --for=condition=Ready artifactregistryrepository test-repo
   ```
3. **Verify Sync to Standby:**

   Confirm the resource has been synced to the standby cluster by the Syncer.
   ```bash
   kubectl --context=<ANOTHER_CONTEXT> get artifactregistryrepository test-repo
   ```
4. **Simulate Failover:**

   Scale down the current leader cluster's KCC deployment.
5. **Verify Reconciliation in New Leader:**

   After the standby cluster takes over leadership, verify that it successfully reconciles the resource.
   ```bash
   kubectl --context=<ANOTHER_CONTEXT> wait --for=condition=Ready artifactregistryrepository mcl-test-repo --timeout=5m
   ```

### Failover with Service-Generated ID Resource (EssentialContactsContact)

This test verifies that resources with IDs generated by the GCP service (stored in `status`) are correctly handled and do not result in duplicate resources or errors after failover.

1. **Apply Resource:**
   ```yaml
   # essentialcontact.yaml
   apiVersion: essentialcontacts.cnrm.cloud.google.com/v1beta1
   kind: EssentialContactsContact
   metadata:
     name: test-contact
   spec:
     projectRef:
       external: <TEST_PROJECT_ID>
     languageTag: en
     notificationCategorySubscriptions:
       - BILLING
     # email needs to be unique within the project
     email: "test-email@test.com"
   ```
   ```bash
   kubectl apply -f essentialcontact.yaml
   ```
2. **Verify Ready and Status:**

   Verify the resource is `Ready` and has a service-generated name in its status.
   ```bash
   kubectl get essentialcontactscontact test-contact -o yaml
   ```
3. **Trigger Failover:**

   Scale down the current leader.
4. **Verify Consistency:**

   Confirm that the new leader reconciles the resource using the existing service-generated ID from the synced status, and the resource remains `Ready`.

### Failover with Unadoptable Resource (DataflowFlexTemplateJob)

Some resources like `DataflowFlexTemplateJob` have unique reconciliation patterns. This test ensures they are handled correctly.

1. **Apply Dataflow Job:**

   Apply a `DataflowFlexTemplateJob` and wait for it to be reconciled in the leader cluster.
2. **Trigger Failover:**

   Scale down the leader.
3. **Verify Behavior:**

   Verify that the new leader correctly observes the existing job and does not attempt to trigger a new job unnecessarily (unless intended by the resource's logic).

### (optional) Failover with Other Resources


## Known Issues

- **Sync Latency:** There might be a slight delay between a resource being updated in the leader cluster and it being synced to the standby cluster.
- **Lease Expiration Time:** The failover time is governed by the lease duration and renewal interval configured in the MCL controller.
