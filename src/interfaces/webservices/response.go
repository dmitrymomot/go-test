package webservices

type response struct {
	Data     interface{} `json:"data,omitempty"`
	Messages []string    `json:"messages,omitempty"`
	Error    *err        `json:"error,omitempty"`
}

type data map[string]interface{}

type err struct {
	Code     int         `json:"code"`
	Title    string      `json:"title"`
	Detail   interface{} `json:"detail,omitempty"`
	Internal error       `json:"-"`
}
