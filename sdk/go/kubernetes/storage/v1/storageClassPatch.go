// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Patch resources are used to modify existing Kubernetes resources by using
// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
// StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.
//
// StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.
type StorageClassPatch struct {
	pulumi.CustomResourceState

	// AllowVolumeExpansion shows whether the storage class allow volume expand
	AllowVolumeExpansion pulumi.BoolPtrOutput `pulumi:"allowVolumeExpansion"`
	// Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
	AllowedTopologies corev1.TopologySelectorTermPatchArrayOutput `pulumi:"allowedTopologies"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
	MountOptions pulumi.StringArrayOutput `pulumi:"mountOptions"`
	// Parameters holds the parameters for the provisioner that should create volumes of this storage class.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// Provisioner indicates the type of the provisioner.
	Provisioner pulumi.StringPtrOutput `pulumi:"provisioner"`
	// Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
	ReclaimPolicy pulumi.StringPtrOutput `pulumi:"reclaimPolicy"`
	// VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
	VolumeBindingMode pulumi.StringPtrOutput `pulumi:"volumeBindingMode"`
}

// NewStorageClassPatch registers a new resource with the given unique name, arguments, and options.
func NewStorageClassPatch(ctx *pulumi.Context,
	name string, args *StorageClassPatchArgs, opts ...pulumi.ResourceOption) (*StorageClassPatch, error) {
	if args == nil {
		args = &StorageClassPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("storage.k8s.io/v1")
	args.Kind = pulumi.StringPtr("StorageClass")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:storage.k8s.io/v1beta1:StorageClassPatch"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageClassPatch
	err := ctx.RegisterResource("kubernetes:storage.k8s.io/v1:StorageClassPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageClassPatch gets an existing StorageClassPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageClassPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageClassPatchState, opts ...pulumi.ResourceOption) (*StorageClassPatch, error) {
	var resource StorageClassPatch
	err := ctx.ReadResource("kubernetes:storage.k8s.io/v1:StorageClassPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageClassPatch resources.
type storageClassPatchState struct {
}

type StorageClassPatchState struct {
}

func (StorageClassPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageClassPatchState)(nil)).Elem()
}

type storageClassPatchArgs struct {
	// AllowVolumeExpansion shows whether the storage class allow volume expand
	AllowVolumeExpansion *bool `pulumi:"allowVolumeExpansion"`
	// Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
	AllowedTopologies []corev1.TopologySelectorTermPatch `pulumi:"allowedTopologies"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
	MountOptions []string `pulumi:"mountOptions"`
	// Parameters holds the parameters for the provisioner that should create volumes of this storage class.
	Parameters map[string]string `pulumi:"parameters"`
	// Provisioner indicates the type of the provisioner.
	Provisioner *string `pulumi:"provisioner"`
	// Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
	ReclaimPolicy *string `pulumi:"reclaimPolicy"`
	// VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
	VolumeBindingMode *string `pulumi:"volumeBindingMode"`
}

// The set of arguments for constructing a StorageClassPatch resource.
type StorageClassPatchArgs struct {
	// AllowVolumeExpansion shows whether the storage class allow volume expand
	AllowVolumeExpansion pulumi.BoolPtrInput
	// Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
	AllowedTopologies corev1.TopologySelectorTermPatchArrayInput
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	// Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
	MountOptions pulumi.StringArrayInput
	// Parameters holds the parameters for the provisioner that should create volumes of this storage class.
	Parameters pulumi.StringMapInput
	// Provisioner indicates the type of the provisioner.
	Provisioner pulumi.StringPtrInput
	// Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
	ReclaimPolicy pulumi.StringPtrInput
	// VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
	VolumeBindingMode pulumi.StringPtrInput
}

func (StorageClassPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageClassPatchArgs)(nil)).Elem()
}

type StorageClassPatchInput interface {
	pulumi.Input

	ToStorageClassPatchOutput() StorageClassPatchOutput
	ToStorageClassPatchOutputWithContext(ctx context.Context) StorageClassPatchOutput
}

func (*StorageClassPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageClassPatch)(nil)).Elem()
}

func (i *StorageClassPatch) ToStorageClassPatchOutput() StorageClassPatchOutput {
	return i.ToStorageClassPatchOutputWithContext(context.Background())
}

func (i *StorageClassPatch) ToStorageClassPatchOutputWithContext(ctx context.Context) StorageClassPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageClassPatchOutput)
}

// StorageClassPatchArrayInput is an input type that accepts StorageClassPatchArray and StorageClassPatchArrayOutput values.
// You can construct a concrete instance of `StorageClassPatchArrayInput` via:
//
//	StorageClassPatchArray{ StorageClassPatchArgs{...} }
type StorageClassPatchArrayInput interface {
	pulumi.Input

	ToStorageClassPatchArrayOutput() StorageClassPatchArrayOutput
	ToStorageClassPatchArrayOutputWithContext(context.Context) StorageClassPatchArrayOutput
}

type StorageClassPatchArray []StorageClassPatchInput

func (StorageClassPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StorageClassPatch)(nil)).Elem()
}

func (i StorageClassPatchArray) ToStorageClassPatchArrayOutput() StorageClassPatchArrayOutput {
	return i.ToStorageClassPatchArrayOutputWithContext(context.Background())
}

