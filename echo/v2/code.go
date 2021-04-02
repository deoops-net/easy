package v2

type Resp struct {
	Code string `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	// 成功码，我们这里主要用3个
	StatusCodeOK                   = "20000" // RFC 7231, 6.3.1
	// 资源已创建过
	// The HTTP 201 Created success status response code indicates that the request has succeeded and has led to the creation of a resource. The new resource is effectively created before this response is sent back and the new resource is returned in the body of the message, its location being either the URL of the request, or the content of the Location header.
	// The common use case of this status code is as the result of a POST request.
	StatusCodeCreated              = "20001" // RFC 7231, 6.3.2
	// 异步请求
	// StatusCodeAccepted             = "20002" // RFC 7231, 6.3.3
	// StatusCodeNonAuthoritativeInfo = "20003" // RFC 7231, 6.3.4
	// 没有相应的数据，在spec中主要用来标识 cache 和更新时无需跳转其他页面
	StatusCodeNoContent            = "20004" // RFC 7231, 6.3.5
	// StatusCodeResetContent         = "205" // RFC 7231, 6.3.6
	// StatusCodePartialContent       = "206" // RFC 7233, 4.1
	// StatusCodeMultiStatusCode          = "207" // RFC 4918, 11.1
	// StatusCodeAlreadyReported      = "208" // RFC 5842, 7.1
	// StatusCodeIMUsed               = "226" // RFC 3229, 10.4.1

	// 重定向相关，暂时都用不到
	// StatusCodeMultipleChoices   = "300" // RFC 7231, 6.4.1
	// StatusCodeMovedPermanently  = "301" // RFC 7231, 6.4.2
	// StatusCodeFound             = "302" // RFC 7231, 6.4.3
	// StatusCodeSeeOther          = "303" // RFC 7231, 6.4.4
	// StatusCodeNotModified       = "304" // RFC 7232, 4.1
	// StatusCodeUseProxy          = "305" // RFC 7231, 6.4.5
	// _                       = "306" // RFC 7231, 6.4.6 (Unused)
	// StatusCodeTemporaryRedirect = "307" // RFC 7231, 6.4.7
	// StatusCodePermanentRedirect = "308" // RFC 7538, 3

	// 权限相关
	// 请求解析出错
	StatusCodeBadRequest                   = "40000" // RFC 7231, 6.5.1
	// 认证失败
	StatusCodeUnauthorized                 = "40001" // RFC 7235, 3.1
	// session失效
	//StatusCodePaymentRequired              = "40002" // RFC 7231, 6.5.2
	// 缺少必要参数
	StatusCodeParamsInvalid = "40003"
	// 未找到资源
	StatusCodeNotFound                     = "40004" // RFC 7231, 6.5.4
	// 权限不足
	StatusCodeMethodNotAllowed             = "40005" // RFC 7231, 6.5.5
	// StatusCodeNotAcceptable                = "406" // RFC 7231, 6.5.6
	// StatusCodeProxyAuthRequired            = "407" // RFC 7235, 3.2
	// StatusCodeRequestTimeout               = "408" // RFC 7231, 6.5.7
	// StatusCodeConflict                     = "409" // RFC 7231, 6.5.8
	// StatusCodeGone                         = "410" // RFC 7231, 6.5.9
	// StatusCodeLengthRequired               = "411" // RFC 7231, 6.5.10
	// StatusCodePreconditionFailed           = "412" // RFC 7232, 4.2
	// StatusCodeRequestEntityTooLarge        = "413" // RFC 7231, 6.5.11
	// StatusCodeRequestURITooLong            = "414" // RFC 7231, 6.5.12
	// StatusCodeUnsupportedMediaType         = "415" // RFC 7231, 6.5.13
	// StatusCodeRequestedRangeNotSatisfiable = "416" // RFC 7233, 4.4
	// StatusCodeExpectationFailed            = "417" // RFC 7231, 6.5.14
	// StatusCodeTeapot                       = "418" // RFC 7168, 2.3.3
	// StatusCodeMisdirectedRequest           = "421" // RFC 7540, 9.1.2
	// StatusCodeUnprocessableEntity          = "422" // RFC 4918, 11.2
	// StatusCodeLocked                       = "423" // RFC 4918, 11.3
	// StatusCodeFailedDependency             = "424" // RFC 4918, 11.4
	// StatusCodeTooEarly                     = "425" // RFC 8470, 5.2.
	// StatusCodeUpgradeRequired              = "426" // RFC 7231, 6.5.15
	// StatusCodePreconditionRequired         = "428" // RFC 6585, 3
	// StatusCodeTooManyRequests              = "429" // RFC 6585, 4
	// StatusCodeRequestHeaderFieldsTooLarge  = "431" // RFC 6585, 5
	// StatusCodeUnavailableForLegalReasons   = "451" // RFC 7725, 3

	// 内部错误
	// 数据库异常
	StatusCodeDatabaseError = "50000" // RFC 7231, 6.6.1
	// IO 异常
	StatusCodeIOError = "50001" // RFC 7231, 6.6.2
	// 依赖服务异常
	StatusCodeDependenceError = "50002" // RFC 7231, 6.6.3
	// 资源忙
	StatusCodeResourceBusyError = "50003" // RFC 7231, 6.6.4
	// StatusCodeGatewayTimeout                = "504" // RFC 7231, 6.6.5
	// StatusCodeHTTPVersionNotSupported       = "505" // RFC 7231, 6.6.6
	// StatusCodeVariantAlsoNegotiates         = "506" // RFC 2295, 8.1
	// StatusCodeInsufficientStorage           = "507" // RFC 4918, 11.5
	// StatusCodeLoopDetected                  = "508" // RFC 5842, 7.2
	// StatusCodeNotExtended                   = "510" // RFC 2774, 7
	// StatusCodeNetworkAuthenticationRequired = "511" // RFC 6585, 6
)

type HTTPOK struct {
	Msg  string `json:"msg" default:"ok"`
	Code string    `json:"code" default:"20000"`
}

type HTTPErr struct {
	Msg  string `json:"msg"`
	Code string    `json:"code" default:"-1" example:"-1"`
}

type Result struct {
	Msg  string      `json:"msg"`
	Code string         `json:"code"`
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

//func DefaultOK() HTTPOK {
//	return HTTPOK{
//		Msg:  "ok",
//		Code: 0,
//	}
//}
