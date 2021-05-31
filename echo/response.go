package echo

import (
	"net/http"
	"strconv"

	"github.com/deoops-net/easy/consts"
	"github.com/labstack/echo/v4"
)

// 实在懒得手敲 response
func JOK(c echo.Context) error {
	return c.JSON(http.StatusOK, consts.HTTPOK{
		Msg:  "ok",
		Code: 0,
	})
}

func JErr500(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, consts.HTTPErr{
		Msg:  err.Error(),
		Code: -1,
	})
}

func JErrRep(c echo.Context, err error, result interface{}) error {
	return c.JSON(http.StatusInternalServerError, consts.Result{
		Msg:  err.Error(),
		Code: -1,
		Data: result,
	})
}

func JErr400(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, consts.HTTPErr{
		Msg:  err.Error(),
		Code: -1,
	})
}

func JErr401(c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, consts.HTTPErr{
		Msg:  "session check failed, please relogin",
		Code: -1,
	})
}

func JErr403(c echo.Context) error {
	return c.JSON(http.StatusForbidden, consts.HTTPErr{
		Msg:  "not allowed",
		Code: -1,
	})
}

func JRep(c echo.Context, result interface{}) error {
	return c.JSON(http.StatusOK, consts.Result{
		Msg:  "ok",
		Code: 0,
		Data: result,
	})
}

func JPageRep(c echo.Context, result consts.PaginationData) error {
	return c.JSON(http.StatusOK, consts.PaginationResult{
		HTTPOK: consts.DefaultOK(),
		Data:   result,
	})
}

// PageInfo 从请求中获取分页信息
func PageInfo(c echo.Context) (page, size int64) {
	p := c.QueryParam("page")
	s := c.QueryParam("size")
	if p == "" {
		page = 1
	} else {
		page, _ = strconv.ParseInt(p, 10, 64)
	}

	if s == "" {
		size = 20
	} else {
		size, _ = strconv.ParseInt(s, 10, 64)
	}
	return
}
