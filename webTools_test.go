package myTools

import (
	"testing"
	"time"
)

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
	if x := Base64Decode(str, false); string(x) != string(data) {
		t.Errorf("Base64Decode(" + str + ") = " + string(x) + ", want " + string(data))
	}
}

func TestBase64UrlEncode(t *testing.T) {
	str := "YWJjMTIzIT8kKiYoKSctPUB-"
	data := []byte("abc123!?$*&()'-=@~")
	if x:= Base64Encode(data, true); x != str  {
		t.Errorf("Base64Encode(" + string(data) + ") = " + x + ", want " + str)
	}	
}

func TestBase64UrlDecode(t *testing.T) {
	str := "JEBtcGxlICsgZGF0YSZ0b0VuYzBkZQ=="
	data := []byte("$@mple + data&toEnc0de")
	if x := Base64Decode(str, true); string(x) != string(data) {
		t.Errorf("Base64Decode(" + str + ") = " + string(x) + ", want " + string(data))
	}
}

func TestValidateJsonValid(t *testing.T) {
	validJson := []byte("{\"sample\":\"data\"}")
	if x, err := ValidateJson(validJson); x == nil || err != nil {
		t.Errorf("ValidateJson(" + string(validJson) + ") = nil")
	}
}

func TestValidateJsonInvalid(t *testing.T) {
	invalidJson := []byte("{\"sample\":\"data\",\"missing\"}")
	if x, err := ValidateJson(invalidJson); x != nil  || err == nil {
		t.Errorf("ValidateJson(" + string(invalidJson) + ") = " + string(x) + ", want nil")
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

