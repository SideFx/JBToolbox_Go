//----------------------------------------------------------------------------------------------------------
// Base64 self test, (w) 2025 by Jan Buchholz
//----------------------------------------------------------------------------------------------------------

package JBSelftest

import (
	"JBToolbox/JBCodec"
	"JBToolbox/JBCommon"
)

func JBBase64Selftest() {
	println("\n-------------- Base64 Selftest --------------")
	println("INPUT text  : ", JBCommon.Text3)
	base64enc := JBCodec.Base64Encode([]byte(JBCommon.Text3))
	println("Encoded text: ", base64enc)
	base64dec := JBCodec.Base64Decode(base64enc)
	println("Decoded text: ", string(base64dec))
}
