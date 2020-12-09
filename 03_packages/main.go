package main

import (
	"fmt"
	"github.com/mattmc96/go_crash_course/03_packages/strutil"
	. "math"
)

func main() {
	fmt.Println(Floor(2.7))
	fmt.Println(Ceil(2.7))
	fmt.Println(Sqrt(16))
	fmt.Println(strutil.Reverse("What the hell man"))

}
