package chapter01

// URLify replace space by %20
func URLify(runes []rune, trueLength int) []rune {
	spaceCount := int(0)
	for i := 0; i < trueLength; i++ {
		if runes[i] == ' ' {
			spaceCount++
		}
	}
	index := trueLength + spaceCount*2
	for i := trueLength - 1; i >= 0; i-- {
		if runes[i] == ' ' {
			runes[index-1] = '0'
			runes[index-2] = '2'
			runes[index-3] = '%'
			index -= 3
		} else {
			runes[index-1] = runes[i]
			index--
		}
	}
	return runes
}
