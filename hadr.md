# Draft Fishfood Guide

The setup is intended for a self-install use case and requires the following to be provisioned:

## Setup environment

The setup involves configuring KCC for multicluster leader election and installing necessary components: KCC, MCL, and Syncer.

#### Install KCC and test CRDs in both clusters
1. **Prepare test clusters**
   1. Clone the Config Connector repository
      ```bash
      git clone git@github.com:GoogleCloudPlatform/k8s-config-connector.git
      cd k8s-config-connector
      ```
   2. (optional) Setup local environment

      Feel free to skip below steps if you have done them before.
      ```bash
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

      Execute the script to create two GKE clusters (with Workload Identity enabled) in your specified locations. 
      It also handles authentication, enables required services, creates a dedicated GCP Service Account, 
      and sets up the required IAM bindings to allow Config Connector to manage resources via Workload Identity. 
      
      [Cluster setup script](cluster-setup.sh)
      ```bash
      ./cluster-setup.sh 
      # Default values are: --cluster1 test-cluster-1 --location1 us-west1 --cluster2 test-cluster-2 --location2 us-central1
      # Run  ./cluster-setup.sh --cluster1 <name> --location1 <location> --cluster2 <name> --location2 <location> with custom parameters
      ```
   4. Identify cluster context names:
      ```bash
      kubectl config get-contexts
      ```
      ```bash
      CURRENT   NAME                                                  
                gke_test-project_us-central1_test-cluster-2       
      *         gke_test-project_us-west1_test-cluster-1
      ```
2. **KCC Installation:**
   1. Deploy the KCC workloads

      ```bash
      gcloud storage cp gs://configconnector-operator/latest/release-bundle.tar.gz release-bundle.tar.gz
      tar zxvf release-bundle.tar.gz
      
      kubectl apply -f operator-system/configconnector-operator.yaml
      kubectl --context=<ANOTHER_CONTEXT> apply -f operator-system/configconnector-operator.yaml
      ```
   2. Configure the operator as cluster mode

      Note: The desired mode should be namespaced. However, we are still developing HA/DR for namespaced mode,
      we'll be using with cluster mode for today's session.
      ```bash
      # configconnector.yaml
      apiVersion: core.cnrm.cloud.google.com/v1beta1
      kind: ConfigConnector
      metadata:
        # the name is restricted to ensure that there is only ConfigConnector resource installed in your cluster
        name: configconnector.core.cnrm.cloud.google.com
      spec:
        mode: cluster
        stateIntoSpec: Absent
        googleServiceAccount: "<TEST_SERVICE_ACCOUNT>@<TEST_PROJECT_ID>.iam.gserviceaccount.com"
      ```
      ```bash
      kubectl apply -f configconnector.yaml
      kubectl --context=<ANOTHER_CONTEXT> apply -f configconnector.yaml
      ```
   2. Verify your installation
      ```bash
      kubectl wait -n cnrm-system --for=condition=Ready pod --all
      # If Config Connector is installed correctly, the output is similar to the following:
      # pod/cnrm-controller-manager-0 condition met
      ```
   3. Create an identity 
      ```bash
      # Create an Google service account. It is recommended to use the default service account `cnrm-system`.
      # Skip this step if you want to use an existing service account.
      gcloud iam service-accounts create <TEST_SERVICE_ACCOUNT>
      
      # Give the GCP Service Account elevated permissions on your project.
      gcloud projects add-iam-policy-binding "<TEST_PROJECT_ID>" \
      --member="serviceAccount:<TEST_SERVICE_ACCOUNT>@<TEST_PROJECT_ID>.iam.gserviceaccount.com" \
      --role="roles/owner"
   
      # Create a GCP IAM Policy Binding between the GCP Service Account and the
      # Kubernetes Service Account.
      gcloud iam service-accounts add-iam-policy-binding "<TEST_SERVICE_ACCOUNT>@<TEST_PROJECT_ID>.iam.gserviceaccount.com" \
      --member="serviceAccount:<TEST_PROJECT_ID>.svc.id.goog[cnrm-system/cnrm-controller-manager]" \
      --role="roles/iam.workloadIdentityUser"
      ```
   4. (optional) Annotate the Kubernetes Service Account in both clusters

      If you are using a GSA other than the default "cnrm-system", you'll need to annotate the KSA with your GSA in **both clusters**.
      ```bash
      kubectl annotate serviceaccount cnrm-controller-manager \
      -n <TEST_SERVICE_ACCOUNT> \
      iam.gke.io/gcp-service-account=<TEST_SERVICE_ACCOUNT>@<TEST_PROJECT_ID>.iam.gserviceaccount.com \
      --overwrite
      
      kubectl --context=<ANOTHER_CONTEXT> annotate serviceaccount cnrm-controller-manager \
      -n <TEST_SERVICE_ACCOUNT> \
      iam.gke.io/gcp-service-account=<TEST_SERVICE_ACCOUNT>@<TEST_PROJECT_ID>.iam.gserviceaccount.com \
      --overwrite
      ```
   5. Verify KSA bindings
      ```bash
      kubectl describe serviceaccount cnrm-controller-manager -n cnrm-system
      # Confirm annotation contains correct GSA
      # iam.gke.io/gcp-service-account: <TEST_SERVICE_ACCOUNT>@<TEST_PROJECT_ID>.iam.gserviceaccount.com
      ```
#### Configure Inter-Cluster Secrets
Assume we have two clusters named `test-cluster-1(us-west1)` and `test-cluster-2(us-central1)`, `test-cluster-1` is the current cluster.

1. Identify cluster context names:
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
   kubectl config view --context=<CLUSTER1_CONTEXT_NAME> --minify --flatten > test-cluster-1-kubeconfig.yaml
   kubectl config view --context=<CLUSTER2_CONTEXT_NAME> --minify --flatten > test-cluster-2-kubeconfig.yaml
   # Replace <CLUSTER1_CONTEXT_NAME> and <CLUSTER2_CONTEXT_NAME> with the name found above, i.e. gke_test-project_us-central1_test-cluster-2
   # --minify: Only includes the information for that specific context.
   # --flatten: Embeds the certificate data directly into the file so it doesn't rely on external file paths.
   ```
