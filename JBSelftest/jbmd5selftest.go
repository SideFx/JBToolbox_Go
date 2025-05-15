//----------------------------------------------------------------------------------------------------------
// MD5 self test, (w) 2025 by Jan Buchholz
//----------------------------------------------------------------------------------------------------------

package JBSelftest

import (
	"JBToolbox/JBCommon"
	"JBToolbox/JBCrypto"
)

func Md5Selftest(inp string, md5 string) {
	tmp := inp
	if len(tmp) == 0 {
		tmp = "(empty)"
	}
	println("INPUT text  : ", tmp)
	println("Expected Md5: ", md5)
	res := JBCrypto.ComputeMd5([]byte(inp))
	println("Computed Md5: ", JBCommon.ByteArrayToHexString(res))
}

func JBMd5Selftest() {
	println("\n---------------- Md5 Selftest ----------------")
	Md5Selftest("", "d41d8cd98f00b204e9800998ecf8427e")
	Md5Selftest("a", "0cc175b9c0f1b6a831c399e269772661")
	Md5Selftest("abc", "900150983cd24fb0d6963f7d28e17f72")
	Md5Selftest("message digest", "f96b697d7cb7938d525a2f31aaf161d0")
	Md5Selftest("abcdefghijklmnopqrstuvwxyz", "c3fcd3d76192e4007dfb496cca67e13b")
	Md5Selftest("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789", "d174ab98d277d9f5a5611c2c9f419d9f")
	Md5Selftest("12345678901234567890123456789012345678901234567890123456789012345678901234567890", "57edf4a22be3c955ac49da2e2107b67a")
	Md5Selftest("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz3456", "6831fa90115bb9a54fbcd4f9fee0b5c4")
	Md5Selftest("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz345", "bc40505cc94a43b7ff3e2ac027325233")
	Md5Selftest("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz34567", "fa94b73a6f072a0239b52acacfbcf9fa")
	Md5Selftest("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz345678901234", "bd201eae17f29568927414fa326f1267")
	Md5Selftest("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz34567890123", "80063db1e6b70a2e91eac903f0e46b85")
}
