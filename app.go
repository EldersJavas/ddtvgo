// Created by EldersJavas(EldersJavas&gmail.com)

package ddtvgo

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func (a *App) RunCmd(request *Request) (interface{}, error) {
	timenow := time.Now().UTC().Unix()
	request.Url = a.APISite + request.Name
	request.Header = &Header{Cmd: request.Name,
		Sig: SHA1(fmt.Sprintf("accesskeyid=%s;accesskeysecret=%s;cmd=%s;time=%s;",
			a.AccessKeyId,
			a.AccessKeySecret,
			strings.ToLower(request.Name),
			strconv.FormatInt(timenow, 10))),
		Accesskeyid: a.AccessKeyId,
		Time:        timenow}
	log.Println(fmt.Sprintf("accesskeyid=%s;accesskeysecret=%s;cmd=%s;time=%s;",
		a.AccessKeyId,
		a.AccessKeySecret,
		strings.ToLower(request.Name),
		strconv.FormatInt(timenow, 10)))
	switch request.RequestMode {
	case Post:
		r, err := PostC(request)
		if err != nil {
			return nil, err
		}
		jsonbyte, err := JsonIo2Text(r.Body)
		if err != nil {
			return nil, err
		}

		return jsonbyte, nil
	case Get:
		return nil, nil
	}
	return nil, nil
}
