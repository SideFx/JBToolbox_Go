//----------------------------------------------------------------------------------------------------------
// CRC64 (reversed ECMA-182 & reversed ISO) Go implementation based on Wikipedia's description
// https://en.wikipedia.org/wiki/Cyclic_redundancy_check
// (w) 2025 by Jan Buchholz
//----------------------------------------------------------------------------------------------------------

package JBCrypto

import "JBToolbox/JBCommon"

const ecmaPoly uint64 = 0xc96c5795d7870f42
const isoPoly uint64 = 0xd800000000000000

var crc64Ecma [256]uint64
var crc64Iso [256]uint64

func init() {
	crc64Ecma = makeTable(ecmaPoly)
	crc64Iso = makeTable(isoPoly)
}

func makeTable(poly uint64) [256]uint64 {
	var table [256]uint64
	var i, j int
	var r uint64
	for i = 0; i < 256; i++ {
		r = uint64(i)
		for j = 0; j < 8; j++ {
			if r&1 == 1 {
				r = (r >> 1) ^ poly
			} else {
				r >>= 1
			}
		}
		table[i] = r
	}
	return table
}

func Crc64Ecma(payload []byte) uint64 {
	var result = JBCommon.MaxUint64
	for _, b := range payload {
		result = (result >> 8) ^ crc64Ecma[byte(result)^b]
	}
	return ^result
}

func Crc64Iso(payload []byte) uint64 {
	var result = JBCommon.MaxUint64
	for _, b := range payload {
		result = (result >> 8) ^ crc64Iso[byte(result)^b]
	}
	return ^result
}
