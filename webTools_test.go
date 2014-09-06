package myTools

import (
	"strings"
	"testing"
	"time"
)


// Basic Unit Tests

func TestBase64Encode(t *testing.T) {
	str := "YWJjMTIzIT8kKiYoKSctPUB+"
	data := []byte("abc123!?$*&()'-=@~")
	if x:= Base64Encode(data, false); x != str  {
		t.Errorf("Base64Encode(" + string(data) + ") = " + x + ", want " + str)
	}	
}

func TestBase64Decode(t *testing.T) {
	str := "JEBtcGxlICsgZGF0YSZ0b0VuYzBkZQ=="
	data := []byte("$@mple + data&toEnc0de")
	if x, err := Base64Decode(str, false); string(x) != string(data) || err != nil {
		t.Errorf("Base64Decode(" + str + ") = " + string(x) + ", want " + string(data))
	}
}

func TestBase64DecodeBadData(t *testing.T) {
	str := "Jdsadsad$$#@CsgZGF0YSZ0b0VuYzBkZQ=="
	if x, err := Base64Decode(str, false); x != nil || err == nil {
		t.Errorf("Base64Decode(" + str + ") = " + string(x) + ", want nil")
	}
}

func TestBase64UrlEncode(t *testing.T) {
	encoded := "URL+encode%3A%2F%2Fthis"
	data := "URL encode://this"
	if x:= UrlEncode(data); x != encoded  {
		t.Errorf("UrlEncode(" + data + ") = " + x + ", want " + encoded)
	}	
}

func TestBase64UrlDecode(t *testing.T) {
	encoded := "URL%20encode%3A%2F%2Fthis"
	data := "URL encode://this"
	if x, err := UrlDecode(encoded); x != data || err != nil {
		t.Errorf("UrlDecode(" + encoded + ") = " + x + ", want " + data)
	}
}

func TestValidateJsonValid(t *testing.T) {
	validJson := []byte("{\"sample\":\"data\"}")
	if x, err := ValidateJson(validJson, true); x == nil || err != nil {
		t.Errorf("ValidateJson(" + string(validJson) + ") = nil")
	}
}

func TestValidateJsonInvalid(t *testing.T) {
	invalidJson := []byte("{\"sample\":\"data\",\"missing\"}")
	if x, err := ValidateJson(invalidJson, false); x != nil  || err == nil {
		t.Errorf("ValidateJson(" + string(invalidJson) + ") = " + removeWhiteSpace(string(x)) + ", want nil")
	}
}

func TestJsonNegativeFilter(t *testing.T) {
	jsonIn := []byte("{\"keep\":\"goodData\",\"remove\":\"bad data\"}")
	jsonOut := "{\"keep\":\"goodData\"}"
	filter := []string{"remove"}

	if x, err := JsonNegativeFilter(jsonIn, filter, true); removeWhiteSpace(string(x)) != jsonOut || err != nil {
		t.Errorf("JsonNegativeFilter(" + string(jsonIn) + "," + arrayToString(filter) + ") = " + removeWhiteSpace(string(x)) + ", want " + string(jsonOut))
	}
}

func TestDeepJsonNegativeFilter(t *testing.T) {
	jsonIn := []byte("{\"keep\":\"goodData\",\"remove\":{\"subOne\":\"bad data\",\"subTwo\":true}}")
	jsonOut := "{\"keep\":\"goodData\",\"remove\":{\"subTwo\":true}}"
	filter := []string{"remove.subOne"}

	if x, err := JsonNegativeFilter(jsonIn, filter, false); removeWhiteSpace(string(x)) != jsonOut || err != nil {
		t.Errorf("JsonNegativeFilter(" + string(jsonIn) + "," + arrayToString(filter) + ") = " + removeWhiteSpace(string(x)) + ", want " + string(jsonOut))
	}
}

