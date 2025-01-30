package trie

type trieNode struct {
	childNodes trieNodes
	isEnd      bool
}

type trieNodes map[rune]*trieNode

type Trie trieNode

func MakeTrie(values []string) Trie {
	trie := Trie(trieNode{trieNodes{}, false})
	trieTn := trieNode(trie)

	for _, val := range values {
		tn := &trieTn

		for _, char := range val {
			if nextTn, ok := tn.childNodes[char]; ok {
				tn = nextTn
			} else {
				nextTn := trieNode{trieNodes{}, false}
				tn.childNodes[char] = &nextTn
				tn = &nextTn
			}
		}

		tn.isEnd = true
	}

	return trie
}

func IsInTrie(trie Trie, val string) bool {
	tn := trieNode(trie)
	for _, char := range val {
		if nextTn, ok := tn.childNodes[char]; ok {
			tn = *nextTn
		} else {
			return false
		}
	}

	return tn.isEnd
}
