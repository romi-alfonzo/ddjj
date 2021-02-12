package extract

import (
	"encoding/json"
	"fmt"
	"github.com/InstIDEA/ddjj/parser/declaration"
	"time"
)

type ParserData struct {
	Message []string                 `json:"message"`
	Status  int                      `json:"status"`
	Data    *declaration.Declaration `json:"data"`
	Raw     []string                 `json:"raw"`
}

type ExpectedValue int

const (
	EVdate ExpectedValue = iota
	EValphaNum
	EVnum
)

func (parser ParserData) AddMessage(msg string) {
	parser.Message = append(parser.Message, msg)
}

func (parser ParserData) AddError(msg error) {
	parser.Message = append(parser.Message, msg.Error())
}

func Create() ParserData {
	return ParserData{
		Message: make([]string, 0),
		Status:  0,
		Data:    nil,
		Raw:     make([]string, 0),
	}
}

func (parser ParserData) SetData(d *declaration.Declaration) {
	parser.Data = d
}

func (parser ParserData) RawData(s string) {
	parser.Raw = append(parser.Raw, s)
}

func (parser ParserData) ParserPrint() {
	b, err := json.MarshalIndent(parser, "", "\t")
	if err != nil {
		fmt.Println("{ message: null, status: 0, data: null, raw: null }")
		return
	}
	fmt.Println(string(b))
}

func (parser ParserData) Check(val string, err error) string {
	if err != nil {
		parser.AddError(err)
	}
	return val
}

func (parser ParserData) CheckInt(val int, err error) int {
	if err != nil {
		parser.AddError(err)
	}
	return val
}

func (parser ParserData) CheckTime(val time.Time, err error) time.Time {
	if err != nil {
		parser.AddError(err)
	}
	return val
}
