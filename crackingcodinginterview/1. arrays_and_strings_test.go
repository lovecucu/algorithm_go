package crackingcodinginterview

import "testing"

func TestIsUnique(t *testing.T) {
	if !isUnique("abc") || isUnique("aabc") || isUnique("leetcode") {
		t.Error(`TestIsUnique failed`)
	}
}

func TestCheckPermutation(t *testing.T) {
	if !CheckPermutation("abc", "bca") || CheckPermutation("abc", "bad") || CheckPermutation("abb", "aab") {
		t.Error(`TestCheckPermutation failed`)
	}
}

func TestReplaceSpaces(t *testing.T) {
	if replaceSpaces("Mr John Smith    ", 13) != "Mr%20John%20Smith" || replaceSpaces("                ", 5) != "%20%20%20%20%20" {
		t.Error(`TestReplaceSpaces failed`)
	}
}

func TestCanPermutePalindrome(t *testing.T) {
	if !canPermutePalindrome("tactcoa") || canPermutePalindrome("aabc") {
		t.Error(`TestCanPermutePalindrome failed`)
	}
}

func TestOneEditAway(t *testing.T) {
	if !oneEditAway("pale", "ple") || oneEditAway("pales", "ple") {
		t.Error(`TestOneEditAway failed`)
	}
}

func TestCompressString(t *testing.T) {
	if compressString("aabcccccaaa") != "a2b1c5a3" || compressString("abbccd") != "abbccd" {
		t.Error(`TestcompressString failed`)
	}
}
