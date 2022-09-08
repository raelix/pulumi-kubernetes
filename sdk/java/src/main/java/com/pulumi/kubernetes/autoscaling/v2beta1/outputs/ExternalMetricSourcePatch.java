// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.autoscaling.v2beta1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.kubernetes.meta.v1.outputs.LabelSelectorPatch;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ExternalMetricSourcePatch {
    /**
     * @return metricName is the name of the metric in question.
     * 
     */
    private @Nullable String metricName;
    /**
     * @return metricSelector is used to identify a specific time series within a given metric.
     * 
     */
    private @Nullable LabelSelectorPatch metricSelector;
    /**
     * @return targetAverageValue is the target per-pod value of global metric (as a quantity). Mutually exclusive with TargetValue.
     * 
     */
    private @Nullable String targetAverageValue;
    /**
     * @return targetValue is the target value of the metric (as a quantity). Mutually exclusive with TargetAverageValue.
     * 
     */
    private @Nullable String targetValue;

    private ExternalMetricSourcePatch() {}
    /**
     * @return metricName is the name of the metric in question.
     * 
     */
    public Optional<String> metricName() {
        return Optional.ofNullable(this.metricName);
    }
    /**
     * @return metricSelector is used to identify a specific time series within a given metric.
     * 
     */
    public Optional<LabelSelectorPatch> metricSelector() {
        return Optional.ofNullable(this.metricSelector);
    }
    /**
     * @return targetAverageValue is the target per-pod value of global metric (as a quantity). Mutually exclusive with TargetValue.
     * 
     */
    public Optional<String> targetAverageValue() {
        return Optional.ofNullable(this.targetAverageValue);
    }
    /**
     * @return targetValue is the target value of the metric (as a quantity). Mutually exclusive with TargetAverageValue.
     * 
     */
    public Optional<String> targetValue() {
        return Optional.ofNullable(this.targetValue);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ExternalMetricSourcePatch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String metricName;
        private @Nullable LabelSelectorPatch metricSelector;
        private @Nullable String targetAverageValue;
        private @Nullable String targetValue;
        public Builder() {}
        public Builder(ExternalMetricSourcePatch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.metricName = defaults.metricName;
    	      this.metricSelector = defaults.metricSelector;
    	      this.targetAverageValue = defaults.targetAverageValue;
    	      this.targetValue = defaults.targetValue;
        }

        @CustomType.Setter
        public Builder metricName(@Nullable String metricName) {
            this.metricName = metricName;
            return this;
        }
        @CustomType.Setter
        public Builder metricSelector(@Nullable LabelSelectorPatch metricSelector) {
            this.metricSelector = metricSelector;
            return this;
        }
        @CustomType.Setter
        public Builder targetAverageValue(@Nullable String targetAverageValue) {
            this.targetAverageValue = targetAverageValue;
            return this;
        }
        @CustomType.Setter
        public Builder targetValue(@Nullable String targetValue) {
            this.targetValue = targetValue;
            return this;
        }
        public ExternalMetricSourcePatch build() {
            final var o = new ExternalMetricSourcePatch();
            o.metricName = metricName;
            o.metricSelector = metricSelector;
            o.targetAverageValue = targetAverageValue;
            o.targetValue = targetValue;
            return o;
        }
    }
}