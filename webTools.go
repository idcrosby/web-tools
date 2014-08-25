package myTools

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"strings"
	"time"
)

func Base64Encode(encode []byte, url bool) string {
	var str string
	if (url) {
		str = base64.URLEncoding.EncodeToString(encode)
	} else {
		str = base64.StdEncoding.EncodeToString(encode)
	}
	
	return str
}

func Base64Decode(decode string, url bool) (buf []byte, err error) {
	if (url) {
		buf, err = base64.URLEncoding.DecodeString(decode)
	} else {
		buf, err = base64.StdEncoding.DecodeString(decode)
	}
	if err != nil {
		buf = nil
	}
	return
}

func ValidateJson(bytes []byte) (buf []byte, err error) {
	var f interface{}
	err = json.Unmarshal(bytes, &f)
	if err != nil {
		return nil, err
	}
	buf, err = json.MarshalIndent(&f, "", "   ")
	check(err)
	return
}

func FilterJson(bytes []byte, filter []string) (buf []byte, err error) {
	var f interface{}
	err = json.Unmarshal(bytes, &f)
	if err != nil {
		return nil, err
	}
	// Access the data's underlying interface
	m := f.(map[string]interface{})

	for _,element := range filter {
		subEls := strings.Split(element,".")
		node := m
		for index,sub := range subEls {
			//Check if last element
			if (index >= (len(subEls) -1)) {
				delete(node, sub)
			} else {
				node = m[sub].(map[string]interface{})
			}
		}
		delete(m, element)
	}
	buf, err = json.MarshalIndent(&m, "", "   ")
	check(err)
	return
}

func Md5Hash(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func Sha1Hash(data []byte) string {
	h := sha1.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func ConvertTimeFromEpoch(epoch int64) time.Time {
	return time.Unix(epoch, 0)
}

func ConvertTimeToEpoch(convert time.Time) int64 {
	return convert.Unix()
}

func check(err error) { if err != nil { panic(err) } }