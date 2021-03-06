// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================
package parameters

import "github.com/gogf/gf/os/gtime"

// RecordNotFound 根据条件判断数据是否存在
// 有数据返回false
// 没数据 true
func RecordNotFound(where ...interface{}) bool {
	return Model.RecordNotFound(where...)
}

func (m *arModel) RecordNotFound(where ...interface{}) bool {
	r, err := m.M.FindOne(where...)
	if r == nil && err == nil {
		return true
	}
	return false
}

// Custom Model
type Parameters struct {
	Id         uint        `json:"id"`         // 自增ID
	CreateAt   *gtime.Time `json:"createAt"`   // 创建时间
	UpdateAt   *gtime.Time `json:"updateAt"`   // 更新时间
	DeleteAt   *gtime.Time `json:"deleteAt"`   // 删除时间
	BaseMenuId int         `json:"baseMenuId"` // BaseMenu的ID
	Key        string      `json:"key"`        // 地址栏携带参数的key
	Type       string      `json:"type"`       // 地址栏携带参数为params还是query
	Value      string      `json:"value"`      // 地址栏携带参数的值
}
