package main

import "fmt"

func main() {
	var batman string = "I am batman"
	fmt.Println(batman)

	var superman string = "I am superman"
	fmt.Println(superman)

	thor := "Iam thor"
	fmt.Println(thor)

	// thorRating := 3.0 //For dynamic variables
	// fmt.Println("%v, %T", thorRating, thorRating)
	// var Ironman, CaptAmerica string = "Iam Ironman", "I am capt America"

	var (
		spiderman = "Iam spiderman"
		age       = 18
		power     = "web slings,spidey sense"
		antman    = "Hai iam antman"
	)

	fmt.Println(spiderman, age, antman, power)

}
