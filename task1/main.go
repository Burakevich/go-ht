package main

// Specify Filter function here

func Filter(s []int, fn func(int, int) bool) []int {
	var p []int
	for i, v := range s {
		if fn(v, i) {
			p = append(p, v)
		}
	}
	return p
}

func main() {

}
