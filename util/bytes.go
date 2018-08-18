package util

func Uint64ToBytes(num uint64) []byte {
	result := make([]byte, 8)

	for i := 7; i >= 0; i-- {

		rawByte := byte(num & 255)
		result[i] = rawByte

		num >>= 8
	}

	return result
}

func BytesToUint64(rawBytes []byte) uint64 {
	var num uint64 = 0

	for i := 0; i < 8; i++ {
		num <<= 8
		num += uint64(rawBytes[i])
	}

	return num
}