3. Verify the file
   ```bash
   kubectl --kubeconfig=test-cluster-1-kubeconfig.yaml get nodes
   kubectl --kubeconfig=test-cluster-2-kubeconfig.yaml get nodes
   # If this command works, `test-cluster-1-kubeconfig.yaml` and `test-cluster-2-kubeconfig.yaml` is ready to be used.
   ```
4. Create the Kubeconfig Secret
   ```bash
    # Ensure secret name is the cluster name.
   kubectl --context=<CLUSTER_1_CONTEXT> create secret generic test-cluster-2 \
   --from-file=kubeconfig=test-cluster-2-kubeconfig.yaml -n cnrm-system
   
   kubectl --context=<CLUSTER_2_CONTEXT> create secret generic test-cluster-1 \
   --from-file=kubeconfig=test-cluster-1-kubeconfig.yaml -n cnrm-system
   ```    
Ensuring both clusters have cross-cluster access is a prerequisite for HA/DR. 
This allows the standby cluster to sync resources from the remote leader during failover.

#### Configure GCP Service Account and GCS Bucket
1. (optional) Create a GCS bucket
   
   Skip this step if you want to use an existing GCS bucket.
   ```bash
   gcloud storage buckets create gs://<TEST_BUCKET>
   # GCS bucket must have a unique name.
   ```
2. Grant the GSA the roles/storage.admin role on the GCS bucket.
   ```bash
   gcloud storage buckets add-iam-policy-binding gs://<TEST_BUCKET> \
   --member="serviceAccount:<TEST_SA>@<TEST_PROJECT_ID>.iam.gserviceaccount.com" \
   --role="roles/storage.admin"
   ```
