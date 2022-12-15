// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

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
// Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system.
type EventPatch struct {
	pulumi.CustomResourceState

	// What action was taken/failed regarding to the regarding object.
	Action pulumi.StringPtrOutput `pulumi:"action"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedCount pulumi.IntPtrOutput `pulumi:"deprecatedCount"`
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedFirstTimestamp pulumi.StringPtrOutput `pulumi:"deprecatedFirstTimestamp"`
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedLastTimestamp pulumi.StringPtrOutput `pulumi:"deprecatedLastTimestamp"`
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedSource corev1.EventSourcePatchPtrOutput `pulumi:"deprecatedSource"`
	// Required. Time when this Event was first observed.
	EventTime pulumi.StringPtrOutput `pulumi:"eventTime"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput          `pulumi:"kind"`
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// Optional. A human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
	Note pulumi.StringPtrOutput `pulumi:"note"`
	// Why the action was taken.
	Reason pulumi.StringPtrOutput `pulumi:"reason"`
	// The object this Event is about. In most cases it's an Object reporting controller implements. E.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
	Regarding corev1.ObjectReferencePatchPtrOutput `pulumi:"regarding"`
	// Optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
	Related corev1.ObjectReferencePatchPtrOutput `pulumi:"related"`
	// Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	ReportingController pulumi.StringPtrOutput `pulumi:"reportingController"`
	// ID of the controller instance, e.g. `kubelet-xyzf`.
	ReportingInstance pulumi.StringPtrOutput `pulumi:"reportingInstance"`
	// Data about the Event series this event represents or nil if it's a singleton Event.
	Series EventSeriesPatchPtrOutput `pulumi:"series"`
	// Type of this event (Normal, Warning), new types could be added in the future.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewEventPatch registers a new resource with the given unique name, arguments, and options.
func NewEventPatch(ctx *pulumi.Context,
	name string, args *EventPatchArgs, opts ...pulumi.ResourceOption) (*EventPatch, error) {
	if args == nil {
		args = &EventPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("events.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("Event")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:core/v1:EventPatch"),
		},
		{
			Type: pulumi.String("kubernetes:events.k8s.io/v1:EventPatch"),
		},
	})
	opts = append(opts, aliases)
	var resource EventPatch
	err := ctx.RegisterResource("kubernetes:events.k8s.io/v1beta1:EventPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventPatch gets an existing EventPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventPatchState, opts ...pulumi.ResourceOption) (*EventPatch, error) {
	var resource EventPatch
	err := ctx.ReadResource("kubernetes:events.k8s.io/v1beta1:EventPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventPatch resources.
type eventPatchState struct {
}

type EventPatchState struct {
}

func (EventPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventPatchState)(nil)).Elem()
}

type eventPatchArgs struct {
	// What action was taken/failed regarding to the regarding object.
	Action *string `pulumi:"action"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedCount *int `pulumi:"deprecatedCount"`
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedFirstTimestamp *string `pulumi:"deprecatedFirstTimestamp"`
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedLastTimestamp *string `pulumi:"deprecatedLastTimestamp"`
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedSource *corev1.EventSourcePatch `pulumi:"deprecatedSource"`
	// Required. Time when this Event was first observed.
	EventTime *string `pulumi:"eventTime"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string                 `pulumi:"kind"`
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// Optional. A human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
	Note *string `pulumi:"note"`
	// Why the action was taken.
	Reason *string `pulumi:"reason"`
	// The object this Event is about. In most cases it's an Object reporting controller implements. E.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
	Regarding *corev1.ObjectReferencePatch `pulumi:"regarding"`
	// Optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
	Related *corev1.ObjectReferencePatch `pulumi:"related"`
	// Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	ReportingController *string `pulumi:"reportingController"`
	// ID of the controller instance, e.g. `kubelet-xyzf`.
	ReportingInstance *string `pulumi:"reportingInstance"`
	// Data about the Event series this event represents or nil if it's a singleton Event.
	Series *EventSeriesPatch `pulumi:"series"`
	// Type of this event (Normal, Warning), new types could be added in the future.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a EventPatch resource.
type EventPatchArgs struct {
	// What action was taken/failed regarding to the regarding object.
	Action pulumi.StringPtrInput
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedCount pulumi.IntPtrInput
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedFirstTimestamp pulumi.StringPtrInput
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedLastTimestamp pulumi.StringPtrInput
	// Deprecated field assuring backward compatibility with core.v1 Event type
	DeprecatedSource corev1.EventSourcePatchPtrInput
	// Required. Time when this Event was first observed.
	EventTime pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPatchPtrInput
	// Optional. A human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
	Note pulumi.StringPtrInput
	// Why the action was taken.
	Reason pulumi.StringPtrInput
	// The object this Event is about. In most cases it's an Object reporting controller implements. E.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
	Regarding corev1.ObjectReferencePatchPtrInput
	// Optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
	Related corev1.ObjectReferencePatchPtrInput
	// Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	ReportingController pulumi.StringPtrInput
	// ID of the controller instance, e.g. `kubelet-xyzf`.
	ReportingInstance pulumi.StringPtrInput
	// Data about the Event series this event represents or nil if it's a singleton Event.
	Series EventSeriesPatchPtrInput
	// Type of this event (Normal, Warning), new types could be added in the future.
	Type pulumi.StringPtrInput
}

func (EventPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventPatchArgs)(nil)).Elem()
}

