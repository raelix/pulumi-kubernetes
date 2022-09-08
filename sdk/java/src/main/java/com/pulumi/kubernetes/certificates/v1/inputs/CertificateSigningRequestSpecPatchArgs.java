// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.certificates.v1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * CertificateSigningRequestSpec contains the certificate request.
 * 
 */
public final class CertificateSigningRequestSpecPatchArgs extends com.pulumi.resources.ResourceArgs {

    public static final CertificateSigningRequestSpecPatchArgs Empty = new CertificateSigningRequestSpecPatchArgs();

    /**
     * expirationSeconds is the requested duration of validity of the issued certificate. The certificate signer may issue a certificate with a different validity duration so a client must check the delta between the notBefore and and notAfter fields in the issued certificate to determine the actual duration.
     * 
     * The v1.22+ in-tree implementations of the well-known Kubernetes signers will honor this field as long as the requested duration is not greater than the maximum duration they will honor per the --cluster-signing-duration CLI flag to the Kubernetes controller manager.
     * 
     * Certificate signers may not honor this field for various reasons:
     * 
     *   1. Old signer that is unaware of the field (such as the in-tree
     *      implementations prior to v1.22)
     *   2. Signer whose configured maximum is shorter than the requested duration
     *   3. Signer whose configured minimum is longer than the requested duration
     * 
     * The minimum valid value for expirationSeconds is 600, i.e. 10 minutes.
     * 
     */
    @Import(name="expirationSeconds")
    private @Nullable Output<Integer> expirationSeconds;

    /**
     * @return expirationSeconds is the requested duration of validity of the issued certificate. The certificate signer may issue a certificate with a different validity duration so a client must check the delta between the notBefore and and notAfter fields in the issued certificate to determine the actual duration.
     * 
     * The v1.22+ in-tree implementations of the well-known Kubernetes signers will honor this field as long as the requested duration is not greater than the maximum duration they will honor per the --cluster-signing-duration CLI flag to the Kubernetes controller manager.
     * 
     * Certificate signers may not honor this field for various reasons:
     * 
     *   1. Old signer that is unaware of the field (such as the in-tree
     *      implementations prior to v1.22)
     *   2. Signer whose configured maximum is shorter than the requested duration
     *   3. Signer whose configured minimum is longer than the requested duration
     * 
     * The minimum valid value for expirationSeconds is 600, i.e. 10 minutes.
     * 
     */
    public Optional<Output<Integer>> expirationSeconds() {
        return Optional.ofNullable(this.expirationSeconds);
    }

    /**
     * extra contains extra attributes of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
     * 
     */
    @Import(name="extra")
    private @Nullable Output<Map<String,List<String>>> extra;

    /**
     * @return extra contains extra attributes of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
     * 
     */
    public Optional<Output<Map<String,List<String>>>> extra() {
        return Optional.ofNullable(this.extra);
    }

    /**
     * groups contains group membership of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
     * 
     */
    @Import(name="groups")
    private @Nullable Output<List<String>> groups;

    /**
     * @return groups contains group membership of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
     * 
     */
    public Optional<Output<List<String>>> groups() {
        return Optional.ofNullable(this.groups);
    }

    /**
     * request contains an x509 certificate signing request encoded in a &#34;CERTIFICATE REQUEST&#34; PEM block. When serialized as JSON or YAML, the data is additionally base64-encoded.
     * 
     */
    @Import(name="request")
    private @Nullable Output<String> request;

    /**
     * @return request contains an x509 certificate signing request encoded in a &#34;CERTIFICATE REQUEST&#34; PEM block. When serialized as JSON or YAML, the data is additionally base64-encoded.
     * 
     */
    public Optional<Output<String>> request() {
        return Optional.ofNullable(this.request);
    }

