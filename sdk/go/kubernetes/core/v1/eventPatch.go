// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Event is a report of an event somewhere in the cluster.  Events have a limited retention time and triggers and messages may evolve with time.  Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason.  Events should be treated as informative, best-effort, supplemental data.
type EventPatch struct {
	pulumi.CustomResourceState

	// What action was taken/failed regarding to the Regarding object.
	Action pulumi.StringPtrOutput `pulumi:"action"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// The number of times this event has occurred.
	Count pulumi.IntPtrOutput `pulumi:"count"`
	// Time when this Event was first observed.
	EventTime pulumi.StringPtrOutput `pulumi:"eventTime"`
	// The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)
	FirstTimestamp pulumi.StringPtrOutput `pulumi:"firstTimestamp"`
	// The object that this event is about.
	InvolvedObject ObjectReferencePatchPtrOutput `pulumi:"involvedObject"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The time at which the most recent occurrence of this event was recorded.
	LastTimestamp pulumi.StringPtrOutput `pulumi:"lastTimestamp"`
	// A human-readable description of the status of this operation.
	Message pulumi.StringPtrOutput `pulumi:"message"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// This should be a short, machine understandable string that gives the reason for the transition into the object's current status.
	Reason pulumi.StringPtrOutput `pulumi:"reason"`
	// Optional secondary object for more complex actions.
	Related ObjectReferencePatchPtrOutput `pulumi:"related"`
	// Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	ReportingComponent pulumi.StringPtrOutput `pulumi:"reportingComponent"`
	// ID of the controller instance, e.g. `kubelet-xyzf`.
	ReportingInstance pulumi.StringPtrOutput `pulumi:"reportingInstance"`
	// Data about the Event series this event represents or nil if it's a singleton Event.
	Series EventSeriesPatchPtrOutput `pulumi:"series"`
	// The component reporting this event. Should be a short machine understandable string.
	Source EventSourcePatchPtrOutput `pulumi:"source"`
	// Type of this event (Normal, Warning), new types could be added in the future
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewEventPatch registers a new resource with the given unique name, arguments, and options.
func NewEventPatch(ctx *pulumi.Context,
	name string, args *EventPatchArgs, opts ...pulumi.ResourceOption) (*EventPatch, error) {
	if args == nil {
		args = &EventPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("Event")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:events.k8s.io/v1:EventPatch"),
		},
		{
			Type: pulumi.String("kubernetes:events.k8s.io/v1beta1:EventPatch"),
		},
	})
	opts = append(opts, aliases)
	var resource EventPatch
	err := ctx.RegisterResource("kubernetes:core/v1:EventPatch", name, args, &resource, opts...)
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
	err := ctx.ReadResource("kubernetes:core/v1:EventPatch", name, id, state, &resource, opts...)
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
	// What action was taken/failed regarding to the Regarding object.
	Action *string `pulumi:"action"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// The number of times this event has occurred.
	Count *int `pulumi:"count"`
	// Time when this Event was first observed.
	EventTime *string `pulumi:"eventTime"`
	// The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)
	FirstTimestamp *string `pulumi:"firstTimestamp"`
	// The object that this event is about.
	InvolvedObject *ObjectReferencePatch `pulumi:"involvedObject"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// The time at which the most recent occurrence of this event was recorded.
	LastTimestamp *string `pulumi:"lastTimestamp"`
	// A human-readable description of the status of this operation.
	Message *string `pulumi:"message"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// This should be a short, machine understandable string that gives the reason for the transition into the object's current status.
	Reason *string `pulumi:"reason"`
	// Optional secondary object for more complex actions.
	Related *ObjectReferencePatch `pulumi:"related"`
	// Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	ReportingComponent *string `pulumi:"reportingComponent"`
	// ID of the controller instance, e.g. `kubelet-xyzf`.
	ReportingInstance *string `pulumi:"reportingInstance"`
	// Data about the Event series this event represents or nil if it's a singleton Event.
	Series *EventSeriesPatch `pulumi:"series"`
	// The component reporting this event. Should be a short machine understandable string.
	Source *EventSourcePatch `pulumi:"source"`
	// Type of this event (Normal, Warning), new types could be added in the future
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a EventPatch resource.
type EventPatchArgs struct {
	// What action was taken/failed regarding to the Regarding object.
	Action pulumi.StringPtrInput
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// The number of times this event has occurred.
	Count pulumi.IntPtrInput
	// Time when this Event was first observed.
	EventTime pulumi.StringPtrInput
	// The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)
	FirstTimestamp pulumi.StringPtrInput
	// The object that this event is about.
	InvolvedObject ObjectReferencePatchPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// The time at which the most recent occurrence of this event was recorded.
	LastTimestamp pulumi.StringPtrInput
	// A human-readable description of the status of this operation.
	Message pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	// This should be a short, machine understandable string that gives the reason for the transition into the object's current status.
	Reason pulumi.StringPtrInput
	// Optional secondary object for more complex actions.
	Related ObjectReferencePatchPtrInput
	// Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	ReportingComponent pulumi.StringPtrInput
	// ID of the controller instance, e.g. `kubelet-xyzf`.
	ReportingInstance pulumi.StringPtrInput
	// Data about the Event series this event represents or nil if it's a singleton Event.
	Series EventSeriesPatchPtrInput
	// The component reporting this event. Should be a short machine understandable string.
	Source EventSourcePatchPtrInput
	// Type of this event (Normal, Warning), new types could be added in the future
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
//          EventPatchArray{ EventPatchArgs{...} }
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
//          EventPatchMap{ "key": EventPatchArgs{...} }
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

// What action was taken/failed regarding to the Regarding object.
func (o EventPatchOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventPatch) pulumi.StringPtrOutput { return v.Action }).(pulumi.StringPtrOutput)
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o EventPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// The number of times this event has occurred.
func (o EventPatchOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EventPatch) pulumi.IntPtrOutput { return v.Count }).(pulumi.IntPtrOutput)
}

