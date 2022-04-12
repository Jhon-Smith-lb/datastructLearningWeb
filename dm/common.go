package dm

type Resp struct {
	Code int64                  `json:"code"`
	Data map[string]interface{} `json:"data"`
	Msg  string                 `json:"msg"`
}

const (
	HTTP_OK int64 = 200
)
