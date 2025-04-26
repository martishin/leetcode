package isomorphic_strings_205

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapping := make(map[byte]byte)
	reverse := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		sc := s[i]
		tc := t[i]

		mc, mok := mapping[sc]
		rc, rok := reverse[tc]

		if !mok && !rok {
			mapping[sc] = tc
			reverse[tc] = sc
		} else if !((mok && mc == tc) && (rok && rc == sc)) {
			return false
		}
	}

	return true
}

func Test() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
}
