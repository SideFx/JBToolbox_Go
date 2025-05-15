//----------------------------------------------------------------------------------------------------------
// CRC64 self test
// Results were checked against https://toolkitbay.com/tkb/tool/CRC-64
// (w) 2025 by Jan Buchholz
//----------------------------------------------------------------------------------------------------------

package JBSelftest

import (
	"JBToolbox/JBCommon"
	"JBToolbox/JBCrypto"
)

func Crc64EcmaSelftest(inp string, crc uint64) {
	println("INPUT text    : ", inp)
	println("Expected Crc64: ", JBCommon.Uint64ToHexString(crc))
	res := JBCrypto.Crc64Ecma([]byte(inp))
	println("Computed Crc64: ", JBCommon.Uint64ToHexString(res))
}

func Crc64IsoSelftest(inp string, crc uint64) {
	println("INPUT text    : ", inp)
	println("Expected Crc64: ", JBCommon.Uint64ToHexString(crc))
	res := JBCrypto.Crc64Iso([]byte(inp))
	println("Computed Crc64: ", JBCommon.Uint64ToHexString(res))
}

func JBCrc64Selftest() {
	println("\n------------ Crc64 ECMA-182 Selftest ------------")
	Crc64EcmaSelftest(JBCommon.Text1, 0x55fbc0cb189f41a9)
	Crc64EcmaSelftest("123456789", 0x995dc9bbdf1939fa)
	Crc64EcmaSelftest(JBCommon.Text3, 0x230E6ABE5DA53740)
	println("\n--------------- Crc64 ISO Selftest --------------")
	Crc64IsoSelftest(JBCommon.Text1, 0x25ef5aea3cba058b)
	Crc64IsoSelftest("123456789", 0xb90956c775a41001)
	Crc64IsoSelftest(JBCommon.Text3, 0x675BAF1A4B03CCF9)
}
