package ransomenote

import (
	"fmt"
	"strings"
)

func Run() {
	ransomNote := "aa"
	magazine := "aab"

	fmt.Printf("ransomNote: %s\nmagazine: %s\ncanConstruct: %t\n", ransomNote, magazine, canConstruct(ransomNote, magazine))
}

func canConstruct(ransomNote string, magazine string) bool {
	for i := 0; i < len(ransomNote); i++ {
		if !strings.Contains(magazine, string(ransomNote[i])) {
			return false
		}
		magazine = strings.Replace(magazine, string(ransomNote[i]), "", 1)
	}
	return true
}
