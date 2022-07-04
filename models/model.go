package models

type ResponseResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func OkWithData(message string, data interface{}) *ResponseResult {
	return &ResponseResult{0, message, data}
}

type ResponseData struct {
	Date string `xml:"date" json:"date"`
	Item []Item `xml:"item" json:"item"`
}

type Item struct {
	Fullname    string `xml:"fullname" json:"fullname"`       //АВСТРАЛИЙСКИЙ ДОЛЛАР
	Title       string `xml:"title" json:"title"`             //AUD
	Description string `xml:"description" json:"description"` //331.35
}
