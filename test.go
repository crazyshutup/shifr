package main

import (
	_"net/http"
	"net/http"
	"log"
	"html/template"
	"strconv"
	"fmt"
)

type micro struct {
	a    int
	b    int
	c    int
	text string
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	/*info := micro{
		a: int(r.PostForm.Get("a")),
		b: int(r.PostForm.Get("b")),
		c: int(r.PostForm.Get("c")),
		text: string(r.PostForm.Get("text")),
	}*/eod, _ := strconv.Atoi(r.PostForm.Get("encryption"))
	//eod := int(r.PostForm.Get("encryption"))
	if eod == 1 {
		println("shifrovnie")
	} else {
		println("deshifrovanie")
	}
	testTemplate, err := template.ParseFiles("templates/test.html")
	//type1 := micro{5, "kek", "cheburek"}
	if err != nil {
		log.Println(err)
		return
	}
	if err := testTemplate.Execute(w, nil); err != nil {
		log.Println(err)
		return
	}

}

const N = 255

func encrypt(a, b, c int, str string) (result string) {
	for i := 0; i < len(str); i++{
		e := str[i]
		k := (a * i * i) + (b * i) + c
		result += string((int32(k)) + (int32(e)) % N)
		fmt.Println(k, string((int32(k)) + (int32(e)) % N))
	}
	return
}

func decrypt(a, b, c int, str string) (result string) {
	for i := 0; i < len(str); i++ {
		e := str[i]
		k := (a * i * i) + (b * i) + c
		for (int32(e))-(int32(k)) < 0 {
			k -= N
		}
		result += string((int32(e)) - (int32(k))%N)
		fmt.Println(i, k, (int32(e))-(int32(k))%N)
	}
	return
}
func main() {
	fmt.Println(encrypt(1, 2, 3, "Galkinpidor, fucking russian gay, you have gay, ssuka"))
	fmt.Println(decrypt(1, 2, 3, encrypt(1, 2, 3, "Galkinpidor, fucking russian gay, you have gay, ssuka")))
	fmt.Println([]byte("Галкин пидор"))
	//http.HandleFunc("/users/", testHandler)
	//err := http.ListenAndServe(":4006", nil)
	//if err != nil {
	//	panic(err)
	//}
}
