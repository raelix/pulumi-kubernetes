// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.flowcontrol.v1beta3.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.kubernetes.flowcontrol.v1beta3.inputs.LimitResponsePatchArgs;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * LimitedPriorityLevelConfiguration specifies how to handle requests that are subject to limits. It addresses two issues:
 *   - How are requests for this priority level limited?
 *   - What should be done with requests that exceed the limit?
 * 
 */
public final class LimitedPriorityLevelConfigurationPatchArgs extends com.pulumi.resources.ResourceArgs {

    public static final LimitedPriorityLevelConfigurationPatchArgs Empty = new LimitedPriorityLevelConfigurationPatchArgs();

    /**
     * `borrowingLimitPercent`, if present, configures a limit on how many seats this priority level can borrow from other priority levels. The limit is known as this level&#39;s BorrowingConcurrencyLimit (BorrowingCL) and is a limit on the total number of seats that this level may borrow at any one time. This field holds the ratio of that limit to the level&#39;s nominal concurrency limit. When this field is non-nil, it must hold a non-negative integer and the limit is calculated as follows.
     * 
     * BorrowingCL(i) = round( NominalCL(i) * borrowingLimitPercent(i)/100.0 )
     * 
     * The value of this field can be more than 100, implying that this priority level can borrow a number of seats that is greater than its own nominal concurrency limit (NominalCL). When this field is left `nil`, the limit is effectively infinite.
     * 
     */
    @Import(name="borrowingLimitPercent")
    private @Nullable Output<Integer> borrowingLimitPercent;

    /**
     * @return `borrowingLimitPercent`, if present, configures a limit on how many seats this priority level can borrow from other priority levels. The limit is known as this level&#39;s BorrowingConcurrencyLimit (BorrowingCL) and is a limit on the total number of seats that this level may borrow at any one time. This field holds the ratio of that limit to the level&#39;s nominal concurrency limit. When this field is non-nil, it must hold a non-negative integer and the limit is calculated as follows.
     * 
     * BorrowingCL(i) = round( NominalCL(i) * borrowingLimitPercent(i)/100.0 )
     * 
     * The value of this field can be more than 100, implying that this priority level can borrow a number of seats that is greater than its own nominal concurrency limit (NominalCL). When this field is left `nil`, the limit is effectively infinite.
     * 
     */
    public Optional<Output<Integer>> borrowingLimitPercent() {
        return Optional.ofNullable(this.borrowingLimitPercent);
    }

    /**
     * `lendablePercent` prescribes the fraction of the level&#39;s NominalCL that can be borrowed by other priority levels. The value of this field must be between 0 and 100, inclusive, and it defaults to 0. The number of seats that other levels can borrow from this level, known as this level&#39;s LendableConcurrencyLimit (LendableCL), is defined as follows.
     * 
     * LendableCL(i) = round( NominalCL(i) * lendablePercent(i)/100.0 )
     * 
     */
    @Import(name="lendablePercent")
    private @Nullable Output<Integer> lendablePercent;

    /**
     * @return `lendablePercent` prescribes the fraction of the level&#39;s NominalCL that can be borrowed by other priority levels. The value of this field must be between 0 and 100, inclusive, and it defaults to 0. The number of seats that other levels can borrow from this level, known as this level&#39;s LendableConcurrencyLimit (LendableCL), is defined as follows.
     * 
     * LendableCL(i) = round( NominalCL(i) * lendablePercent(i)/100.0 )
     * 
     */
    public Optional<Output<Integer>> lendablePercent() {
        return Optional.ofNullable(this.lendablePercent);
    }

    /**
     * `limitResponse` indicates what to do with requests that can not be executed right now
     * 
     */
    @Import(name="limitResponse")
    private @Nullable Output<LimitResponsePatchArgs> limitResponse;

    /**
     * @return `limitResponse` indicates what to do with requests that can not be executed right now
     * 
     */
    public Optional<Output<LimitResponsePatchArgs>> limitResponse() {
        return Optional.ofNullable(this.limitResponse);
    }

