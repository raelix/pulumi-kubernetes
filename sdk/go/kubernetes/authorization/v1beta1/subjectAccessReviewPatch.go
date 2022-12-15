// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Patch resources are used to modify existing Kubernetes resources by using
// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
// SubjectAccessReview checks whether or not a user or group can perform an action.
type SubjectAccessReviewPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput          `pulumi:"kind"`
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// Spec holds information about the request being evaluated
	Spec SubjectAccessReviewSpecPatchPtrOutput `pulumi:"spec"`
	// Status is filled in by the server and indicates whether the request is allowed or not
	Status SubjectAccessReviewStatusPatchPtrOutput `pulumi:"status"`
}

// NewSubjectAccessReviewPatch registers a new resource with the given unique name, arguments, and options.
func NewSubjectAccessReviewPatch(ctx *pulumi.Context,
	name string, args *SubjectAccessReviewPatchArgs, opts ...pulumi.ResourceOption) (*SubjectAccessReviewPatch, error) {
	if args == nil {
		args = &SubjectAccessReviewPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("authorization.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("SubjectAccessReview")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:authorization.k8s.io/v1:SubjectAccessReviewPatch"),
		},
	})
	opts = append(opts, aliases)
	var resource SubjectAccessReviewPatch
	err := ctx.RegisterResource("kubernetes:authorization.k8s.io/v1beta1:SubjectAccessReviewPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubjectAccessReviewPatch gets an existing SubjectAccessReviewPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubjectAccessReviewPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubjectAccessReviewPatchState, opts ...pulumi.ResourceOption) (*SubjectAccessReviewPatch, error) {
	var resource SubjectAccessReviewPatch
	err := ctx.ReadResource("kubernetes:authorization.k8s.io/v1beta1:SubjectAccessReviewPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubjectAccessReviewPatch resources.
type subjectAccessReviewPatchState struct {
}

type SubjectAccessReviewPatchState struct {
}

func (SubjectAccessReviewPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*subjectAccessReviewPatchState)(nil)).Elem()
}

type subjectAccessReviewPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string                 `pulumi:"kind"`
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// Spec holds information about the request being evaluated
	Spec *SubjectAccessReviewSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a SubjectAccessReviewPatch resource.
type SubjectAccessReviewPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPatchPtrInput
	// Spec holds information about the request being evaluated
	Spec SubjectAccessReviewSpecPatchPtrInput
}

func (SubjectAccessReviewPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subjectAccessReviewPatchArgs)(nil)).Elem()
}

type SubjectAccessReviewPatchInput interface {
	pulumi.Input

	ToSubjectAccessReviewPatchOutput() SubjectAccessReviewPatchOutput
	ToSubjectAccessReviewPatchOutputWithContext(ctx context.Context) SubjectAccessReviewPatchOutput
}

func (*SubjectAccessReviewPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**SubjectAccessReviewPatch)(nil)).Elem()
}

func (i *SubjectAccessReviewPatch) ToSubjectAccessReviewPatchOutput() SubjectAccessReviewPatchOutput {
	return i.ToSubjectAccessReviewPatchOutputWithContext(context.Background())
}

func (i *SubjectAccessReviewPatch) ToSubjectAccessReviewPatchOutputWithContext(ctx context.Context) SubjectAccessReviewPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubjectAccessReviewPatchOutput)
}

// SubjectAccessReviewPatchArrayInput is an input type that accepts SubjectAccessReviewPatchArray and SubjectAccessReviewPatchArrayOutput values.
// You can construct a concrete instance of `SubjectAccessReviewPatchArrayInput` via:
//
//	SubjectAccessReviewPatchArray{ SubjectAccessReviewPatchArgs{...} }
type SubjectAccessReviewPatchArrayInput interface {
	pulumi.Input

	ToSubjectAccessReviewPatchArrayOutput() SubjectAccessReviewPatchArrayOutput
	ToSubjectAccessReviewPatchArrayOutputWithContext(context.Context) SubjectAccessReviewPatchArrayOutput
}

type SubjectAccessReviewPatchArray []SubjectAccessReviewPatchInput

func (SubjectAccessReviewPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubjectAccessReviewPatch)(nil)).Elem()
}

func (i SubjectAccessReviewPatchArray) ToSubjectAccessReviewPatchArrayOutput() SubjectAccessReviewPatchArrayOutput {
	return i.ToSubjectAccessReviewPatchArrayOutputWithContext(context.Background())
}

func (i SubjectAccessReviewPatchArray) ToSubjectAccessReviewPatchArrayOutputWithContext(ctx context.Context) SubjectAccessReviewPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubjectAccessReviewPatchArrayOutput)
}