    /**
     * signerName indicates the requested signer, and is a qualified name.
     * 
     * List/watch requests for CertificateSigningRequests can filter on this field using a &#34;spec.signerName=NAME&#34; fieldSelector.
     * 
     * Well-known Kubernetes signers are:
     *  1. &#34;kubernetes.io/kube-apiserver-client&#34;: issues client certificates that can be used to authenticate to kube-apiserver.
     *       Requests for this signer are never auto-approved by kube-controller-manager, can be issued by the &#34;csrsigning&#34; controller in kube-controller-manager.
     *  2. &#34;kubernetes.io/kube-apiserver-client-kubelet&#34;: issues client certificates that kubelets use to authenticate to kube-apiserver.
     *       Requests for this signer can be auto-approved by the &#34;csrapproving&#34; controller in kube-controller-manager, and can be issued by the &#34;csrsigning&#34; controller in kube-controller-manager.
     *  3. &#34;kubernetes.io/kubelet-serving&#34; issues serving certificates that kubelets use to serve TLS endpoints, which kube-apiserver can connect to securely.
     *       Requests for this signer are never auto-approved by kube-controller-manager, and can be issued by the &#34;csrsigning&#34; controller in kube-controller-manager.
     * 
     * More details are available at https://k8s.io/docs/reference/access-authn-authz/certificate-signing-requests/#kubernetes-signers
     * 
     * Custom signerNames can also be specified. The signer defines:
     *  1. Trust distribution: how trust (CA bundles) are distributed.
     *  2. Permitted subjects: and behavior when a disallowed subject is requested.
     *  3. Required, permitted, or forbidden x509 extensions in the request (including whether subjectAltNames are allowed, which types, restrictions on allowed values) and behavior when a disallowed extension is requested.
     *  4. Required, permitted, or forbidden key usages / extended key usages.
     *  5. Expiration/certificate lifetime: whether it is fixed by the signer, configurable by the admin.
     *  6. Whether or not requests for CA certificates are allowed.
     * 
     */
    @Import(name="signerName")
    private @Nullable Output<String> signerName;

    /**
     * @return signerName indicates the requested signer, and is a qualified name.
     * 
     * List/watch requests for CertificateSigningRequests can filter on this field using a &#34;spec.signerName=NAME&#34; fieldSelector.
     * 
     * Well-known Kubernetes signers are:
     *  1. &#34;kubernetes.io/kube-apiserver-client&#34;: issues client certificates that can be used to authenticate to kube-apiserver.
     *       Requests for this signer are never auto-approved by kube-controller-manager, can be issued by the &#34;csrsigning&#34; controller in kube-controller-manager.
     *  2. &#34;kubernetes.io/kube-apiserver-client-kubelet&#34;: issues client certificates that kubelets use to authenticate to kube-apiserver.
     *       Requests for this signer can be auto-approved by the &#34;csrapproving&#34; controller in kube-controller-manager, and can be issued by the &#34;csrsigning&#34; controller in kube-controller-manager.
     *  3. &#34;kubernetes.io/kubelet-serving&#34; issues serving certificates that kubelets use to serve TLS endpoints, which kube-apiserver can connect to securely.
     *       Requests for this signer are never auto-approved by kube-controller-manager, and can be issued by the &#34;csrsigning&#34; controller in kube-controller-manager.
     * 
     * More details are available at https://k8s.io/docs/reference/access-authn-authz/certificate-signing-requests/#kubernetes-signers
     * 
     * Custom signerNames can also be specified. The signer defines:
     *  1. Trust distribution: how trust (CA bundles) are distributed.
     *  2. Permitted subjects: and behavior when a disallowed subject is requested.
     *  3. Required, permitted, or forbidden x509 extensions in the request (including whether subjectAltNames are allowed, which types, restrictions on allowed values) and behavior when a disallowed extension is requested.
     *  4. Required, permitted, or forbidden key usages / extended key usages.
     *  5. Expiration/certificate lifetime: whether it is fixed by the signer, configurable by the admin.
     *  6. Whether or not requests for CA certificates are allowed.
     * 
     */
    public Optional<Output<String>> signerName() {
        return Optional.ofNullable(this.signerName);
    }

    /**
     * uid contains the uid of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
     * 
     */
    @Import(name="uid")
    private @Nullable Output<String> uid;

    /**
     * @return uid contains the uid of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
     * 
     */
    public Optional<Output<String>> uid() {
        return Optional.ofNullable(this.uid);
    }

