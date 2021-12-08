// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// If present, the name of the kubeconfig cluster to use.
func GetCluster(ctx *pulumi.Context) string {
	return config.Get(ctx, "kubernetes:cluster")
}

// If present, the name of the kubeconfig context to use.
func GetContext(ctx *pulumi.Context) string {
	return config.Get(ctx, "kubernetes:context")
}

// BETA FEATURE - If present and set to true, enable server-side diff calculations.
// This feature is in developer preview, and is disabled by default.
//
// This config can be specified in the following ways, using this precedence:
// 1. This `enableDryRun` parameter.
// 2. The `PULUMI_K8S_ENABLE_DRY_RUN` environment variable.
func GetEnableDryRun(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "kubernetes:enableDryRun")
}

// BETA FEATURE - If present and set to true, replace CRDs on update rather than patching.
// This feature is in developer preview, and is disabled by default.
//
// This config can be specified in the following ways, using this precedence:
// 1. This `enableReplaceCRD` parameter.
// 2. The `PULUMI_K8S_ENABLE_REPLACE_CRD` environment variable.
func GetEnableReplaceCRD(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "kubernetes:enableReplaceCRD")
}

// The contents of a kubeconfig file or the path to a kubeconfig file. If this is set, this config will be used instead of $KUBECONFIG.
func GetKubeconfig(ctx *pulumi.Context) string {
	return config.Get(ctx, "kubernetes:kubeconfig")
}

// If present, the default namespace to use. This flag is ignored for cluster-scoped resources.
//
// A namespace can be specified in multiple places, and the precedence is as follows:
// 1. `.metadata.namespace` set on the resource.
// 2. This `namespace` parameter.
// 3. `namespace` set for the active context in the kubeconfig.
func GetNamespace(ctx *pulumi.Context) string {
	return config.Get(ctx, "kubernetes:namespace")
}

// BETA FEATURE - If present, render resource manifests to this directory. In this mode, resources will not
// be created on a Kubernetes cluster, but the rendered manifests will be kept in sync with changes
// to the Pulumi program. This feature is in developer preview, and is disabled by default.
//
// Note that some computed Outputs such as status fields will not be populated
// since the resources are not created on a Kubernetes cluster. These Output values will remain undefined,
// and may result in an error if they are referenced by other resources. Also note that any secret values
// used in these resources will be rendered in plaintext to the resulting YAML.
func GetRenderYamlToDirectory(ctx *pulumi.Context) string {
	return config.Get(ctx, "kubernetes:renderYamlToDirectory")
}

// If present and set to true, suppress apiVersion deprecation warnings from the CLI.
//
// This config can be specified in the following ways, using this precedence:
// 1. This `suppressDeprecationWarnings` parameter.
// 2. The `PULUMI_K8S_SUPPRESS_DEPRECATION_WARNINGS` environment variable.
func GetSuppressDeprecationWarnings(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "kubernetes:suppressDeprecationWarnings")
}

// If present and set to true, suppress unsupported Helm hook warnings from the CLI.
//
// This config can be specified in the following ways, using this precedence:
// 1. This `suppressHelmHookWarnings` parameter.
// 2. The `PULUMI_K8S_SUPPRESS_HELM_HOOK_WARNINGS` environment variable.
func GetSuppressHelmHookWarnings(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "kubernetes:suppressHelmHookWarnings")
}
