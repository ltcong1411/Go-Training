package demo_string

import (
	"fmt"
	"strings"
)

func Demo1() {
	var ch rune = 'a'
	fmt.Printf("%d\n", ch)
	fmt.Printf("%c\n", ch)
}

func Demo2() {
	s := "abc"
	fmt.Println("len: ", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("\t%d:%c", s[i], s[i])
	}
}

func Demo3() {
	s := "   abc   "
	fmt.Println("len: ", len(s))
	s = strings.TrimSpace(s)
	fmt.Println("len: ", len(s))
}

func Demo4() {
	s := "1254abc889efd45"
	fmt.Println(s)
	result := strings.TrimFunc(s, func(c rune) bool {
		return c >= 48 && c <= 57
	})
	fmt.Println(result)
}

func Demo5() {
	s1 := "abc"
	s1 = strings.ToUpper(s1)
	fmt.Println(s1)

	s2 := "ABC"
	s2 = strings.ToLower(s2)
	fmt.Println(s2)
}

func Demo6() {
	s1 := "cong"
	keyword := "o"
	result1 := strings.HasPrefix(s1, keyword)
	fmt.Println("result 1: ", result1)

	result2 := strings.HasSuffix(s1, keyword)
	fmt.Println("result 2: ", result2)

	result3 := strings.Contains(s1, keyword)
	fmt.Println("result 2: ", result3)
}

func Demo7() {
	name1 := "Nghia"
	name2 := "Cong"
	result := strings.Compare(name1, name2)
	fmt.Println(result) // khac => 1, giong => 0
}

func SearchKeyWord(keyword string, ds [5]string) []string {
	a := []string{}
	for _, v := range ds {
		if strings.Contains(strings.ToLower(v), strings.ToLower(keyword)) {
			a = append(a, v)
		}
	}
	return a
}