#### Install the Multi-Cluster Lease (MCL) controller and its CRD in both clusters
1. Build and push the controller image
   ```bash
   git clone git@github.com:gke-labs/multicluster-leader-election.git
   cd multicluster-leader-election
   export IMG=gcr.io/<TEST_PROJECT_ID>/multiclusterlease/controller:latest
   cd config/manager && kustomize edit set image controller=${IMG}
   cd ../..
   make docker-build && make docker-push
   ```
2. Apply RBAC permissions
   
   Ensure the controller has necessary permissions to manage MultiClusterLease object.
   Example [rbac-mcl.yaml](rbac-mcl.yaml).
   ```yaml
   # rbac-mcl.yaml
   # Update name and namespace in with the KSA and namespace if it's not cnrm-system/cnrm-controller-manager
   apiVersion: rbac.authorization.k8s.io/v1
   kind: ClusterRoleBinding
   metadata:
     name: mcl-manager-rolebinding
   roleRef:
     apiGroup: rbac.authorization.k8s.io
     kind: ClusterRole
     name: mcl-manager-role
   subjects:
   - kind: ServiceAccount
     name: cnrm-controller-manager    >>>> default KSA name
     namespace: cnrm-system    >>>>> default KSA namespace
   ```
   ```bash
   kubectl apply -f rbac-mcl.yaml
   kubectl --context=<ANOTHER_CONTEXT_NAME> apply -f rbac-mcl.yaml
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
#### Enable HA/DR experiment in both clusters
```bash
kubectl config use-context <CLUSTER1_CONTEXT>
```
1. Update the ConfigConnector.yaml file in both clusters to integrate the multiclusterlease mechanism.
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
   kubectl describe MultiClusterLease test-lease -n cnrm-system
   kubectl --context=<ANOTHER_CONTEXT_NAME> describe MultiClusterLease test-lease -n cnrm-system
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
2. (optional) Check controller logs

   Please review the controller logs if your objects are in an unexpected or error state.
   Reach out to acpana@ or yuhou@ to confirm if a bug report/ticket is needed.

## Syncer Setup in Standby Cluster

Note: Syncer integration in KCC is still under development, so we have to manually compelete below steps for today's session.
Only apply this in the **standby** cluster.

```bash
# switch context to the standby cluster
kubectl config use-context <ANOTHER_CONTEXT>
```

#### Install the Syncer controller and its CRD
1. Build and push the controller image
   ```bash
   git clone git@github.com:gke-labs/kube-etl.git
   cd syncer
   # image name is default to gcr.io/<TEST_PROJECT_ID>/krmsyncer/controller:latest
   make docker-build && make docker-push
   ```
2. Apply RBAC permissions

   Ensure the controller has necessary permissions to manage KRMSyncer object.
   Example [rbac-syncer.yaml](rbac-syncer.yaml).
   ```yaml
   # rbac-syncer.yaml
   # Update name and namespace in with the KSA and namespace if it's not cnrm-system/cnrm-controller-manager
   apiVersion: rbac.authorization.k8s.io/v1
   kind: ClusterRoleBinding
   metadata:
     name: syncer-manager-rolebinding
   roleRef:
     apiGroup: rbac.authorization.k8s.io
     kind: ClusterRole
     name: syncer-manager-role
   subjects:
   - kind: ServiceAccount
     name: cnrm-controller-manager    >>>> default KSA name
     namespace: cnrm-system    >>>>> default KSA namespace
   ```
   ```bash
   kubectl apply -f rbac-syncer.yaml
   ```
3. Deploy the Syncer controller
   ```bash
   # syncercontroller.yaml
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
   kubectl apply -f syncercontroller.yaml
   ```
4. Install KRMSyncer CRD
   ```bash
   kubectl apply -f config/crd/syncer.gkelabs.io_krmsyncers.yaml
   ```
5. Apply KRMSyncer configuration
   ```yaml
   apiVersion: syncer.gkelabs.io/v1alpha1
   kind: KRMSyncer
   metadata:
     name: test-syncer
     namespace: cnrm-system
   spec:
     destination:
       clusterConfig:
         kubeConfigSecretRef:
           name: test-cluster-1    >>>> use the name(secret name) of the leader cluster. Change to test-cluster-2 if it's the current leader.
     rules:
       - group: artifactregistry.cnrm.cloud.google.com
         kind: ArtifactRegistryRepository
         version: v1beta1
       - group: essentialcontacts.cnrm.cloud.google.com
         kind: EssentialContactsContact
         version: v1beta1
       - group: dataflow.cnrm.cloud.google.com
         kind: DataflowFlexTemplateJob
         version: v1beta1
     suspend: false
   ```
6. Switch back to leader cluster

   The Syncer controller in the standby cluster will then begin reconciling the KRMSyncer object,
   which synchronizes specified resources from the leader cluster. This ensures that if a failover occurs,
   and the standby cluster now becomes the new leader, KCC controller can seamlessly resume resource reconciliation.
   We will verify this process in the next section. Before that, let's switch back to the current leader.
   ```bash
   kubectl config use-context <LEADER_CONTEXT>
   ```
## Test Cases
The testing is divided into the following scenarios to verify HA/DR functionality and resource reconciliation after failover.

### Enable GCP service API
```bash
gcloud services enable artifactregistry.googleapis.com \
                       dataflow.googleapis.com \
                       essentialcontacts.googleapis.com
