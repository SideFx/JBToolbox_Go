//----------------------------------------------------------------------------------------------------------
// Blowfish speed test, (w) 2025 by Jan Buchholz
//----------------------------------------------------------------------------------------------------------

package JBSelftest

import (
	"JBToolbox/JBCommon"
	"JBToolbox/JBCrypto"
	"time"
)

func JBBlowfishECBSpeedtest() {
	const rounds = 1000
	const key = "Batteries not included"
	words := JBCommon.WordsInString(JBCommon.Text2)
	println()
	println(rounds, "x Blowfish ECB encryption/decryption of a plain text containing",
		words, "words:")
	start := time.Now()
	bfish := JBCrypto.BlowfishFactory([]byte(key))
	for i := 0; i < rounds; i++ {
		enc := bfish.EncryptECB([]byte(JBCommon.Text2))
		_ = bfish.DecryptECB(enc)
	}
	elapsed := time.Since(start)
	println("Runtime measured: ", elapsed.Microseconds(), "microseconds")
}

func JBBlowfishCBCSpeedtest() {
	const rounds = 1000
	const key = "Batteries not included"
	ivec := []byte{0x01, 0xaa, 0xff, 0x55, 0x23, 0x91, 0x45, 0xee}
	words := JBCommon.WordsInString(JBCommon.Text2)
	println()
	println(rounds, "x Blowfish CBC encryption/decryption of a plain text containing",
		words, "words:")
	start := time.Now()
	bfish := JBCrypto.BlowfishFactory([]byte(key))
	for i := 0; i < rounds; i++ {
		enc := bfish.EncryptCBC([]byte(JBCommon.Text2), ivec)
		_ = bfish.DecryptCBC(enc, ivec)
	}
	elapsed := time.Since(start)
	println("Runtime measured: ", elapsed.Microseconds(), "microseconds")
}