type EventPatchInput interface {
	pulumi.Input

	ToEventPatchOutput() EventPatchOutput
	ToEventPatchOutputWithContext(ctx context.Context) EventPatchOutput
}

func (*EventPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**EventPatch)(nil)).Elem()
}

func (i *EventPatch) ToEventPatchOutput() EventPatchOutput {
	return i.ToEventPatchOutputWithContext(context.Background())
}

func (i *EventPatch) ToEventPatchOutputWithContext(ctx context.Context) EventPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventPatchOutput)
}

// EventPatchArrayInput is an input type that accepts EventPatchArray and EventPatchArrayOutput values.
// You can construct a concrete instance of `EventPatchArrayInput` via:
//
//	EventPatchArray{ EventPatchArgs{...} }
type EventPatchArrayInput interface {
	pulumi.Input

	ToEventPatchArrayOutput() EventPatchArrayOutput
	ToEventPatchArrayOutputWithContext(context.Context) EventPatchArrayOutput
}

type EventPatchArray []EventPatchInput

func (EventPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventPatch)(nil)).Elem()
}

func (i EventPatchArray) ToEventPatchArrayOutput() EventPatchArrayOutput {
	return i.ToEventPatchArrayOutputWithContext(context.Background())
}

func (i EventPatchArray) ToEventPatchArrayOutputWithContext(ctx context.Context) EventPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventPatchArrayOutput)
}

// EventPatchMapInput is an input type that accepts EventPatchMap and EventPatchMapOutput values.
// You can construct a concrete instance of `EventPatchMapInput` via:
//
//	EventPatchMap{ "key": EventPatchArgs{...} }
type EventPatchMapInput interface {
	pulumi.Input

	ToEventPatchMapOutput() EventPatchMapOutput
	ToEventPatchMapOutputWithContext(context.Context) EventPatchMapOutput
}

type EventPatchMap map[string]EventPatchInput

func (EventPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventPatch)(nil)).Elem()
}

func (i EventPatchMap) ToEventPatchMapOutput() EventPatchMapOutput {
	return i.ToEventPatchMapOutputWithContext(context.Background())
}

func (i EventPatchMap) ToEventPatchMapOutputWithContext(ctx context.Context) EventPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventPatchMapOutput)
}

type EventPatchOutput struct{ *pulumi.OutputState }

func (EventPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventPatch)(nil)).Elem()
}

func (o EventPatchOutput) ToEventPatchOutput() EventPatchOutput {
	return o
}

func (o EventPatchOutput) ToEventPatchOutputWithContext(ctx context.Context) EventPatchOutput {
	return o
}

// What action was taken/failed regarding to the regarding object.
func (o EventPatchOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventPatch) pulumi.StringPtrOutput { return v.Action }).(pulumi.StringPtrOutput)
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o EventPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Deprecated field assuring backward compatibility with core.v1 Event type
func (o EventPatchOutput) DeprecatedCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EventPatch) pulumi.IntPtrOutput { return v.DeprecatedCount }).(pulumi.IntPtrOutput)
}

// Deprecated field assuring backward compatibility with core.v1 Event type
func (o EventPatchOutput) DeprecatedFirstTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventPatch) pulumi.StringPtrOutput { return v.DeprecatedFirstTimestamp }).(pulumi.StringPtrOutput)
}