```

### Apply test resources to both clusters

Simulate the initial state.

1. ArtifactRegistryRepository
   
   Example [artifactregistry.yaml](artifactregistry.yaml)
   ```yaml
   # artifactregistry.yaml
   apiVersion: artifactregistry.cnrm.cloud.google.com/v1beta1
   kind: ArtifactRegistryRepository
   metadata:
     name: test-repo
   spec:
     location: us-central1
     format: DOCKER
     description: "test repository description"
   ```
   ```bash
   kubectl apply -f artifactregistry.yaml
   kubectl --context=<ANOTHER_CONTEXT> apply -f artifactregistry.yaml
   ```
2. EssentialContactsContact

   Example [essentialcontact.yaml](essentialcontact.yaml)   
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
   kubectl --context=<ANOTHER_CONTEXT> apply -f essentialcontact.yaml
   ```
3. DataflowFlexTemplateJob

   Example [dataflowflextemplatejob.yaml](dataflowflextemplatejob.yaml)
   ```yaml
   # dataflowflextemplatejob.yaml
   apiVersion: storage.cnrm.cloud.google.com/v1beta1
   kind: StorageBucket
   metadata:
     annotations:
       cnrm.cloud.google.com/force-destroy: "true"
       cnrm.cloud.google.com/reconcile-interval-in-seconds: "0" # Avoid time-dependencies
     name: storagebucket
   ---
   apiVersion: dataflow.cnrm.cloud.google.com/v1beta1
   kind: DataflowFlexTemplateJob
   metadata:
     annotations:
       cnrm.cloud.google.com/on-delete: "cancel"
     name: test-job
     labels:
       label-one: "value-one"  
   spec:
     region: us-central1
     # This is a public, Google-maintained Dataflow Job flex template of a batch job
     containerSpecGcsPath: gs://dataflow-templates/2022-10-03-00_RC00/flex/File_Format_Conversion
     parameters:
       autoscalingAlgorithm: AUTOSCALING_ALGORITHM_NONE
       inputFileFormat: csv
       outputFileFormat: avro
       # This is maintained by us.
       inputFileSpec: gs://config-connector-samples/dataflowflextemplate/numbertest.csv
       outputBucket: gs://storagebucket
       # This is maintained by us.
       schema: gs://config-connector-samples/dataflowflextemplate/numbers.avsc
   ```
   ```bash
   kubectl apply -f dataflowflextemplatejob.yaml
   kubectl --context=<ANOTHER_CONTEXT> apply -f dataflowflextemplatejob.yaml
   ```