func TestDeepMultiJsonNegativeFilter(t *testing.T) {
	jsonIn := []byte("{\"keep\":\"goodData\",\"remove\":{\"subOne\":\"bad data\",\"subTwo\":true,\"subThree\":{\"deepRemove\":false}}}")
	jsonOut := "{\"remove\":{\"subThree\":{\"deepRemove\":false},\"subTwo\":true}}"
	filter := []string{"remove.subOne","keep","remove.subThree.missing"}

	if x, err := JsonNegativeFilter(jsonIn, filter, true); removeWhiteSpace(string(x)) != jsonOut || err != nil {
		t.Errorf("JsonNegativeFilter(" + string(jsonIn) + "," + arrayToString(filter) + ") = " + removeWhiteSpace(string(x)) + ", want " + string(jsonOut))
	}
}

func TestDeepMultiJsonNegativeFilterTwo(t *testing.T) {
	jsonIn := []byte("{\"keep\":\"goodData\",\"remove\":{\"subOne\":\"bad data\",\"subTwo\":true,\"subThree\":{\"deepRemove\":false}}}")
	jsonOut := "{\"remove\":{\"subThree\":{},\"subTwo\":true}}"
	filter := []string{"remove.subOne","keep","remove.subThree.deepRemove"}

	if x, err := JsonNegativeFilter(jsonIn, filter, true); removeWhiteSpace(string(x)) != jsonOut || err != nil {
		t.Errorf("JsonNegativeFilter(" + string(jsonIn) + "," + arrayToString(filter) + ") = " + removeWhiteSpace(string(x)) + ", want " + string(jsonOut))
	}
}

func TestBadJsonNegativeFilter(t *testing.T) {
	jsonIn := []byte("{\"keep\":\"goodData\",\"remove\":bad data}")
	filter := []string{"remove"}

	if x, err := JsonNegativeFilter(jsonIn, filter, false); x != nil || err == nil {
		t.Errorf("JsonNegativeFilter(" + string(jsonIn) + "," + arrayToString(filter) + ") = " + removeWhiteSpace(string(x)) + ", want nil")
	}
}

func TestJsonPositiveFilterSingle(t *testing.T) {
	jsonIn := []byte("{\"keep\":\"goodData\",\"remove\":\"bad data\"}")
	jsonOut := "{\"keep\":\"goodData\"}"
	filter := []string{"keep"}

	if x, err := JsonPositiveFilter(jsonIn, filter, false); removeWhiteSpace(string(x)) != jsonOut || err != nil {
		t.Errorf("JsonPositiveFilter(" + string(jsonIn) + "," + arrayToString(filter) + ") = " + removeWhiteSpace(string(x)) + ", want " + string(jsonOut))
	}
}

func TestDeepJsonPositiveFilter(t *testing.T) {
	jsonIn := []byte("{\"keepMe\":\"goodData\",\"remove\":{\"subOne\":\"bad data\",\"subTwo\":true}}")
	jsonOut := "{\"keepMe\":\"goodData\",\"remove\":{\"subTwo\":true}}"
	filter := []string{"keepMe","remove.subTwo"}

	if x, err := JsonPositiveFilter(jsonIn, filter, true); removeWhiteSpace(string(x)) != jsonOut || err != nil {
		t.Errorf("JsonPositiveFilter(" + string(jsonIn) + "," + arrayToString(filter) + ") = " + removeWhiteSpace(string(x)) + ", want " + string(jsonOut))
	}
}

func TestDeepMultiJsonPositiveFilter(t *testing.T) {
	jsonIn := []byte("{\"keep\":\"goodData\",\"remove\":{\"subOne\":\"bad_data\",\"subTwo\":true}}")
	jsonOut := "{\"keep\":\"goodData\",\"remove\":{\"subOne\":\"bad_data\",\"subTwo\":true}}"
	filter := []string{"remove","keep"}

	if x, err := JsonPositiveFilter(jsonIn, filter, true); removeWhiteSpace(string(x)) != jsonOut || err != nil {
		t.Errorf("JsonPositiveFilter(" + string(jsonIn) + "," + arrayToString(filter) + ") = " + removeWhiteSpace(string(x)) + ", want " + string(jsonOut))
	}
}

