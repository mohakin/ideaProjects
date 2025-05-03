// this file should be used as a learning forum for GoLang. I will try to comment in each new section being covered here.
package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5 // immutable value is stored
	// investmentAmount, years, expectedReturnRate := 0 ,10.0, 5.5 // Not always preferred due to readability issues.
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5
	// var can be resigned or 'mutable'.
	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

// Working With Functions and Values:
// ----------------------------------------------------------------
// package = Every GO file must have a package inside of it. When writing GO code you split your code across packages. You must have at least one package
// per GO program. You can have multiple packages in one project. A single package can be split across multiple files. You can have multiple packages in one GO file.
// packages are there to help organize your code. You can have multiple files for a single package. You can use this to import and export features across your files.
// so that your individual files stay lean.
//----------------------------------------------------------------
// import = this is how you bring in/import packages.
//----------------------------------------------------------------
// func = A block of code that can be executed by calling it.
// ----------------------------------------------------------------
// print = Is a function, which can just be thought of as a command that we are calling. Print is outputting to the cmd line.
// print is a builtin command. Strings need () and "" only, no single quotes ''. Back ticks are accepted. ``.
// ----------------------------------------------------------------
// fmt = This is part of GO's standard library. GO comes with a large standard library.
// ----------------------------------------------------------------
// END.

// The Importance Of The "main" Package:
// ----------------------------------------------------------------
// You can theoretically can use any name that you want. Different packages need different names after all. Main is a special package name.
// Main is used as the main entry point of the application that we're writing. This matter because we will not always run our code as it is here.
// This is convenient during development but that's not the only way to execute the code. And it's also not how the code will be executed when written for
// production and make it available for others. Those people may not even have GO on their system. Typically, you would run go.build in your project folder.
// This will then tell GO to build an executable file so that it can run on systems that don't have Go on their system.
// ----------------------------------------------------------------

// Understanding Go Modules & Building Go Programs:
// ----------------------------------------------------------------
// We get an error in the terminal about not having a main module. One module consists of multiple packages. In many cases a Go project is a Go module.
// We could consider this app a module. We have to run a specific command to show that this is a module. In order to do this we need to run a command called
// ' go mod init '. We must give our module a name or a path where it can be found. By running 'go mod init' with a path name like 'example.com/first-app',
// it will put a go.mod file in our folder for us.
// For future references do not go straight into your terminal and start running the commands.
// Use the builtin terminal on IntelliJ for that so that everything in your folder matches what you're
// learning from. For now though just know that 'example.com/first-app' is 'app' on my end.
// This is why we needed a module and a main package. We needed this to find a main entry point for our application.
// Main is a reserved package name that lets Go know where it should start.
// ----------------------------------------------------------------

// The Main Func is Important!:
// ----------------------------------------------------------------
// The main package is not the only thing named 'main'. There is also the main func(tion). This also must be named. It is a function that must be named main.
// Go will call and execute that function and therefore that code in that function when the program starts. This is also so that Go knows where to start when
// the application starts. Go does not simply start at the top of the file like some other languages. Then execute from top to bottom.
// The main application code must be wrapped in such a function, you are allowed and forced to have the package and import not wrapped in a function. There
// are a few things that are allowed to not be wrapped in a function. But for the most part all of your code must be wrapped in functions. If you're using
// multiple files you must not use 'main' for another function and only the function where you want to start. You can only have one.
// If you're building a function or package for a library you don't need to have a main function since it's meant to be imported to other applications.
// ----------------------------------------------------------------
// These are core components that make up Go components.

// Values and types:
// ----------------------------------------------------------------
// Basic types - (int, a number without a decimal.) (float64, a number with decimals.) (string, a text value.) (bool, true || false.)
// There are also niche basic types as well. (unit, an unsigned type that is strictly not a negative int.) (int32, a 32-bit integer.)
// (rune, an alias for int32, it represents a Unicode point and is used when dealing with Unicode characters.) (unit32, a 32-bit unsigned int.)
// (int64, a 64-bit signed integer.) There is also int8 and unit8 that work the same way but with 8-bits instead.
// null values are used if no value is explicitly specified. int = 0, float64 = 0.0, string = " ", bool = false.
// ----------------------------------------------------------------
// Using Alternative Variable Declaration Styles:
// ----------------------------------------------------------------
// Instead of calling out variable like 'var value = 5.5' where this value would be inferred as a float64 or 'var value float64 = 10',
// we can simply say 'value := 5.5'. This is a shortcut from Go that allows automatically declare, define and infer the type. This will allow Go to automatically infer the type that it is and is recommended for these types that you know you want,
// to be this type.
// You can also declare multiple variables of the same type in one line just like in VBA. 'var value, return, overallValue float64 = 1000, 10, 100000'.
// You can also declare multiple types in the same line. var investmentAmount, years = 1000, "10". You have no just declared an int and a string in the same line.
// This may come to be very useful later on, but you must let Go infer the type. If you declare the type they must all be of that type.
// ----------------------------------------------------------------
// Understanding The Importance of Variables:
// ----------------------------------------------------------------
// When using the fmt.Scan function(input) we can use a variable that we have already declared by using it and a pointer to it. The pointer used is &.
// So in this case it would look like 'fmt.Scan(&foo)'. This would point to the foo variable. Pointers will be talked about later on.
// ----------------------------------------------------------------
// Improved User Input Fetching:
// ----------------------------------------------------------------
// In a case where you want your variable to be null or open, you must always use the var syntax. 'years' does not work as Go doesn't know what you're writing about.
// So you will need to write out 'var years float64' instead. It does need to know what type will be eventually stored in the variable. So we must set the type.
// fmt.Scan() is a great function for single string values but not so much for multi string values. How we deal with/handle this will be brought up later.
// ----------------------------------------------------------------
//