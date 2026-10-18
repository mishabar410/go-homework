package unicode_name

import "unicode"

func AnalyzeName(name string) (byteCount, runeCount, latinCount, cyrillicCount int, positions []int) {
	byteCount = len(name)
	positions = []int{}

	for i, r := range name {
		runeCount++
		positions = append(positions, i)

		if unicode.IsLetter(r) && unicode.Is(unicode.Latin, r) {
			latinCount++
		} else if unicode.IsLetter(r) && unicode.Is(unicode.Cyrillic, r) {
			cyrillicCount++
		}
	}

	return byteCount, runeCount, latinCount, cyrillicCount, positions
}