    /**
     * usages specifies a set of key usages requested in the issued certificate.
     * 
     * Requests for TLS client certificates typically request: &#34;digital signature&#34;, &#34;key encipherment&#34;, &#34;client auth&#34;.
     * 
     * Requests for TLS serving certificates typically request: &#34;key encipherment&#34;, &#34;digital signature&#34;, &#34;server auth&#34;.
     * 
     * Valid values are:
     *  &#34;signing&#34;, &#34;digital signature&#34;, &#34;content commitment&#34;,
     *  &#34;key encipherment&#34;, &#34;key agreement&#34;, &#34;data encipherment&#34;,
     *  &#34;cert sign&#34;, &#34;crl sign&#34;, &#34;encipher only&#34;, &#34;decipher only&#34;, &#34;any&#34;,
     *  &#34;server auth&#34;, &#34;client auth&#34;,
     *  &#34;code signing&#34;, &#34;email protection&#34;, &#34;s/mime&#34;,
     *  &#34;ipsec end system&#34;, &#34;ipsec tunnel&#34;, &#34;ipsec user&#34;,
     *  &#34;timestamping&#34;, &#34;ocsp signing&#34;, &#34;microsoft sgc&#34;, &#34;netscape sgc&#34;
     * 
     */
    @Import(name="usages")
    private @Nullable Output<List<String>> usages;

    /**
     * @return usages specifies a set of key usages requested in the issued certificate.
     * 
     * Requests for TLS client certificates typically request: &#34;digital signature&#34;, &#34;key encipherment&#34;, &#34;client auth&#34;.
     * 
     * Requests for TLS serving certificates typically request: &#34;key encipherment&#34;, &#34;digital signature&#34;, &#34;server auth&#34;.
     * 
     * Valid values are:
     *  &#34;signing&#34;, &#34;digital signature&#34;, &#34;content commitment&#34;,
     *  &#34;key encipherment&#34;, &#34;key agreement&#34;, &#34;data encipherment&#34;,
     *  &#34;cert sign&#34;, &#34;crl sign&#34;, &#34;encipher only&#34;, &#34;decipher only&#34;, &#34;any&#34;,
     *  &#34;server auth&#34;, &#34;client auth&#34;,
     *  &#34;code signing&#34;, &#34;email protection&#34;, &#34;s/mime&#34;,
     *  &#34;ipsec end system&#34;, &#34;ipsec tunnel&#34;, &#34;ipsec user&#34;,
     *  &#34;timestamping&#34;, &#34;ocsp signing&#34;, &#34;microsoft sgc&#34;, &#34;netscape sgc&#34;
     * 
     */
    public Optional<Output<List<String>>> usages() {
        return Optional.ofNullable(this.usages);
    }

    /**
     * username contains the name of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return username contains the name of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private CertificateSigningRequestSpecPatchArgs() {}

    private CertificateSigningRequestSpecPatchArgs(CertificateSigningRequestSpecPatchArgs $) {
        this.expirationSeconds = $.expirationSeconds;
        this.extra = $.extra;
        this.groups = $.groups;
        this.request = $.request;
        this.signerName = $.signerName;
        this.uid = $.uid;
        this.usages = $.usages;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CertificateSigningRequestSpecPatchArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CertificateSigningRequestSpecPatchArgs $;

        public Builder() {
            $ = new CertificateSigningRequestSpecPatchArgs();
        }

        public Builder(CertificateSigningRequestSpecPatchArgs defaults) {
            $ = new CertificateSigningRequestSpecPatchArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param expirationSeconds expirationSeconds is the requested duration of validity of the issued certificate. The certificate signer may issue a certificate with a different validity duration so a client must check the delta between the notBefore and and notAfter fields in the issued certificate to determine the actual duration.
         * 
         * The v1.22+ in-tree implementations of the well-known Kubernetes signers will honor this field as long as the requested duration is not greater than the maximum duration they will honor per the --cluster-signing-duration CLI flag to the Kubernetes controller manager.
         * 
         * Certificate signers may not honor this field for various reasons:
         * 
         *   1. Old signer that is unaware of the field (such as the in-tree
         *      implementations prior to v1.22)
         *   2. Signer whose configured maximum is shorter than the requested duration
         *   3. Signer whose configured minimum is longer than the requested duration
         * 
         * The minimum valid value for expirationSeconds is 600, i.e. 10 minutes.
         * 
         * @return builder
         * 
         */
        public Builder expirationSeconds(@Nullable Output<Integer> expirationSeconds) {
            $.expirationSeconds = expirationSeconds;
            return this;
        }

