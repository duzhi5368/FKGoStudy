package FKLib

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_HttpWrapper(t *testing.T) {
	pHttpWrapper := NewHttpWrapper()
	pHttpWrapper.GET("/user/:id", _Test_HttpWrapper_ShowUser)
	pHttpWrapper.GET("/foo/*", _Test_HttpWrapper_CatchFoo)
	pHttpWrapper.m_pNotFoundFunc = _Test_HttpWrapper_NotFound
	pHttpWrapper.Listen(":8008")
}

func _Test_HttpWrapper_ShowUser(tagWriter http.ResponseWriter, pRequest *http.Request) {
	var strID string
	// 获取该请求的映射表
	mapVars := GetVars(pRequest)
	strID = mapVars["id"]

	// 也可能只有一个
	strID = GetVar(pRequest, "id")

	fmt.Fprintf(tagWriter, "User ID: %s", strID)
}

func _Test_HttpWrapper_CatchFoo(tagWriter http.ResponseWriter, pRequest *http.Request) {
	tagWriter.Write([]byte("Gotta catch 'em all"))
}

func _Test_HttpWrapper_NotFound(tagWriter http.ResponseWriter, pRequest *http.Request) {
	tagWriter.Write([]byte("Nothing to see here"))
}
