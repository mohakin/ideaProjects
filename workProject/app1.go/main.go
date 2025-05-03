// this file should be used as a learning forum for GOlang. I will try to comment in each new section being covered here.
package main

import "fmt"

func main() {
	fmt.Print("Hello World")

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
// func =
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
// This is convenient during development but thats not the only way to execute the code. And it's also not how the code will be executed when written for
// production and make it available for others. Those people may not even have GO on their system. Typically you would run go.build in your project folder.
// This will then tell GO to build an executable file so that it can run on systems that don't have Go on their system.
// ----------------------------------------------------------------

// Understanding Go Modules & Building Go Programs:
// ----------------------------------------------------------------
// We get an error in the terminal about not having a main module. One module consists of multiple packages. In many cases a Go project is a Go module.
// We could consider this app a module. We have to run a specific command to show that this is a module. In order to do this we need to run a command called
// ' go mod init '. We must give our module a name or a path where it can be found. By running 'go mod init' with a path name like 'example.com/first-app',
// it will put a go.mod file in our folder for us.
