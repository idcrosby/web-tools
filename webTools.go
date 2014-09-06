package myTools

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"reflect"
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

func UrlEncode(encode string) string {
	return url.QueryEscape(encode)
}

func UrlDecode(decode string) (output string, err error) {

	output, err = url.QueryUnescape(decode)

	if err != nil {
		output = ""
	}

	return
}

func ValidateJson(bytes []byte, pretty bool) (buf []byte, err error) {
	var f interface{}
	err = json.Unmarshal(bytes, &f)
	if err != nil {
		return nil, err
	}
	if pretty {
		buf, err = json.MarshalIndent(&f, "", "   ")
	} else {
		buf, err = json.Marshal(&f)
	}
	check(err)
	return
}

func JsonNegativeFilter(bytes []byte, filter []string, pretty bool) (buf []byte, err error) {
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
	if pretty {
		buf, err = json.MarshalIndent(&m, "", "   ")
	} else {
		buf, err = json.Marshal(&m)
	}
	check(err)
	return
}

func BuildJsonStructure(bytes []byte) (buf []byte, err error) {
	var f interface{}
	err = json.Unmarshal(bytes, &f)
	if err != nil {
		return nil, err
	}
	// Access the data's underlying interface
	m := f.(map[string]interface{})

	var result map[string]interface{}
	result = make(map[string]interface{})

	for key,value := range m {
		var valueType string
		// determine type of value
		switch value.(type) {
			case int, float64:
				valueType = "number"
			case string:
				valueType = "string"
			case bool:
				valueType = "boolean"
			case nil:
				valueType = "null"
			case map[string]interface{}:
				valueType = "object"
			case []interface{}:
				valueType = "array"
			default:
				valueType = "unknown"
		}
		result[key] = valueType
	}
	buf, err = json.MarshalIndent(&result, "", "   ")
	check(err)
	return
}

func JsonPositiveFilter(bytes []byte, filter []string, pretty bool) (buf []byte, err error) {

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
	if pretty {
		buf, err = json.MarshalIndent(&result, "", "   ")
	} else {
		buf, err = json.Marshal(&result)
	}
	check(err)
	return
}

func JsonCompare(jsonOne []byte, jsonTwo []byte) (buf []byte, err error) {
	// var result map[string]interface{}
	// result = make(map[string]interface{})

	var f,g interface{}
	err = json.Unmarshal(jsonOne, &f)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonTwo, &g)
	if err != nil {
		return nil, err
	}
	// Access the data's underlying interface
	m1 := f.(map[string]interface{})
	m2 := g.(map[string]interface{})

	result := CompareJson(m1, m2)

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

func CompareJson(m1 map[string]interface{}, m2 map[string]interface{}) map[string]interface{} {
	var result map[string]interface{}
	result = make(map[string]interface{})

	for k, el := range m1 {
		if el2,ok := m2[k]; ok {
			if reflect.ValueOf(el).Kind() == reflect.Map {
				subCompare := CompareJson(el.(map[string]interface{}), el2.(map[string]interface{}))
				if len(subCompare) > 0 {
					result[k] = subCompare
				}
			} else if reflect.ValueOf(el).Kind() == reflect.Slice {
				if !CompareSlices(el.([]interface{}), el2.([]interface{})) {
					result[k + "_1"] = el
					result[k + "_2"] = el2
				}
			} else {
				if el != el2 {
					result[k + "_1"] = el //"diff" //el + "/" + el2
					result[k + "_2"] = el2
				}
			}
		} else {
			result[k] = el
		}
	}
	return result
}

// Compares two slices, returns true only if they contain the same elements in the same order
func CompareSlices(s1 []interface{}, s2 []interface{}) bool {

	if len(s1) != len(s2) {
		return false
	}

	for ix,el := range s1 {
		if s2[ix] != el {
			return false
		}
	}

	return true
}

func check(err error) { if err != nil { panic(err) } }