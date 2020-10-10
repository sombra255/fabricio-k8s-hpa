package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	calcRaiz()
	fmt.Fprintf(w, msgCodeEducation())
}

func calcRaiz() float64 {
	x := 0.0001
	for i := 0; i <= 1000000; i++ {
		x += math.Sqrt(x)
	}
	return x
}

func msgCodeEducation() string {
	return "Code.education Rocks!"
}
