package utils

// 用于分页

type Resp struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Rows  interface{} `json:"rows"`
	Total interface{} `json:"total"`
}
