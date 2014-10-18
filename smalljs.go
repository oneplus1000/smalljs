package smalljs

import (
	"bytes"
	"fmt"
	"io/ioutil"
	//"regexp
	"strings"
	"unicode/utf8"
)

type SmallJs struct {
	//stringRgx *regexp.Regexp
}

func NewSmallJs() *SmallJs {
	var smj SmallJs
	return &smj
}

func (me *SmallJs) Make(srcs []string, dest string) error {
	fmt.Println("start\n")
	err := me.init()
	if err != nil {
		return err
	}
	var buff bytes.Buffer
	for _, src := range srcs {
		jsRaw, err := ioutil.ReadFile(src)
		if err != nil {
			return err
		}
		jsRaw, err = me.RemoveCommentAndDebugger(jsRaw)
		if err != nil {
			return err
		}
		jsRaw, err = me.ReduceSpace(jsRaw)
		if err != nil {
			return err
		}

		buff.Write(jsRaw)
	}
	err = ioutil.WriteFile(dest, buff.Bytes(), 0666)
	if err != nil {
		return err
	}
	return nil
}

func (me *SmallJs) init() error {
	//	var err error

	return nil
}

func (me *SmallJs) ReduceSpace(jsRaw []byte) ([]byte, error) {
	js := string(jsRaw)
	/*js = strings.Replace(js, "\r\n", " ", -1)
	js = strings.Replace(js, "\n", " ", -1)
	*/
	var buff bytes.Buffer
	lines := strings.Split(js, "\n")
	for _, line := range lines {
		buff.WriteString(strings.TrimSpace(line))
		buff.WriteString(" ")
	}

	return buff.Bytes(), nil
}

func (me *SmallJs) RemoveCommentAndDebugger(jsRaw []byte) ([]byte, error) {

	i := 0
	max := utf8.RuneCount(jsRaw)

	inCommentOneline := false
	inCommentMultiline := false
	inStr1 := false
	inStr2 := false
	inReg := false

	var comments []Range // comment and debugger
	indexOfcomments := 0
	for i < max {
		//fmt.Printf("%#U\n", rune(jsRaw[i]))
		if !(inCommentOneline || inCommentMultiline || inStr1 || inStr2 || inReg) { //start
			if me.CheckRuneFromBytes(jsRaw, i, "//") {
				//fmt.Println("//")
				inCommentOneline = true
				var tmp Range
				comments = append(comments, tmp)
				comments[indexOfcomments].Start = i
				i = i + 2
				continue
			} else if me.CheckRuneFromBytes(jsRaw, i, "/*") {
				//fmt.Println("/*")
				inCommentMultiline = true
				var tmp Range
				comments = append(comments, tmp)
				comments[indexOfcomments].Start = i
				i = i + 2
				continue
			} else if me.CheckRuneFromBytes(jsRaw, i, "/") && !me.CheckRuneFromBytes(jsRaw, i+1, "/") {
				//fmt.Println("/")
				inReg = true
				i = i + 1
				continue
			} else if me.CheckRuneFromBytes(jsRaw, i, "'") {
				//fmt.Println("\"")
				inStr1 = true
				i = i + 1
				continue
			} else if me.CheckRuneFromBytes(jsRaw, i, "\"") {
				//fmt.Println("\"")
				inStr2 = true
				i = i + 1
				continue
			} else {
				if me.CheckRuneFromBytes(jsRaw, i, "debugger") {
					//fmt.Println("debugger")
					var tmp Range
					comments = append(comments, tmp)
					comments[indexOfcomments].Start = i
					i = i + utf8.RuneCountInString("debugger")
					comments[indexOfcomments].End = i
					indexOfcomments++
					continue
				}
			}

		} else { //end
			if inCommentOneline {
				if me.CheckRuneFromBytes(jsRaw, i, "\n") {
					inCommentOneline = false
					//fmt.Println("end")
					i = i + 1
					comments[indexOfcomments].End = i
					indexOfcomments++
					continue
				}
			} else if inCommentMultiline {
				if me.CheckRuneFromBytes(jsRaw, i, "*/") {
					inCommentMultiline = false
					//fmt.Println("end")
					i = i + 2
					comments[indexOfcomments].End = i
					indexOfcomments++
					continue
				}
			} else if inReg {
				if me.CheckRuneFromBytes(jsRaw, i, "\n") || me.CheckRuneFromBytes(jsRaw, i, ";") || me.CheckRuneFromBytes(jsRaw, i, ",") {
					inReg = false
					//fmt.Println("end")
					i = i + 1
					continue
				}
			} else if inStr1 {
				if me.CheckRuneFromBytes(jsRaw, i, "'") {
					inStr1 = false
					//fmt.Println("end")
					i = i + 1
					continue
				}
			} else if inStr2 {
				if me.CheckRuneFromBytes(jsRaw, i, "\"") {
					inStr2 = false
					//fmt.Println("end")
					i = i + 1
					continue
				}
			}
		}

		i++
	}

	var buff bytes.Buffer

	i = 0
	for _, comment := range comments {
		buff.Write(jsRaw[i:comment.Start])
		buff.WriteString(" ")
		i = comment.End
	}
	buff.Write(jsRaw[i:])

	return buff.Bytes(), nil
}

func (me *SmallJs) CheckRuneFromBytes(jsRaw []byte, index int, ch string) bool {
	var matched = true
	i := index
	for _, ru := range ch {
		if ru != rune(jsRaw[i]) {
			matched = false
			break
		}
		i++
	}

	if index-1 >= 0 {
		if me.CheckRune(rune(jsRaw[index-1]), "\\") {
			matched = false
		}
	}

	return matched
}

func (me *SmallJs) CheckRune(ru rune, ch string) bool {
	if val, _ := utf8.DecodeRuneInString(ch); val == ru {
		return true
	}
	return false
}

type Range struct {
	Start int
	End   int
}
