// Created by EldersJavas(EldersJavas&gmail.com)

package ddtvgo

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

// SHA1 Get SHA1
func SHA1(s string) string {
	t := sha1.New()
	_, err := io.WriteString(t, s)
	if err != nil {
		log.Fatalln("SHA1 Error")
	}
	return fmt.Sprintf("%x", t.Sum(nil))
}

func GetFullURL(app App, request Request) string {
	return fmt.Sprintf("%s/api/%s", app.APISite, request.Name)
}

func JsonIo2Text(bt io.ReadCloser) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(bt)
	if err != nil {
		return nil, err
	}
	out := buf.Bytes()
	return out, nil
}
