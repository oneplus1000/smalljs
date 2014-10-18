smalljs
=======

basic javascript minify write in golang.

##feature

-remove all comments

-remove new line
 

  ```go
  package main
  
  import (
  	"fmt"
  	"github.com/oneplus1000/smalljs"
  )
  
  func main() {
  	smj := smalljs.NewSmallJs()
  	err := smj.Make([]string{
  		"js/comment.js",
  	},
  		"js/result_comment.js",
  	)
  	if err != nil {
  		fmt.Printf("%s\n", err.Error())
  	}
  }
  ```
