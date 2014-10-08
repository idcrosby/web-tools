web-tools
=========

## DESCRIPTION

Collection of my commonly used tools.
Written in Go.

## DETAILS

`func Base64Encode(encode []byte, url bool) string`

Returns a Base 64 encoded string of the given bytes, url param indicates to use URL specific encoding.

`func Base64Decode(decode string, url bool) []byte`

Returns decoded byte array for the given encoded string, url param indicates to use URL specific encoding.
Invalid input returns nil.

`func ValidateJson(bytes []byte) (buf []byte, err error)`

Formats and returns the given JSON data. Invalid JSON will return nil and a descriptive error.

`func JsonNegativeFilter(bytes []byte, filter []string) (buf []byte, err error)`

Removes elements of the given JSON data with names listed in the filter array (use dot separated names for nested fields)

`func BuildJsonStructure(bytes []byte) (buf []byte, err error)`

Return A JSON object representing the structure of the JSON object passed in, i.e. each field with it's corresponding JSON type

`func JsonPositiveFilter(bytes []byte, filter []string) (buf []byte, err error)`

Keeps only the elements of the given JSON data with names listed in the filter array. (Top level fields only)

`func JsonCompare(jsonOne []byte, jsonTwo []byte) (buf []byte, err error)`

Compares two JSON objects, and returnsa JSON object containing any differing elements.

`func Md5Hash(data []byte) string`

Returns an MD5 hash string of the given data.

`func Sha1Hash(data []byte) string`

Returns a SHA-1 hash string of the given data.

`func ConvertTimeFromEpoch(epoch int64) time.Time`

Returns the given Time as Unix Epoch time stamp (in seconds).

`func ConvertTimeToEpoch(convert time.Time) int64`

Convers the given Unix Epoch time stamp to Time type.


## TODO List

- Automate tests/coverage/performance(siege)
	- Travis CI?
- JSON filtering
	- Improve filtering performance (streaming)
	- Augment JSON comparison
- XML <-> JSON conversion?
- HATEOS Support/Validator/Creation/Expansion
- OAuth stub?
- Tests
	- individual checks
- Add Logging