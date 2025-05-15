//----------------------------------------------------------------------------------------------------------
// Ascii85 Go implementation based on Wikipedia's web site
// https://en.wikipedia.org/wiki/Ascii85
// (w) 2025 by Jan Buchholz
//----------------------------------------------------------------------------------------------------------

package JBCodec

const cZero byte = 'z'
const cBase uint32 = 33

func Ascii85Encode(payload []byte) []byte {
	var encodedData []byte = nil
	var number uint32 = 0
	var i = 0
	var j int
	var pad = 0
	var enc [5]byte
	var dataLen = len(payload)
	var c byte
	for i < dataLen {
		for j = 0; j < 4; j++ {
			if i < dataLen {
				c = payload[i]
				i++
			} else {
				c = 0x00
				pad++
			}
			number = (number << 8) | uint32(c)
		}
		if number > 0 {
			enc[4] = byte((number % 85) + cBase)
			number /= 85
			enc[3] = byte((number % 85) + cBase)
			number /= 85
			enc[2] = byte((number % 85) + cBase)
			number /= 85
			enc[1] = byte((number % 85) + cBase)
			number /= 85
			enc[0] = byte((number % 85) + cBase)
			for j = 0; j < (4-pad)+1; j++ {
				encodedData = append(encodedData, enc[j])
			}
		} else {
			encodedData = append(encodedData, cZero)
		}
	}
	return encodedData
}

func Ascii85Decode(payload []byte) []byte {
	var decodedData []byte = nil
	var number uint32 = 0
	var i = 0
	var j, cnt int
	var dataLen = len(payload)
	for i < dataLen {
		if payload[i] == cZero {
			for j = 0; j < 4; j++ {
				decodedData = append(decodedData, 0x00)
			}
			i++
		} else {
			number = 0
			cnt = 0
			for j = 0; j < 5; j++ {
				if i < dataLen {
					if (payload[i] < 33) || (payload[i] > 117) {
						break
					}
					number = number*85 + (uint32(payload[i]) - cBase)
					i++
					cnt++
				} else {
					number = number*85 + 84
				}
			}
			if cnt >= 2 {
				decodedData = append(decodedData, byte(number>>24))
			}
			if cnt >= 3 {
				decodedData = append(decodedData, byte(number>>16))
			}
			if cnt >= 4 {
				decodedData = append(decodedData, byte(number>>8))
			}
			if cnt == 5 {
				decodedData = append(decodedData, byte(number))
			}
		}
	}
	return decodedData
}
