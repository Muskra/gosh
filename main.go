package main

// program imports
import (
  // standard library
  "bufio"
  "fmt"
  "os"
  "strings"
  // customs
  "gsh/std"
)

// main script, it loops into an infinite command prompt.
func main() {
	// program loop
	for {
    defer fmt.Println("syntax error!")
    // variables
    var prompt string = "";
    var commandList []string;

  	fmt.Printf("$_ ")
		// read the user's input
    // we use pointers in ReadingInput() to improve performances
    ReadingInput(&prompt)
    // loops if the user typed enter
    if prompt == "\n" {
      continue;
    }
    // handling the "\" character which means the command is multiline
    for bl := strings.HasSuffix(strings.TrimSuffix(prompt, "\n"), "\\"); bl == true; {
      defer fmt.Println("syntax error!")
      fmt.Print("> ")
      ReadingInput(&prompt)
      // a check is required to break the loop
      if bl := strings.HasSuffix(strings.TrimSuffix(prompt, "\n"), "\\"); bl == false {
        break;
      }
    }
    // this is where we call our next function
    commandList = SentenceSplitter(prompt)
    //fmt.Println("")
    if len(commandList) > 0 {
      for _, cmd := range commandList {
        fmt.Printf("{%s}", cmd)
      }
    }
    fmt.Println("")
  }
}
// ReadingInput
// input reader
func ReadingInput(xPrompt *string) {
  // last thing to do when we can't handle anything
	reader := bufio.NewReader(os.Stdin)
	yPrompt, _ := reader.ReadString('\n')
  *xPrompt = *xPrompt+yPrompt
}
// this function is made to split the sentence into words, where each special characters, commands and other artifacts are isolated
// NEED TO IMPLEMENT:
//   * IOE because all args are commands
//   * string handling
func SentenceSplitter(sentence string) []string {
  // defines the ret variable
  var ret []string;
  // split our sentence by " " (spaces)
  xSentence := strings.Fields(sentence)
  // here we lookup throught all words from the sentence and return a list of each words separated from their special characters
  for _, wrd := range xSentence {
    wrd = strings.TrimSpace(wrd)
    //wrd = strings.TrimPrefix(strings.TrimSuffix(wrd, " "), " ")
    spltWrd := SpecialCharLookup(wrd)
    if len(spltWrd) >= 1 {
      for _, xWrd := range spltWrd {
        ret = append(ret, xWrd)
      }
    }
  }
  return ret;
}

func SpecialCharLookup(word string) []string {
  // declarations
  // buffer
  var buff string;
  // return variable
  var ret []string;
  // boolean checks, used as true to declare if the previous char was special or not
  var bl bool;
  // if there is at least 2 chars
  if len(word) > 1 {
    // we loop throught the string
    for index, char := range word {
      // here we check wether Census returned (as true) that the char is actually special
      if std.Census(2, string(char)) == true && bl == false {
        // here, we also checks if the special char combined with the next one is a special char too
        if index < len(word)-1 && std.Census(2, string(char)+string(word[index+1])) == true {
          ret = append(ret, buff, string(char)+string(word[index+1]))
          bl = true
        } else {
          ret = append(ret, buff, string(char))
        }
        buff = ""
        // in case we have a special char, but the bool was set, we just reset the buffer
      } else if std.Census(2, string(char)) == true && bl == true {
        buff = ""
        bl = false
        // when we are at the last char of the string
      } else if index == len(word)-1 && bl == false {
        ret = append(ret, buff+string(char))
        // if nothing was triggered, just append buffer
      } else {
        buff = buff+string(char)
      }
    }
    // we have only one char in the string
  } else if len(word) == 1 {
    ret = append(ret, word)
    // error trigger
  } else {
    fmt.Println("parse error!")
  }
  return ret
}
