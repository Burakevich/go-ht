package main

// Specify Filter function here


func Filter(s []int, fn func(int, int) bool) []int {
	var p []int 
	for _, v := range s {
		if fn(v, 1) {
			p = append(p, v)
		}
	}
	return p
}

func main() {

}
