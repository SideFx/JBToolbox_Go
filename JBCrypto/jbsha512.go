//----------------------------------------------------------------------------------------------------------
// Sha512 Go implementation based on Wikipedia's pseudo code
// https://en.wikipedia.org/wiki/SHA-2
// (w) 2025 by Jan Buchholz
//----------------------------------------------------------------------------------------------------------

package JBCrypto

import "JBToolbox/JBCommon"

const sha512Rounds uint64 = 80
const sha512BlockSize uint64 = 128
const sha512BitCountSize uint64 = 16

var k512Array = [...]uint64{
	0x428a2f98d728ae22, 0x7137449123ef65cd, 0xb5c0fbcfec4d3b2f, 0xe9b5dba58189dbbc,
	0x3956c25bf348b538, 0x59f111f1b605d019, 0x923f82a4af194f9b, 0xab1c5ed5da6d8118,
	0xd807aa98a3030242, 0x12835b0145706fbe, 0x243185be4ee4b28c, 0x550c7dc3d5ffb4e2,
	0x72be5d74f27b896f, 0x80deb1fe3b1696b1, 0x9bdc06a725c71235, 0xc19bf174cf692694,
	0xe49b69c19ef14ad2, 0xefbe4786384f25e3, 0x0fc19dc68b8cd5b5, 0x240ca1cc77ac9c65,
	0x2de92c6f592b0275, 0x4a7484aa6ea6e483, 0x5cb0a9dcbd41fbd4, 0x76f988da831153b5,
	0x983e5152ee66dfab, 0xa831c66d2db43210, 0xb00327c898fb213f, 0xbf597fc7beef0ee4,
	0xc6e00bf33da88fc2, 0xd5a79147930aa725, 0x06ca6351e003826f, 0x142929670a0e6e70,
	0x27b70a8546d22ffc, 0x2e1b21385c26c926, 0x4d2c6dfc5ac42aed, 0x53380d139d95b3df,
	0x650a73548baf63de, 0x766a0abb3c77b2a8, 0x81c2c92e47edaee6, 0x92722c851482353b,
	0xa2bfe8a14cf10364, 0xa81a664bbc423001, 0xc24b8b70d0f89791, 0xc76c51a30654be30,
	0xd192e819d6ef5218, 0xd69906245565a910, 0xf40e35855771202a, 0x106aa07032bbd1b8,
	0x19a4c116b8d2d0c8, 0x1e376c085141ab53, 0x2748774cdf8eeb99, 0x34b0bcb5e19b48a8,
	0x391c0cb3c5c95a63, 0x4ed8aa4ae3418acb, 0x5b9cca4f7763e373, 0x682e6ff3d6b2b8a3,
	0x748f82ee5defb2fc, 0x78a5636f43172f60, 0x84c87814a1f0ab72, 0x8cc702081a6439ec,
	0x90befffa23631e28, 0xa4506cebde82bde9, 0xbef9a3f7b2c67915, 0xc67178f2e372532b,
	0xca273eceea26619c, 0xd186b8c721c0c207, 0xeada7dd6cde0eb1e, 0xf57d4f7fee6ed178,
	0x06f067aa72176fba, 0x0a637dc5a2c898a6, 0x113f9804bef90dae, 0x1b710b35131c471b,
	0x28db77f523047d84, 0x32caab7b40c72493, 0x3c9ebe0a15c9bebc, 0x431d67c49c100d4c,
	0x4cc5d4becb3e42b6, 0x597f299cfc657e2a, 0x5fcb6fab3ad6faec, 0x6c44198c4a475817,
}

func ComputeSha512(payload []byte) []byte {
	var result []byte = nil
	var l = uint64(len(payload))
	var ha, hb, hc, hd, he, hf, hg, hh, sl, sh, sc, sm, t0, t1, z, n, i, j uint64
	var hArray = [...]uint64{
		0x6a09e667f3bcc908, 0xbb67ae8584caa73b, 0x3c6ef372fe94f82b, 0xa54ff53a5f1d36f1,
		0x510e527fade682d1, 0x9b05688c2b3e6c1f, 0x1f83d9abfb41bd6b, 0x5be0cd19137e2179,
	}
	var dArray [sha512Rounds]uint64
	var sd = l/sha512BlockSize + 1
	if (l % sha512BlockSize) >= (sha512BlockSize - sha512BitCountSize) {
		sd++
	}
	sl, sh = l<<3, l>>61
	var mc = make([]byte, l+1)
	copy(mc, payload)
	mc[l] = 0x80
	for i = l + 1; i < sd*sha512BlockSize-sha512BitCountSize; i++ {
		mc = append(mc, 0x00)
	}
	mc = append(mc, byte(sh>>56))
	mc = append(mc, byte(sh>>48))
	mc = append(mc, byte(sh>>40))
	mc = append(mc, byte(sh>>32))
	mc = append(mc, byte(sh>>24))
	mc = append(mc, byte(sh>>16))
	mc = append(mc, byte(sh>>8))
	mc = append(mc, byte(sh))
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
		for i = 0; i < 16; i++ {
			dArray[i] = uint64(mc[z])<<56 | uint64(mc[z+1])<<48 | uint64(mc[z+2])<<40 | uint64(mc[z+3])<<32 |
				uint64(mc[z+4])<<24 | uint64(mc[z+5])<<16 | uint64(mc[z+6])<<8 | uint64(mc[z+7])
			z += 8
		}
		for i = 16; i < sha512Rounds; i++ {
			sl = JBCommon.Ror64(dArray[i-15], 1) ^ JBCommon.Ror64(dArray[i-15], 8) ^ (dArray[i-15] >> 7)
			sh = JBCommon.Ror64(dArray[i-2], 19) ^ JBCommon.Ror64(dArray[i-2], 61) ^ (dArray[i-2] >> 6)
			dArray[i] = dArray[i-16] + sl + dArray[i-7] + sh
		}
		for i = 0; i < sha512Rounds; i++ {
			sh = JBCommon.Ror64(he, 14) ^ JBCommon.Ror64(he, 18) ^ JBCommon.Ror64(he, 41)
			sc = (he & hf) ^ ((^he) & hg)
			t0 = hh + sh + sc + k512Array[i] + dArray[i]
			sl = JBCommon.Ror64(ha, 28) ^ JBCommon.Ror64(ha, 34) ^ JBCommon.Ror64(ha, 39)
			sm = (ha & hb) ^ (ha & hc) ^ (hb & hc)
			t1 = sl + sm
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
	for j = 0; j < uint64(len(hArray)); j++ {
		result = append(result, byte(hArray[j]>>56))
		result = append(result, byte(hArray[j]>>48))
		result = append(result, byte(hArray[j]>>40))
		result = append(result, byte(hArray[j]>>32))
		result = append(result, byte(hArray[j]>>24))
		result = append(result, byte(hArray[j]>>16))
		result = append(result, byte(hArray[j]>>8))
		result = append(result, byte(hArray[j]))
	}
	return result
}
