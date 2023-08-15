// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.core.v1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.kubernetes.core.v1.outputs.ContainerStatusPatch;
import com.pulumi.kubernetes.core.v1.outputs.HostIPPatch;
import com.pulumi.kubernetes.core.v1.outputs.PodConditionPatch;
import com.pulumi.kubernetes.core.v1.outputs.PodIPPatch;
import com.pulumi.kubernetes.core.v1.outputs.PodResourceClaimStatusPatch;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PodStatusPatch {
    /**
     * @return Current service state of pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
     * 
     */
    private @Nullable List<PodConditionPatch> conditions;
    /**
     * @return The list has one entry per container in the manifest. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status
     * 
     */
    private @Nullable List<ContainerStatusPatch> containerStatuses;
    /**
     * @return Status for any ephemeral containers that have run in this pod.
     * 
     */
    private @Nullable List<ContainerStatusPatch> ephemeralContainerStatuses;
    /**
     * @return hostIP holds the IP address of the host to which the pod is assigned. Empty if the pod has not started yet. A pod can be assigned to a node that has a problem in kubelet which in turns mean that HostIP will not be updated even if there is a node is assigned to pod
     * 
     */
    private @Nullable String hostIP;
    /**
     * @return hostIPs holds the IP addresses allocated to the host. If this field is specified, the first entry must match the hostIP field. This list is empty if the pod has not started yet. A pod can be assigned to a node that has a problem in kubelet which in turns means that HostIPs will not be updated even if there is a node is assigned to this pod.
     * 
     */
    private @Nullable List<HostIPPatch> hostIPs;
    /**
     * @return The list has one entry per init container in the manifest. The most recent successful init container will have ready = true, the most recently started container will have startTime set. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status
     * 
     */
    private @Nullable List<ContainerStatusPatch> initContainerStatuses;
    /**
     * @return A human readable message indicating details about why the pod is in this condition.
     * 
     */
    private @Nullable String message;
    /**
     * @return nominatedNodeName is set only when this pod preempts other pods on the node, but it cannot be scheduled right away as preemption victims receive their graceful termination periods. This field does not guarantee that the pod will be scheduled on this node. Scheduler may decide to place the pod elsewhere if other nodes become available sooner. Scheduler may also decide to give the resources on this node to a higher priority pod that is created after preemption. As a result, this field may be different than PodSpec.nodeName when the pod is scheduled.
     * 
     */
    private @Nullable String nominatedNodeName;
    /**
     * @return The phase of a Pod is a simple, high-level summary of where the Pod is in its lifecycle. The conditions array, the reason and message fields, and the individual container status arrays contain more detail about the pod&#39;s status. There are five possible phase values:
     * 
     * Pending: The pod has been accepted by the Kubernetes system, but one or more of the container images has not been created. This includes time before being scheduled as well as time spent downloading images over the network, which could take a while. Running: The pod has been bound to a node, and all of the containers have been created. At least one container is still running, or is in the process of starting or restarting. Succeeded: All containers in the pod have terminated in success, and will not be restarted. Failed: All containers in the pod have terminated, and at least one container has terminated in failure. The container either exited with non-zero status or was terminated by the system. Unknown: For some reason the state of the pod could not be obtained, typically due to an error in communicating with the host of the pod.
     * 
     * More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-phase
     * 
     */
    private @Nullable String phase;
    /**
     * @return podIP address allocated to the pod. Routable at least within the cluster. Empty if not yet allocated.
     * 
     */
    private @Nullable String podIP;
    /**
     * @return podIPs holds the IP addresses allocated to the pod. If this field is specified, the 0th entry must match the podIP field. Pods may be allocated at most 1 value for each of IPv4 and IPv6. This list is empty if no IPs have been allocated yet.
     * 
     */
    private @Nullable List<PodIPPatch> podIPs;
    /**
     * @return The Quality of Service (QOS) classification assigned to the pod based on resource requirements See PodQOSClass type for available QOS classes More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-qos/#quality-of-service-classes
     * 
     */
    private @Nullable String qosClass;
    /**
     * @return A brief CamelCase message indicating details about why the pod is in this state. e.g. &#39;Evicted&#39;
     * 
     */
    private @Nullable String reason;
    /**
     * @return Status of resources resize desired for pod&#39;s containers. It is empty if no resources resize is pending. Any changes to container resources will automatically set this to &#34;Proposed&#34;
     * 
     */
    private @Nullable String resize;
    /**
     * @return Status of resource claims.
     * 
     */
    private @Nullable List<PodResourceClaimStatusPatch> resourceClaimStatuses;
    /**
     * @return RFC 3339 date and time at which the object was acknowledged by the Kubelet. This is before the Kubelet pulled the container image(s) for the pod.
     * 
     */
    private @Nullable String startTime;

    private PodStatusPatch() {}
    /**
     * @return Current service state of pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
     * 
     */
    public List<PodConditionPatch> conditions() {
        return this.conditions == null ? List.of() : this.conditions;
    }
    /**
     * @return The list has one entry per container in the manifest. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status
     * 
     */
    public List<ContainerStatusPatch> containerStatuses() {
        return this.containerStatuses == null ? List.of() : this.containerStatuses;
    }
    /**
     * @return Status for any ephemeral containers that have run in this pod.
     * 
     */
    public List<ContainerStatusPatch> ephemeralContainerStatuses() {
        return this.ephemeralContainerStatuses == null ? List.of() : this.ephemeralContainerStatuses;
    }
    /**
     * @return hostIP holds the IP address of the host to which the pod is assigned. Empty if the pod has not started yet. A pod can be assigned to a node that has a problem in kubelet which in turns mean that HostIP will not be updated even if there is a node is assigned to pod
     * 
     */
    public Optional<String> hostIP() {
        return Optional.ofNullable(this.hostIP);
    }
    /**
     * @return hostIPs holds the IP addresses allocated to the host. If this field is specified, the first entry must match the hostIP field. This list is empty if the pod has not started yet. A pod can be assigned to a node that has a problem in kubelet which in turns means that HostIPs will not be updated even if there is a node is assigned to this pod.
     * 
     */
    public List<HostIPPatch> hostIPs() {
        return this.hostIPs == null ? List.of() : this.hostIPs;
    }
    /**
     * @return The list has one entry per init container in the manifest. The most recent successful init container will have ready = true, the most recently started container will have startTime set. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status
     * 
     */
    public List<ContainerStatusPatch> initContainerStatuses() {
        return this.initContainerStatuses == null ? List.of() : this.initContainerStatuses;
    }
    /**
     * @return A human readable message indicating details about why the pod is in this condition.
     * 
     */
    public Optional<String> message() {
        return Optional.ofNullable(this.message);
    }
    /**
     * @return nominatedNodeName is set only when this pod preempts other pods on the node, but it cannot be scheduled right away as preemption victims receive their graceful termination periods. This field does not guarantee that the pod will be scheduled on this node. Scheduler may decide to place the pod elsewhere if other nodes become available sooner. Scheduler may also decide to give the resources on this node to a higher priority pod that is created after preemption. As a result, this field may be different than PodSpec.nodeName when the pod is scheduled.
     * 
     */
    public Optional<String> nominatedNodeName() {
        return Optional.ofNullable(this.nominatedNodeName);
    }
    /**
     * @return The phase of a Pod is a simple, high-level summary of where the Pod is in its lifecycle. The conditions array, the reason and message fields, and the individual container status arrays contain more detail about the pod&#39;s status. There are five possible phase values:
     * 
     * Pending: The pod has been accepted by the Kubernetes system, but one or more of the container images has not been created. This includes time before being scheduled as well as time spent downloading images over the network, which could take a while. Running: The pod has been bound to a node, and all of the containers have been created. At least one container is still running, or is in the process of starting or restarting. Succeeded: All containers in the pod have terminated in success, and will not be restarted. Failed: All containers in the pod have terminated, and at least one container has terminated in failure. The container either exited with non-zero status or was terminated by the system. Unknown: For some reason the state of the pod could not be obtained, typically due to an error in communicating with the host of the pod.
     * 
     * More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-phase
     * 
     */
    public Optional<String> phase() {
        return Optional.ofNullable(this.phase);
    }
    /**
     * @return podIP address allocated to the pod. Routable at least within the cluster. Empty if not yet allocated.
     * 
     */
    public Optional<String> podIP() {
        return Optional.ofNullable(this.podIP);
    }
    /**
     * @return podIPs holds the IP addresses allocated to the pod. If this field is specified, the 0th entry must match the podIP field. Pods may be allocated at most 1 value for each of IPv4 and IPv6. This list is empty if no IPs have been allocated yet.
     * 
     */
    public List<PodIPPatch> podIPs() {
        return this.podIPs == null ? List.of() : this.podIPs;
    }
    /**
     * @return The Quality of Service (QOS) classification assigned to the pod based on resource requirements See PodQOSClass type for available QOS classes More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-qos/#quality-of-service-classes
     * 
     */
    public Optional<String> qosClass() {
        return Optional.ofNullable(this.qosClass);
    }
    /**
     * @return A brief CamelCase message indicating details about why the pod is in this state. e.g. &#39;Evicted&#39;
     * 
     */
    public Optional<String> reason() {
        return Optional.ofNullable(this.reason);
    }
    /**
     * @return Status of resources resize desired for pod&#39;s containers. It is empty if no resources resize is pending. Any changes to container resources will automatically set this to &#34;Proposed&#34;
     * 
     */
    public Optional<String> resize() {
        return Optional.ofNullable(this.resize);
    }
    /**
     * @return Status of resource claims.
     * 
     */
    public List<PodResourceClaimStatusPatch> resourceClaimStatuses() {
        return this.resourceClaimStatuses == null ? List.of() : this.resourceClaimStatuses;
    }
    /**
     * @return RFC 3339 date and time at which the object was acknowledged by the Kubelet. This is before the Kubelet pulled the container image(s) for the pod.
     * 
     */
    public Optional<String> startTime() {
        return Optional.ofNullable(this.startTime);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PodStatusPatch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<PodConditionPatch> conditions;
        private @Nullable List<ContainerStatusPatch> containerStatuses;
        private @Nullable List<ContainerStatusPatch> ephemeralContainerStatuses;
        private @Nullable String hostIP;
        private @Nullable List<HostIPPatch> hostIPs;
        private @Nullable List<ContainerStatusPatch> initContainerStatuses;
        private @Nullable String message;
        private @Nullable String nominatedNodeName;
        private @Nullable String phase;
        private @Nullable String podIP;
        private @Nullable List<PodIPPatch> podIPs;
        private @Nullable String qosClass;
        private @Nullable String reason;
        private @Nullable String resize;
        private @Nullable List<PodResourceClaimStatusPatch> resourceClaimStatuses;
        private @Nullable String startTime;
        public Builder() {}
        public Builder(PodStatusPatch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.conditions = defaults.conditions;
    	      this.containerStatuses = defaults.containerStatuses;
    	      this.ephemeralContainerStatuses = defaults.ephemeralContainerStatuses;
    	      this.hostIP = defaults.hostIP;
    	      this.hostIPs = defaults.hostIPs;
    	      this.initContainerStatuses = defaults.initContainerStatuses;
    	      this.message = defaults.message;
    	      this.nominatedNodeName = defaults.nominatedNodeName;
    	      this.phase = defaults.phase;
    	      this.podIP = defaults.podIP;
    	      this.podIPs = defaults.podIPs;
    	      this.qosClass = defaults.qosClass;
    	      this.reason = defaults.reason;
    	      this.resize = defaults.resize;
    	      this.resourceClaimStatuses = defaults.resourceClaimStatuses;
    	      this.startTime = defaults.startTime;
        }

        @CustomType.Setter
        public Builder conditions(@Nullable List<PodConditionPatch> conditions) {
            this.conditions = conditions;
            return this;
        }
        public Builder conditions(PodConditionPatch... conditions) {
            return conditions(List.of(conditions));
        }
        @CustomType.Setter
        public Builder containerStatuses(@Nullable List<ContainerStatusPatch> containerStatuses) {
            this.containerStatuses = containerStatuses;
            return this;
        }
        public Builder containerStatuses(ContainerStatusPatch... containerStatuses) {
            return containerStatuses(List.of(containerStatuses));
        }
        @CustomType.Setter
        public Builder ephemeralContainerStatuses(@Nullable List<ContainerStatusPatch> ephemeralContainerStatuses) {
            this.ephemeralContainerStatuses = ephemeralContainerStatuses;
            return this;
        }
        public Builder ephemeralContainerStatuses(ContainerStatusPatch... ephemeralContainerStatuses) {
            return ephemeralContainerStatuses(List.of(ephemeralContainerStatuses));
        }
        @CustomType.Setter
        public Builder hostIP(@Nullable String hostIP) {
            this.hostIP = hostIP;
            return this;
        }
        @CustomType.Setter
        public Builder hostIPs(@Nullable List<HostIPPatch> hostIPs) {
            this.hostIPs = hostIPs;
            return this;
        }
        public Builder hostIPs(HostIPPatch... hostIPs) {
            return hostIPs(List.of(hostIPs));
        }
        @CustomType.Setter
        public Builder initContainerStatuses(@Nullable List<ContainerStatusPatch> initContainerStatuses) {
            this.initContainerStatuses = initContainerStatuses;
            return this;
        }
        public Builder initContainerStatuses(ContainerStatusPatch... initContainerStatuses) {
            return initContainerStatuses(List.of(initContainerStatuses));
        }
        @CustomType.Setter
        public Builder message(@Nullable String message) {
            this.message = message;
            return this;
        }
        @CustomType.Setter
        public Builder nominatedNodeName(@Nullable String nominatedNodeName) {
            this.nominatedNodeName = nominatedNodeName;
            return this;
        }
        @CustomType.Setter
        public Builder phase(@Nullable String phase) {
            this.phase = phase;
            return this;
        }
        @CustomType.Setter
        public Builder podIP(@Nullable String podIP) {
            this.podIP = podIP;
            return this;
        }
        @CustomType.Setter
        public Builder podIPs(@Nullable List<PodIPPatch> podIPs) {
            this.podIPs = podIPs;
            return this;
        }
        public Builder podIPs(PodIPPatch... podIPs) {
            return podIPs(List.of(podIPs));
        }
        @CustomType.Setter
        public Builder qosClass(@Nullable String qosClass) {
            this.qosClass = qosClass;
            return this;
        }
        @CustomType.Setter
        public Builder reason(@Nullable String reason) {
            this.reason = reason;
            return this;
        }
        @CustomType.Setter
        public Builder resize(@Nullable String resize) {
            this.resize = resize;
            return this;
        }
        @CustomType.Setter
        public Builder resourceClaimStatuses(@Nullable List<PodResourceClaimStatusPatch> resourceClaimStatuses) {
            this.resourceClaimStatuses = resourceClaimStatuses;
            return this;
        }
        public Builder resourceClaimStatuses(PodResourceClaimStatusPatch... resourceClaimStatuses) {
            return resourceClaimStatuses(List.of(resourceClaimStatuses));
        }
        @CustomType.Setter
        public Builder startTime(@Nullable String startTime) {
            this.startTime = startTime;
            return this;
        }
        public PodStatusPatch build() {
            final var o = new PodStatusPatch();
            o.conditions = conditions;
            o.containerStatuses = containerStatuses;
            o.ephemeralContainerStatuses = ephemeralContainerStatuses;
            o.hostIP = hostIP;
            o.hostIPs = hostIPs;
            o.initContainerStatuses = initContainerStatuses;
            o.message = message;
            o.nominatedNodeName = nominatedNodeName;
            o.phase = phase;
            o.podIP = podIP;
            o.podIPs = podIPs;
            o.qosClass = qosClass;
            o.reason = reason;
            o.resize = resize;
            o.resourceClaimStatuses = resourceClaimStatuses;
            o.startTime = startTime;
            return o;
        }
    }
}
