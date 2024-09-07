// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2

package types

type GetNavigationRequest struct {
}

type GetNavigationResponse struct {
	Code    int64             `json:"code" dc:"code 码"`
	Message string            `json:"message" dc:"消息"`
	Data    []NavigationModel `json:"data" dc:"数据"`
}

type NavigationModel struct {
	Id         int64             `json:"id" dc:"ID"`
	Title      string            `json:"title" dc:"标题"`
	Sort       int64             `json:"sort" dc:"排序"`
	SubNavList []NavigationModel `json:"subNavList" dc:"子菜单"`
}
