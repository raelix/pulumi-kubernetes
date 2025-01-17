// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/internal"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Patch resources are used to modify existing Kubernetes resources by using
// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
// ResourceQuota sets aggregate quota restrictions enforced per namespace
type ResourceQuotaPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// Spec defines the desired quota. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec ResourceQuotaSpecPatchPtrOutput `pulumi:"spec"`
	// Status defines the actual enforced quota and its current usage. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status ResourceQuotaStatusPatchPtrOutput `pulumi:"status"`
}

// NewResourceQuotaPatch registers a new resource with the given unique name, arguments, and options.
func NewResourceQuotaPatch(ctx *pulumi.Context,
	name string, args *ResourceQuotaPatchArgs, opts ...pulumi.ResourceOption) (*ResourceQuotaPatch, error) {
	if args == nil {
		args = &ResourceQuotaPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("ResourceQuota")
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResourceQuotaPatch
	err := ctx.RegisterResource("kubernetes:core/v1:ResourceQuotaPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceQuotaPatch gets an existing ResourceQuotaPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceQuotaPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceQuotaPatchState, opts ...pulumi.ResourceOption) (*ResourceQuotaPatch, error) {
	var resource ResourceQuotaPatch
	err := ctx.ReadResource("kubernetes:core/v1:ResourceQuotaPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceQuotaPatch resources.
type resourceQuotaPatchState struct {
}

type ResourceQuotaPatchState struct {
}

func (ResourceQuotaPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceQuotaPatchState)(nil)).Elem()
}

type resourceQuotaPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// Spec defines the desired quota. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *ResourceQuotaSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a ResourceQuotaPatch resource.
type ResourceQuotaPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	// Spec defines the desired quota. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec ResourceQuotaSpecPatchPtrInput
}

func (ResourceQuotaPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceQuotaPatchArgs)(nil)).Elem()
}

type ResourceQuotaPatchInput interface {
	pulumi.Input

	ToResourceQuotaPatchOutput() ResourceQuotaPatchOutput
	ToResourceQuotaPatchOutputWithContext(ctx context.Context) ResourceQuotaPatchOutput
}

func (*ResourceQuotaPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceQuotaPatch)(nil)).Elem()
}

func (i *ResourceQuotaPatch) ToResourceQuotaPatchOutput() ResourceQuotaPatchOutput {
	return i.ToResourceQuotaPatchOutputWithContext(context.Background())
}

func (i *ResourceQuotaPatch) ToResourceQuotaPatchOutputWithContext(ctx context.Context) ResourceQuotaPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceQuotaPatchOutput)
}

// ResourceQuotaPatchArrayInput is an input type that accepts ResourceQuotaPatchArray and ResourceQuotaPatchArrayOutput values.
// You can construct a concrete instance of `ResourceQuotaPatchArrayInput` via:
//
//	ResourceQuotaPatchArray{ ResourceQuotaPatchArgs{...} }
type ResourceQuotaPatchArrayInput interface {
	pulumi.Input

	ToResourceQuotaPatchArrayOutput() ResourceQuotaPatchArrayOutput
	ToResourceQuotaPatchArrayOutputWithContext(context.Context) ResourceQuotaPatchArrayOutput
}

type ResourceQuotaPatchArray []ResourceQuotaPatchInput

func (ResourceQuotaPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceQuotaPatch)(nil)).Elem()
}

func (i ResourceQuotaPatchArray) ToResourceQuotaPatchArrayOutput() ResourceQuotaPatchArrayOutput {
	return i.ToResourceQuotaPatchArrayOutputWithContext(context.Background())
}

func (i ResourceQuotaPatchArray) ToResourceQuotaPatchArrayOutputWithContext(ctx context.Context) ResourceQuotaPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceQuotaPatchArrayOutput)
}

// ResourceQuotaPatchMapInput is an input type that accepts ResourceQuotaPatchMap and ResourceQuotaPatchMapOutput values.
// You can construct a concrete instance of `ResourceQuotaPatchMapInput` via:
//
//	ResourceQuotaPatchMap{ "key": ResourceQuotaPatchArgs{...} }
type ResourceQuotaPatchMapInput interface {
	pulumi.Input

	ToResourceQuotaPatchMapOutput() ResourceQuotaPatchMapOutput
	ToResourceQuotaPatchMapOutputWithContext(context.Context) ResourceQuotaPatchMapOutput
}