### Make an Update on Standby

1. ArtifactRegistryRepository
   ```yaml
   # artifactregistry.yaml
   apiVersion: artifactregistry.cnrm.cloud.google.com/v1beta1
   kind: ArtifactRegistryRepository
   metadata:
     name: test-repo
   spec:
     location: us-central1
     format: DOCKER
     description: "test repository description updated" >>>>> update this field
   ```
   ```bash
   kubectl apply -f artifactregistry.yaml
   ```
2. EssentialContactsContact
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
       - LEGAL  >>>>> update this field 
     # email needs to be unique within the project
     email: "test-email@test.com"
   ```
   ```bash
   kubectl apply -f essentialcontact.yaml
   ```
3. DataflowFlexTemplateJob
   
   Resource is immutable.

### Verify No Reconciliation on Standby
   
   Ensure the resources remain unchanged
   ```bash
   gcloud artifacts repositories describe test-repo --location=us-west1
   gcloud essential-contacts describe <GENERATED_ID>
   # dataflowflextemplatejob is immutable
   ```

### Verify Reconciliation on Leader
   ```bash
   kubectl apply -f artifactregistry.yaml
   kubectl apply -f essentialcontact.yaml
   # dataflowflextemplatejob is immutable
   ```
   Ensure the resources have been updated
   ```bash
   gcloud artifacts repositories describe test-repo --location=us-central1
   gcloud essential-contacts describe <GENERATED_ID>
   # dataflowflextemplatejob is immutable
   ```

### Trigger a failover

Scale down the replica of leader cluster's manager.
```bash
kubectl scale statefulset cnrm-controller-manager -n cnrm-system --replicas=0
```
Wait for 15-20s, verify previous standby becomes the new leader.
```bash
kubectl describe MultiClusterLease test-lease -n cnrm-system | grep  "Global Holder Identity"
kubectl --context=<ANOTHER_CONTEXT_NAME> describe MultiClusterLease test-lease -n cnrm-system | grep  "Global Holder Identity"
```
Check `multiclusterlease-controller-manager` log if the output is unexpected.
### Verify Reconciliation in New Leader

```bash
kubectl config use-context <NEW_LEADER_CONTEXT>
````

1. ArtifactRegistryRepository

   Verify that the new leader successfully reconciles the resource with the update.
   ```bash
   kubectl get artifactregistryrepository test-repo -o yaml
   ```

2. EssentialContactsContact

   Verify that the new leader reconciles the resource with the update using the existing service-generated ID from the synced status. 
   The resource remains `UpToDate`, and does not attempt to trigger a new job unnecessarily.
   ```bash
   kubectl get essentialcontactscontact test-contact -o yaml
   ```
   ```bash
   gcloud essential-contacts list --project=<TEST_PROJECT_ID>
   # there should be only one contact resource we just created
   ```

3. DataflowFlexTemplateJob

   Verify that the new leader correctly observes the existing job and reconciles the resource with the update.
   The resource remains `UpToDate`, and does not attempt to trigger a new job unnecessarily.
   ```bash
   kubectl get dataflowflextemplatejob test-job -o yaml
   ```
   ```bash
   gcloud dataflow jobs list --region="us-central1"
   # there should be only one job resource we just created
   ```

### (optional) Failover with Other Resources

## (optional) Cleanup
Feel free to remove any test resources, deployments, clusters.

## Known Issues

- **Sync Latency:** There might be a slight delay between a resource being updated in the leader cluster and it being synced to the standby cluster.
- **Lease Expiration Time:** The failover time is governed by the lease duration and renewal interval configured in the MCL controller.