func (i StorageClassPatchArray) ToStorageClassPatchArrayOutputWithContext(ctx context.Context) StorageClassPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageClassPatchArrayOutput)
}

// StorageClassPatchMapInput is an input type that accepts StorageClassPatchMap and StorageClassPatchMapOutput values.
// You can construct a concrete instance of `StorageClassPatchMapInput` via:
//
//	StorageClassPatchMap{ "key": StorageClassPatchArgs{...} }
type StorageClassPatchMapInput interface {
	pulumi.Input

	ToStorageClassPatchMapOutput() StorageClassPatchMapOutput
	ToStorageClassPatchMapOutputWithContext(context.Context) StorageClassPatchMapOutput
}

type StorageClassPatchMap map[string]StorageClassPatchInput

func (StorageClassPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StorageClassPatch)(nil)).Elem()
}

func (i StorageClassPatchMap) ToStorageClassPatchMapOutput() StorageClassPatchMapOutput {
	return i.ToStorageClassPatchMapOutputWithContext(context.Background())
}

func (i StorageClassPatchMap) ToStorageClassPatchMapOutputWithContext(ctx context.Context) StorageClassPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageClassPatchMapOutput)
}

type StorageClassPatchOutput struct{ *pulumi.OutputState }

func (StorageClassPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageClassPatch)(nil)).Elem()
}

func (o StorageClassPatchOutput) ToStorageClassPatchOutput() StorageClassPatchOutput {
	return o
}

func (o StorageClassPatchOutput) ToStorageClassPatchOutputWithContext(ctx context.Context) StorageClassPatchOutput {
	return o
}

// AllowVolumeExpansion shows whether the storage class allow volume expand
func (o StorageClassPatchOutput) AllowVolumeExpansion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageClassPatch) pulumi.BoolPtrOutput { return v.AllowVolumeExpansion }).(pulumi.BoolPtrOutput)
}

// Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
func (o StorageClassPatchOutput) AllowedTopologies() corev1.TopologySelectorTermPatchArrayOutput {
	return o.ApplyT(func(v *StorageClassPatch) corev1.TopologySelectorTermPatchArrayOutput { return v.AllowedTopologies }).(corev1.TopologySelectorTermPatchArrayOutput)
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o StorageClassPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageClassPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o StorageClassPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageClassPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o StorageClassPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *StorageClassPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
func (o StorageClassPatchOutput) MountOptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StorageClassPatch) pulumi.StringArrayOutput { return v.MountOptions }).(pulumi.StringArrayOutput)
}

// Parameters holds the parameters for the provisioner that should create volumes of this storage class.
func (o StorageClassPatchOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StorageClassPatch) pulumi.StringMapOutput { return v.Parameters }).(pulumi.StringMapOutput)
}

// Provisioner indicates the type of the provisioner.
func (o StorageClassPatchOutput) Provisioner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageClassPatch) pulumi.StringPtrOutput { return v.Provisioner }).(pulumi.StringPtrOutput)
}

// Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
func (o StorageClassPatchOutput) ReclaimPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageClassPatch) pulumi.StringPtrOutput { return v.ReclaimPolicy }).(pulumi.StringPtrOutput)
}

// VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
func (o StorageClassPatchOutput) VolumeBindingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageClassPatch) pulumi.StringPtrOutput { return v.VolumeBindingMode }).(pulumi.StringPtrOutput)
}

type StorageClassPatchArrayOutput struct{ *pulumi.OutputState }

func (StorageClassPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StorageClassPatch)(nil)).Elem()
}

func (o StorageClassPatchArrayOutput) ToStorageClassPatchArrayOutput() StorageClassPatchArrayOutput {
	return o
}

func (o StorageClassPatchArrayOutput) ToStorageClassPatchArrayOutputWithContext(ctx context.Context) StorageClassPatchArrayOutput {
	return o
}

func (o StorageClassPatchArrayOutput) Index(i pulumi.IntInput) StorageClassPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StorageClassPatch {
		return vs[0].([]*StorageClassPatch)[vs[1].(int)]
	}).(StorageClassPatchOutput)
}

type StorageClassPatchMapOutput struct{ *pulumi.OutputState }

func (StorageClassPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StorageClassPatch)(nil)).Elem()
}

func (o StorageClassPatchMapOutput) ToStorageClassPatchMapOutput() StorageClassPatchMapOutput {
	return o
}

func (o StorageClassPatchMapOutput) ToStorageClassPatchMapOutputWithContext(ctx context.Context) StorageClassPatchMapOutput {
	return o
}

func (o StorageClassPatchMapOutput) MapIndex(k pulumi.StringInput) StorageClassPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StorageClassPatch {
		return vs[0].(map[string]*StorageClassPatch)[vs[1].(string)]
	}).(StorageClassPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageClassPatchInput)(nil)).Elem(), &StorageClassPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageClassPatchArrayInput)(nil)).Elem(), StorageClassPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageClassPatchMapInput)(nil)).Elem(), StorageClassPatchMap{})
	pulumi.RegisterOutputType(StorageClassPatchOutput{})
	pulumi.RegisterOutputType(StorageClassPatchArrayOutput{})
	pulumi.RegisterOutputType(StorageClassPatchMapOutput{})
}
