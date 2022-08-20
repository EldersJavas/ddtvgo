// Created by EldersJavas(EldersJavas&gmail.com)

package ddtvgo

import "encoding/json"

//-------------------------
//请求头结构体

func UnmarshalHeader(data []byte) (Header, error) {
	var r Header
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Header) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Header struct {
	Time        int64  `json:"time,omitempty"`
	Cmd         string `json:"cmd,omitempty"`
	Sig         string `json:"sig,omitempty"`
	Accesskeyid int64  `json:"accesskeyid,omitempty"`
}

//  -------------------------

//Request 请求结构
type Request struct {
	Name              string
	Header            *Header
	RequestMode       RequestMode
	RequestReturnType RequestReturnType
}

func NewRequest(name string, requestMode RequestMode, returnType RequestReturnType) *Request {
	r := &Header{Cmd: name}
	return &Request{Name: name, Header: r, RequestMode: requestMode, RequestReturnType: returnType}
}

//  -------------------------
