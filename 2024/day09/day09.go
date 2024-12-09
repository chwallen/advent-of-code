package day09

import (
	"slices"
)

type block struct {
	start  int
	length int
	id     int
}

func PartOne(lines []string, extras ...any) any {
	fileBlocks, emptyBlocks := parseBlocks(lines, false)

	emptyBlockIndex := 0
	for fileIndex, file := range slices.Backward(fileBlocks) {
		emptyBlock := emptyBlocks[emptyBlockIndex]

		if file.start < emptyBlock.start {
			break
		}

		fileBlocks[fileIndex].start = emptyBlock.start
		emptyBlocks[emptyBlockIndex].start = file.start
		emptyBlockIndex += 1
	}

	return calculateChecksum(fileBlocks)
}

func PartTwo(lines []string, extras ...any) any {
	fileBlocks, emptyBlocks := parseBlocks(lines, true)

	for fileIndex, file := range slices.Backward(fileBlocks) {
		for emptyBlockIndex, emptyBlock := range emptyBlocks {
			if emptyBlock.start >= file.start {
				break
			}
			if emptyBlock.length >= file.length {
				fileBlocks[fileIndex].start = emptyBlock.start
				emptyBlocks[emptyBlockIndex].start += file.length
				emptyBlocks[emptyBlockIndex].length -= file.length
				break
			}
		}
	}

	return calculateChecksum(fileBlocks)
}

func parseBlocks(lines []string, compact bool) (fileBlocks, emptyBlocks []block) {
	line := lines[0]
	size := (len(line) + 1) / 2
	if !compact {
		size *= 5
	}
	fileBlocks = make([]block, 0, size)
	emptyBlocks = make([]block, 0, size)

	id := 0
	start := 0
	for i, char := range line {
		length := int(char - '0')
		if i%2 == 0 {
			if compact {
				fileBlocks = append(fileBlocks, block{start, length, id})
			} else {
				for j := range length {
					fileBlocks = append(fileBlocks, block{start + j, 1, id})
				}
			}
			id += 1
		} else if compact {
			emptyBlocks = append(emptyBlocks, block{start, length, -1})
		} else {
			for j := range length {
				emptyBlocks = append(emptyBlocks, block{start + j, 1, -1})
			}
		}
		start += length
	}

	return fileBlocks, emptyBlocks
}

func calculateChecksum(fileBlocks []block) int {
	checksum := 0
	for _, fileBlock := range fileBlocks {
		for i := 0; i < fileBlock.length; i++ {
			checksum += fileBlock.id * (fileBlock.start + i)
		}
	}
	return checksum
}
