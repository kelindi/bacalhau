// Code generated by "stringer -type=ExecutionStateType --trimprefix=ExecutionState --output execution_state_string.go"; DO NOT EDIT.

package model

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ExecutionStateNew-0]
	_ = x[ExecutionStateAskForBid-1]
	_ = x[ExecutionStateAskForBidAccepted-2]
	_ = x[ExecutionStateAskForBidRejected-3]
	_ = x[ExecutionStateBidAccepted-4]
	_ = x[ExecutionStateBidRejected-5]
	_ = x[ExecutionStateResultProposed-6]
	_ = x[ExecutionStateResultAccepted-7]
	_ = x[ExecutionStateResultRejected-8]
	_ = x[ExecutionStateCompleted-9]
	_ = x[ExecutionStateFailed-10]
	_ = x[ExecutionStateCanceled-11]
}

const _ExecutionStateType_name = "NewAskForBidAskForBidAcceptedAskForBidRejectedBidAcceptedBidRejectedWaitingVerificationResultAcceptedResultRejectedCompletedFailedCancelled"

var _ExecutionStateType_index = [...]uint8{0, 3, 12, 29, 46, 57, 68, 87, 101, 115, 124, 130, 139}

func (i ExecutionStateType) String() string {
	if i < 0 || i >= ExecutionStateType(len(_ExecutionStateType_index)-1) {
		return "ExecutionStateType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ExecutionStateType_name[_ExecutionStateType_index[i]:_ExecutionStateType_index[i+1]]
}
