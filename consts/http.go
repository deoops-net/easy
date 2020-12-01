package consts

type HTTPOK struct {
	Msg  string `json:"msg" default:"ok"`
	Code int    `json:"code" default:"0"`
}

type HTTPErr struct {
	Msg  string `json:"msg"`
	Code int    `json:"code" default:"-1" example:"-1"`
}

type Result struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type PaginationResult struct {
	HTTPOK
	Data PaginationData `json:"data"`
}

type PaginationData struct {
	Tatal  int64       `json:"total"`
	Page   int64       `json:"page"`
	Size   int64       `json:"size"`
	Result interface{} `json:"result"`
}

func DefaultOK() HTTPOK {
	return HTTPOK{
		Msg:  "ok",
		Code: 0,
	}
}
