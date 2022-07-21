package main

func Iftrue(n int) {
	var a []string
	for i := 0; i < n; i++ {
		if true {
			a = append(a, "test")
		} else {
		}
	}
}

func Iffalse(n int) {
	var a []string
	for i := 0; i < n; i++ {
		if false {
		} else {
			a = append(a, "test")
		}
	}
}
