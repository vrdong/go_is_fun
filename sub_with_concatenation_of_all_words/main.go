package main

func main() {

}

//Input: s = "barfoothefoobarman", words = ["foo","bar"]
//Output: [0,9]

//1 <= s.length <= 10^4  10000
//s consists of lower-case English letters.
//1 <= words.length <= 5000 từ
//1 <= words[i].length <= 30 từ dài 30
//words[i] consists of lower-case English letters.

0 3 9 12
3 9
0 12
func findSubstring(s string, words []string) []int {
	var result []int
	wordPoint := make([]int, len(words))
	count := 0
	for i := 0; i < len(words); i++ {
		if wordPoint[i] == 0 {
			listWord := make([]int, len(words))
			for j := i; j < len(words); j++ {
				if words[j] == words[i]{
					count += 1
					listWord = append(listWord, j)
				}
			}
			for j := 0; j < len(listWord); j ++ {

			}
		}
	}
	return result
}
