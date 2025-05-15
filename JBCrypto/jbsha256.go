//----------------------------------------------------------------------------------------------------------
// Sha256 Go implementation based on Wikipedia's pseudo code
// https://en.wikipedia.org/wiki/SHA-2
// (w) 2025 by Jan Buchholz
//----------------------------------------------------------------------------------------------------------

package JBCrypto

import "JBToolbox/JBCommon"

const sha256Rounds uint32 = 64
const sha256BlockSize uint32 = 64
const sha256BitCountSize uint32 = 8

var k256Array = [...]uint32{
	0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
	0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
	0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
	0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
	0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
	0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
	0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
	0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2,
}

func ComputeSha256(payload []byte) []byte {
	var result []byte = nil
	var ha, hb, hc, hd, he, hf, hg, hh, t0, t1, i, j, n, z uint32
	var l = uint32(len(payload))
	var sd = l/sha256BlockSize + 1
	if (l % sha256BlockSize) >= (sha256BlockSize - sha256BitCountSize) {
		sd++
	}
	var sl = uint64(l) << 3
	var hArray = [...]uint32{
		0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a, 0x510e527f, 0x9b05688c, 0x1f83d9ab, 0x5be0cd19,
	}
	var dArray [sha256Rounds]uint32
	var mc = make([]byte, l+1)
	copy(mc, payload)
	mc[l] = 0x80
	for i = l + 1; i < sd*sha256BlockSize-sha256BitCountSize; i++ {
		mc = append(mc, 0x00)
	}
	mc = append(mc, byte(sl>>56))
	mc = append(mc, byte(sl>>48))
	mc = append(mc, byte(sl>>40))
	mc = append(mc, byte(sl>>32))
	mc = append(mc, byte(sl>>24))
	mc = append(mc, byte(sl>>16))
	mc = append(mc, byte(sl>>8))
	mc = append(mc, byte(sl))
	z = 0
	for n = 0; n < sd; n++ {
		ha = hArray[0]
		hb = hArray[1]
		hc = hArray[2]
		hd = hArray[3]
		he = hArray[4]
		hf = hArray[5]
		hg = hArray[6]
		hh = hArray[7]
		for j = 0; j < 16; j++ {
			dArray[j] = (uint32(mc[z]) << 24) | (uint32(mc[z+1]) << 16) | (uint32(mc[z+2]) << 8) | uint32(mc[z+3])
			z += 4
		}
		for j = 16; j < sha256Rounds; j++ {
			dArray[j] = dArray[j-16] + (JBCommon.Ror32(dArray[j-15], 7) ^ JBCommon.Ror32(dArray[j-15], 18) ^
				(dArray[j-15] >> 3)) + dArray[j-7] + (JBCommon.Ror32(dArray[j-2], 17) ^
				JBCommon.Ror32(dArray[j-2], 19) ^ (dArray[j-2] >> 10))
		}
		for j = 0; j < sha256Rounds; j++ {
			t0 = hh + k256Array[j] + dArray[j] + ((he & hf) ^ ((^he) & hg)) +
				(JBCommon.Ror32(he, 6) ^ JBCommon.Ror32(he, 11) ^ JBCommon.Ror32(he, 25))
			t1 = (JBCommon.Ror32(ha, 2) ^ JBCommon.Ror32(ha, 13) ^ JBCommon.Ror32(ha, 22)) +
				((ha & hb) ^ (ha & hc) ^ (hb & hc))
			hh = hg
			hg = hf
			hf = he
			he = hd + t0
			hd = hc
			hc = hb
			hb = ha
			ha = t0 + t1
		}
		hArray[0] += ha
		hArray[1] += hb
		hArray[2] += hc
		hArray[3] += hd
		hArray[4] += he
		hArray[5] += hf
		hArray[6] += hg
		hArray[7] += hh
	}
	for j = 0; j < uint32(len(hArray)); j++ {
		result = append(result, byte(hArray[j]>>24))
		result = append(result, byte(hArray[j]>>16))
		result = append(result, byte(hArray[j]>>8))
		result = append(result, byte(hArray[j]))
	}
	return result
}
