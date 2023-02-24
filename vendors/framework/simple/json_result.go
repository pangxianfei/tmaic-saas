package simple

type JsonResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

func Json(code int, message string, data interface{}, success bool) *JsonResult {
	return &JsonResult{
		Code:    code,
		Success: success,
		Message: message,
		Data:    data,
	}
}

func JsonData(data interface{}) *JsonResult {
	return &JsonResult{
		Code:    0,
		Success: true,
		Data:    data,
	}
}

func JsonItemList(data []interface{}) *JsonResult {
	return &JsonResult{
		Code:    0,
		Success: true,
		Data:    data,
	}
}

func JsonPageData(results interface{}, page *Paging) *JsonResult {
	return JsonData(&PageResult{
		Results: results,
		Page:    page,
	})
}

func JsonCursorData(results interface{}, cursor string, hasMore bool) *JsonResult {
	return JsonData(&CursorResult{
		Results: results,
		Cursor:  cursor,
		HasMore: hasMore,
	})
}

func JsonSuccess() *JsonResult {
	return &JsonResult{
		Code:    0,
		Success: true,
		Data:    nil,
	}
}

func JsonError(err *CodeError) *JsonResult {
	return &JsonResult{
		Code:    err.Code,
		Success: true,
		Message: err.Message,
		Data:    err.Data,
	}
}

func JsonErrorMsg(message string) *JsonResult {
	return &JsonResult{
		Code:    0,
		Success: false,
		Message: message,
		Data:    nil,
	}
}

func JsonCode(code int, message string) *JsonResult {
	return &JsonResult{
		Code:    code,
		Success: false,
		Message: message,
		Data:    nil,
	}
}

func JsonErrorData(code int, message string, data interface{}) *JsonResult {
	return &JsonResult{
		Code:    code,
		Success: false,
		Message: message,
		Data:    data,
	}
}

type RspBuilder struct {
	Data map[string]interface{}
}

func NewEmptyRspBuilder() *RspBuilder {
	return &RspBuilder{Data: make(map[string]interface{})}
}

func NewRspBuilder(obj interface{}) *RspBuilder {
	return NewRspBuilderExcludes(obj)
}

func NewRspBuilderExcludes(obj interface{}, excludes ...string) *RspBuilder {
	return &RspBuilder{Data: StructToMap(obj, excludes...)}
}

func (builder *RspBuilder) Put(key string, value interface{}) *RspBuilder {
	builder.Data[key] = value
	return builder
}

func (builder *RspBuilder) Build() map[string]interface{} {
	return builder.Data
}

func (builder *RspBuilder) JsonResult() *JsonResult {
	return JsonData(builder.Data)
}
