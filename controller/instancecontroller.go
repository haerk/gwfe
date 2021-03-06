/**
 * @Author: lzw5399
 * @Date: 2021/1/16 21:43
 * @Desc: 流程实例
 */
package controller

import (
	"net/http"
	"strconv"

	"workflow/global/response"
	"workflow/model/request"
	"workflow/service"

	"github.com/labstack/echo/v4"
)

var (
	instanceSvc service.InstanceService = service.NewInstanceService()
)

// 创建新的实例
func StartProcessInstance(c echo.Context) error {
	var r request.InstanceRequest
	if err := c.Bind(&r); err != nil {
		return response.Failed(c, http.StatusBadRequest)
	}

	result, err := instanceSvc.Start(&r)
	if err != nil {
		return response.FailWithMsg(c, int(result), err)
	}

	return response.OkWithData(c, result)
}

// 获取一个实例
func GetProcessInstance(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return response.Failed(c, http.StatusBadRequest)
	}

	instance, err := instanceSvc.Get(uint(id))
	if err != nil {
		return response.Failed(c, http.StatusNotFound)
	}

	return response.OkWithData(c, instance)
}

// process instance list
func ListProcessInstances(c echo.Context) error {
	// 从queryString获取分页参数
	var r request.PagingRequest
	if err := c.Bind(&r); err != nil {
		return response.Failed(c, http.StatusBadRequest)
	}

	instances, err := instanceSvc.List(&r)
	if err != nil {
		return response.Failed(c, http.StatusInternalServerError)
	}

	return response.OkWithData(c, instances)
}

// 获取流程实例中的变量
func GetInstanceVariable(c echo.Context) error {
	var r request.GetVariableRequest
	if err := c.Bind(&r); err != nil {
		return response.Failed(c, http.StatusBadRequest)
	}

	resp, err := instanceSvc.GetVariable(&r)
	if err != nil {
		return response.FailWithMsg(c, http.StatusInternalServerError, err)
	}

	return response.OkWithData(c, resp)
}

func GetInstanceVariableList(c echo.Context) error {
	var r request.GetVariableListRequest
	if err := c.Bind(&r); err != nil {
		return response.Failed(c, http.StatusBadRequest)
	}
	variables, err := instanceSvc.ListVariables(&r)
	if err != nil {
		return response.FailWithMsg(c, http.StatusInternalServerError, err)
	}

	return response.OkWithData(c, variables)
}