// SubjectAccessReviewPatchMapInput is an input type that accepts SubjectAccessReviewPatchMap and SubjectAccessReviewPatchMapOutput values.
// You can construct a concrete instance of `SubjectAccessReviewPatchMapInput` via:
//
//	SubjectAccessReviewPatchMap{ "key": SubjectAccessReviewPatchArgs{...} }
type SubjectAccessReviewPatchMapInput interface {
	pulumi.Input

	ToSubjectAccessReviewPatchMapOutput() SubjectAccessReviewPatchMapOutput
	ToSubjectAccessReviewPatchMapOutputWithContext(context.Context) SubjectAccessReviewPatchMapOutput
}

type SubjectAccessReviewPatchMap map[string]SubjectAccessReviewPatchInput

func (SubjectAccessReviewPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubjectAccessReviewPatch)(nil)).Elem()
}

func (i SubjectAccessReviewPatchMap) ToSubjectAccessReviewPatchMapOutput() SubjectAccessReviewPatchMapOutput {
	return i.ToSubjectAccessReviewPatchMapOutputWithContext(context.Background())
}

func (i SubjectAccessReviewPatchMap) ToSubjectAccessReviewPatchMapOutputWithContext(ctx context.Context) SubjectAccessReviewPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubjectAccessReviewPatchMapOutput)
}

type SubjectAccessReviewPatchOutput struct{ *pulumi.OutputState }

func (SubjectAccessReviewPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubjectAccessReviewPatch)(nil)).Elem()
}

func (o SubjectAccessReviewPatchOutput) ToSubjectAccessReviewPatchOutput() SubjectAccessReviewPatchOutput {
	return o
}

func (o SubjectAccessReviewPatchOutput) ToSubjectAccessReviewPatchOutputWithContext(ctx context.Context) SubjectAccessReviewPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o SubjectAccessReviewPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubjectAccessReviewPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o SubjectAccessReviewPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubjectAccessReviewPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o SubjectAccessReviewPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *SubjectAccessReviewPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// Spec holds information about the request being evaluated
func (o SubjectAccessReviewPatchOutput) Spec() SubjectAccessReviewSpecPatchPtrOutput {
	return o.ApplyT(func(v *SubjectAccessReviewPatch) SubjectAccessReviewSpecPatchPtrOutput { return v.Spec }).(SubjectAccessReviewSpecPatchPtrOutput)
}

// Status is filled in by the server and indicates whether the request is allowed or not
func (o SubjectAccessReviewPatchOutput) Status() SubjectAccessReviewStatusPatchPtrOutput {
	return o.ApplyT(func(v *SubjectAccessReviewPatch) SubjectAccessReviewStatusPatchPtrOutput { return v.Status }).(SubjectAccessReviewStatusPatchPtrOutput)
}

type SubjectAccessReviewPatchArrayOutput struct{ *pulumi.OutputState }

func (SubjectAccessReviewPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubjectAccessReviewPatch)(nil)).Elem()
}

func (o SubjectAccessReviewPatchArrayOutput) ToSubjectAccessReviewPatchArrayOutput() SubjectAccessReviewPatchArrayOutput {
	return o
}

func (o SubjectAccessReviewPatchArrayOutput) ToSubjectAccessReviewPatchArrayOutputWithContext(ctx context.Context) SubjectAccessReviewPatchArrayOutput {
	return o
}

func (o SubjectAccessReviewPatchArrayOutput) Index(i pulumi.IntInput) SubjectAccessReviewPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SubjectAccessReviewPatch {
		return vs[0].([]*SubjectAccessReviewPatch)[vs[1].(int)]
	}).(SubjectAccessReviewPatchOutput)
}

type SubjectAccessReviewPatchMapOutput struct{ *pulumi.OutputState }

func (SubjectAccessReviewPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubjectAccessReviewPatch)(nil)).Elem()
}

func (o SubjectAccessReviewPatchMapOutput) ToSubjectAccessReviewPatchMapOutput() SubjectAccessReviewPatchMapOutput {
	return o
}

func (o SubjectAccessReviewPatchMapOutput) ToSubjectAccessReviewPatchMapOutputWithContext(ctx context.Context) SubjectAccessReviewPatchMapOutput {
	return o
}

func (o SubjectAccessReviewPatchMapOutput) MapIndex(k pulumi.StringInput) SubjectAccessReviewPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SubjectAccessReviewPatch {
		return vs[0].(map[string]*SubjectAccessReviewPatch)[vs[1].(string)]
	}).(SubjectAccessReviewPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SubjectAccessReviewPatchInput)(nil)).Elem(), &SubjectAccessReviewPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubjectAccessReviewPatchArrayInput)(nil)).Elem(), SubjectAccessReviewPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubjectAccessReviewPatchMapInput)(nil)).Elem(), SubjectAccessReviewPatchMap{})
	pulumi.RegisterOutputType(SubjectAccessReviewPatchOutput{})
	pulumi.RegisterOutputType(SubjectAccessReviewPatchArrayOutput{})
	pulumi.RegisterOutputType(SubjectAccessReviewPatchMapOutput{})
}
