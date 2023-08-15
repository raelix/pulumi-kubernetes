// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.admissionregistration.v1alpha1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.kubernetes.admissionregistration.v1alpha1.outputs.AuditAnnotationPatch;
import com.pulumi.kubernetes.admissionregistration.v1alpha1.outputs.MatchConditionPatch;
import com.pulumi.kubernetes.admissionregistration.v1alpha1.outputs.MatchResourcesPatch;
import com.pulumi.kubernetes.admissionregistration.v1alpha1.outputs.ParamKindPatch;
import com.pulumi.kubernetes.admissionregistration.v1alpha1.outputs.ValidationPatch;
import com.pulumi.kubernetes.admissionregistration.v1alpha1.outputs.VariablePatch;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ValidatingAdmissionPolicySpecPatch {
    /**
     * @return auditAnnotations contains CEL expressions which are used to produce audit annotations for the audit event of the API request. validations and auditAnnotations may not both be empty; a least one of validations or auditAnnotations is required.
     * 
     */
    private @Nullable List<AuditAnnotationPatch> auditAnnotations;
    /**
     * @return failurePolicy defines how to handle failures for the admission policy. Failures can occur from CEL expression parse errors, type check errors, runtime errors and invalid or mis-configured policy definitions or bindings.
     * 
     * A policy is invalid if spec.paramKind refers to a non-existent Kind. A binding is invalid if spec.paramRef.name refers to a non-existent resource.
     * 
     * failurePolicy does not define how validations that evaluate to false are handled.
     * 
     * When failurePolicy is set to Fail, ValidatingAdmissionPolicyBinding validationActions define how failures are enforced.
     * 
     * Allowed values are Ignore or Fail. Defaults to Fail.
     * 
     */
    private @Nullable String failurePolicy;
    /**
     * @return MatchConditions is a list of conditions that must be met for a request to be validated. Match conditions filter requests that have already been matched by the rules, namespaceSelector, and objectSelector. An empty list of matchConditions matches all requests. There are a maximum of 64 match conditions allowed.
     * 
     * If a parameter object is provided, it can be accessed via the `params` handle in the same manner as validation expressions.
     * 
     * The exact matching logic is (in order):
     *   1. If ANY matchCondition evaluates to FALSE, the policy is skipped.
     *   2. If ALL matchConditions evaluate to TRUE, the policy is evaluated.
     *   3. If any matchCondition evaluates to an error (but none are FALSE):
     *      - If failurePolicy=Fail, reject the request
     *      - If failurePolicy=Ignore, the policy is skipped
     * 
     */
    private @Nullable List<MatchConditionPatch> matchConditions;
    /**
     * @return MatchConstraints specifies what resources this policy is designed to validate. The AdmissionPolicy cares about a request if it matches _all_ Constraints. However, in order to prevent clusters from being put into an unstable state that cannot be recovered from via the API ValidatingAdmissionPolicy cannot match ValidatingAdmissionPolicy and ValidatingAdmissionPolicyBinding. Required.
     * 
     */
    private @Nullable MatchResourcesPatch matchConstraints;
    /**
     * @return ParamKind specifies the kind of resources used to parameterize this policy. If absent, there are no parameters for this policy and the param CEL variable will not be provided to validation expressions. If ParamKind refers to a non-existent kind, this policy definition is mis-configured and the FailurePolicy is applied. If paramKind is specified but paramRef is unset in ValidatingAdmissionPolicyBinding, the params variable will be null.
     * 
     */
    private @Nullable ParamKindPatch paramKind;
    /**
     * @return Validations contain CEL expressions which is used to apply the validation. Validations and AuditAnnotations may not both be empty; a minimum of one Validations or AuditAnnotations is required.
     * 
     */
    private @Nullable List<ValidationPatch> validations;
    /**
     * @return Variables contain definitions of variables that can be used in composition of other expressions. Each variable is defined as a named CEL expression. The variables defined here will be available under `variables` in other expressions of the policy except MatchConditions because MatchConditions are evaluated before the rest of the policy.
     * 
     * The expression of a variable can refer to other variables defined earlier in the list but not those after. Thus, Variables must be sorted by the order of first appearance and acyclic.
     * 
     */
    private @Nullable List<VariablePatch> variables;

    private ValidatingAdmissionPolicySpecPatch() {}
    /**
     * @return auditAnnotations contains CEL expressions which are used to produce audit annotations for the audit event of the API request. validations and auditAnnotations may not both be empty; a least one of validations or auditAnnotations is required.
     * 
     */
    public List<AuditAnnotationPatch> auditAnnotations() {
        return this.auditAnnotations == null ? List.of() : this.auditAnnotations;
    }
    /**
     * @return failurePolicy defines how to handle failures for the admission policy. Failures can occur from CEL expression parse errors, type check errors, runtime errors and invalid or mis-configured policy definitions or bindings.
     * 
     * A policy is invalid if spec.paramKind refers to a non-existent Kind. A binding is invalid if spec.paramRef.name refers to a non-existent resource.
     * 
     * failurePolicy does not define how validations that evaluate to false are handled.
     * 
     * When failurePolicy is set to Fail, ValidatingAdmissionPolicyBinding validationActions define how failures are enforced.
     * 
     * Allowed values are Ignore or Fail. Defaults to Fail.
     * 
     */
    public Optional<String> failurePolicy() {
        return Optional.ofNullable(this.failurePolicy);
    }
    /**
     * @return MatchConditions is a list of conditions that must be met for a request to be validated. Match conditions filter requests that have already been matched by the rules, namespaceSelector, and objectSelector. An empty list of matchConditions matches all requests. There are a maximum of 64 match conditions allowed.
     * 
     * If a parameter object is provided, it can be accessed via the `params` handle in the same manner as validation expressions.
     * 
     * The exact matching logic is (in order):
     *   1. If ANY matchCondition evaluates to FALSE, the policy is skipped.
     *   2. If ALL matchConditions evaluate to TRUE, the policy is evaluated.
     *   3. If any matchCondition evaluates to an error (but none are FALSE):
     *      - If failurePolicy=Fail, reject the request
     *      - If failurePolicy=Ignore, the policy is skipped
     * 
     */
    public List<MatchConditionPatch> matchConditions() {
        return this.matchConditions == null ? List.of() : this.matchConditions;
    }
    /**
     * @return MatchConstraints specifies what resources this policy is designed to validate. The AdmissionPolicy cares about a request if it matches _all_ Constraints. However, in order to prevent clusters from being put into an unstable state that cannot be recovered from via the API ValidatingAdmissionPolicy cannot match ValidatingAdmissionPolicy and ValidatingAdmissionPolicyBinding. Required.
     * 
     */
    public Optional<MatchResourcesPatch> matchConstraints() {
        return Optional.ofNullable(this.matchConstraints);
    }
    /**
     * @return ParamKind specifies the kind of resources used to parameterize this policy. If absent, there are no parameters for this policy and the param CEL variable will not be provided to validation expressions. If ParamKind refers to a non-existent kind, this policy definition is mis-configured and the FailurePolicy is applied. If paramKind is specified but paramRef is unset in ValidatingAdmissionPolicyBinding, the params variable will be null.
     * 
     */
    public Optional<ParamKindPatch> paramKind() {
        return Optional.ofNullable(this.paramKind);
    }
    /**
     * @return Validations contain CEL expressions which is used to apply the validation. Validations and AuditAnnotations may not both be empty; a minimum of one Validations or AuditAnnotations is required.
     * 
     */
    public List<ValidationPatch> validations() {
        return this.validations == null ? List.of() : this.validations;
    }
    /**
     * @return Variables contain definitions of variables that can be used in composition of other expressions. Each variable is defined as a named CEL expression. The variables defined here will be available under `variables` in other expressions of the policy except MatchConditions because MatchConditions are evaluated before the rest of the policy.
     * 
     * The expression of a variable can refer to other variables defined earlier in the list but not those after. Thus, Variables must be sorted by the order of first appearance and acyclic.
     * 
     */
    public List<VariablePatch> variables() {
        return this.variables == null ? List.of() : this.variables;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ValidatingAdmissionPolicySpecPatch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<AuditAnnotationPatch> auditAnnotations;
        private @Nullable String failurePolicy;
        private @Nullable List<MatchConditionPatch> matchConditions;
        private @Nullable MatchResourcesPatch matchConstraints;
        private @Nullable ParamKindPatch paramKind;
        private @Nullable List<ValidationPatch> validations;
        private @Nullable List<VariablePatch> variables;
        public Builder() {}
        public Builder(ValidatingAdmissionPolicySpecPatch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.auditAnnotations = defaults.auditAnnotations;
    	      this.failurePolicy = defaults.failurePolicy;
    	      this.matchConditions = defaults.matchConditions;
    	      this.matchConstraints = defaults.matchConstraints;
    	      this.paramKind = defaults.paramKind;
    	      this.validations = defaults.validations;
    	      this.variables = defaults.variables;
        }

        @CustomType.Setter
        public Builder auditAnnotations(@Nullable List<AuditAnnotationPatch> auditAnnotations) {
            this.auditAnnotations = auditAnnotations;
            return this;
        }
        public Builder auditAnnotations(AuditAnnotationPatch... auditAnnotations) {
            return auditAnnotations(List.of(auditAnnotations));
        }
        @CustomType.Setter
        public Builder failurePolicy(@Nullable String failurePolicy) {
            this.failurePolicy = failurePolicy;
            return this;
        }
        @CustomType.Setter
        public Builder matchConditions(@Nullable List<MatchConditionPatch> matchConditions) {
            this.matchConditions = matchConditions;
            return this;
        }
        public Builder matchConditions(MatchConditionPatch... matchConditions) {
            return matchConditions(List.of(matchConditions));
        }
        @CustomType.Setter
        public Builder matchConstraints(@Nullable MatchResourcesPatch matchConstraints) {
            this.matchConstraints = matchConstraints;
            return this;
        }
        @CustomType.Setter
        public Builder paramKind(@Nullable ParamKindPatch paramKind) {
            this.paramKind = paramKind;
            return this;
        }
        @CustomType.Setter
        public Builder validations(@Nullable List<ValidationPatch> validations) {
            this.validations = validations;
            return this;
        }
        public Builder validations(ValidationPatch... validations) {
            return validations(List.of(validations));
        }
        @CustomType.Setter
        public Builder variables(@Nullable List<VariablePatch> variables) {
            this.variables = variables;
            return this;
        }
        public Builder variables(VariablePatch... variables) {
            return variables(List.of(variables));
        }
        public ValidatingAdmissionPolicySpecPatch build() {
            final var o = new ValidatingAdmissionPolicySpecPatch();
            o.auditAnnotations = auditAnnotations;
            o.failurePolicy = failurePolicy;
            o.matchConditions = matchConditions;
            o.matchConstraints = matchConstraints;
            o.paramKind = paramKind;
            o.validations = validations;
            o.variables = variables;
            return o;
        }
    }
}
