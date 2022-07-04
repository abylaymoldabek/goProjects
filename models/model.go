package models

// import "encoding/xml"

type ResponseData struct {
	Date    string   `xml:"date" json:"date"`
	Item    []Item   `xml:"item" json:"item"`
}

type Item struct {
	Fullname    string   `xml:"fullname" json:"fullname"`  //АВСТРАЛИЙСКИЙ ДОЛЛАР
	Title       string   `xml:"title" json:"title"`  //AUD
	Description string   `xml:"description" json:"description"`  //331.35
}

// type ParsingData struct {
// 	Rates struct {
// 		XMLName     xml.Name `xml:"rates" json:"-"`
// 		Generator   string   `xml:"generator" json:"generator"`
// 		Title       string   `xml:"title" json:"title"`
// 		Link        string   `xml:"link" json:"link"`
// 		Description string   `xml:"description" json:"description"`
// 		Copyright   string   `xml:"copyright" json:"copyright"`
// 		Date        string   `xml:"date" json:"date"`
// 		Item        struct {
// 			XMLName     xml.Name `xml:"item" json:"-"`
// 			Fullname    string   `xml:"fullname" json:"fullname"`
// 			Title       string   `xml:"title" json:"title"`
// 			Description string   `xml:"description" json:"description"`
// 			Quant       string   `xml:"quant" json:"quant"`
// 			Index       string   `xml:"index" json:"index"`
// 			Change      string   `xml:"change" json:"change"`
// 		}
// 	}
// }
