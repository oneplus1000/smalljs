smalljs
=======

Basic javascript minify wrote with GO language.

##Feature

- Remove all comments.

- Reduce unnecssary space in javascript code.
 
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
  		"js/file1.js",
  		"js/file2.js",
  	},
  		"js/result.js",
  	)
  	if err != nil {
  		fmt.Printf("%s\n", err.Error())
  	}
  }
  ```
