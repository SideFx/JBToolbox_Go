//----------------------------------------------------------------------------------------------------------
// Constants & common functions
// (w) 2025 by Jan Buchholz
//----------------------------------------------------------------------------------------------------------

package JBCommon

import "strings"

const Uint64Bits = 64
const Uint32Bits = 32
const MaxUint64 = ^uint64(0)

var hex = [...]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

func Rol32(x uint32, n int) uint32 {
	return (x << n) | (x >> (Uint32Bits - n))
}

func Ror32(x uint32, n int) uint32 {
	return (x >> n) | (x << (Uint32Bits - n))
}

func Ror64(x uint64, n int) uint64 {
	return (x >> n) | (x << (Uint64Bits - n))
}

func ByteArrayToHexString(bytes []byte) string {
	var result = ""
	for _, b := range bytes {
		result += string(hex[(b>>4)&0x0f])
		result += string(hex[b&0x0f])
	}
	return result
}

func Uint32ToHexString(value uint32) string {
	var result = ""
	var b byte
	j := 24
	for i := 0; i < 4; i++ {
		b = byte(value >> j)
		result += string(hex[(b>>4)&0xf])
		result += string(hex[b&0xf])
		j -= 8
	}
	return result
}

func Uint64ToHexString(value uint64) string {
	var result = ""
	var b byte
	j := 56
	for i := 0; i < 8; i++ {
		b = byte(value >> j)
		result += string(hex[(b>>4)&0xf])
		result += string(hex[b&0xf])
		j -= 8
	}
	return result
}

func WordsInString(str string) int {
	return len(strings.Fields(str))
}
