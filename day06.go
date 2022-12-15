package adventOfCode2022

func FindStartOfPacket(input string, frameSize int) int {
	for i := frameSize; i < len(input); i++ {
		frame := input[i-frameSize : i]
		if AllLettersUnique(frame) {
			return i
		}
	}
	return -1
}

func FindStartOfPacketFromFile(fileName string, frameSize int) int {
	start := -1
	ForLinesIn(fileName, func(line string) {
		start = FindStartOfPacket(line, frameSize)
	})
	return start
}
