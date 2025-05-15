//----------------------------------------------------------------------------------------------------------
// JBToolbox Go edition, (w) 2025 Jan Buchholz
//----------------------------------------------------------------------------------------------------------

package main

import (
	"JBToolbox/JBSelftest"
	"fmt"
)

func main() {
	var key string
	println("\n-------------------------\nJB Toolbox Demo:\n-------------------------")
	println(`(a) Adler32 selftest`)
	println(`(b) Crc64 selftest`)
	println(`(c) Md5 selftest`)
	println(`(d) Sha256 selftest`)
	println(`(e) Sha512 selftest`)
	println(`(f) Ascii85 selftest`)
	println(`(g) Base64 selftest`)
	println(`(h) Blowfish selftest`)
	println(`(i) Twofish selftest`)
	println(`(k) Blowfish ECB speed test`)
	println(`(l) Blowfish CBC speed test`)
	println(`(m) Twofish ECB speed test`)
	println(`(n) Twofish CBC speed test`)
	println(`(x) Exit`)
	for key != "x" {
		key = ""
		print("\nPlease enter an option: ")
		_, _ = fmt.Scanln(&key)
		switch key {
		case "a":
			JBSelftest.JBAdler32Selftest()
		case "b":
			JBSelftest.JBCrc64Selftest()
		case "c":
			JBSelftest.JBMd5Selftest()
		case "d":
			JBSelftest.JBSha256Selftest()
		case "e":
			JBSelftest.JBSha512Selftest()
		case "f":
			JBSelftest.JBAscii85Selftest()
		case "g":
			JBSelftest.JBBase64Selftest()
		case "h":
			JBSelftest.JBBlowfishSelftest()
		case "i":
			JBSelftest.JBTwofishSelftest()
		case "k":
			JBSelftest.JBBlowfishECBSpeedtest()
		case "l":
			JBSelftest.JBBlowfishCBCSpeedtest()
		case "m":
			JBSelftest.JBTwofishECBSpeedtest()
		case "n":
			JBSelftest.JBTwofishCBCSpeedtest()
		case "x":
			println(`Bye.`)
		default:
			println("Invalid input (press 'x' to exit)")
		}
	}
}
