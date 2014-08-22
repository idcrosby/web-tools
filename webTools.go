package myTools

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
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

func Base64Decode(decode string, url bool) []byte {
	var data []byte
	var err error
	if (url) {
		data, err = base64.URLEncoding.DecodeString(decode)
	} else {
		data, err = base64.StdEncoding.DecodeString(decode)
	}
	if err != nil {
		fmt.Println("Error Decoding:", err)
		return nil
	}
	return data
}

func ValidateJson(bytes []byte) (buf []byte, err error) {
	var f interface{}
	err = json.Unmarshal(bytes, &f)
	if err != nil {
		fmt.Println("Error reading JSON: ", err)
		return nil, err
	}
	buf, err = json.MarshalIndent(&f, "", "   ")
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