// Time when this Event was first observed.
func (o EventPatchOutput) EventTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventPatch) pulumi.StringPtrOutput { return v.EventTime }).(pulumi.StringPtrOutput)
}

// The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)
func (o EventPatchOutput) FirstTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventPatch) pulumi.StringPtrOutput { return v.FirstTimestamp }).(pulumi.StringPtrOutput)
}

// The object that this event is about.
func (o EventPatchOutput) InvolvedObject() ObjectReferencePatchPtrOutput {
	return o.ApplyT(func(v *EventPatch) ObjectReferencePatchPtrOutput { return v.InvolvedObject }).(ObjectReferencePatchPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o EventPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The time at which the most recent occurrence of this event was recorded.
func (o EventPatchOutput) LastTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventPatch) pulumi.StringPtrOutput { return v.LastTimestamp }).(pulumi.StringPtrOutput)
}

// A human-readable description of the status of this operation.
func (o EventPatchOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventPatch) pulumi.StringPtrOutput { return v.Message }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o EventPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *EventPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// This should be a short, machine understandable string that gives the reason for the transition into the object's current status.
func (o EventPatchOutput) Reason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventPatch) pulumi.StringPtrOutput { return v.Reason }).(pulumi.StringPtrOutput)
}

// Optional secondary object for more complex actions.
func (o EventPatchOutput) Related() ObjectReferencePatchPtrOutput {
	return o.ApplyT(func(v *EventPatch) ObjectReferencePatchPtrOutput { return v.Related }).(ObjectReferencePatchPtrOutput)
}

// Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
func (o EventPatchOutput) ReportingComponent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventPatch) pulumi.StringPtrOutput { return v.ReportingComponent }).(pulumi.StringPtrOutput)
}

// ID of the controller instance, e.g. `kubelet-xyzf`.
func (o EventPatchOutput) ReportingInstance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventPatch) pulumi.StringPtrOutput { return v.ReportingInstance }).(pulumi.StringPtrOutput)
}

// Data about the Event series this event represents or nil if it's a singleton Event.
func (o EventPatchOutput) Series() EventSeriesPatchPtrOutput {
	return o.ApplyT(func(v *EventPatch) EventSeriesPatchPtrOutput { return v.Series }).(EventSeriesPatchPtrOutput)
}

// The component reporting this event. Should be a short machine understandable string.
func (o EventPatchOutput) Source() EventSourcePatchPtrOutput {
	return o.ApplyT(func(v *EventPatch) EventSourcePatchPtrOutput { return v.Source }).(EventSourcePatchPtrOutput)
}

// Type of this event (Normal, Warning), new types could be added in the future
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