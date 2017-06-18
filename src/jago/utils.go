package jago


func u2b(u1s []u1) []uint8 {
	bytes := make([]uint8, len(u1s))
	for i := 0; i < len(bytes); i++ {
		bytes[i] = uint8(u1s[i])
	}
	return bytes
}

func b2u(bytes []uint8) []u1 {
	bs := make([]u1, len(bytes))
	for i := 0; i < len(bytes); i++ {
		bs[i] = u1(bytes[i])
	}
	return bs
}

func u2s(u1s []u1) string {
	return string(u2b(u1s))
}

func bytes2uint16(bytes []uint8) uint16 {
	return uint16((bytes[0] << 8) | bytes[1])
}