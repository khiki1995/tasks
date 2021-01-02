package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func main() {
	candidates := make(map[string]int)
	var voters []string
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 4096)
	read, err := file.Read(buf)
	if err != nil {
		log.Printf("Error in reading file! error = %v", err)
	}

	data := string(buf[:read])
	arrTxt := strings.Split(data, "\n")

	for _, str := range arrTxt {
		arrRow := strings.Split(str, ",")
		if stringInSlice(arrRow[0], voters) {
			fmt.Printf("voter: %v is already voted\n", arrRow[0])
			continue
		}
		voters = append(voters, arrRow[0])
		candidates[arrRow[1]] = candidates[arrRow[1]] + 1
	}

	p := make(PairList, len(candidates))

	i := 0
	for k, v := range candidates {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(p))
	for _, k := range p {
		fmt.Printf("%v\t%v\n", k.Key, k.Value)
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