        /**
         * @param expirationSeconds expirationSeconds is the requested duration of validity of the issued certificate. The certificate signer may issue a certificate with a different validity duration so a client must check the delta between the notBefore and and notAfter fields in the issued certificate to determine the actual duration.
         * 
         * The v1.22+ in-tree implementations of the well-known Kubernetes signers will honor this field as long as the requested duration is not greater than the maximum duration they will honor per the --cluster-signing-duration CLI flag to the Kubernetes controller manager.
         * 
         * Certificate signers may not honor this field for various reasons:
         * 
         *   1. Old signer that is unaware of the field (such as the in-tree
         *      implementations prior to v1.22)
         *   2. Signer whose configured maximum is shorter than the requested duration
         *   3. Signer whose configured minimum is longer than the requested duration
         * 
         * The minimum valid value for expirationSeconds is 600, i.e. 10 minutes.
         * 
         * @return builder
         * 
         */
        public Builder expirationSeconds(Integer expirationSeconds) {
            return expirationSeconds(Output.of(expirationSeconds));
        }

        /**
         * @param extra extra contains extra attributes of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
         * 
         * @return builder
         * 
         */
        public Builder extra(@Nullable Output<Map<String,List<String>>> extra) {
            $.extra = extra;
            return this;
        }

        /**
         * @param extra extra contains extra attributes of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
         * 
         * @return builder
         * 
         */
        public Builder extra(Map<String,List<String>> extra) {
            return extra(Output.of(extra));
        }

        /**
         * @param groups groups contains group membership of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
         * 
         * @return builder
         * 
         */
        public Builder groups(@Nullable Output<List<String>> groups) {
            $.groups = groups;
            return this;
        }

        /**
         * @param groups groups contains group membership of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
         * 
         * @return builder
         * 
         */
        public Builder groups(List<String> groups) {
            return groups(Output.of(groups));
        }

        /**
         * @param groups groups contains group membership of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
         * 
         * @return builder
         * 
         */
        public Builder groups(String... groups) {
            return groups(List.of(groups));
        }

        /**
         * @param request request contains an x509 certificate signing request encoded in a &#34;CERTIFICATE REQUEST&#34; PEM block. When serialized as JSON or YAML, the data is additionally base64-encoded.
         * 
         * @return builder
         * 
         */
        public Builder request(@Nullable Output<String> request) {
            $.request = request;
            return this;
        }

        /**
         * @param request request contains an x509 certificate signing request encoded in a &#34;CERTIFICATE REQUEST&#34; PEM block. When serialized as JSON or YAML, the data is additionally base64-encoded.
         * 
         * @return builder
         * 
         */
        public Builder request(String request) {
            return request(Output.of(request));
        }

