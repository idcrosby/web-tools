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

func JsonNegativeFilter(bytes []byte, filter []string) (buf []byte, err error) {
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
				node = node[sub].(map[string]interface{})
			}
		}
		delete(m, element)
	}
	buf, err = json.MarshalIndent(&m, "", "   ")
	check(err)
	return
}

func JsonPositiveFilter(bytes []byte, filter []string) (buf []byte, err error) {

	var f interface{}
	err = json.Unmarshal(bytes, &f)
	if err != nil {
		return nil, err
	}
	// Access the data's underlying interface
	m := f.(map[string]interface{})

	var result map[string]interface{}
	result = make(map[string]interface{})

	for _,element := range filter {
		subEls := strings.Split(element,".")
		node := m

		for index,sub := range subEls {
			//Check if last element
			if (index >= (len(subEls) -1)) {
				// Check if element exists
				if el,ok := node[sub]; ok {
					result = MergeJson(result, subEls, el)
				}
			} else {
				node = m[sub].(map[string]interface{})
			}
		}
	}
	buf, err = json.MarshalIndent(&result, "", "   ")
	check(err)
	return
}

// TODO Deprecated - Keeping for potential performance comparison
// func JsonPositiveFilterSingle(bytes []byte, filter []string) (buf []byte, err error) {
// 	var f interface{}
// 	err = json.Unmarshal(bytes, &f)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Access the data's underlying interface
// 	m := f.(map[string]interface{})

// 	var found bool
// 	// For each field in JSON, check if contained in filter
// 	for k := range m {
// 		for _,el := range filter {
// 			if el == k {
// 				found = true
// 				// skip
// 				break;
// 			}
// 		}
// 		if !found {
// 			delete(m, k)		
// 		}
// 		found = false
// 	}
// 	buf, err = json.MarshalIndent(&m, "", "   ")
// 	check(err)
// 	return
// }

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


// Utility Functions

func MergeJson(input map[string]interface{}, subEls []string, el interface{}) map[string]interface{} {

	result := input

	// If last element, add 
	// TODO currently will overwrite, add flag to potentially skip
	if len(subEls) == 1 {
		result[subEls[0]] = el
		return result
	}

	var node map[string]interface{}
	if _,ok := result[subEls[0]]; ok {
		node = result[subEls[0]].(map[string]interface{})
	} else {
		node = make(map[string]interface{})
	}	
	result[subEls[0]] = MergeJson(node, subEls[1:], el)

	return result
}

func check(err error) { if err != nil { panic(err) } }