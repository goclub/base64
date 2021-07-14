package xbase64

import "encoding/base64"

func StdEncode (src []byte) (base64Bytes []byte) {
	enc := base64.StdEncoding
	base64Bytes = make([]byte, enc.EncodedLen(len(src)))
	enc.Encode(base64Bytes, src)
	return base64Bytes
}
func StdDecode(src []byte) (base64Bytes []byte, err error) {
	enc := base64.StdEncoding
	dbuf := make([]byte, enc.DecodedLen(len(src)))
	n, err := enc.Decode(dbuf, src)
	return dbuf[:n], err
}
func URLEncode (src []byte) (base64Bytes []byte) {
	enc := base64.URLEncoding
	base64Bytes = make([]byte, enc.EncodedLen(len(src)))
	enc.Encode(base64Bytes, src)
	return base64Bytes
}
func URLDecode(src []byte) (base64Bytes []byte, err error) {
	enc := base64.URLEncoding
	dbuf := make([]byte, enc.DecodedLen(len(src)))
	n, err := enc.Decode(dbuf, src)
	return dbuf[:n], err
}