    /**
     * `nominalConcurrencyShares` (NCS) contributes to the computation of the NominalConcurrencyLimit (NominalCL) of this level. This is the number of execution seats available at this priority level. This is used both for requests dispatched from this priority level as well as requests dispatched from other priority levels borrowing seats from this level. The server&#39;s concurrency limit (ServerCL) is divided among the Limited priority levels in proportion to their NCS values:
     * 
     * NominalCL(i)  = ceil( ServerCL * NCS(i) / sum_ncs ) sum_ncs = sum[priority level k] NCS(k)
     * 
     * Bigger numbers mean a larger nominal concurrency limit, at the expense of every other priority level. This field has a default value of 30.
     * 
     */
    @Import(name="nominalConcurrencyShares")
    private @Nullable Output<Integer> nominalConcurrencyShares;

    /**
     * @return `nominalConcurrencyShares` (NCS) contributes to the computation of the NominalConcurrencyLimit (NominalCL) of this level. This is the number of execution seats available at this priority level. This is used both for requests dispatched from this priority level as well as requests dispatched from other priority levels borrowing seats from this level. The server&#39;s concurrency limit (ServerCL) is divided among the Limited priority levels in proportion to their NCS values:
     * 
     * NominalCL(i)  = ceil( ServerCL * NCS(i) / sum_ncs ) sum_ncs = sum[priority level k] NCS(k)
     * 
     * Bigger numbers mean a larger nominal concurrency limit, at the expense of every other priority level. This field has a default value of 30.
     * 
     */
    public Optional<Output<Integer>> nominalConcurrencyShares() {
        return Optional.ofNullable(this.nominalConcurrencyShares);
    }

    private LimitedPriorityLevelConfigurationPatchArgs() {}

