// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Core.V1
{

    /// <summary>
    /// HTTPHeader describes a custom header to be used in HTTP probes
    /// </summary>
    public class HTTPHeaderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The header field value
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public HTTPHeaderArgs()
        {
        }
        public static new HTTPHeaderArgs Empty => new HTTPHeaderArgs();
    }
}
