//----------------------------------------------------------------------------------------------------------
// Ascii85 self test, (w) 2025 by Jan Buchholz
//----------------------------------------------------------------------------------------------------------

package JBSelftest

import (
	"JBToolbox/JBCodec"
	"JBToolbox/JBCommon"
)

func JBAscii85Selftest() {
	println("\n-------------- Ascii85 Selftest --------------")
	println("INPUT text  : ", JBCommon.Text3)
	enc := JBCodec.Ascii85Encode([]byte(JBCommon.Text3))
	println("Encoded text: ", string(enc))
	dec := JBCodec.Ascii85Decode(enc)
	println("Decoded text: ", string(dec))
}
