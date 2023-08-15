// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Authentication.V1
{

    /// <summary>
    /// SelfSubjectReviewStatus is filled by the kube-apiserver and sent back to a user.
    /// </summary>
    [OutputType]
    public sealed class SelfSubjectReviewStatusPatch
    {
        /// <summary>
        /// User attributes of the user making this request.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Authentication.V1.UserInfoPatch UserInfo;

        [OutputConstructor]
        private SelfSubjectReviewStatusPatch(Pulumi.Kubernetes.Types.Outputs.Authentication.V1.UserInfoPatch userInfo)
        {
            UserInfo = userInfo;
        }
    }
}
