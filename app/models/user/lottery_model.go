// Package user 存放用户 Model 相关逻辑
package user

import (
	"gohub/app/models"
)

// User 用户模型
type Lottery struct {
	models.BaseModel

	Name   string `json:"name,omitempty"`
	Status string `json:"-"`

	models.CommonTimestampsField
}
