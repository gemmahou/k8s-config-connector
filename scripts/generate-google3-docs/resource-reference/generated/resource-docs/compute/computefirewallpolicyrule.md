{# AUTOGENERATED. DO NOT EDIT. #}

{% extends "config-connector/_base.html" %}

{% block page_title %}ComputeFirewallPolicyRule{% endblock %}
{% block body %}

<table>
<thead>
<tr>
<th><strong>Property</strong></th>
<th><strong>Value</strong></th>
</tr>
</thead>
<tbody>
<tr>
<td>{{gcp_name_short}} Service Name</td>
<td>Compute Engine</td>
</tr>
<tr>
<td>{{gcp_name_short}} Service Documentation</td>
<td><a href="/compute/docs/">/compute/docs/</a></td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Name</td>
<td>v1.firewallPolicies</td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Documentation</td>
<td><a href="/compute/docs/reference/rest/v1/firewallPolicies">/compute/docs/reference/rest/v1/firewallPolicies</a></td>
</tr>
<tr>
<td>{{product_name_short}} Resource Short Names</td>
<td>gcpcomputefirewallpolicyrule<br>gcpcomputefirewallpolicyrules<br>computefirewallpolicyrule</td>
</tr>
<tr>
<td>{{product_name_short}} Service Name</td>
<td>compute.googleapis.com</td>
</tr>
<tr>
<td>{{product_name_short}} Resource Fully Qualified Name</td>
<td>computefirewallpolicyrules.compute.cnrm.cloud.google.com</td>
</tr>

<tr>
    <td>Can Be Referenced by IAMPolicy/IAMPolicyMember</td>
    <td>No</td>
</tr>


<tr>
<td>{{product_name_short}} Default Average Reconcile Interval In Seconds</td>
<td>600</td>
</tr>
</tbody>
</table>

## Custom Resource Definition Properties



### Spec
#### Schema
```yaml
action: string
description: string
direction: string
disabled: boolean
enableLogging: boolean
firewallPolicyRef:
  external: string
  name: string
  namespace: string
match:
  destAddressGroups:
  - string
  destFqdns:
  - string
  destIPRanges:
  - string
  destRegionCodes:
  - string
  destThreatIntelligences:
  - string
  layer4Configs:
  - ipProtocol: string
    ports:
    - string
  srcAddressGroups:
  - string
  srcFqdns:
  - string
  srcIPRanges:
  - string
  srcRegionCodes:
  - string
  srcThreatIntelligences:
  - string
priority: integer
targetResources:
- external: string
  name: string
  namespace: string
targetServiceAccounts:
- external: string
  name: string
  namespace: string
```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td>
            <p><code>action</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The Action to perform when the client connection triggers the rule. Valid actions are "allow", "deny" and "goto_next".{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>description</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}An optional description for this resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>direction</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The direction in which this rule applies. Possible values: INGRESS, EGRESS{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>disabled</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>enableLogging</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "goto_next" rules.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>firewallPolicyRef</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Immutable.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>firewallPolicyRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}A reference to an externally managed ComputeFirewallPolicy resource. Should be in the format `locations/global/firewallPolicies/{{firewallPolicyID}}`.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>firewallPolicyRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The `name` field of a `ComputeFirewallPolicy` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>firewallPolicyRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The `namespace` field of a `ComputeFirewallPolicy` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.destAddressGroups</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Address groups which should be matched against the traffic destination. Maximum number of destination address groups is 10. Destination address groups is only supported in Egress rules.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.destAddressGroups[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.destFqdns</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Domain names that will be used to match against the resolved domain name of destination of traffic. Can only be specified if DIRECTION is egress.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.destFqdns[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.destIPRanges</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}CIDR IP address range. Maximum number of destination CIDR IP ranges allowed is 256.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.destIPRanges[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.destRegionCodes</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}The Unicode country codes whose IP addresses will be used to match against the source of traffic. Can only be specified if DIRECTION is egress.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.destRegionCodes[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.destThreatIntelligences</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Name of the Google Cloud Threat Intelligence list.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.destThreatIntelligences[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.layer4Configs</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Pairs of IP protocols and ports that the rule should match.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.layer4Configs[]</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.layer4Configs[].ipProtocol</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The IP protocol to which this rule applies. The protocol type is required when creating a firewall rule. This value can either be one of the following well known protocol strings (`tcp`, `udp`, `icmp`, `esp`, `ah`, `ipip`, `sctp`), or the IP protocol number.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.layer4Configs[].ports</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}An optional list of ports to which this rule applies. This field is only applicable for UDP or TCP protocol. Each entry must be either an integer or a range. If not specified, this rule applies to connections through any port. Example inputs include: ``.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.layer4Configs[].ports[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.srcAddressGroups</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Address groups which should be matched against the traffic source. Maximum number of source address groups is 10. Source address groups is only supported in Ingress rules.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.srcAddressGroups[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.srcFqdns</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Domain names that will be used to match against the resolved domain name of source of traffic. Can only be specified if DIRECTION is ingress.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.srcFqdns[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.srcIPRanges</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}CIDR IP address range. Maximum number of source CIDR IP ranges allowed is 256.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.srcIPRanges[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.srcRegionCodes</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}The Unicode country codes whose IP addresses will be used to match against the source of traffic. Can only be specified if DIRECTION is ingress.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.srcRegionCodes[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.srcThreatIntelligences</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Name of the Google Cloud Threat Intelligence list.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>match.srcThreatIntelligences[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>priority</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Immutable. An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>targetResources</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>targetResources[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>targetResources[].external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}A reference to an externally managed Compute Network resource. Should be in the format `projects/{{projectID}}/global/networks/{{network}}`.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>targetResources[].name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The `name` field of a `ComputeNetwork` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>targetResources[].namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The `namespace` field of a `ComputeNetwork` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>targetServiceAccounts</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>targetServiceAccounts[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>targetServiceAccounts[].external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The `email` field of an `IAMServiceAccount` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>targetServiceAccounts[].name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>targetServiceAccounts[].namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>


<p>* Field is required when parent field is specified</p>


### Status
#### Schema
```yaml
conditions:
- lastTransitionTime: string
  message: string
  reason: string
  status: string
  type: string
externalRef: string
kind: string
observedGeneration: integer
ruleTupleCount: integer
```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>conditions</code></td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Conditions represent the latest available observations of the object's current state.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[]</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].lastTransitionTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Last time the condition transitioned from one status to another.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].message</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Human-readable message indicating details about last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].reason</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Unique, one-word, CamelCase reason for the condition's last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].status</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Status is the status of the condition. Can be True, False, Unknown.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].type</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Type is the type of the condition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>externalRef</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}A unique Config Connector specifier for the resource in GCP.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>kind</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Type of the resource. Always `compute#firewallPolicyRule` for firewall policy rules{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedGeneration</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>ruleTupleCount</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Calculation of the complexity of a single firewall policy rule.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>

## Sample YAML(s)

### Typical Use Case
```yaml
# Copyright 2020 Google LLC
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

apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeFirewallPolicyRule
metadata:
  name: firewallpolicyrule-sample
spec:
  action: "deny"
  description: "A Firewall Policy Rule"
  direction: "INGRESS"
  disabled: false
  enableLogging: false
  firewallPolicyRef:
    name: firewallpolicyrule-dep
  match:
    layer4Configs:
    - ipProtocol: "tcp"
    srcIPRanges:
    - "10.100.0.1/32"
  priority: 9000
  targetResources:
    - name: firewallpolicyrule-dep
  targetServiceAccounts:
    - name: firewallpolicyrule-dep
---
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeFirewallPolicy
metadata:
  name: firewallpolicyrule-dep
spec:
  organizationRef:
    # Replace "${ORG_ID?}" with the numeric ID for your organization
    external: "organizations/${ORG_ID?}"
  # ComputeFirewallPolicy shortNames must be unique in the organization in
  # which the firewall policy is created
  shortName: ${PROJECT_ID?}-short
  description: "A basic organization firewall policy"
---
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeNetwork
metadata:
  name: firewallpolicyrule-dep
spec:
  routingMode: REGIONAL
  autoCreateSubnetworks: false
---
apiVersion: iam.cnrm.cloud.google.com/v1beta1
kind: IAMServiceAccount
metadata:
  annotations:
    # Replace ${PROJECT_ID?} with your project ID.
    cnrm.cloud.google.com/project-id: "${PROJECT_ID?}"
  name: firewallpolicyrule-dep
```


Note: If you have any trouble with instantiating the resource, refer to <a href="/config-connector/docs/troubleshooting">Troubleshoot Config Connector</a>.

{% endblock %}
