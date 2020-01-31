package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var th = []int{0, 1, 2, 3, 4, 5}

func SingleHash(in, out chan interface{}) {
	for i := range in {
		data := fmt.Sprintf("%v", i)
		fmt.Println("Incomming data is ----", data)
		result := DataSignerCrc32(data) + "~" + DataSignerCrc32(DataSignerMd5(data))
		fmt.Println("Result is --", result)
		out <- result
	}
}

func MultiHash(in, out chan interface{}) {
	data := <-in
	res := ""
	for n := range th {
		str_th := strconv.FormatInt(int64(n), 10)
		result := DataSignerCrc32(str_th + data.(string))
		res = res + result
	}
	out <- res
}

func CombineResults(in, out chan interface{}) {
	results := []string{}
	for v := range in {
		fmt.Println(v)
		data := <-in
		results = append(results, data.(string))
	}
	sort.Strings(results)
	final := strings.Join(results, "_")
	out <- final
}

func main() {
	fmt.Println("Started...")
}

func ExecutePipeline(jobs ...job) {
	in := make(chan interface{})
	fmt.Println(jobs)
	for _, job := range jobs {
		out := make(chan interface{})
		fmt.Println("Starting job --- ", job)
		fmt.Println("In --", in, "Out --", out)
		go job(in, out)
		in = out
	}
}
