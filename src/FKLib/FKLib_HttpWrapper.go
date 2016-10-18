package FKLib

//--------------------------------------------------------------
import (
	"fmt"
	"net/http"
)

//--------------------------------------------------------------
type SFKHttpWrapperHandler struct {
	m_strPatt  string
	m_sliParts []string
	m_bWild    bool
	http.HandlerFunc
}
type SHttpWrapper struct {
	m_mapHandlers   map[string][]*SFKHttpWrapperHandler
	m_pNotFoundFunc http.HandlerFunc
}

//--------------------------------------------------------------
var s_RequestsVars = map[*http.Request]map[string]string{}

//--------------------------------------------------------------
// 创建一个SHttpWrapper实体对象
func NewHttpWrapper() *SHttpWrapper {
	return &SHttpWrapper{make(map[string][]*SFKHttpWrapperHandler), nil}
}

//--------------------------------------------------------------
// 获取一个request请求的变量
func GetVars(pRequest *http.Request) map[string]string {
	if sliRets, ok := s_RequestsVars[pRequest]; ok {
		return sliRets
	}
	return nil
}

//--------------------------------------------------------------
// 获取一个request请求的指定变量
func GetVar(pRequest *http.Request, strVarKey string) string {
	var strRet string
	if Vars := GetVars(pRequest); Vars != nil {
		strRet, _ = Vars[strVarKey]
	}
	return strRet
}

//--------------------------------------------------------------
// 监听一个端口
func (thisHttpWrapper *SHttpWrapper) Listen(strPort string) {
	fmt.Printf("Listening port: %s\n", strPort[1:])
	http.ListenAndServe(strPort, thisHttpWrapper)
}

//--------------------------------------------------------------
// 尝试链接一个地址
func (pHander *SFKHttpWrapperHandler) _Try(sliUsegs []string) (bool, map[string]string) {
	if len(pHander.m_sliParts) != len(sliUsegs) && !pHander.m_bWild {
		return false, nil
	}

	tmpVars := map[string]string{}

	for nTmpIndex, strTmpPart := range pHander.m_sliParts {
		if strTmpPart == "*" {
			continue
		}
		if strTmpPart != "" && strTmpPart[0] == ':' {
			tmpVars[strTmpPart[1:]] = sliUsegs[nTmpIndex]
			continue
		}
		if strTmpPart != sliUsegs[nTmpIndex] {
			return false, nil
		}
	}

	return true, tmpVars
}

//--------------------------------------------------------------
// HTTP服务
func (thisHttpWrapper *SHttpWrapper) ServeHTTP(tagWriter http.ResponseWriter, pRequest *http.Request) {
	nUrlLen := len(pRequest.URL.Path)
	// 重定向Url
	if nUrlLen > 1 && pRequest.URL.Path[nUrlLen-1:] == "/" {
		http.Redirect(tagWriter, pRequest, pRequest.URL.Path[:nUrlLen-1], 301)
		return
	}
	// Split the URL into segments.
	// 分割Url为字符碎片
	sliSegments := SplitString(TrimString(pRequest.URL.Path, "/"), "/")
	// Map over the registered handlers for
	// the current request (if there is any).
	for _, handler := range thisHttpWrapper.m_mapHandlers[pRequest.Method] {
		// Try and match the pattern
		if handler.m_strPatt == pRequest.URL.Path {
			handler.ServeHTTP(tagWriter, pRequest)
			return
		}
		// Compare pattern segments to URL.
		if ok, v := handler._Try(sliSegments); ok {
			s_RequestsVars[pRequest] = v
			handler.ServeHTTP(tagWriter, pRequest)
			delete(s_RequestsVars, pRequest)
			return
		}
	}
	// 是否有自定义404返回
	if thisHttpWrapper.m_pNotFoundFunc != nil {
		tagWriter.WriteHeader(404)
		thisHttpWrapper.m_pNotFoundFunc.ServeHTTP(tagWriter, pRequest)
		return
	}
	// 默认返回404
	http.NotFound(tagWriter, pRequest)
}

//--------------------------------------------------------------
// 添加一个处理函数
func (thisHttpWrapper *SHttpWrapper) _Add(strMeth, strPatt string, pHandlerFunc http.HandlerFunc) {
	handler := &SFKHttpWrapperHandler{
		strPatt,
		SplitString(TrimString(strPatt, "/"), "/"),
		strPatt[len(strPatt)-1:] == "*",
		pHandlerFunc,
	}
	thisHttpWrapper.m_mapHandlers[strMeth] = append(
		thisHttpWrapper.m_mapHandlers[strMeth],
		handler,
	)
}

//--------------------------------------------------------------
// GET
func (thisHttpWrapper *SHttpWrapper) GET(strPatt string, pHandlerFunc http.HandlerFunc) {
	thisHttpWrapper._Add("GET", strPatt, pHandlerFunc)
}

//--------------------------------------------------------------
// HEAD
func (thisHttpWrapper *SHttpWrapper) HEAD(strPatt string, pHandlerFunc http.HandlerFunc) {
	thisHttpWrapper._Add("HEAD", strPatt, pHandlerFunc)
}

//--------------------------------------------------------------
// PUT
func (thisHttpWrapper *SHttpWrapper) PUT(strPatt string, pHandlerFunc http.HandlerFunc) {
	thisHttpWrapper._Add("PUT", strPatt, pHandlerFunc)
}

//--------------------------------------------------------------
// POST
func (thisHttpWrapper *SHttpWrapper) POST(strPatt string, pHandlerFunc http.HandlerFunc) {
	thisHttpWrapper._Add("POST", strPatt, pHandlerFunc)
}

//--------------------------------------------------------------
// DELETE
func (thisHttpWrapper *SHttpWrapper) DELETE(strPatt string, pHandlerFunc http.HandlerFunc) {
	thisHttpWrapper._Add("DELETE", strPatt, pHandlerFunc)
}

//--------------------------------------------------------------
// OPTIONS
func (thisHttpWrapper *SHttpWrapper) OPTIONS(strPatt string, pHandlerFunc http.HandlerFunc) {
	thisHttpWrapper._Add("OPTIONS", strPatt, pHandlerFunc)
}

//--------------------------------------------------------------
// PATCH
func (thisHttpWrapper *SHttpWrapper) PATCH(strPatt string, pHandlerFunc http.HandlerFunc) {
	thisHttpWrapper._Add("PATCH", strPatt, pHandlerFunc)
}

//--------------------------------------------------------------
