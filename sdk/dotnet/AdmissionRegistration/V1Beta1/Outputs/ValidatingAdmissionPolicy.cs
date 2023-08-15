// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.AdmissionRegistration.V1Beta1
{

    /// <summary>
    /// ValidatingAdmissionPolicy describes the definition of an admission validation policy that accepts or rejects an object without changing it.
    /// </summary>
    [OutputType]
    public sealed class ValidatingAdmissionPolicy
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        public readonly string ApiVersion;
        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta Metadata;
        /// <summary>
        /// Specification of the desired behavior of the ValidatingAdmissionPolicy.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.AdmissionRegistration.V1Beta1.ValidatingAdmissionPolicySpec Spec;
        /// <summary>
        /// The status of the ValidatingAdmissionPolicy, including warnings that are useful to determine if the policy behaves in the expected way. Populated by the system. Read-only.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.AdmissionRegistration.V1Beta1.ValidatingAdmissionPolicyStatus Status;

        [OutputConstructor]
        private ValidatingAdmissionPolicy(
            string apiVersion,

            string kind,

            Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta metadata,

            Pulumi.Kubernetes.Types.Outputs.AdmissionRegistration.V1Beta1.ValidatingAdmissionPolicySpec spec,

            Pulumi.Kubernetes.Types.Outputs.AdmissionRegistration.V1Beta1.ValidatingAdmissionPolicyStatus status)
        {
            ApiVersion = apiVersion;
            Kind = kind;
            Metadata = metadata;
            Spec = spec;
            Status = status;
        }
    }
}
