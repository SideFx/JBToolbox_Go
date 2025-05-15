//----------------------------------------------------------------------------------------------------------
// Twofish speed test, (w) 2025 by Jan Buchholz
//----------------------------------------------------------------------------------------------------------

package JBSelftest

import (
	"JBToolbox/JBCommon"
	"JBToolbox/JBCrypto"
	"time"
)

func JBTwofishECBSpeedtest() {
	const rounds = 1000
	const key = "Batteries not included"
	words := JBCommon.WordsInString(JBCommon.Text2)
	println()
	println(rounds, "x Twofish ECB encryption/decryption of a plain text containing",
		words, "words:")
	start := time.Now()
	tfish := JBCrypto.TwofishFactory([]byte(key))
	for i := 0; i < rounds; i++ {
		enc := tfish.EncryptECB([]byte(JBCommon.Text2))
		_ = tfish.DecryptECB(enc)
	}
	elapsed := time.Since(start)
	println("Runtime measured: ", elapsed.Microseconds(), "microseconds")
}

func JBTwofishCBCSpeedtest() {
	const rounds = 1000
	const key = "Batteries not included"
	ivec := []byte{0x01, 0xaa, 0xff, 0x55, 0x23, 0x91, 0x45, 0xee}
	words := JBCommon.WordsInString(JBCommon.Text2)
	println()
	println(rounds, "x Twofish CBC encryption/decryption of a plain text containing",
		words, "words:")
	start := time.Now()
	tfish := JBCrypto.TwofishFactory([]byte(key))
	for i := 0; i < rounds; i++ {
		enc := tfish.EncryptCBC([]byte(JBCommon.Text2), ivec)
		_ = tfish.DecryptCBC(enc, ivec)
	}
	elapsed := time.Since(start)
	println("Runtime measured: ", elapsed.Microseconds(), "microseconds")
}
