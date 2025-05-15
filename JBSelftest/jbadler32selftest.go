//----------------------------------------------------------------------------------------------------------
// Adler32 self test
// Some test vectors were taken from
// https://mines.lumpylumpy.com/Electronics/Computers/Software/Cpp/Lib/Adler32.h
// (w) 2025 by Jan Buchholz
//----------------------------------------------------------------------------------------------------------

package JBSelftest

import (
	"JBToolbox/JBCommon"
	"JBToolbox/JBCrypto"
)

func Adler32Selftest(inp []byte, adler32 uint32) {
	println("INPUT hex       : ", JBCommon.ByteArrayToHexString(inp))
	println("Expected Adler32: ", JBCommon.Uint32ToHexString(adler32))
	res := JBCrypto.Adler32(inp)
	println("Computed Adler32: ", JBCommon.Uint32ToHexString(res))
}

func JBAdler32Selftest() {
	println("\n-------------- Adler32 Selftest --------------")
	Adler32Selftest([]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07}, 0x005c001d)
	Adler32Selftest([]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}, 0x02b80079)
	Adler32Selftest([]byte{0x41, 0x41, 0x41, 0x41}, 0x028e0105)
	Adler32Selftest([]byte{0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42}, 0x09500211)
	Adler32Selftest([]byte{0x43, 0x43, 0x43, 0x43, 0x43, 0x43, 0x43, 0x43, 0x43, 0x43, 0x43, 0x43, 0x43, 0x43, 0x43, 0x43}, 0x23a80431)
	Adler32Selftest([]byte("Hello"), 0x058c01f5)
	Adler32Selftest([]byte("Neon"), 0x03b70191)
	Adler32Selftest([]byte(JBCommon.Text1), 0x88c0622d)
}
