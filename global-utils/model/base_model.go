package model

type Response struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data,omitempty"`
	Error      *ErrorLog   `json:"error,omitempty"`
	Pages      int64       `json:"pages,omitempty"`
	NumItems   int64       `json:"num_items,omitempty"`
	Total      int64       `json:"total,omitempty"`
	StatCode   int         `json:"stat_code,omitempty"`
	StatMsg    string      `json:"stat_msg,omitempty"`
	Message    string      `json:"message,omitempty"`
}

type ResponseChannel struct {
	Data  interface{} `json:"data,omitempty"`
	Error error       `json:"errors,omitempty"`
}

type ErrorLog struct {
	Line              string      `json:"line,omitempty"`
	Filename          string      `json:"filename,omitempty"`
	Function          string      `json:"function,omitempty"`
	Message           interface{} `json:"message,omitempty"`
	SystemMessage     string      `json:"system_message,omitempty"`
	Url               string      `json:"url,omitempty"`
	Method            string      `json:"method,omitempty"`
	Fields            interface{} `json:"fields,omitempty"`
	ConsumerTopic     string      `json:"consumer_topic,omitempty"`
	ConsumerPartition int         `json:"consumer_partition,omitempty"`
	ConsumerName      string      `json:"consumer_name,omitempty"`
	ConsumerOffset    int64       `json:"consumer_offset,omitempty"`
	ConsumerKey       string      `json:"consumer_key,omitempty"`
	Err               error       `json:"-"`
	StatusCode        int         `json:"-"`
}

type ResponseNotFound struct {
	StatusCode int          `json:"status_code"`
	Data       any          `json:"data"`
	Error      ErrorMessage `json:"error"`
	Total      int          `json:"total"`
}
type ErrorMessage struct {
	Message       string `json:"message"`
	SystemMessage string `json:"system_message"`
}

type JWTAccessTokenPayload struct {
	Id       string `json:"id,omitempty" bson:"id,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"  `
	Token    string `json:"token,omitempty" `
	Role     string `json:"role,omitempty"`
}

type JWTAccessTokenPayloadChan struct {
	JWTAccessTokenPayload *JWTAccessTokenPayload `json:"jwt_access_token_payload,omitempty"`
	Error                 error                  `json:"error,omitempty"`
	ErrorLog              *ErrorLog              `json:"error_log,omitempty"`
}
