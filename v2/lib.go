// package workspace_lib is an exercise from [Learning Go] on [Oreilly]. The book can also be found on [Amazon].
// This example is from [chapter 10] exercises 1-3.
//
// [Learning Go]: https://learning.oreilly.com/library/view/learning-go-2nd/9781098139285/
// [Oreilly]: https://www.oreilly.com/
// [Amazon]: https://www.amazon.com/Learning-Go-Idiomatic-Real-World-Programming/dp/1492077216
// [chapter 10]: https://learning.oreilly.com/library/view/learning-go-2nd/9781098139285/ch10.html#id351
package workspace_lib

//
// Exercise 1:
// Create a module in your own public repository. This module has a single function named Add with two int parameters and one int return value.
// This function adds the two parameters together and returns them. Make this version v1.0.0.
//
// Exercise 2:
// Add godoc comments to your module that describe the package and the Add function.
// Be sure to include a link to https://www.mathsisfun.com/numbers/addition.html in your Add function godoc. Make this version v1.0.1.
//
// Exercise 3:
// Change Add to make it generic. Import the golang.org/x/exp/constraints package. Combine the Integer and Float types in that package to create an interface called Number.
// Rewrite Add to take in two parameters of type Number and return a value of type Number. Version your module again.
// Because this is a backward-breaking change, this should be v2.0.0 of your module.

// A Number this allows the user to call the [Add].
type Number interface {
	~int | ~float32 | ~float64
}

// Add takes two [Number] and adds the to together. To explain how addition works please take a look at [Addition].
//
// [Addition]: https://www.mathsisfun.com/numbers/addition.html
func Add[N Number](var1, var2 N) N {
	return var1 + var2
}

// Dang this is kinda amazing!
