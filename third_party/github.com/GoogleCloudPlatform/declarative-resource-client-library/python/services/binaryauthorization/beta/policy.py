# Copyright 2024 Google LLC. All Rights Reserved.
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
from connector import channel
from google3.cloud.graphite.mmv2.services.google.binary_authorization import policy_pb2
from google3.cloud.graphite.mmv2.services.google.binary_authorization import (
    policy_pb2_grpc,
)

from typing import List


class Policy(object):
    def __init__(
        self,
        admission_whitelist_patterns: list = None,
        cluster_admission_rules: dict = None,
        kubernetes_namespace_admission_rules: dict = None,
        kubernetes_service_account_admission_rules: dict = None,
        istio_service_identity_admission_rules: dict = None,
        default_admission_rule: dict = None,
        description: str = None,
        global_policy_evaluation_mode: str = None,
        self_link: str = None,
        project: str = None,
        update_time: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.admission_whitelist_patterns = admission_whitelist_patterns
        self.cluster_admission_rules = cluster_admission_rules
        self.kubernetes_namespace_admission_rules = kubernetes_namespace_admission_rules
        self.kubernetes_service_account_admission_rules = (
            kubernetes_service_account_admission_rules
        )
        self.istio_service_identity_admission_rules = (
            istio_service_identity_admission_rules
        )
        self.default_admission_rule = default_admission_rule
        self.description = description
        self.global_policy_evaluation_mode = global_policy_evaluation_mode
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = policy_pb2_grpc.BinaryauthorizationBetaPolicyServiceStub(
            channel.Channel()
        )
        request = policy_pb2.ApplyBinaryauthorizationBetaPolicyRequest()
        if PolicyAdmissionWhitelistPatternsArray.to_proto(
            self.admission_whitelist_patterns
        ):
            request.resource.admission_whitelist_patterns.extend(
                PolicyAdmissionWhitelistPatternsArray.to_proto(
                    self.admission_whitelist_patterns
                )
            )
        if Primitive.to_proto(self.cluster_admission_rules):
            request.resource.cluster_admission_rules = Primitive.to_proto(
                self.cluster_admission_rules
            )

        if Primitive.to_proto(self.kubernetes_namespace_admission_rules):
            request.resource.kubernetes_namespace_admission_rules = Primitive.to_proto(
                self.kubernetes_namespace_admission_rules
            )

        if Primitive.to_proto(self.kubernetes_service_account_admission_rules):
            request.resource.kubernetes_service_account_admission_rules = (
                Primitive.to_proto(self.kubernetes_service_account_admission_rules)
            )

        if Primitive.to_proto(self.istio_service_identity_admission_rules):
            request.resource.istio_service_identity_admission_rules = (
                Primitive.to_proto(self.istio_service_identity_admission_rules)
            )

        if PolicyDefaultAdmissionRule.to_proto(self.default_admission_rule):
            request.resource.default_admission_rule.CopyFrom(
                PolicyDefaultAdmissionRule.to_proto(self.default_admission_rule)
            )
        else:
            request.resource.ClearField("default_admission_rule")
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if PolicyGlobalPolicyEvaluationModeEnum.to_proto(
            self.global_policy_evaluation_mode
        ):
            request.resource.global_policy_evaluation_mode = (
                PolicyGlobalPolicyEvaluationModeEnum.to_proto(
                    self.global_policy_evaluation_mode
                )
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyBinaryauthorizationBetaPolicy(request)
        self.admission_whitelist_patterns = (
            PolicyAdmissionWhitelistPatternsArray.from_proto(
                response.admission_whitelist_patterns
            )
        )
        self.cluster_admission_rules = Primitive.from_proto(
            response.cluster_admission_rules
        )
        self.kubernetes_namespace_admission_rules = Primitive.from_proto(
            response.kubernetes_namespace_admission_rules
        )
        self.kubernetes_service_account_admission_rules = Primitive.from_proto(
            response.kubernetes_service_account_admission_rules
        )
        self.istio_service_identity_admission_rules = Primitive.from_proto(
            response.istio_service_identity_admission_rules
        )
        self.default_admission_rule = PolicyDefaultAdmissionRule.from_proto(
            response.default_admission_rule
        )
        self.description = Primitive.from_proto(response.description)
        self.global_policy_evaluation_mode = (
            PolicyGlobalPolicyEvaluationModeEnum.from_proto(
                response.global_policy_evaluation_mode
            )
        )
        self.self_link = Primitive.from_proto(response.self_link)
        self.project = Primitive.from_proto(response.project)
        self.update_time = Primitive.from_proto(response.update_time)

    def delete(self):
        stub = policy_pb2_grpc.BinaryauthorizationBetaPolicyServiceStub(
            channel.Channel()
        )
        request = policy_pb2.DeleteBinaryauthorizationBetaPolicyRequest()
        request.service_account_file = self.service_account_file
        if PolicyAdmissionWhitelistPatternsArray.to_proto(
            self.admission_whitelist_patterns
        ):
            request.resource.admission_whitelist_patterns.extend(
                PolicyAdmissionWhitelistPatternsArray.to_proto(
                    self.admission_whitelist_patterns
                )
            )
        if Primitive.to_proto(self.cluster_admission_rules):
            request.resource.cluster_admission_rules = Primitive.to_proto(
                self.cluster_admission_rules
            )

        if Primitive.to_proto(self.kubernetes_namespace_admission_rules):
            request.resource.kubernetes_namespace_admission_rules = Primitive.to_proto(
                self.kubernetes_namespace_admission_rules
            )

        if Primitive.to_proto(self.kubernetes_service_account_admission_rules):
            request.resource.kubernetes_service_account_admission_rules = (
                Primitive.to_proto(self.kubernetes_service_account_admission_rules)
            )

        if Primitive.to_proto(self.istio_service_identity_admission_rules):
            request.resource.istio_service_identity_admission_rules = (
                Primitive.to_proto(self.istio_service_identity_admission_rules)
            )

        if PolicyDefaultAdmissionRule.to_proto(self.default_admission_rule):
            request.resource.default_admission_rule.CopyFrom(
                PolicyDefaultAdmissionRule.to_proto(self.default_admission_rule)
            )
        else:
            request.resource.ClearField("default_admission_rule")
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if PolicyGlobalPolicyEvaluationModeEnum.to_proto(
            self.global_policy_evaluation_mode
        ):
            request.resource.global_policy_evaluation_mode = (
                PolicyGlobalPolicyEvaluationModeEnum.to_proto(
                    self.global_policy_evaluation_mode
                )
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteBinaryauthorizationBetaPolicy(request)

    @classmethod
    def list(self, service_account_file=""):
        stub = policy_pb2_grpc.BinaryauthorizationBetaPolicyServiceStub(
            channel.Channel()
        )
        request = policy_pb2.ListBinaryauthorizationBetaPolicyRequest()
        request.service_account_file = service_account_file
        return stub.ListBinaryauthorizationBetaPolicy(request).items

    def to_proto(self):
        resource = policy_pb2.BinaryauthorizationBetaPolicy()
        if PolicyAdmissionWhitelistPatternsArray.to_proto(
            self.admission_whitelist_patterns
        ):
            resource.admission_whitelist_patterns.extend(
                PolicyAdmissionWhitelistPatternsArray.to_proto(
                    self.admission_whitelist_patterns
                )
            )
        if Primitive.to_proto(self.cluster_admission_rules):
            resource.cluster_admission_rules = Primitive.to_proto(
                self.cluster_admission_rules
            )
        if Primitive.to_proto(self.kubernetes_namespace_admission_rules):
            resource.kubernetes_namespace_admission_rules = Primitive.to_proto(
                self.kubernetes_namespace_admission_rules
            )
        if Primitive.to_proto(self.kubernetes_service_account_admission_rules):
            resource.kubernetes_service_account_admission_rules = Primitive.to_proto(
                self.kubernetes_service_account_admission_rules
            )
        if Primitive.to_proto(self.istio_service_identity_admission_rules):
            resource.istio_service_identity_admission_rules = Primitive.to_proto(
                self.istio_service_identity_admission_rules
            )
        if PolicyDefaultAdmissionRule.to_proto(self.default_admission_rule):
            resource.default_admission_rule.CopyFrom(
                PolicyDefaultAdmissionRule.to_proto(self.default_admission_rule)
            )
        else:
            resource.ClearField("default_admission_rule")
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if PolicyGlobalPolicyEvaluationModeEnum.to_proto(
            self.global_policy_evaluation_mode
        ):
            resource.global_policy_evaluation_mode = (
                PolicyGlobalPolicyEvaluationModeEnum.to_proto(
                    self.global_policy_evaluation_mode
                )
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class PolicyAdmissionWhitelistPatterns(object):
    def __init__(self, name_pattern: str = None):
        self.name_pattern = name_pattern

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = policy_pb2.BinaryauthorizationBetaPolicyAdmissionWhitelistPatterns()
        if Primitive.to_proto(resource.name_pattern):
            res.name_pattern = Primitive.to_proto(resource.name_pattern)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PolicyAdmissionWhitelistPatterns(
            name_pattern=Primitive.from_proto(resource.name_pattern),
        )


class PolicyAdmissionWhitelistPatternsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PolicyAdmissionWhitelistPatterns.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PolicyAdmissionWhitelistPatterns.from_proto(i) for i in resources]


class PolicyClusterAdmissionRules(object):
    def __init__(
        self,
        evaluation_mode: str = None,
        require_attestations_by: list = None,
        enforcement_mode: str = None,
    ):
        self.evaluation_mode = evaluation_mode
        self.require_attestations_by = require_attestations_by
        self.enforcement_mode = enforcement_mode

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = policy_pb2.BinaryauthorizationBetaPolicyClusterAdmissionRules()
        if PolicyClusterAdmissionRulesEvaluationModeEnum.to_proto(
            resource.evaluation_mode
        ):
            res.evaluation_mode = (
                PolicyClusterAdmissionRulesEvaluationModeEnum.to_proto(
                    resource.evaluation_mode
                )
            )
        if Primitive.to_proto(resource.require_attestations_by):
            res.require_attestations_by.extend(
                Primitive.to_proto(resource.require_attestations_by)
            )
        if PolicyClusterAdmissionRulesEnforcementModeEnum.to_proto(
            resource.enforcement_mode
        ):
            res.enforcement_mode = (
                PolicyClusterAdmissionRulesEnforcementModeEnum.to_proto(
                    resource.enforcement_mode
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PolicyClusterAdmissionRules(
            evaluation_mode=PolicyClusterAdmissionRulesEvaluationModeEnum.from_proto(
                resource.evaluation_mode
            ),
            require_attestations_by=Primitive.from_proto(
                resource.require_attestations_by
            ),
            enforcement_mode=PolicyClusterAdmissionRulesEnforcementModeEnum.from_proto(
                resource.enforcement_mode
            ),
        )


class PolicyClusterAdmissionRulesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PolicyClusterAdmissionRules.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PolicyClusterAdmissionRules.from_proto(i) for i in resources]


class PolicyKubernetesNamespaceAdmissionRules(object):
    def __init__(
        self,
        evaluation_mode: str = None,
        require_attestations_by: list = None,
        enforcement_mode: str = None,
    ):
        self.evaluation_mode = evaluation_mode
        self.require_attestations_by = require_attestations_by
        self.enforcement_mode = enforcement_mode

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            policy_pb2.BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRules()
        )
        if PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum.to_proto(
            resource.evaluation_mode
        ):
            res.evaluation_mode = (
                PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum.to_proto(
                    resource.evaluation_mode
                )
            )
        if Primitive.to_proto(resource.require_attestations_by):
            res.require_attestations_by.extend(
                Primitive.to_proto(resource.require_attestations_by)
            )
        if PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum.to_proto(
            resource.enforcement_mode
        ):
            res.enforcement_mode = (
                PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum.to_proto(
                    resource.enforcement_mode
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PolicyKubernetesNamespaceAdmissionRules(
            evaluation_mode=PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum.from_proto(
                resource.evaluation_mode
            ),
            require_attestations_by=Primitive.from_proto(
                resource.require_attestations_by
            ),
            enforcement_mode=PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum.from_proto(
                resource.enforcement_mode
            ),
        )


class PolicyKubernetesNamespaceAdmissionRulesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PolicyKubernetesNamespaceAdmissionRules.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            PolicyKubernetesNamespaceAdmissionRules.from_proto(i) for i in resources
        ]


class PolicyKubernetesServiceAccountAdmissionRules(object):
    def __init__(
        self,
        evaluation_mode: str = None,
        require_attestations_by: list = None,
        enforcement_mode: str = None,
    ):
        self.evaluation_mode = evaluation_mode
        self.require_attestations_by = require_attestations_by
        self.enforcement_mode = enforcement_mode

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            policy_pb2.BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRules()
        )
        if PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum.to_proto(
            resource.evaluation_mode
        ):
            res.evaluation_mode = (
                PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum.to_proto(
                    resource.evaluation_mode
                )
            )
        if Primitive.to_proto(resource.require_attestations_by):
            res.require_attestations_by.extend(
                Primitive.to_proto(resource.require_attestations_by)
            )
        if PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum.to_proto(
            resource.enforcement_mode
        ):
            res.enforcement_mode = PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum.to_proto(
                resource.enforcement_mode
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PolicyKubernetesServiceAccountAdmissionRules(
            evaluation_mode=PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum.from_proto(
                resource.evaluation_mode
            ),
            require_attestations_by=Primitive.from_proto(
                resource.require_attestations_by
            ),
            enforcement_mode=PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum.from_proto(
                resource.enforcement_mode
            ),
        )


class PolicyKubernetesServiceAccountAdmissionRulesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            PolicyKubernetesServiceAccountAdmissionRules.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            PolicyKubernetesServiceAccountAdmissionRules.from_proto(i)
            for i in resources
        ]


class PolicyIstioServiceIdentityAdmissionRules(object):
    def __init__(
        self,
        evaluation_mode: str = None,
        require_attestations_by: list = None,
        enforcement_mode: str = None,
    ):
        self.evaluation_mode = evaluation_mode
        self.require_attestations_by = require_attestations_by
        self.enforcement_mode = enforcement_mode

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            policy_pb2.BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRules()
        )
        if PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum.to_proto(
            resource.evaluation_mode
        ):
            res.evaluation_mode = (
                PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum.to_proto(
                    resource.evaluation_mode
                )
            )
        if Primitive.to_proto(resource.require_attestations_by):
            res.require_attestations_by.extend(
                Primitive.to_proto(resource.require_attestations_by)
            )
        if PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum.to_proto(
            resource.enforcement_mode
        ):
            res.enforcement_mode = (
                PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum.to_proto(
                    resource.enforcement_mode
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PolicyIstioServiceIdentityAdmissionRules(
            evaluation_mode=PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum.from_proto(
                resource.evaluation_mode
            ),
            require_attestations_by=Primitive.from_proto(
                resource.require_attestations_by
            ),
            enforcement_mode=PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum.from_proto(
                resource.enforcement_mode
            ),
        )


class PolicyIstioServiceIdentityAdmissionRulesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PolicyIstioServiceIdentityAdmissionRules.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            PolicyIstioServiceIdentityAdmissionRules.from_proto(i) for i in resources
        ]


class PolicyDefaultAdmissionRule(object):
    def __init__(
        self,
        evaluation_mode: str = None,
        require_attestations_by: list = None,
        enforcement_mode: str = None,
    ):
        self.evaluation_mode = evaluation_mode
        self.require_attestations_by = require_attestations_by
        self.enforcement_mode = enforcement_mode

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = policy_pb2.BinaryauthorizationBetaPolicyDefaultAdmissionRule()
        if PolicyDefaultAdmissionRuleEvaluationModeEnum.to_proto(
            resource.evaluation_mode
        ):
            res.evaluation_mode = PolicyDefaultAdmissionRuleEvaluationModeEnum.to_proto(
                resource.evaluation_mode
            )
        if Primitive.to_proto(resource.require_attestations_by):
            res.require_attestations_by.extend(
                Primitive.to_proto(resource.require_attestations_by)
            )
        if PolicyDefaultAdmissionRuleEnforcementModeEnum.to_proto(
            resource.enforcement_mode
        ):
            res.enforcement_mode = (
                PolicyDefaultAdmissionRuleEnforcementModeEnum.to_proto(
                    resource.enforcement_mode
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return PolicyDefaultAdmissionRule(
            evaluation_mode=PolicyDefaultAdmissionRuleEvaluationModeEnum.from_proto(
                resource.evaluation_mode
            ),
            require_attestations_by=Primitive.from_proto(
                resource.require_attestations_by
            ),
            enforcement_mode=PolicyDefaultAdmissionRuleEnforcementModeEnum.from_proto(
                resource.enforcement_mode
            ),
        )


class PolicyDefaultAdmissionRuleArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [PolicyDefaultAdmissionRule.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [PolicyDefaultAdmissionRule.from_proto(i) for i in resources]


class PolicyClusterAdmissionRulesEvaluationModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return policy_pb2.BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum.Value(
            "BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return policy_pb2.BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum.Name(
            resource
        )[
            len(
                "BinaryauthorizationBetaPolicyClusterAdmissionRulesEvaluationModeEnum"
            ) :
        ]


class PolicyClusterAdmissionRulesEnforcementModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return policy_pb2.BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum.Value(
            "BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return policy_pb2.BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum.Name(
            resource
        )[
            len(
                "BinaryauthorizationBetaPolicyClusterAdmissionRulesEnforcementModeEnum"
            ) :
        ]


class PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return policy_pb2.BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum.Value(
            "BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return policy_pb2.BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum.Name(
            resource
        )[
            len(
                "BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum"
            ) :
        ]


class PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return policy_pb2.BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum.Value(
            "BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return policy_pb2.BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum.Name(
            resource
        )[
            len(
                "BinaryauthorizationBetaPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum"
            ) :
        ]


class PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return policy_pb2.BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum.Value(
            "BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return policy_pb2.BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum.Name(
            resource
        )[
            len(
                "BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum"
            ) :
        ]


class PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return policy_pb2.BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum.Value(
            "BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return policy_pb2.BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum.Name(
            resource
        )[
            len(
                "BinaryauthorizationBetaPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum"
            ) :
        ]


class PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return policy_pb2.BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum.Value(
            "BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return policy_pb2.BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum.Name(
            resource
        )[
            len(
                "BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum"
            ) :
        ]


class PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return policy_pb2.BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum.Value(
            "BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return policy_pb2.BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum.Name(
            resource
        )[
            len(
                "BinaryauthorizationBetaPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum"
            ) :
        ]


class PolicyDefaultAdmissionRuleEvaluationModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return policy_pb2.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum.Value(
            "BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return policy_pb2.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum.Name(
            resource
        )[
            len("BinaryauthorizationBetaPolicyDefaultAdmissionRuleEvaluationModeEnum") :
        ]


class PolicyDefaultAdmissionRuleEnforcementModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return policy_pb2.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum.Value(
            "BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return policy_pb2.BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum.Name(
            resource
        )[
            len(
                "BinaryauthorizationBetaPolicyDefaultAdmissionRuleEnforcementModeEnum"
            ) :
        ]


class PolicyGlobalPolicyEvaluationModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return policy_pb2.BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum.Value(
            "BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return (
            policy_pb2.BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum.Name(
                resource
            )[len("BinaryauthorizationBetaPolicyGlobalPolicyEvaluationModeEnum") :]
        )


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
