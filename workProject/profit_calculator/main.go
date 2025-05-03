package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	//var investmentAmount float64
	//var years float64
	//expectedReturnRate := 5.5

	//fmt.Println("Welcome to your personal profit calculator!")
	outputText("Welcome to your personal profit calculator!")

	//fmt.Print("Please enter your revenue: ")
	outputText("Please enter your revenue: ")
	fmt.Scan(&revenue)

	//fmt.Print("Please enter your expenses: ")
	outputText("Please enter your expenses: ")
	fmt.Scan(&expenses)

	//fmt.Print("Please enter your tax rate: ")
	outputText("Please enter your tax rate: ")
	fmt.Scan(&taxRate)

	//ebt := revenue - expenses
	//profit := ebt * (1 - taxRate/100)
	//ratio := ebt / profit

	newEBT, newProfit, newRatio := returnedValue(revenue, expenses, taxRate)

	fmt.Printf("You're earnings before taxes are $%.2f.\n", newEBT)
	fmt.Printf("You're earnings after taxes are $%.2f.\n", newProfit)
	fmt.Printf("Your ratio is %.2f%%.", newRatio)

	//futureValue, futureRealValue :=calculateFutureValue(investmentAmount, expectedReturnRate, years)
	//fmt.Print(futureValue, futureRealValue)

}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	//return fv, rfv
	return
}

func returnedValue(revenueOut, expensesOut, taxRateOut float64) (float64, float64, float64) {

	ebt := revenueOut - expensesOut
	profit := ebt * (1 - taxRateOut/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

// The only thing that was off for me in this little project was the formula for profit and realizing last
// second about how the printf statements should work. I figured this part out by myself, we are not using
// printf statements yet. %f = float, %s = string, %i = int, %v = value or constant and so on and so forth. Very simple.
// I'm still not sure how we get all the 0's down to just 2 decimal places. Maybe an 8-bit float?
// ----------------------------------------------------------------
// 								 Formatting Strings(Text) - Basics
// ----------------------------------------------------------------
// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
// Now we're going over printf xD. xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
// One way to format this would be like this, 'fmt.Println("Future Value:", futureValue').
// I was not aware of this way, and it makes sense looking at it.
// Printf was covered and so far I'm following along pretty well with it.
// ----------------------------------------------------------------
// 									Formatting Floats in Strings:
// ----------------------------------------------------------------
// If you need to use a % sign you will need to use 2 of them. It should look like this %%.
// For precision with the printf(or F) we can do this. %.2f, or %9f, %9.2f, %9.0f and so on and so forth.
// In the cases above 9 is the width and the .X is how many numeral places we want to go into for the decimal point.
// ----------------------------------------------------------------
// 										Creating Formatted Strings:
// ----------------------------------------------------------------
// fmt.Sprint(f and ln)() to make a new text value.
// For instance. formattedFV := fmt.Sprintf("Future value %.1f", futureValue).
// I can then pass this onto my print statement and shorten it up quite a bit.
// I.E. fmt.Print(formattedFV).
// ----------------------------------------------------------------
// 										Building Multiline Strings:
// ----------------------------------------------------------------
// You can not separate lines in a print statement with ""(double quotes). If we want to do that we will
// have to use the back tick(``). Using this we can now line break without needing \n or Println,
// all you need to do is just put an empty line in print statement.
// ----------------------------------------------------------------
// 										Understanding Functions:
// ----------------------------------------------------------------
// The idea behind a function is to have a block of code that is only used when it is called.
// The main function is different due to it being used by Go as a place to start at.
// When creating a new function you need to pass in at least one argument. For instance (func outputText(text string){fmt,Print()}). You also need to put,
// what the parameter needs to be, in our case we want/need a string. You can add as many arguments as you want to it, they can be multiple types.
// In order to use multiple arguments you need to add a 'comma(,)' between each argument.
// Example: func outputText(text string, text2 string). You can also write it as (text, text2 string). Just to shorten things up a little bit.
// If we are using the above style though they must all be strings, if you want it to be something different you'll have to do it like before with,
// (text string, number int) etc...
// ----------------------------------------------------------------
// 						Functions Return Values & Variable Scope:
// ----------------------------------------------------------------
// Functions don't need to accept any arguments if not needed. They can also produce and return values.
// The math. function works this very way. Think of math.Pow() which is used to calculate the power of something.
// The 'return' keyword(builtin) will tell Go that this value should be returned as a result by this function.
// You can return multiple values by separating them with a comma(,). For instance look at the calculateFutureValue function above.
// In Go any variables that are defined in a function are scoped to that function. Meaning they are only available to that function.
// You can declare any variable or constant declarations outside a function making it scoped to the whole file I.E. making them global.
// Everything else must be inside a function.
// When you start defining a function you must declare what type of value it's returning when using the 'return' keyword.
// For example: func calculateFutureValue(investmentAmount, expectedReturnRate, years float64)(float64, float64) {}
// In this function we have a return statement for 2 values that we want to be returned by the function. These will both be float64,
// so we will have to add 2 float64 after the initial arguments are passed to the function. This has to do with Go being statically typed.
// You can tie a function's returned value to a variable by writing a variable name and then := func(v interface{}) or whatever your func is.
// If you have multiple items being returned you'll need that many variables for it. So in our case above we have a return statement for,
// 2 values. We will end up doing this: 'futureValue, futureRealValue :=calculateFutureValue(investmentAmount, expectedReturnRate, years)'.
// ----------------------------------------------------------------
// 								An Alternative Return Value Syntax:
// ----------------------------------------------------------------
// Instead of using the 'return' keyword. You can go to where you have your return value types, in our case it's (float64, float64).
// You can add your variables to those so that it looks like (fv float64, rfv float64). Then you don't need the := and you can just use =.
// If you're doing it like this then you don't need 'return fv, rfv'. You can just use return and Go will return the variables that you listed,
// next to your value types. Personally I like doing it the original way better as of now.
// ----------------------------------------------------------------
// 								 Exercise - Working with Functions:
// ----------------------------------------------------------------
// I misunderstood the first part about making the user input function one thing. I called the function multiple times instead.
// Other than that everything else is correct and it worked as expected.
// ----------------------------------------------------------------
// 									 Onwards to Control Structures:
// ----------------------------------------------------------------
// A new project folder will be made for this. With a new go innit path as well. This will continue over there.
