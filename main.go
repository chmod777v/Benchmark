package main

import "fmt"

func copy1(in []string) []string {
	var out []string
	for _, s := range in {
		out = append(out, s)
	}
	return out
}
func copy2(in []string) []string {
	out := make([]string, len(in))
	for i, s := range in {
		out[i] = s
	}
	return out
}
func copy3(in []string) []string {
	out := make([]string, len(in))
	copy(out, in)
	return out
}

func main() {
	input := []string{"1", "2", "3"}
	out := copy2(input)
	fmt.Println(out)
}
