package main

import (
	"fmt"

	"github.com/pier101/leango-nomade/study_Dictionary/mydict"
)

func main() {
	//================SEARCH================
	dictionary := mydict.Dictionary{"first": "First word"}
	// definition, err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	//================ADD==================
	//word 추가
	err := dictionary.Add("hello", "Greeting")
	if err != nil {
		fmt.Println(err)
	}

	//word 검색
	definition, _ := dictionary.Search("hello")
	fmt.Println(definition)

	//word 중복 추가
	err2 := dictionary.Add("hello", "Greeting")
	if err2 != nil {
		fmt.Println(err2)
	}

	//============UPDATE/DELETE=============
	baseWord := "one"
	dictionary.Add(baseWord, "first")
	fmt.Println("-------------------")
	word1, _ := dictionary.Search(baseWord)
	fmt.Println(word1)
	dictionary.Delete(baseWord)

	// err3 := dictionary.Update(baseWord, "second")
	// if err3 != nil {
	// 	fmt.Println(err3)
	// }
	fmt.Println("-------------------")
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("key: ", baseWord, "\nvalue : ", word)
	}

}
