// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.AdmissionRegistration.V1Beta1
{
    /// <summary>
    /// ValidatingAdmissionPolicyBinding binds the ValidatingAdmissionPolicy with paramerized resources. ValidatingAdmissionPolicyBinding and parameter CRDs together define how cluster administrators configure policies for clusters.
    /// 
    /// For a given admission request, each binding will cause its policy to be evaluated N times, where N is 1 for policies/bindings that don't use params, otherwise N is the number of parameters selected by the binding.
    /// 
    /// The CEL expressions of a policy must have a computed CEL cost below the maximum CEL budget. Each evaluation of the policy is given an independent CEL cost budget. Adding/removing policies, bindings, or params can not affect whether a given (policy, binding, param) combination is within its own CEL budget.
    /// </summary>
    [KubernetesResourceType("kubernetes:admissionregistration.k8s.io/v1beta1:ValidatingAdmissionPolicyBinding")]
    public partial class ValidatingAdmissionPolicyBinding : KubernetesResource
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Output("apiVersion")]
        public Output<string> ApiVersion { get; private set; } = null!;

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
        /// </summary>
        [Output("metadata")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta> Metadata { get; private set; } = null!;

        /// <summary>
        /// Specification of the desired behavior of the ValidatingAdmissionPolicyBinding.
        /// </summary>
        [Output("spec")]
        public Output<Pulumi.Kubernetes.Types.Outputs.AdmissionRegistration.V1Beta1.ValidatingAdmissionPolicyBindingSpec> Spec { get; private set; } = null!;


        /// <summary>
        /// Create a ValidatingAdmissionPolicyBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ValidatingAdmissionPolicyBinding(string name, Pulumi.Kubernetes.Types.Inputs.AdmissionRegistration.V1Beta1.ValidatingAdmissionPolicyBindingArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:admissionregistration.k8s.io/v1beta1:ValidatingAdmissionPolicyBinding", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }
        internal ValidatingAdmissionPolicyBinding(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:admissionregistration.k8s.io/v1beta1:ValidatingAdmissionPolicyBinding", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private ValidatingAdmissionPolicyBinding(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("kubernetes:admissionregistration.k8s.io/v1beta1:ValidatingAdmissionPolicyBinding", name, null, MakeResourceOptions(options, id))
        {
        }

        private static Pulumi.Kubernetes.Types.Inputs.AdmissionRegistration.V1Beta1.ValidatingAdmissionPolicyBindingArgs? MakeArgs(Pulumi.Kubernetes.Types.Inputs.AdmissionRegistration.V1Beta1.ValidatingAdmissionPolicyBindingArgs? args)
        {
            args ??= new Pulumi.Kubernetes.Types.Inputs.AdmissionRegistration.V1Beta1.ValidatingAdmissionPolicyBindingArgs();
            args.ApiVersion = "admissionregistration.k8s.io/v1beta1";
            args.Kind = "ValidatingAdmissionPolicyBinding";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "kubernetes:admissionregistration.k8s.io/v1alpha1:ValidatingAdmissionPolicyBinding"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ValidatingAdmissionPolicyBinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ValidatingAdmissionPolicyBinding Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ValidatingAdmissionPolicyBinding(name, id, options);
        }
    }
}
namespace Pulumi.Kubernetes.Types.Inputs.AdmissionRegistration.V1Beta1
{

    public class ValidatingAdmissionPolicyBindingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Input("apiVersion")]
        public Input<string>? ApiVersion { get; set; }

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
        /// </summary>
        [Input("metadata")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs>? Metadata { get; set; }

        /// <summary>
        /// Specification of the desired behavior of the ValidatingAdmissionPolicyBinding.
        /// </summary>
        [Input("spec")]
        public Input<Pulumi.Kubernetes.Types.Inputs.AdmissionRegistration.V1Beta1.ValidatingAdmissionPolicyBindingSpecArgs>? Spec { get; set; }

        public ValidatingAdmissionPolicyBindingArgs()
        {
        }
        public static new ValidatingAdmissionPolicyBindingArgs Empty => new ValidatingAdmissionPolicyBindingArgs();
    }
}
