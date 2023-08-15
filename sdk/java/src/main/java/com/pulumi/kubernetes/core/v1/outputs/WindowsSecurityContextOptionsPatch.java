// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.core.v1.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class WindowsSecurityContextOptionsPatch {
    /**
     * @return GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.
     * 
     */
    private @Nullable String gmsaCredentialSpec;
    /**
     * @return GMSACredentialSpecName is the name of the GMSA credential spec to use.
     * 
     */
    private @Nullable String gmsaCredentialSpecName;
    /**
     * @return HostProcess determines if a container should be run as a &#39;Host Process&#39; container. All of a Pod&#39;s containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers). In addition, if HostProcess is true then HostNetwork must also be set to true.
     * 
     */
    private @Nullable Boolean hostProcess;
    /**
     * @return The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
     * 
     */
    private @Nullable String runAsUserName;

    private WindowsSecurityContextOptionsPatch() {}
    /**
     * @return GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.
     * 
     */
    public Optional<String> gmsaCredentialSpec() {
        return Optional.ofNullable(this.gmsaCredentialSpec);
    }
    /**
     * @return GMSACredentialSpecName is the name of the GMSA credential spec to use.
     * 
     */
    public Optional<String> gmsaCredentialSpecName() {
        return Optional.ofNullable(this.gmsaCredentialSpecName);
    }
    /**
     * @return HostProcess determines if a container should be run as a &#39;Host Process&#39; container. All of a Pod&#39;s containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers). In addition, if HostProcess is true then HostNetwork must also be set to true.
     * 
     */
    public Optional<Boolean> hostProcess() {
        return Optional.ofNullable(this.hostProcess);
    }
    /**
     * @return The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
     * 
     */
    public Optional<String> runAsUserName() {
        return Optional.ofNullable(this.runAsUserName);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(WindowsSecurityContextOptionsPatch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String gmsaCredentialSpec;
        private @Nullable String gmsaCredentialSpecName;
        private @Nullable Boolean hostProcess;
        private @Nullable String runAsUserName;
        public Builder() {}
        public Builder(WindowsSecurityContextOptionsPatch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.gmsaCredentialSpec = defaults.gmsaCredentialSpec;
    	      this.gmsaCredentialSpecName = defaults.gmsaCredentialSpecName;
    	      this.hostProcess = defaults.hostProcess;
    	      this.runAsUserName = defaults.runAsUserName;
        }

        @CustomType.Setter
        public Builder gmsaCredentialSpec(@Nullable String gmsaCredentialSpec) {
            this.gmsaCredentialSpec = gmsaCredentialSpec;
            return this;
        }
        @CustomType.Setter
        public Builder gmsaCredentialSpecName(@Nullable String gmsaCredentialSpecName) {
            this.gmsaCredentialSpecName = gmsaCredentialSpecName;
            return this;
        }
        @CustomType.Setter
        public Builder hostProcess(@Nullable Boolean hostProcess) {
            this.hostProcess = hostProcess;
            return this;
        }
        @CustomType.Setter
        public Builder runAsUserName(@Nullable String runAsUserName) {
            this.runAsUserName = runAsUserName;
            return this;
        }
        public WindowsSecurityContextOptionsPatch build() {
            final var o = new WindowsSecurityContextOptionsPatch();
            o.gmsaCredentialSpec = gmsaCredentialSpec;
            o.gmsaCredentialSpecName = gmsaCredentialSpecName;
            o.hostProcess = hostProcess;
            o.runAsUserName = runAsUserName;
            return o;
        }
    }
}
