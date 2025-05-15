//----------------------------------------------------------------------------------------------------------
// Twofish self test, (w) 2025 by Jan Buchholz
//----------------------------------------------------------------------------------------------------------

package JBSelftest

import (
	"JBToolbox/JBCommon"
	"JBToolbox/JBCrypto"
	"slices"
)

func TwofishSelfTest() bool {
	const twofishBlocksize uint32 = 16
	const twofishKeysize uint32 = 32
	var ct49 = [...]byte{
		0x37, 0xfe, 0x26, 0xff, 0x1c, 0xf6, 0x61, 0x75, 0xf5, 0xdd, 0xf4, 0xc3, 0x3b, 0x97, 0xa2, 0x05,
	}
	const rounds uint32 = 49
	var tfkey = make([]byte, twofishKeysize)
	var pt, ptx, ct = make([]byte, twofishBlocksize), make([]byte, twofishBlocksize), make([]byte, twofishBlocksize)
	var i, j uint32
	var twofish JBCrypto.Twofish
	for i = 0; i < twofishKeysize; i++ {
		tfkey[i] = 0
	}
	for i = 0; i < twofishBlocksize; i++ {
		pt[i] = 0
	}
	for i = 0; i < rounds; i++ {
		twofish = JBCrypto.TwofishFactory(tfkey)
		copy(ptx, pt)
		twofish.EncryptBase(pt)
		copy(ct, pt)
		twofish.DecryptBase(pt)
		if !slices.Equal(pt[:], ptx[:]) {
			return false
		}
		for j = twofishKeysize / 2; j < twofishKeysize; j++ {
			tfkey[j] = tfkey[j-twofishKeysize/2]
		}
		for j = 0; j < twofishBlocksize; j++ {
			tfkey[j] = ptx[j]
		}
		copy(pt, ct)
	}
	return slices.Equal(ct[:], ct49[:])
}

func JBTwofishSelftest() {
	var res string
	if TwofishSelfTest() {
		res = "PASSED"
	} else {
		res = "FAILED"
	}
	const key = "Batteries not included"
	ivec := []byte{0x01, 0xaa, 0xff, 0x55, 0x23, 0x91, 0x45, 0xee, 0xaf, 0xfe}
	println("\n---------------- Twofish Selftest ----------------")
	println("... using 49 rounds test function: ", res)
	println("\n---------------- Twofish ECB Test ----------------")
	tfish := JBCrypto.TwofishFactory([]byte(key))
	println("Encryption key: ", key)
	println("Plain text    : ", JBCommon.Text1)
	enc := tfish.EncryptECB([]byte(JBCommon.Text1))
	println("Encrypted text: ", JBCommon.ByteArrayToHexString(enc))
	dec := tfish.DecryptECB(enc)
	println("Decrypted text: ", string(dec))
	println("\n---------------- Twofish CBC Test ----------------")
	println("Encryption key  : ", key)
	println("Init. CBC vector: ", JBCommon.ByteArrayToHexString(ivec))
	println("Plain text      : ", JBCommon.Text1)
	enc = tfish.EncryptCBC([]byte(JBCommon.Text1), ivec)
	println("Encrypted text  : ", JBCommon.ByteArrayToHexString(enc))
	dec = tfish.DecryptCBC(enc, ivec)
	println("Decrypted text  : ", string(dec))
}
