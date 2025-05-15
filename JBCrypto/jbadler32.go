//----------------------------------------------------------------------------------------------------------
// Adler32 Go implementation based on Wikipedia's C code
// https://en.wikipedia.org/wiki/Adler-32
// (w) 2025 by Jan Buchholz
//----------------------------------------------------------------------------------------------------------

package JBCrypto

func Adler32(payload []byte) uint32 {
	const modAdler32 uint32 = 65521
	var a uint32 = 1
	var b uint32 = 0
	var c byte
	for _, c = range payload {
		a += uint32(c)
		if a >= modAdler32 {
			a -= modAdler32
		}
		b += a
		if b >= modAdler32 {
			b -= modAdler32
		}
	}
	return (b << 16) | a
}