        /**
         * @param signerName signerName indicates the requested signer, and is a qualified name.
         * 
         * List/watch requests for CertificateSigningRequests can filter on this field using a &#34;spec.signerName=NAME&#34; fieldSelector.
         * 
         * Well-known Kubernetes signers are:
         *  1. &#34;kubernetes.io/kube-apiserver-client&#34;: issues client certificates that can be used to authenticate to kube-apiserver.
         *       Requests for this signer are never auto-approved by kube-controller-manager, can be issued by the &#34;csrsigning&#34; controller in kube-controller-manager.
         *  2. &#34;kubernetes.io/kube-apiserver-client-kubelet&#34;: issues client certificates that kubelets use to authenticate to kube-apiserver.
         *       Requests for this signer can be auto-approved by the &#34;csrapproving&#34; controller in kube-controller-manager, and can be issued by the &#34;csrsigning&#34; controller in kube-controller-manager.
         *  3. &#34;kubernetes.io/kubelet-serving&#34; issues serving certificates that kubelets use to serve TLS endpoints, which kube-apiserver can connect to securely.
         *       Requests for this signer are never auto-approved by kube-controller-manager, and can be issued by the &#34;csrsigning&#34; controller in kube-controller-manager.
         * 
         * More details are available at https://k8s.io/docs/reference/access-authn-authz/certificate-signing-requests/#kubernetes-signers
         * 
         * Custom signerNames can also be specified. The signer defines:
         *  1. Trust distribution: how trust (CA bundles) are distributed.
         *  2. Permitted subjects: and behavior when a disallowed subject is requested.
         *  3. Required, permitted, or forbidden x509 extensions in the request (including whether subjectAltNames are allowed, which types, restrictions on allowed values) and behavior when a disallowed extension is requested.
         *  4. Required, permitted, or forbidden key usages / extended key usages.
         *  5. Expiration/certificate lifetime: whether it is fixed by the signer, configurable by the admin.
         *  6. Whether or not requests for CA certificates are allowed.
         * 
         * @return builder
         * 
         */
        public Builder signerName(@Nullable Output<String> signerName) {
            $.signerName = signerName;
            return this;
        }

        /**
         * @param signerName signerName indicates the requested signer, and is a qualified name.
         * 
         * List/watch requests for CertificateSigningRequests can filter on this field using a &#34;spec.signerName=NAME&#34; fieldSelector.
         * 
         * Well-known Kubernetes signers are:
         *  1. &#34;kubernetes.io/kube-apiserver-client&#34;: issues client certificates that can be used to authenticate to kube-apiserver.
         *       Requests for this signer are never auto-approved by kube-controller-manager, can be issued by the &#34;csrsigning&#34; controller in kube-controller-manager.
         *  2. &#34;kubernetes.io/kube-apiserver-client-kubelet&#34;: issues client certificates that kubelets use to authenticate to kube-apiserver.
         *       Requests for this signer can be auto-approved by the &#34;csrapproving&#34; controller in kube-controller-manager, and can be issued by the &#34;csrsigning&#34; controller in kube-controller-manager.
         *  3. &#34;kubernetes.io/kubelet-serving&#34; issues serving certificates that kubelets use to serve TLS endpoints, which kube-apiserver can connect to securely.
         *       Requests for this signer are never auto-approved by kube-controller-manager, and can be issued by the &#34;csrsigning&#34; controller in kube-controller-manager.
         * 
         * More details are available at https://k8s.io/docs/reference/access-authn-authz/certificate-signing-requests/#kubernetes-signers
         * 
         * Custom signerNames can also be specified. The signer defines:
         *  1. Trust distribution: how trust (CA bundles) are distributed.
         *  2. Permitted subjects: and behavior when a disallowed subject is requested.
         *  3. Required, permitted, or forbidden x509 extensions in the request (including whether subjectAltNames are allowed, which types, restrictions on allowed values) and behavior when a disallowed extension is requested.
         *  4. Required, permitted, or forbidden key usages / extended key usages.
         *  5. Expiration/certificate lifetime: whether it is fixed by the signer, configurable by the admin.
         *  6. Whether or not requests for CA certificates are allowed.
         * 
         * @return builder
         * 
         */
        public Builder signerName(String signerName) {
            return signerName(Output.of(signerName));
        }

        /**
         * @param uid uid contains the uid of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
         * 
         * @return builder
         * 
         */
        public Builder uid(@Nullable Output<String> uid) {
            $.uid = uid;
            return this;
        }

        /**
         * @param uid uid contains the uid of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
         * 
         * @return builder
         * 
         */
        public Builder uid(String uid) {
            return uid(Output.of(uid));
        }

