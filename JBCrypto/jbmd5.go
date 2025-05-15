//----------------------------------------------------------------------------------------------------------
// MD5 Go port based on Wikipedia's pseudo code
// https://en.wikipedia.org/wiki/MD5
// (w) 2025 by Jan Buchholz
//----------------------------------------------------------------------------------------------------------

package JBCrypto

import "JBToolbox/JBCommon"

const md5Rounds uint32 = 64
const md5BlockSize uint32 = 64
const md5BitCountSize uint32 = 8

var sArray = []int{
	7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22,
	5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20,
	4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23,
	6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21,
}

var kArray = [...]uint32{
	0xd76aa478, 0xe8c7b756, 0x242070db, 0xc1bdceee, 0xf57c0faf, 0x4787c62a, 0xa8304613, 0xfd469501,
	0x698098d8, 0x8b44f7af, 0xffff5bb1, 0x895cd7be, 0x6b901122, 0xfd987193, 0xa679438e, 0x49b40821,
	0xf61e2562, 0xc040b340, 0x265e5a51, 0xe9b6c7aa, 0xd62f105d, 0x02441453, 0xd8a1e681, 0xe7d3fbc8,
	0x21e1cde6, 0xc33707d6, 0xf4d50d87, 0x455a14ed, 0xa9e3e905, 0xfcefa3f8, 0x676f02d9, 0x8d2a4c8a,
	0xfffa3942, 0x8771f681, 0x6d9d6122, 0xfde5380c, 0xa4beea44, 0x4bdecfa9, 0xf6bb4b60, 0xbebfbc70,
	0x289b7ec6, 0xeaa127fa, 0xd4ef3085, 0x04881d05, 0xd9d4d039, 0xe6db99e5, 0x1fa27cf8, 0xc4ac5665,
	0xf4292244, 0x432aff97, 0xab9423a7, 0xfc93a039, 0x655b59c3, 0x8f0ccc92, 0xffeff47d, 0x85845dd1,
	0x6fa87e4f, 0xfe2ce6e0, 0xa3014314, 0x4e0811a1, 0xf7537e82, 0xbd3af235, 0x2ad7d2bb, 0xeb86d391,
}

func ComputeMd5(payload []byte) []byte {
	var result []byte = nil
	var chunks [16]uint32
	var rArray = []uint32{0x67452301, 0xefcdab89, 0x98badcfe, 0x10325476}
	var a, b, c, d, i, k, g, f uint32
	var l = uint32(len(payload))
	var bits = uint64(l) << 3
	var mc = make([]byte, l+1)
	copy(mc, payload)
	mc[l] = 0x80
	for uint32(len(mc))%md5BlockSize != (md5BlockSize - md5BitCountSize) {
		mc = append(mc, 0)
	}
	for i = 0; i < md5BitCountSize; i++ {
		mc = append(mc, byte(bits>>(i<<3)))
	}
	l = uint32(len(mc))
	k = 0
	for k < l {
		a = rArray[0]
		b = rArray[1]
		c = rArray[2]
		d = rArray[3]
		for i = 0; i < 16; i++ {
			chunks[i] = uint32(mc[k+3])<<24 | uint32(mc[k+2])<<16 | uint32(mc[k+1])<<8 | uint32(mc[k])
			k += 4
		}
		for i = 0; i < md5Rounds; i++ {
			switch i >> 4 {
			case 0:
				f = d ^ (b & (c ^ d))
				g = i % 16
				f = f + a + kArray[i] + chunks[g]
				a = d
				d = c
				c = b
				b = b + JBCommon.Rol32(f, sArray[i])
			case 1:
				f = c ^ (d & (b ^ c))
				g = (5*i + 1) % 16
				f = f + a + kArray[i] + chunks[g]
				a = d
				d = c
				c = b
				b = b + JBCommon.Rol32(f, sArray[i])
			case 2:
				f = b ^ c ^ d
				g = (3*i + 5) % 16
				f = f + a + kArray[i] + chunks[g]
				a = d
				d = c
				c = b
				b = b + JBCommon.Rol32(f, sArray[i])
			case 3:
				f = c ^ (b | (^d))
				g = (7 * i) % 16
				f = f + a + kArray[i] + chunks[g]
				a = d
				d = c
				c = b
				b = b + JBCommon.Rol32(f, sArray[i])
			}
		}
		rArray[0] += a
		rArray[1] += b
		rArray[2] += c
		rArray[3] += d
	}
	for i = 0; i < 4; i++ {
		result = append(result, byte(rArray[i]))
		result = append(result, byte(rArray[i]>>8))
		result = append(result, byte(rArray[i]>>16))
		result = append(result, byte(rArray[i]>>24))
	}
	return result
}
