// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Beta2
{

    /// <summary>
    /// PriorityLevelConfigurationSpec specifies the configuration of a priority level.
    /// </summary>
    [OutputType]
    public sealed class PriorityLevelConfigurationSpec
    {
        /// <summary>
        /// `exempt` specifies how requests are handled for an exempt priority level. This field MUST be empty if `type` is `"Limited"`. This field MAY be non-empty if `type` is `"Exempt"`. If empty and `type` is `"Exempt"` then the default values for `ExemptPriorityLevelConfiguration` apply.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Beta2.ExemptPriorityLevelConfiguration Exempt;
        /// <summary>
        /// `limited` specifies how requests are handled for a Limited priority level. This field must be non-empty if and only if `type` is `"Limited"`.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Beta2.LimitedPriorityLevelConfiguration Limited;
        /// <summary>
        /// `type` indicates whether this priority level is subject to limitation on request execution.  A value of `"Exempt"` means that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the capacity made available to other priority levels.  A value of `"Limited"` means that (a) requests of this priority level _are_ subject to limits and (b) some of the server's limited capacity is made available exclusively to this priority level. Required.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private PriorityLevelConfigurationSpec(
            Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Beta2.ExemptPriorityLevelConfiguration exempt,

            Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Beta2.LimitedPriorityLevelConfiguration limited,

            string type)
        {
            Exempt = exempt;
            Limited = limited;
            Type = type;
        }
    }
}