    private LimitedPriorityLevelConfigurationPatchArgs(LimitedPriorityLevelConfigurationPatchArgs $) {
        this.borrowingLimitPercent = $.borrowingLimitPercent;
        this.lendablePercent = $.lendablePercent;
        this.limitResponse = $.limitResponse;
        this.nominalConcurrencyShares = $.nominalConcurrencyShares;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LimitedPriorityLevelConfigurationPatchArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LimitedPriorityLevelConfigurationPatchArgs $;

        public Builder() {
            $ = new LimitedPriorityLevelConfigurationPatchArgs();
        }

        public Builder(LimitedPriorityLevelConfigurationPatchArgs defaults) {
            $ = new LimitedPriorityLevelConfigurationPatchArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param borrowingLimitPercent `borrowingLimitPercent`, if present, configures a limit on how many seats this priority level can borrow from other priority levels. The limit is known as this level&#39;s BorrowingConcurrencyLimit (BorrowingCL) and is a limit on the total number of seats that this level may borrow at any one time. This field holds the ratio of that limit to the level&#39;s nominal concurrency limit. When this field is non-nil, it must hold a non-negative integer and the limit is calculated as follows.
         * 
         * BorrowingCL(i) = round( NominalCL(i) * borrowingLimitPercent(i)/100.0 )
         * 
         * The value of this field can be more than 100, implying that this priority level can borrow a number of seats that is greater than its own nominal concurrency limit (NominalCL). When this field is left `nil`, the limit is effectively infinite.
         * 
         * @return builder
         * 
         */
        public Builder borrowingLimitPercent(@Nullable Output<Integer> borrowingLimitPercent) {
            $.borrowingLimitPercent = borrowingLimitPercent;
            return this;
        }

        /**
         * @param borrowingLimitPercent `borrowingLimitPercent`, if present, configures a limit on how many seats this priority level can borrow from other priority levels. The limit is known as this level&#39;s BorrowingConcurrencyLimit (BorrowingCL) and is a limit on the total number of seats that this level may borrow at any one time. This field holds the ratio of that limit to the level&#39;s nominal concurrency limit. When this field is non-nil, it must hold a non-negative integer and the limit is calculated as follows.
         * 
         * BorrowingCL(i) = round( NominalCL(i) * borrowingLimitPercent(i)/100.0 )
         * 
         * The value of this field can be more than 100, implying that this priority level can borrow a number of seats that is greater than its own nominal concurrency limit (NominalCL). When this field is left `nil`, the limit is effectively infinite.
         * 
         * @return builder
         * 
         */
        public Builder borrowingLimitPercent(Integer borrowingLimitPercent) {
            return borrowingLimitPercent(Output.of(borrowingLimitPercent));
        }

        /**
         * @param lendablePercent `lendablePercent` prescribes the fraction of the level&#39;s NominalCL that can be borrowed by other priority levels. The value of this field must be between 0 and 100, inclusive, and it defaults to 0. The number of seats that other levels can borrow from this level, known as this level&#39;s LendableConcurrencyLimit (LendableCL), is defined as follows.
         * 
         * LendableCL(i) = round( NominalCL(i) * lendablePercent(i)/100.0 )
         * 
         * @return builder
         * 
         */
        public Builder lendablePercent(@Nullable Output<Integer> lendablePercent) {
            $.lendablePercent = lendablePercent;
            return this;
        }

        /**
         * @param lendablePercent `lendablePercent` prescribes the fraction of the level&#39;s NominalCL that can be borrowed by other priority levels. The value of this field must be between 0 and 100, inclusive, and it defaults to 0. The number of seats that other levels can borrow from this level, known as this level&#39;s LendableConcurrencyLimit (LendableCL), is defined as follows.
         * 
         * LendableCL(i) = round( NominalCL(i) * lendablePercent(i)/100.0 )
         * 
         * @return builder
         * 
         */
        public Builder lendablePercent(Integer lendablePercent) {
            return lendablePercent(Output.of(lendablePercent));
        }

        /**
         * @param limitResponse `limitResponse` indicates what to do with requests that can not be executed right now
         * 
         * @return builder
         * 
         */
        public Builder limitResponse(@Nullable Output<LimitResponsePatchArgs> limitResponse) {
            $.limitResponse = limitResponse;
            return this;
        }

        /**
         * @param limitResponse `limitResponse` indicates what to do with requests that can not be executed right now
         * 
         * @return builder
         * 
         */
        public Builder limitResponse(LimitResponsePatchArgs limitResponse) {
            return limitResponse(Output.of(limitResponse));
        }

        /**
         * @param nominalConcurrencyShares `nominalConcurrencyShares` (NCS) contributes to the computation of the NominalConcurrencyLimit (NominalCL) of this level. This is the number of execution seats available at this priority level. This is used both for requests dispatched from this priority level as well as requests dispatched from other priority levels borrowing seats from this level. The server&#39;s concurrency limit (ServerCL) is divided among the Limited priority levels in proportion to their NCS values:
         * 
         * NominalCL(i)  = ceil( ServerCL * NCS(i) / sum_ncs ) sum_ncs = sum[priority level k] NCS(k)
         * 
         * Bigger numbers mean a larger nominal concurrency limit, at the expense of every other priority level. This field has a default value of 30.
         * 
         * @return builder
         * 
         */
        public Builder nominalConcurrencyShares(@Nullable Output<Integer> nominalConcurrencyShares) {
            $.nominalConcurrencyShares = nominalConcurrencyShares;
            return this;
        }

        /**
         * @param nominalConcurrencyShares `nominalConcurrencyShares` (NCS) contributes to the computation of the NominalConcurrencyLimit (NominalCL) of this level. This is the number of execution seats available at this priority level. This is used both for requests dispatched from this priority level as well as requests dispatched from other priority levels borrowing seats from this level. The server&#39;s concurrency limit (ServerCL) is divided among the Limited priority levels in proportion to their NCS values:
         * 
         * NominalCL(i)  = ceil( ServerCL * NCS(i) / sum_ncs ) sum_ncs = sum[priority level k] NCS(k)
         * 
         * Bigger numbers mean a larger nominal concurrency limit, at the expense of every other priority level. This field has a default value of 30.
         * 
         * @return builder
         * 
         */
        public Builder nominalConcurrencyShares(Integer nominalConcurrencyShares) {
            return nominalConcurrencyShares(Output.of(nominalConcurrencyShares));
        }

        public LimitedPriorityLevelConfigurationPatchArgs build() {
            return $;
        }
    }

}
