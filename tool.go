// Created by EldersJavas(EldersJavas&gmail.com)

package ddtvgo

import (
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
