package xbase64

import "encoding/base64"

func EncodeRawStd(source []byte) (base64Bytes []byte) {
	enc := base64.RawStdEncoding
	base64Bytes = make([]byte, enc.EncodedLen(len(source)))
	enc.Encode(base64Bytes, source)
	return base64Bytes
}
func DecodeRawStd(base64Bytes []byte) (source []byte, err error) {
	enc := base64.RawStdEncoding
	dbuf := make([]byte, enc.DecodedLen(len(base64Bytes)))
	n, err := enc.Decode(dbuf, base64Bytes)
	return dbuf[:n], err
}
func EncodeRawURL(src []byte) (base64Bytes []byte) {
	enc := base64.RawURLEncoding
	base64Bytes = make([]byte, enc.EncodedLen(len(src)))
	enc.Encode(base64Bytes, src)
	return base64Bytes
}
func DecodeRawURL(base64Bytes []byte) (source []byte, err error) {
	enc := base64.RawURLEncoding
	dbuf := make([]byte, enc.DecodedLen(len(base64Bytes)))
	n, err := enc.Decode(dbuf, base64Bytes);
	if err != nil {

		return
	}
	return dbuf[:n], err
}