func TestJsonPositiveFilterAnother(t *testing.T) {
	jsonIn := []byte("{\"keep\":\"goodData\",\"remove\":{\"subOne\":\"bad_data\",\"subTwo\":true}}")
	jsonOut := "{\"keep\":\"goodData\",\"remove\":{\"subOne\":\"bad_data\",\"subTwo\":true}}"
	filter := []string{"remove.subOne","keep","remove.subTwo"}

	if x, err := JsonPositiveFilter(jsonIn, filter, false); removeWhiteSpace(string(x)) != jsonOut || err != nil {
		t.Errorf("JsonPositiveFilter(" + string(jsonIn) + "," + arrayToString(filter) + ") = " + removeWhiteSpace(string(x)) + ", want " + string(jsonOut))
	}
}

func TestBadJsonPositiveFilterSingle(t *testing.T) {
	jsonIn := []byte("{\"keep\":\"goodData\",\"remove\":bad data}")
	filter := []string{"keep"}

	if x, err := JsonPositiveFilter(jsonIn, filter, true); x != nil || err == nil {
		t.Errorf("JsonPositiveFilter(" + string(jsonIn) + "," + arrayToString(filter) + ") = " + removeWhiteSpace(string(x)) + ", want nil")
	}
}

func TestBuildJsonStructureBasic(t *testing.T) {
	jsonIn := []byte("{\"alpha\":\"goodData\",\"beta\":3,\"delta\":true,\"epsilon\":null,\"gamma\":[],\"zeta\":{}}")
	jsonOut := "{\"alpha\":\"string\",\"beta\":\"number\",\"delta\":\"boolean\",\"epsilon\":\"null\",\"gamma\":\"array\",\"zeta\":\"object\"}"
	if x, err := BuildJsonStructure(jsonIn); removeWhiteSpace(string(x)) != jsonOut || err != nil {
		t.Errorf("BuildJsonStructure(" + string(jsonIn) + ") = " + string(x) + ", want " + jsonOut)
	}
}

func TestJsonCompareNoDiff(t *testing.T) {
	jsonIn := []byte("{\"keep\":\"goodData\",\"remove\":\"bad data\"}")
	result := "{}"

	if x, err := JsonCompare(jsonIn, jsonIn); removeWhiteSpace(string(x)) != result || err != nil {
		if err != nil {
			t.Errorf("JsonCompare(" + string(jsonIn) + ", " + string(jsonIn) + ") throwing Error: " + err.Error())
		} else {
			t.Errorf("JsonCompare(" + string(jsonIn) + ", " + string(jsonIn) + ") = " + removeWhiteSpace(string(x)) + ", want " + result)
		}
	}
}

func TestJsonCompareSimple(t *testing.T) {
	jsonOne := []byte("{\"keep\":\"goodData\",\"remove\":\"bad data\",\"third\":false}")
	jsonTwo := []byte("{\"keep\":\"other_data\",\"remove\":\"bad data\"}")
	result := "{\"keep_1\":\"goodData\",\"keep_2\":\"other_data\",\"third\":false}"

	if x, err := JsonCompare(jsonOne, jsonTwo); removeWhiteSpace(string(x)) != result || err != nil {
		if err != nil {
			t.Errorf("JsonCompare(" + string(jsonOne) + ", " + string(jsonTwo) + ") throwing Error: " + err.Error())
		} else {
			t.Errorf("JsonCompare(" + string(jsonOne) + ", " + string(jsonTwo) + ") = " + removeWhiteSpace(string(x)) + ", want " + result)
		}
	}
}

