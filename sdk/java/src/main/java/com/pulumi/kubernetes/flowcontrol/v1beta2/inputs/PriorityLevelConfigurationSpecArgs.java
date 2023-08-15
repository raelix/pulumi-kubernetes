// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.flowcontrol.v1beta2.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.kubernetes.flowcontrol.v1beta2.inputs.ExemptPriorityLevelConfigurationArgs;
import com.pulumi.kubernetes.flowcontrol.v1beta2.inputs.LimitedPriorityLevelConfigurationArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * PriorityLevelConfigurationSpec specifies the configuration of a priority level.
 * 
 */
public final class PriorityLevelConfigurationSpecArgs extends com.pulumi.resources.ResourceArgs {

    public static final PriorityLevelConfigurationSpecArgs Empty = new PriorityLevelConfigurationSpecArgs();

    /**
     * `exempt` specifies how requests are handled for an exempt priority level. This field MUST be empty if `type` is `&#34;Limited&#34;`. This field MAY be non-empty if `type` is `&#34;Exempt&#34;`. If empty and `type` is `&#34;Exempt&#34;` then the default values for `ExemptPriorityLevelConfiguration` apply.
     * 
     */
    @Import(name="exempt")
    private @Nullable Output<ExemptPriorityLevelConfigurationArgs> exempt;

    /**
     * @return `exempt` specifies how requests are handled for an exempt priority level. This field MUST be empty if `type` is `&#34;Limited&#34;`. This field MAY be non-empty if `type` is `&#34;Exempt&#34;`. If empty and `type` is `&#34;Exempt&#34;` then the default values for `ExemptPriorityLevelConfiguration` apply.
     * 
     */
    public Optional<Output<ExemptPriorityLevelConfigurationArgs>> exempt() {
        return Optional.ofNullable(this.exempt);
    }

    /**
     * `limited` specifies how requests are handled for a Limited priority level. This field must be non-empty if and only if `type` is `&#34;Limited&#34;`.
     * 
     */
    @Import(name="limited")
    private @Nullable Output<LimitedPriorityLevelConfigurationArgs> limited;

    /**
     * @return `limited` specifies how requests are handled for a Limited priority level. This field must be non-empty if and only if `type` is `&#34;Limited&#34;`.
     * 
     */
    public Optional<Output<LimitedPriorityLevelConfigurationArgs>> limited() {
        return Optional.ofNullable(this.limited);
    }

    /**
     * `type` indicates whether this priority level is subject to limitation on request execution.  A value of `&#34;Exempt&#34;` means that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the capacity made available to other priority levels.  A value of `&#34;Limited&#34;` means that (a) requests of this priority level _are_ subject to limits and (b) some of the server&#39;s limited capacity is made available exclusively to this priority level. Required.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return `type` indicates whether this priority level is subject to limitation on request execution.  A value of `&#34;Exempt&#34;` means that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the capacity made available to other priority levels.  A value of `&#34;Limited&#34;` means that (a) requests of this priority level _are_ subject to limits and (b) some of the server&#39;s limited capacity is made available exclusively to this priority level. Required.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private PriorityLevelConfigurationSpecArgs() {}

    private PriorityLevelConfigurationSpecArgs(PriorityLevelConfigurationSpecArgs $) {
        this.exempt = $.exempt;
        this.limited = $.limited;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PriorityLevelConfigurationSpecArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PriorityLevelConfigurationSpecArgs $;

        public Builder() {
            $ = new PriorityLevelConfigurationSpecArgs();
        }

        public Builder(PriorityLevelConfigurationSpecArgs defaults) {
            $ = new PriorityLevelConfigurationSpecArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param exempt `exempt` specifies how requests are handled for an exempt priority level. This field MUST be empty if `type` is `&#34;Limited&#34;`. This field MAY be non-empty if `type` is `&#34;Exempt&#34;`. If empty and `type` is `&#34;Exempt&#34;` then the default values for `ExemptPriorityLevelConfiguration` apply.
         * 
         * @return builder
         * 
         */
        public Builder exempt(@Nullable Output<ExemptPriorityLevelConfigurationArgs> exempt) {
            $.exempt = exempt;
            return this;
        }

        /**
         * @param exempt `exempt` specifies how requests are handled for an exempt priority level. This field MUST be empty if `type` is `&#34;Limited&#34;`. This field MAY be non-empty if `type` is `&#34;Exempt&#34;`. If empty and `type` is `&#34;Exempt&#34;` then the default values for `ExemptPriorityLevelConfiguration` apply.
         * 
         * @return builder
         * 
         */
        public Builder exempt(ExemptPriorityLevelConfigurationArgs exempt) {
            return exempt(Output.of(exempt));
        }

        /**
         * @param limited `limited` specifies how requests are handled for a Limited priority level. This field must be non-empty if and only if `type` is `&#34;Limited&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder limited(@Nullable Output<LimitedPriorityLevelConfigurationArgs> limited) {
            $.limited = limited;
            return this;
        }

        /**
         * @param limited `limited` specifies how requests are handled for a Limited priority level. This field must be non-empty if and only if `type` is `&#34;Limited&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder limited(LimitedPriorityLevelConfigurationArgs limited) {
            return limited(Output.of(limited));
        }

        /**
         * @param type `type` indicates whether this priority level is subject to limitation on request execution.  A value of `&#34;Exempt&#34;` means that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the capacity made available to other priority levels.  A value of `&#34;Limited&#34;` means that (a) requests of this priority level _are_ subject to limits and (b) some of the server&#39;s limited capacity is made available exclusively to this priority level. Required.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type `type` indicates whether this priority level is subject to limitation on request execution.  A value of `&#34;Exempt&#34;` means that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the capacity made available to other priority levels.  A value of `&#34;Limited&#34;` means that (a) requests of this priority level _are_ subject to limits and (b) some of the server&#39;s limited capacity is made available exclusively to this priority level. Required.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public PriorityLevelConfigurationSpecArgs build() {
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
