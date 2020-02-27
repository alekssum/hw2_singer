package main

import (
	"crypto/md5"
	"fmt"
	"hash/crc32"
	"io"
	"strings"
)

// ExecutePipeline is a function
// which solve my task
// and this comment is nessesary for this func as demands the IDE
func ExecutePipeline(freeFlowJobs ...job) {

	test := []byte("some text")
	b := md5.Sum(test)

	fmt.Printf("%T\t%v\n", b, b)

	fmt.Println(freeFlowJobs)
}

// SingleHash считает значение crc32(data)+"~"+crc32(md5(data))
// ( конкатенация двух строк через ~),
// где data - то что пришло на вход (по сути - числа из первой функции)
var SingleHash = job(func(in, out chan interface{}) {
	// read from the chanel and convert interface to string
	data := fmt.Sprintf("%s", <-in)

	//
	firstPart := crc32.NewIEEE()
	io.Copy(firstPart, strings.NewReader(data))

	secondPart := crc32.NewIEEE()
	io.Copy(secondPart, strings.NewReader(data))

	// out <- result
})

// MultiHash is a func
var MultiHash = job(func(in, out chan interface{}) {
})

// CombineResults is a func
var CombineResults = job(func(in, out chan interface{}) {
})