type ResourceQuotaPatchMap map[string]ResourceQuotaPatchInput

func (ResourceQuotaPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceQuotaPatch)(nil)).Elem()
}

func (i ResourceQuotaPatchMap) ToResourceQuotaPatchMapOutput() ResourceQuotaPatchMapOutput {
	return i.ToResourceQuotaPatchMapOutputWithContext(context.Background())
}

func (i ResourceQuotaPatchMap) ToResourceQuotaPatchMapOutputWithContext(ctx context.Context) ResourceQuotaPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceQuotaPatchMapOutput)
}

type ResourceQuotaPatchOutput struct{ *pulumi.OutputState }

func (ResourceQuotaPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceQuotaPatch)(nil)).Elem()
}

func (o ResourceQuotaPatchOutput) ToResourceQuotaPatchOutput() ResourceQuotaPatchOutput {
	return o
}

func (o ResourceQuotaPatchOutput) ToResourceQuotaPatchOutputWithContext(ctx context.Context) ResourceQuotaPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ResourceQuotaPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceQuotaPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ResourceQuotaPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceQuotaPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o ResourceQuotaPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *ResourceQuotaPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// Spec defines the desired quota. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o ResourceQuotaPatchOutput) Spec() ResourceQuotaSpecPatchPtrOutput {
	return o.ApplyT(func(v *ResourceQuotaPatch) ResourceQuotaSpecPatchPtrOutput { return v.Spec }).(ResourceQuotaSpecPatchPtrOutput)
}

// Status defines the actual enforced quota and its current usage. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o ResourceQuotaPatchOutput) Status() ResourceQuotaStatusPatchPtrOutput {
	return o.ApplyT(func(v *ResourceQuotaPatch) ResourceQuotaStatusPatchPtrOutput { return v.Status }).(ResourceQuotaStatusPatchPtrOutput)
}

type ResourceQuotaPatchArrayOutput struct{ *pulumi.OutputState }

func (ResourceQuotaPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceQuotaPatch)(nil)).Elem()
}

func (o ResourceQuotaPatchArrayOutput) ToResourceQuotaPatchArrayOutput() ResourceQuotaPatchArrayOutput {
	return o
}

func (o ResourceQuotaPatchArrayOutput) ToResourceQuotaPatchArrayOutputWithContext(ctx context.Context) ResourceQuotaPatchArrayOutput {
	return o
}

func (o ResourceQuotaPatchArrayOutput) Index(i pulumi.IntInput) ResourceQuotaPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourceQuotaPatch {
		return vs[0].([]*ResourceQuotaPatch)[vs[1].(int)]
	}).(ResourceQuotaPatchOutput)
}

type ResourceQuotaPatchMapOutput struct{ *pulumi.OutputState }

func (ResourceQuotaPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceQuotaPatch)(nil)).Elem()
}

func (o ResourceQuotaPatchMapOutput) ToResourceQuotaPatchMapOutput() ResourceQuotaPatchMapOutput {
	return o
}

func (o ResourceQuotaPatchMapOutput) ToResourceQuotaPatchMapOutputWithContext(ctx context.Context) ResourceQuotaPatchMapOutput {
	return o
}

func (o ResourceQuotaPatchMapOutput) MapIndex(k pulumi.StringInput) ResourceQuotaPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourceQuotaPatch {
		return vs[0].(map[string]*ResourceQuotaPatch)[vs[1].(string)]
	}).(ResourceQuotaPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceQuotaPatchInput)(nil)).Elem(), &ResourceQuotaPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceQuotaPatchArrayInput)(nil)).Elem(), ResourceQuotaPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceQuotaPatchMapInput)(nil)).Elem(), ResourceQuotaPatchMap{})
	pulumi.RegisterOutputType(ResourceQuotaPatchOutput{})
	pulumi.RegisterOutputType(ResourceQuotaPatchArrayOutput{})
	pulumi.RegisterOutputType(ResourceQuotaPatchMapOutput{})
}
