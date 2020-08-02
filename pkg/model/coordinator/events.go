package coordinator

import (
	"github.com/gohornet/hornet/pkg/model/hornet"
	"github.com/gohornet/hornet/pkg/model/milestone"
)

// CheckpointCaller is used to signal issued checkpoints.
func CheckpointCaller(handler interface{}, params ...interface{}) {
	handler.(func(checkpointName string, checkpointIndex int, tipIndex int, tipsTotal int, txHash hornet.Hash))(params[0].(string), params[1].(int), params[2].(int), params[3].(int), params[4].(hornet.Hash))
}

// MilestoneCaller is used to signal issued milestones.
func MilestoneCaller(handler interface{}, params ...interface{}) {
	handler.(func(index milestone.Index, tailTxHash hornet.Hash))(params[0].(milestone.Index), params[1].(hornet.Hash))
}
