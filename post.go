// Created by EldersJavas(EldersJavas&gmail.com)

package ddtvgo

import (
	"bytes"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

/*
"form-data":
{
   "time":1641149566,
   "cmd":"Room_Info",
   "sig":"566a322043b6217334bc15f4e6d18973d033aa4b",
   "accesskeyid":1
}


accesskeyid=1;accesskeysecret=2;cmd=room_info;time=1641149566
*/

func PostC(request *Request) (*http.Response, error) {
	client := http.Client{}
	formValues := url.Values{}

	formValues.Set("time", strconv.FormatInt(request.Header.Time, 10))
	formValues.Set("cmd", request.Header.Cmd)
	formValues.Set("sig", request.Header.Sig)
	formValues.Set("accesskeyid", request.Header.Accesskeyid)

	formDataStr := formValues.Encode()
	log.Println(formDataStr)
	formDataBytes := []byte(formDataStr)
	formBytesReader := bytes.NewReader(formDataBytes)
	resqq, err := client.Post(request.Url, "application/x-www-form-urlencoded", formBytesReader)
	if err != nil {
		return resqq, err
	}
	as, _ := JsonIo2Text(resqq.Body)
	log.Println(string(as), resqq.Status)
	return resqq, nil
}
