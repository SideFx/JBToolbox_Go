//----------------------------------------------------------------------------------------------------------
// Base64 Go implementation based on Wikipedia web site
// https://en.wikipedia.org/wiki/Base64
// (w) 2025 by Jan Buchholz
//----------------------------------------------------------------------------------------------------------

package JBCodec

import (
	"bytes"
)

var base64Array = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

const padding = '='

func Base64Encode(payload []byte) string {
	var result = ""
	var l = len(payload)
	var buf uint32
	var i = 0
	var j, k, idx, n int
	for i < l {
		buf = 0
		k = 0
		for j = 0; j < 3; j++ {
			buf <<= 8
			if i+j < l {
				buf |= uint32(payload[i+j])
				k++
			}
		}
		i += 3
		n = 3
		for j = 0; j <= k; j++ {
			idx = int((buf >> (6 * n)) & 0x3f)
			result += string(base64Array[idx])
			n--
		}
		for j = 0; j <= n; j++ {
			result += string(padding)
		}
	}
	return result
}

func Base64Decode(payload string) []byte {
	var result []byte = nil
	var l = len(payload)
	var i = 0
	var j, k, n, idx int
	var buf uint32
	var b byte
	for i < l {
		buf = 0
		k = -1
		for j = 0; j < 4; j++ {
			if i+j < l {
				buf <<= 6
				b = payload[i+j]
				if b != padding {
					idx = bytes.IndexByte(base64Array, b)
					if idx == -1 {
						i = l
						break
					}
					buf |= uint32(idx)
					k++
				}
			}
		}
		i += 4
		n = 16
		for j = 0; j < k; j++ {
			b = byte(buf >> n)
			result = append(result, b)
			n -= 8
		}
	}
	return result
}
