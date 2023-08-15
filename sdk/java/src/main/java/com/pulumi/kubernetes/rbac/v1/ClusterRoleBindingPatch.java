// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.rbac.v1;

import com.pulumi.core.Alias;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kubernetes.Utilities;
import com.pulumi.kubernetes.meta.v1.outputs.ObjectMetaPatch;
import com.pulumi.kubernetes.rbac.v1.ClusterRoleBindingPatchArgs;
import com.pulumi.kubernetes.rbac.v1.outputs.RoleRefPatch;
import com.pulumi.kubernetes.rbac.v1.outputs.SubjectPatch;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Patch resources are used to modify existing Kubernetes resources by using
 * Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
 * one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
 * Conflicts will result in an error by default, but can be forced using the &#34;pulumi.com/patchForce&#34; annotation. See the
 * [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
 * additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
 * ClusterRoleBinding references a ClusterRole, but not contain it.  It can reference a ClusterRole in the global namespace, and adds who information via Subject.
 * 
 */
@ResourceType(type="kubernetes:rbac.authorization.k8s.io/v1:ClusterRoleBindingPatch")
public class ClusterRoleBindingPatch extends com.pulumi.resources.CustomResource {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    @Export(name="apiVersion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> apiVersion;

    /**
     * @return APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    public Output<Optional<String>> apiVersion() {
        return Codegen.optional(this.apiVersion);
    }
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    @Export(name="kind", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kind;

    /**
     * @return Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    public Output<Optional<String>> kind() {
        return Codegen.optional(this.kind);
    }
    /**
     * Standard object&#39;s metadata.
     * 
     */
    @Export(name="metadata", refs={ObjectMetaPatch.class}, tree="[0]")
    private Output</* @Nullable */ ObjectMetaPatch> metadata;

    /**
     * @return Standard object&#39;s metadata.
     * 
     */
    public Output<Optional<ObjectMetaPatch>> metadata() {
        return Codegen.optional(this.metadata);
    }
    /**
     * RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error. This field is immutable.
     * 
     */
    @Export(name="roleRef", refs={RoleRefPatch.class}, tree="[0]")
    private Output</* @Nullable */ RoleRefPatch> roleRef;

    /**
     * @return RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error. This field is immutable.
     * 
     */
    public Output<Optional<RoleRefPatch>> roleRef() {
        return Codegen.optional(this.roleRef);
    }
    /**
     * Subjects holds references to the objects the role applies to.
     * 
     */
    @Export(name="subjects", refs={List.class,SubjectPatch.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SubjectPatch>> subjects;

    /**
     * @return Subjects holds references to the objects the role applies to.
     * 
     */
    public Output<Optional<List<SubjectPatch>>> subjects() {
        return Codegen.optional(this.subjects);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ClusterRoleBindingPatch(String name) {
        this(name, ClusterRoleBindingPatchArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ClusterRoleBindingPatch(String name, @Nullable ClusterRoleBindingPatchArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ClusterRoleBindingPatch(String name, @Nullable ClusterRoleBindingPatchArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:rbac.authorization.k8s.io/v1:ClusterRoleBindingPatch", name, makeArgs(args), makeResourceOptions(options, Codegen.empty()));
    }

    private ClusterRoleBindingPatch(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:rbac.authorization.k8s.io/v1:ClusterRoleBindingPatch", name, null, makeResourceOptions(options, id));
    }

    private static ClusterRoleBindingPatchArgs makeArgs(@Nullable ClusterRoleBindingPatchArgs args) {
        var builder = args == null ? ClusterRoleBindingPatchArgs.builder() : ClusterRoleBindingPatchArgs.builder(args);
        return builder
            .apiVersion("rbac.authorization.k8s.io/v1")
            .kind("ClusterRoleBinding")
            .build();
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .aliases(List.of(
                Output.of(Alias.builder().type("kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRoleBindingPatch").build()),
                Output.of(Alias.builder().type("kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRoleBindingPatch").build())
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static ClusterRoleBindingPatch get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ClusterRoleBindingPatch(name, id, options);
    }
}
