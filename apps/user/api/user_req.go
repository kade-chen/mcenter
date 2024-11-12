package api

import "github.com/kade-chen/mcenter/apps/user"

// NewQueryUserRequest 列表查询请求
func NewQueryUserRequest() *user.QueryUserRequest {
	return &user.QueryUserRequest{
		Page: &user.PageRequest{
			// PageSize: 20,
			// PageNumber:  1,
		},
		// Page:         NewPageRequest(20, 1),
		SkipItems:    false,
		Labels:       map[string]string{},
		UserIds:      []string{},
		ExtraUserIds: []string{},
	}
}

// NewPageRequest 实例化
func NewPageRequest(ps uint, pn uint) *user.PageRequest {
	return &user.PageRequest{
		// PageSize:   uint64(ps),
		// PageNumber: uint64(pn),
	}
}
