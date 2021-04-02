package v2

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// 无返回的200
func JsonOK(c echo.Context) error {
	return c.JSON(http.StatusOK, HTTPOK{
		Msg:  "ok",
		Code: StatusCodeOK,
	})
}

// 带返回的200
func JsonRep(c echo.Context, result interface{}) error {
	return c.JSON(http.StatusOK, Result{
		Msg:  "ok",
		Code: StatusCodeOK,
		Data: result,
	})
}

// 通用JSON错误返回
func JsonError(c echo.Context, code, msg string) error {
	return c.JSON(http.StatusOK, Result{
		Msg:  msg,
		Code: code,
		Data: nil,
	})
}

func JsonNewError(code, msg string) Result {
	return Result{
		Msg:  msg,
		Code: code,
		Data: nil,
	}	
}

func JsonNewOk() Result {
	return Result{
		Msg:  "ok",
		Code: "20000",
		Data: nil,
	}
}

// 参数错误
func JsonParamErr(c echo.Context, msg ...interface{}) error {
	m, ok := (msg[0]).(string)
	if !ok || m == "" {
		m = "params missing or wrong types"
	}
	return c.JSON(http.StatusOK, Result{
		Msg:  m,
		Code: StatusCodeParamsInvalid,
	})
}

// 资源重复创建
func JsonDup(c echo.Context) error {
	return c.JSON(http.StatusOK, Result{
		Msg:  "resource already created",
		Code: StatusCodeCreated,
	})
}

// 数据库异常
func JsonDBErr(c echo.Context) error {
	return c.JSON(http.StatusOK, Result{
		Msg:  "database error",
		Code: StatusCodeDatabaseError,
	})
}

// 系统IO异常
func JsonIOErr(c echo.Context) error {
	return c.JSON(http.StatusOK, Result{
		Msg:  "database error",
		Code: StatusCodeIOError,
	})
}

// 依赖服务出错
func JsonDepErr(c echo.Context) error {
	return c.JSON(http.StatusOK, Result{
		Msg:  "database error",
		Code: StatusCodeDependenceError,
	})
}


// 资源忙
func JsonBusyErr(c echo.Context) error {
	return c.JSON(http.StatusOK, Result{
		Msg:  "database error",
		Code: StatusCodeResourceBusyError,
	})
}

// 分页返回数据
func JsonPageRep(c echo.Context, result PaginationData) error {
	return c.JSON(http.StatusOK, PaginationResult{
		HTTPOK: HTTPOK{
			Msg:  "ok",
			Code: StatusCodeOK,
		},
		Data:   result,
	})
}




//func JErr500(c echo.Context, err error) error {
//	return c.JSON(http.StatusInternalServerError, consts.HTTPErr{
//		Msg:  err.Error(),
//		Code: -1,
//	})
//}
//
//func JErr400(c echo.Context, err error) error {
//	return c.JSON(http.StatusInternalServerError, consts.HTTPErr{
//		Msg:  err.Error(),
//		Code: -1,
//	})
//}
//
//func JErr401(c echo.Context) error {
//	return c.JSON(http.StatusUnauthorized, consts.HTTPErr{
//		Msg:  "session check failed, please relogin",
//		Code: -1,
//	})
//}
//
//func JErr403(c echo.Context) error {
//	return c.JSON(http.StatusForbidden, consts.HTTPErr{
//		Msg:  "not allowed",
//		Code: -1,
//	})
//}
//
//func JRep(c echo.Context, result interface{}) error {
//	return c.JSON(http.StatusOK, consts.Result{
//		Msg:  "ok",
//		Code: 0,
//		Data: result,
//	})
//}
//
//func JPageRep(c echo.Context, result consts.PaginationData) error {
//	return c.JSON(http.StatusOK, consts.PaginationResult{
//		HTTPOK: consts.DefaultOK(),
//		Data:   result,
//	})
//}
//
//// PageInfo 从请求中获取分页信息
//func PageInfo(c echo.Context) (page, size int64) {
//	p := c.QueryParam("page")
//	s := c.QueryParam("size")
//	if p == "" {
//		page = 1
//	} else {
//		page, _ = strconv.ParseInt(p, 10, 64)
//	}
//
//	if s == "" {
//		size = 20
//	} else {
//		size, _ = strconv.ParseInt(s, 10, 64)
//	}
//	return
//}
