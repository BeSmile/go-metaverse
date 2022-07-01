package app

type Response struct {
	Code int         `json:"code" example: "200"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type Page struct {
	List      interface{} `json:"list"`
	Count     int         `json:"count"`
	PageIndex int         `json:"current"`
	PageSize  int         `json:"pageSize"`
}

type PageResponse struct {
	Response
	Data Page
}

func (res *Response) ReturnOk() *Response {
	res.Code = 200
	return res
}

func (res *Response) ReturnError(code int) *Response {
	res.Code = code
	return res
}

func (res *PageResponse) ReturnOK() *PageResponse {
	res.Code = 200
	return res
}
