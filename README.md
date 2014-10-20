smalljs
=======

Basic javascript minify wrote with GO language.

##Feature

- Remove all comments in javascript code.

- Reduce unnecssary space in javascript code.

- Combine multiple javascript files into one file. 

 
##Usage

  ```go
  package main
  
  import (
  	"fmt"
  	"github.com/oneplus1000/smalljs"
  )
  
  func main() {
  	smj := smalljs.NewSmallJs()
  	err := smj.Make([]string{
  		"js/file1.js",   //input 1
  		"js/file2.js",   //input 2
  	},
  		"js/result.js",  //output file
  	)
  	if err != nil {
  		fmt.Printf("%s\n", err.Error())
  	}
  }
  ```
