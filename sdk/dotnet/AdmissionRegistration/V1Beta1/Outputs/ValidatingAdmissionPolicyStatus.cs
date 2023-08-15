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
    /// ValidatingAdmissionPolicyStatus represents the status of an admission validation policy.
    /// </summary>
    [OutputType]
    public sealed class ValidatingAdmissionPolicyStatus
    {
        /// <summary>
        /// The conditions represent the latest available observations of a policy's current state.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Meta.V1.Condition> Conditions;
        /// <summary>
        /// The generation observed by the controller.
        /// </summary>
        public readonly int ObservedGeneration;
        /// <summary>
        /// The results of type checking for each expression. Presence of this field indicates the completion of the type checking.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.AdmissionRegistration.V1Beta1.TypeChecking TypeChecking;

        [OutputConstructor]
        private ValidatingAdmissionPolicyStatus(
            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Meta.V1.Condition> conditions,

            int observedGeneration,

            Pulumi.Kubernetes.Types.Outputs.AdmissionRegistration.V1Beta1.TypeChecking typeChecking)
        {
            Conditions = conditions;
            ObservedGeneration = observedGeneration;
            TypeChecking = typeChecking;
        }
    }
}