func TestJsonCompareComplex(t *testing.T) {
	jsonOne := []byte("{\"keep\":\"goodData\",\"remove\":\"bad data\",\"more\":[4,5],\"same\":[1,2,3]}")
	jsonTwo := []byte("{\"keep\":{\"sub\":true},\"remove\":\"bad data\",\"more\":[4,5,6],\"same\":[1,2,3]}")
	result := "{\"keep_1\":\"goodData\",\"keep_2\":{\"sub\":true},\"more_1\":[4,5],\"more_2\":[4,5,6]}"

	if x, err := JsonCompare(jsonOne, jsonTwo); removeWhiteSpace(string(x)) != result || err != nil {
		if err != nil {
			t.Errorf("JsonCompare(" + string(jsonOne) + ", " + removeWhiteSpace(string(jsonTwo)) + ") throwing Error: " + err.Error())
		} else {
			t.Errorf("JsonCompare(" + string(jsonOne) + ", " + removeWhiteSpace(string(jsonTwo)) + ") = " + removeWhiteSpace(string(x)) + ", want " + result)
		}
	}
}

func TestJsonCompareMoar(t *testing.T) {
	jsonOne := []byte("{\"keep\":\"goodData\",\"remove\":\"bad data\",\"more\":[4,5],\"bajs\":{\"gris\":null},\"anka\":{\"kaja\":1}}")
	jsonTwo := []byte("{\"keep\":{\"sub\":true},\"remove\":\"bad data\",\"more\":[5,4],\"bajs\":{\"gris\":null},\"anka\":{\"kaja\":2}}")
	result := "{\"anka\":{\"kaja_1\":1,\"kaja_2\":2},\"keep_1\":\"goodData\",\"keep_2\":{\"sub\":true},\"more_1\":[4,5],\"more_2\":[5,4]}"

	if x, err := JsonCompare(jsonOne, jsonTwo); removeWhiteSpace(string(x)) != result || err != nil {
		if err != nil {
			t.Errorf("JsonCompare(" + string(jsonOne) + ", " + string(jsonTwo) + ") throwing Error: " + err.Error())
		} else {
			t.Errorf("JsonCompare(" + string(jsonOne) + ", " + string(jsonTwo) + ") = " + removeWhiteSpace(string(x)) + ", want " + result)
		}
	}
}

func TestJsonCompareBadJsonOne(t *testing.T) {
	jsonOne := []byte("{\"keep\":\"goodData\",\"remove\":\"bad data\",\"more\":[4,5]}")
	jsonTwo := []byte("{\"keep\":{\"sub\":true},\"remove\":\"bad data\",\"more\":4,5,6]}")

	if x, err := JsonCompare(jsonOne, jsonTwo); x != nil || err == nil {
		t.Errorf("JsonCompare(" + string(jsonOne) + ", " + string(jsonTwo) + ") = " + removeWhiteSpace(string(x)) + ", want error")
	}
}

func TestJsonCompareBadJsonTwo(t *testing.T) {
	jsonOne := []byte("{\"keep\":\"goodData\",\"remove\":\"bad data,\"more\":[4,5]}")
	jsonTwo := []byte("{\"keep\":{\"sub\":true},\"remove\":\"bad data\",\"more\":[4,5,6]}")

	if x, err := JsonCompare(jsonOne, jsonTwo); x != nil || err == nil {
		t.Errorf("JsonCompare(" + string(jsonOne) + ", " + string(jsonTwo) + ") = " + removeWhiteSpace(string(x)) + ", want error")
	}
}

func TestMd5Hash(t *testing.T) {
	data := []byte("h@$H_th1s *tr1ng")
	hash := "7952b401cd8a52dc40536d7e1b4c6658"
	if x := Md5Hash(data); x != hash {
		t.Errorf("Md5Hash(" + string(data) + ") = " + x + ", want " + hash)
	}
}

