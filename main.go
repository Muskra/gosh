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
// main function
// here you can define a Characterset variable and a boolean variable Test to verify if the syntax you have created is handling normally or not by the program. You can also set the lenght of the words you want to test but be carefull since it can generate performance issues i'm not responsible for any damage done with this
// if the Test variable is false (as by default) program will launch normally
func main() {
  // variables
  var Characterset string = " |&;\""; // doesn't work with "quotes" for some reason
  // value is 6 by default
  var SetLenght int = 6;
  // value is false by default
  var Test bool = false;
  // if Test variable is false we launch program normally
  if Test == false {
    NormalLoop()
  }
  // if Test variable is true we launch the test loop
  if Test == true {
    TestLoop(Characterset, SetLenght)
  }
}
// test loop, can be launched if Test variable is set to true
func TestLoop(characterset string, lenght int) {
  // declaration
  var pmtList []string;
  // generating each sentences from the GenerateCombinations() function with the given Characterset variable
	for combination := range std.GenerateCombinations(characterset, lenght) {
	  pmtList = append(pmtList, combination)
	}
  // program loop
  for _, prompt := range pmtList {
    //defer fmt.Println("syntax error!")
    // variables
    var commandList []string;
    // loops if the user typed enter
    if prompt == "\n" {
      continue;
    }
    // handling the "\" character which means the command is multiline
    for bl := strings.HasSuffix(strings.TrimSuffix(prompt, "\n"), "\\"); bl == true; {
      //defer fmt.Println("syntax error!")
      fmt.Print("> ")
      ReadingInput(&prompt)
      // a check is required to break the loop
      if bl := strings.HasSuffix(strings.TrimSuffix(prompt, "\n"), "\\"); bl == false {
        break;
      }
    }
    fmt.Printf("sentence: \n{%s}", prompt)
    // this is where we call our next function
    commandList = SentenceSplitter(prompt)
    // printout result
    fmt.Println("\nfinal_command: ")
    if len(commandList) > 0 {
      for _, cmd := range commandList {
        fmt.Printf("{%s}", cmd)
      }
    }
    fmt.Println("")
  }
}
// Entrypoint of the program, will be launch as normal behavior
func NormalLoop() {
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
  var xRet []string;
  // used for temporary or buffering
  var temp string;
  // bool variable, used as a check for the presence of "quotes"
  var bl bool;
  // iterate throught the whole sentence and isolate each strings from other words within a slice
  for index, ch := range sentence {
    if string(ch) == "\"" && bl == false {
      ret = append(ret, temp)
      bl = true
      temp = string(ch)
      continue;
    }
    if string(ch) == "\"" && bl == true {
      temp = temp+string(ch)
      ret = append(ret, temp)
      bl = false
      temp = ""
      continue;
    }
    if string(ch) != "\"" && bl == false {
      temp = temp+string(ch)
    }
    if string(ch) != "\"" && bl == true {
      temp = temp+string(ch)
    }
    if index == len(sentence)-1 {
      ret = append(ret, temp)
    }
  }
  // here we are iterating again, but this time throught the slice, and we are isolating special characters, say punctuation from the other words
  // we also won't check strings here since they already are checked beforehand
  for index, word := range ret {
    if len(word) > 1 {
      if strings.Contains(string(word), "\"") == false {
        for _, wrd := range strings.Fields(word) {
          wrd = strings.TrimSpace(wrd)
          spltWrd := SpecialCharLookup(wrd)
          for _, xWrd := range spltWrd {
            xRet = append(xRet, xWrd)
          }
        }
      } else if strings.Contains(string(word), "\"") == true {
        xRet = append(xRet, string(word))
      } else {
        fmt.Println("syntax error!")
      }
    } else if len(word) == 1 || index == len(ret)-1 {
      xRet = append(xRet, word)
    } else {
      fmt.Print("\nparse error!\n")
      break;
    }
  }
  return xRet
}

// checks wether we have special characters or not
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
  } 
  if len(word) == 1 {
    ret = append(ret, word)
    // error trigger
  } /*else {
    fmt.Print("\nparse error!\n")
  }*/
  return ret
}
