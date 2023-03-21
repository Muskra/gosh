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
        fmt.Printf("{%s}{%#v}", cmd, commandList)
      }
    }
    //fmt.Println("")
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
      commandList = SentenceSplitter(strings.TrimSuffix(prompt, "\n"))
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
  var zRet []string;
  // used for temporary or buffering
  var temp string;
  // bool variable, used as a check for the presence of "quotes"
  var bl bool;
  // iterate throught the whole sentence and isolate each strings from other words within a slice
  for index, ch := range sentence {
    // character is a quote
    if string(ch) == "\"" && bl == false && index < len(sentence)-1 {
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
    // character is not a quote
    if string(ch) != "\"" && bl == false {
      temp = temp+string(ch)
    }
    if string(ch) != "\"" && bl == true {
      temp = temp+string(ch)
    }
    // we are at the end of the sentence
    if index == len(sentence)-1 && bl == false && string(ch) != "\"" {
      ret = append(ret, temp)
    }
    // we are at the end of the sentence but we have seen a quote before, the program raise an error because we can't parse the string. this is a syntax error from the user
    // break is not required but for "safety reasons" i prefer to add it explicitely
    if index == len(sentence)-1 && (bl == true || string(ch) == "\"") {
      ret = nil
      fmt.Println("syntax error!")
      break;
    }
  }
  // here we are iterating again, but this time throught the slice, and we are isolating special characters, say punctuation from the other words
  // we also won't check strings here since they already are checked beforehand
  for index, word := range ret {
    if len(word) > 1 {
      // before calling the function respoinsible to lookup the special characters, we verify if the word don't have any quote in it
      if strings.Contains(string(word), "\"") == false {
        for _, wrd := range strings.Fields(word) {
          wrd = strings.TrimSpace(wrd)
          spltWrd := SpecialCharLookup(wrd)
          for _, xWrd := range spltWrd {
            xRet = append(xRet, xWrd)
          }
        }
        // in case we have some quotes we append it without trying to lookup for special characters
      } else if strings.Contains(string(word), "\"") == true {
        xRet = append(xRet, string(word))
        // incase there is something unexpected
      } else {
        xRet = nil
        fmt.Println("syntax error!")
        break;
      }
      // if word isn't an ampty string AND ( word is just one char OR index is equal to word lenght )
    } else if len(word) != 0 && (len(word) == 1 || index == len(ret)-1) {
      xRet = append(xRet, word)
      // word is an empty string, we don't care about it
    } else if len(word) == 0 && word != "\"" {
      continue;
    } else {
      xRet = nil
      fmt.Println("parse error!")
      break;
    }
  }
  // here we just reloop and remove all empty quote, not ideal but it's working and prevent future actions to be flooded by empty useless words
  for _, word := range xRet {
    if len(word) == 0 {
      continue;
    } else {
      zRet = append(zRet, word)
    }
  }
  return zRet
}

// checks wether we have special characters or not
func SpecialCharLookup(word string) []string {
  // declarations
  // buffer
  var buff string;
  // return variable
  var yRet []string;
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
          yRet = append(yRet, buff, string(char)+string(word[index+1]))
          bl = true
        } else {
          yRet = append(yRet, buff, string(char))
        }
        buff = ""
        // in case we have a special char, but the bool was set, we just reset the buffer
      } else if std.Census(2, string(char)) == true && bl == true {
        buff = ""
        bl = false
        // when we are at the last char of the string
      } else if index == len(word)-1 && bl == false {
        yRet = append(yRet, buff+string(char))
        // if nothing was triggered, just append buffer
      } else {
        buff = buff+string(char)
      }
    }
    // we have only one char in the string and is not of len(0)
  } 
  if len(word) == 1 && len(word) != 0 {
    yRet = append(yRet, word)
  }
  return yRet
}