func TestSha1Hash(t *testing.T) {
	data := []byte("h@$H_th1s *tr1ng")
	hash := "832ed14e88a3e412ebf3474c83ce01cf99c90b87"
	if x := Sha1Hash(data); x != hash {
		t.Errorf("Sha1Hash(" + string(data) + ") = " + x + ", want " + hash)
	}
}

func TestConvertTimeToEpoch(t *testing.T) {
	var epoch int64
	epoch = 1408109234
	timeString := "2014-08-15 15:27:14 +0200 CEST"
	myTime, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", timeString)
	if x := ConvertTimeToEpoch(myTime); x != epoch {
		t.Errorf("ConvertTimeToEpoch(" + myTime.String() + ")=%v, want %v", x, epoch)
	}
}

func TestConvertTimeFromEpoch(t *testing.T) {
	var epoch int64
	epoch = 1408109234
	timeString := "2014-08-15 15:27:14 +0200 CEST"
	if x := ConvertTimeFromEpoch(epoch); x.String() != timeString {
		t.Errorf("ConvertTimeFromEpoch(%v)=" + x.String() + ", want " + timeString, epoch)
	}
}


func TestMergeJson(t *testing.T) {

}

// Benchmark Tests

func BenchmarkBase64Encode(b *testing.B) {
	data := []byte("abc123!?$*&()'-=@~")
	for i := 0; i < b.N; i++ {
		Base64Encode(data, false)
	}
}

func BenchmarkBase64Decode(b *testing.B) {
	data := "JEBtcGxlICsgZGF0YSZ0b0VuYzBkZQ=="
	for i := 0; i < b.N; i++ {
		Base64Decode(data, false)
	}
}

func BenchmarkValidateJson(b *testing.B) {
	data := []byte("{\"sample\":\"data\"}")
	for i := 0; i < b.N; i++ {
		ValidateJson(data, false)
	}
}

func BenchmarkJsonNegativeFilter(b *testing.B) {
	data := []byte("{\"keep\":\"goodData\",\"remove\":\"bad data\"}")
	filter := []string{"remove"}
	for i := 0; i < b.N; i++ {
		JsonNegativeFilter(data, filter, false)
	}
}

func BenchmarkJsonPositiveFilter(b *testing.B) {
	data := []byte("{\"keep\":\"goodData\",\"remove\":\"bad data\"}")
	filter := []string{"keep"}
	for i := 0; i < b.N; i++ {
		JsonPositiveFilter(data, filter, false)
	}
}

func BenchmarkJsonCompare(b * testing.B) {
	jsonOne := []byte("{\"keep\":\"goodData\",\"remove\":\"bad data\"}")
	jsonTwo := []byte("{\"keep\":\"other_data\",\"remove\":\"bad data\"}")
	for i := 0; i < b.N; i++ {
		JsonCompare(jsonOne, jsonTwo)
	}
}

func BenchmarkMd5Hash(b *testing.B) {
	data := []byte("h@$H_th1s *tr1ng")
	for i := 0; i < b.N; i++ {
		Md5Hash(data)
	}
}

func BenchmarkSha1Hash(b *testing.B) {
	data := []byte("h@$H_th1s *tr1ng")
	for i := 0; i < b.N; i++ {
		Sha1Hash(data)
	}
}

func BenchmarkConvertTimeToEpoch(b *testing.B) {
	timeString := "2014-08-15 15:27:14 +0200 CEST"
	data, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", timeString)
	for i := 0; i < b.N; i++ {
		ConvertTimeToEpoch(data)
	}
}

func BenchmarkConvertTimeFromEpoch(b *testing.B) {
	var data int64
	data = 1408109234
	for i := 0; i < b.N; i++ {
		ConvertTimeFromEpoch(data)
	}
}

// TODO Create Util function to compare JSON equivalence

// Util Methods

func arrayToString(input []string) (output string) {
	for _, value := range input { output += string(value) }
	return
}

func removeWhiteSpace(in string) (out string) {
	return arrayToString(strings.Fields(in))
}