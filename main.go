package main

import (
	"fmt"
	"io/ioutil"
	"math/bits"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("need two arguments")
		return
	}
	cipher1, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	cipher2, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	ones := 0
	for i := 0; i < len(cipher1); i++ {
		ones += bits.OnesCount(uint(cipher1[i] ^ cipher2[i]))
	}
	fmt.Printf("%.2f%%\n", float64(ones)/(float64(len(cipher1))*8.0)*100)
}
