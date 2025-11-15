package main

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
