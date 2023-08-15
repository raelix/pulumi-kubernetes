// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.core.v1.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class HTTPHeaderPatch {
    /**
     * @return The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.
     * 
     */
    private @Nullable String name;
    /**
     * @return The header field value
     * 
     */
    private @Nullable String value;

    private HTTPHeaderPatch() {}
    /**
     * @return The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return The header field value
     * 
     */
    public Optional<String> value() {
        return Optional.ofNullable(this.value);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(HTTPHeaderPatch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String name;
        private @Nullable String value;
        public Builder() {}
        public Builder(HTTPHeaderPatch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder value(@Nullable String value) {
            this.value = value;
            return this;
        }
        public HTTPHeaderPatch build() {
            final var o = new HTTPHeaderPatch();
            o.name = name;
            o.value = value;
            return o;
        }
    }
}
