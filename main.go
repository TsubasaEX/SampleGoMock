package main

//go:generate mockery --all

import (
	"fmt"
	"greeting"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
)

func main() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL + "/test/fdsfdse")
	if err != nil {
		log.Fatal(err)
	}
	resp, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", resp)
	fmt.Printf("%s", resp)
	fmt.Printf("%s", resp)
	fmt.Printf("%s", resp)
	fmt.Printf("%s", resp)

	fmt.Println(greeting.Hello("Zach"))
}