        /**
         * @param usages usages specifies a set of key usages requested in the issued certificate.
         * 
         * Requests for TLS client certificates typically request: &#34;digital signature&#34;, &#34;key encipherment&#34;, &#34;client auth&#34;.
         * 
         * Requests for TLS serving certificates typically request: &#34;key encipherment&#34;, &#34;digital signature&#34;, &#34;server auth&#34;.
         * 
         * Valid values are:
         *  &#34;signing&#34;, &#34;digital signature&#34;, &#34;content commitment&#34;,
         *  &#34;key encipherment&#34;, &#34;key agreement&#34;, &#34;data encipherment&#34;,
         *  &#34;cert sign&#34;, &#34;crl sign&#34;, &#34;encipher only&#34;, &#34;decipher only&#34;, &#34;any&#34;,
         *  &#34;server auth&#34;, &#34;client auth&#34;,
         *  &#34;code signing&#34;, &#34;email protection&#34;, &#34;s/mime&#34;,
         *  &#34;ipsec end system&#34;, &#34;ipsec tunnel&#34;, &#34;ipsec user&#34;,
         *  &#34;timestamping&#34;, &#34;ocsp signing&#34;, &#34;microsoft sgc&#34;, &#34;netscape sgc&#34;
         * 
         * @return builder
         * 
         */
        public Builder usages(@Nullable Output<List<String>> usages) {
            $.usages = usages;
            return this;
        }

        /**
         * @param usages usages specifies a set of key usages requested in the issued certificate.
         * 
         * Requests for TLS client certificates typically request: &#34;digital signature&#34;, &#34;key encipherment&#34;, &#34;client auth&#34;.
         * 
         * Requests for TLS serving certificates typically request: &#34;key encipherment&#34;, &#34;digital signature&#34;, &#34;server auth&#34;.
         * 
         * Valid values are:
         *  &#34;signing&#34;, &#34;digital signature&#34;, &#34;content commitment&#34;,
         *  &#34;key encipherment&#34;, &#34;key agreement&#34;, &#34;data encipherment&#34;,
         *  &#34;cert sign&#34;, &#34;crl sign&#34;, &#34;encipher only&#34;, &#34;decipher only&#34;, &#34;any&#34;,
         *  &#34;server auth&#34;, &#34;client auth&#34;,
         *  &#34;code signing&#34;, &#34;email protection&#34;, &#34;s/mime&#34;,
         *  &#34;ipsec end system&#34;, &#34;ipsec tunnel&#34;, &#34;ipsec user&#34;,
         *  &#34;timestamping&#34;, &#34;ocsp signing&#34;, &#34;microsoft sgc&#34;, &#34;netscape sgc&#34;
         * 
         * @return builder
         * 
         */
        public Builder usages(List<String> usages) {
            return usages(Output.of(usages));
        }

        /**
         * @param usages usages specifies a set of key usages requested in the issued certificate.
         * 
         * Requests for TLS client certificates typically request: &#34;digital signature&#34;, &#34;key encipherment&#34;, &#34;client auth&#34;.
         * 
         * Requests for TLS serving certificates typically request: &#34;key encipherment&#34;, &#34;digital signature&#34;, &#34;server auth&#34;.
         * 
         * Valid values are:
         *  &#34;signing&#34;, &#34;digital signature&#34;, &#34;content commitment&#34;,
         *  &#34;key encipherment&#34;, &#34;key agreement&#34;, &#34;data encipherment&#34;,
         *  &#34;cert sign&#34;, &#34;crl sign&#34;, &#34;encipher only&#34;, &#34;decipher only&#34;, &#34;any&#34;,
         *  &#34;server auth&#34;, &#34;client auth&#34;,
         *  &#34;code signing&#34;, &#34;email protection&#34;, &#34;s/mime&#34;,
         *  &#34;ipsec end system&#34;, &#34;ipsec tunnel&#34;, &#34;ipsec user&#34;,
         *  &#34;timestamping&#34;, &#34;ocsp signing&#34;, &#34;microsoft sgc&#34;, &#34;netscape sgc&#34;
         * 
         * @return builder
         * 
         */
        public Builder usages(String... usages) {
            return usages(List.of(usages));
        }

        /**
         * @param username username contains the name of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username username contains the name of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public CertificateSigningRequestSpecPatchArgs build() {
            return $;
        }
    }

}