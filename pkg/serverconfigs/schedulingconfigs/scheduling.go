package schedulingconfigs

import (
	"github.com/iwind/TeaGo/maps"
	"github.com/oy1978/EdgeCommon/pkg/serverconfigs/shared"
)

// SchedulingInterface 调度算法接口
type SchedulingInterface interface {
	// HasCandidates 是否有候选对象
	HasCandidates() bool

	// Add 添加候选对象
	Add(candidate ...CandidateInterface)

	// Start 启动
	Start()

	// Next 查找下一个候选对象
	Next(call *shared.RequestCall) CandidateInterface

	// Summary 获取简要信息
	Summary() maps.Map
}

// Scheduling 调度算法基础类
type Scheduling struct {
	Candidates []CandidateInterface
}

// HasCandidates 判断是否有候选对象
func (this *Scheduling) HasCandidates() bool {
	return len(this.Candidates) > 0
}

// Add 添加候选对象
func (this *Scheduling) Add(candidate ...CandidateInterface) {
	this.Candidates = append(this.Candidates, candidate...)
}