// Deprecated field assuring backward compatibility with core.v1 Event type
func (o EventPatchOutput) DeprecatedLastTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventPatch) pulumi.StringPtrOutput { return v.DeprecatedLastTimestamp }).(pulumi.StringPtrOutput)
}

// Deprecated field assuring backward compatibility with core.v1 Event type
func (o EventPatchOutput) DeprecatedSource() corev1.EventSourcePatchPtrOutput {
	return o.ApplyT(func(v *EventPatch) corev1.EventSourcePatchPtrOutput { return v.DeprecatedSource }).(corev1.EventSourcePatchPtrOutput)
}

// Required. Time when this Event was first observed.
func (o EventPatchOutput) EventTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventPatch) pulumi.StringPtrOutput { return v.EventTime }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o EventPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o EventPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *EventPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// Optional. A human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
func (o EventPatchOutput) Note() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventPatch) pulumi.StringPtrOutput { return v.Note }).(pulumi.StringPtrOutput)
}

// Why the action was taken.
func (o EventPatchOutput) Reason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventPatch) pulumi.StringPtrOutput { return v.Reason }).(pulumi.StringPtrOutput)
}

// The object this Event is about. In most cases it's an Object reporting controller implements. E.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
func (o EventPatchOutput) Regarding() corev1.ObjectReferencePatchPtrOutput {
	return o.ApplyT(func(v *EventPatch) corev1.ObjectReferencePatchPtrOutput { return v.Regarding }).(corev1.ObjectReferencePatchPtrOutput)
}

// Optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
func (o EventPatchOutput) Related() corev1.ObjectReferencePatchPtrOutput {
	return o.ApplyT(func(v *EventPatch) corev1.ObjectReferencePatchPtrOutput { return v.Related }).(corev1.ObjectReferencePatchPtrOutput)
}

// Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
func (o EventPatchOutput) ReportingController() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventPatch) pulumi.StringPtrOutput { return v.ReportingController }).(pulumi.StringPtrOutput)
}

// ID of the controller instance, e.g. `kubelet-xyzf`.
func (o EventPatchOutput) ReportingInstance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventPatch) pulumi.StringPtrOutput { return v.ReportingInstance }).(pulumi.StringPtrOutput)
}

// Data about the Event series this event represents or nil if it's a singleton Event.
func (o EventPatchOutput) Series() EventSeriesPatchPtrOutput {
	return o.ApplyT(func(v *EventPatch) EventSeriesPatchPtrOutput { return v.Series }).(EventSeriesPatchPtrOutput)
}

// Type of this event (Normal, Warning), new types could be added in the future.
func (o EventPatchOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventPatch) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type EventPatchArrayOutput struct{ *pulumi.OutputState }

func (EventPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventPatch)(nil)).Elem()
}

func (o EventPatchArrayOutput) ToEventPatchArrayOutput() EventPatchArrayOutput {
	return o
}

func (o EventPatchArrayOutput) ToEventPatchArrayOutputWithContext(ctx context.Context) EventPatchArrayOutput {
	return o
}

func (o EventPatchArrayOutput) Index(i pulumi.IntInput) EventPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EventPatch {
		return vs[0].([]*EventPatch)[vs[1].(int)]
	}).(EventPatchOutput)
}

type EventPatchMapOutput struct{ *pulumi.OutputState }

func (EventPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventPatch)(nil)).Elem()
}

func (o EventPatchMapOutput) ToEventPatchMapOutput() EventPatchMapOutput {
	return o
}

func (o EventPatchMapOutput) ToEventPatchMapOutputWithContext(ctx context.Context) EventPatchMapOutput {
	return o
}

func (o EventPatchMapOutput) MapIndex(k pulumi.StringInput) EventPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EventPatch {
		return vs[0].(map[string]*EventPatch)[vs[1].(string)]
	}).(EventPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EventPatchInput)(nil)).Elem(), &EventPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventPatchArrayInput)(nil)).Elem(), EventPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventPatchMapInput)(nil)).Elem(), EventPatchMap{})
	pulumi.RegisterOutputType(EventPatchOutput{})
	pulumi.RegisterOutputType(EventPatchArrayOutput{})
	pulumi.RegisterOutputType(EventPatchMapOutput{})
}
