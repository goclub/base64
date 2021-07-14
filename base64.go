package xbase64

import "encoding/base64"

func RawStdEncode (src []byte) (base64Bytes []byte) {
	enc := base64.RawStdEncoding
	base64Bytes = make([]byte, enc.EncodedLen(len(src)))
	enc.Encode(base64Bytes, src)
	return base64Bytes
}
func RawStdDecode(src []byte) (base64Bytes []byte, err error) {
	enc := base64.RawStdEncoding
	dbuf := make([]byte, enc.DecodedLen(len(src)))
	n, err := enc.Decode(dbuf, src)
	return dbuf[:n], err
}
func RawURLEncode (src []byte) (base64Bytes []byte) {
	enc := base64.RawURLEncoding
	base64Bytes = make([]byte, enc.EncodedLen(len(src)))
	enc.Encode(base64Bytes, src)
	return base64Bytes
}
func RawURLDecode(src []byte) (base64Bytes []byte, err error) {
	enc := base64.RawURLEncoding
	dbuf := make([]byte, enc.DecodedLen(len(src)))
	n, err := enc.Decode(dbuf, src)
	return dbuf[:n], err
}