// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.AdmissionRegistration.V1Beta1
{

    /// <summary>
    /// ValidatingAdmissionPolicyStatus represents the status of an admission validation policy.
    /// </summary>
    public class ValidatingAdmissionPolicyStatusArgs : global::Pulumi.ResourceArgs
    {
        [Input("conditions")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ConditionArgs>? _conditions;

        /// <summary>
        /// The conditions represent the latest available observations of a policy's current state.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ConditionArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// The generation observed by the controller.
        /// </summary>
        [Input("observedGeneration")]
        public Input<int>? ObservedGeneration { get; set; }

        /// <summary>
        /// The results of type checking for each expression. Presence of this field indicates the completion of the type checking.
        /// </summary>
        [Input("typeChecking")]
        public Input<Pulumi.Kubernetes.Types.Inputs.AdmissionRegistration.V1Beta1.TypeCheckingArgs>? TypeChecking { get; set; }

        public ValidatingAdmissionPolicyStatusArgs()
        {
        }
        public static new ValidatingAdmissionPolicyStatusArgs Empty => new ValidatingAdmissionPolicyStatusArgs();
    }
}
