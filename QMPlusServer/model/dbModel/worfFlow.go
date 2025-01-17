package dbModel

import (
	"github.com/jinzhu/gorm"
	"main/init/qmsql"
)

//工作流属性表
type Workflow struct {
	gorm.Model
	WorkflowNickName    string             `json:"workflowNickName"`    // 工作流名称
	WorkflowName        string             `json:"workflowName"`        // 工作流英文id
	WorkflowDescription string             `json:"workflowDescription"` // 工作流描述
	WorkflowStep        []WorkflowStepInfo `json:"workflowStep"`        // 工作流步骤
}

// 工作流状态表
type WorkflowStepInfo struct {
	gorm.Model
	WorkflowID      uint    `json:"workflowID"`      // 所属工作流ID
	IsStrat         bool    `json:"isStrat"`         // 是否是开始流节点
	StepName        string  `json:"stepName"`        // 工作流名称
	StepNo          float64 `json:"stepNo"`          // 步骤id （第几步）
	StepAuthorityID string  `json:"stepAuthorityID"` // 操作者级别id
	IsEnd           bool    `json:"isEnd"`           // 是否是完结流节点
}

//创建工作流
func (wk *Workflow) Create() error {
	err := qmsql.DEFAULTDB.Create(&wk).Error
	return err
}
