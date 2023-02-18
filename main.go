package main

// program imports
import (
	"fmt"
	"strings"
	"bufio"
	"os"
	//"gsh/std"
)

const SpecialCharacters = "&;|";
var SpecialMultiCharacters = []string{
  "&&",
  "||",
}
// special characters states
type Tempo struct {
  last    bool
  actual  bool
  next    bool
}
// word tokens
type Lexic struct {
  word    bool
  special bool
}
// python's like tuple
type Tuple struct {
  x int
  y int
}
// token with address
type Token struct {
  lexic Lexic
  address Tuple
}

// main script, it loops into an infinite command prompt.
func main() {
	// program loop
	for {
		// read the user's input
		prompt := ReadingInput()
		// parse the user's input
		Lexer(prompt)
	}
}
// ReadingInput
// input reader, it removes the '\n' at every ends
func ReadingInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("$_ ")
	prompt, _ := reader.ReadString('\n')
	// trim only if the string is longer than one character
	if len(prompt) > 1 {
		// removing '\n' character
		prompt = strings.TrimSuffix(prompt, "\n")
	}
	return prompt
}
/*
special characters:
        |   the output of the first command is sent to the second one
        &   launch command in background, the second one is foregrounded
        ;   launch second command after the first one, not affected by error handling
        &&  launch second command after the first one but won't launch if error has been raised
        <   redirect from
        >   redirect to
        \   used to cancel the next character
        ""  consider multiple words as one
        ''  same as quotes
        ``  same as quotes but multilines
        #   comment/ignore the line after the # character
*/
func IsSpecialChar(prompt string, index int) Tempo {
	var ret = Tempo{last: false, actual: false, next: false}
	// we are at the first char of the sentence and there is more after
	if index == 0 && len(prompt) > index+1 {
		ret := Tempo{
			actual: strings.ContainsAny(string(prompt[index]), SpecialCharacters),
			next:   strings.ContainsAny(string(prompt[index+1]), SpecialCharacters),
		}
		return ret
	}
	// we are at the first character and it's the only one
	if index == 0 && len(prompt) == index+1 {
		ret := Tempo{
			actual: strings.ContainsAny(string(prompt[0]), SpecialCharacters),
			next:   false,
		}
		return ret
	}
	// we are not at the first char of the sentence and there is more after
	if index != 0 && len(prompt) > index+1 {
		ret := Tempo{
      last:   strings.ContainsAny(string(prompt[index-1]), SpecialCharacters),
			actual: strings.ContainsAny(string(prompt[index]), SpecialCharacters),
			next:   strings.ContainsAny(string(prompt[index+1]), SpecialCharacters),
		}
		return ret
	}
	// we are not at the first character and there is no more characters after
	if index != 0 && len(prompt) == index+1 {
		ret := Tempo{
      last:   strings.ContainsAny(string(prompt[index-1]), SpecialCharacters),
			actual: strings.ContainsAny(string(prompt[index]), SpecialCharacters),
			next:   false,
		}
		return ret
	}
	return ret
}

func IsSpecialMultiChar(word string) bool {
  var ret bool = false
  for _, mChar := range SpecialMultiCharacters {
    if word == mChar {
      ret = true
      return ret
    }
  }
  return ret
}

// lexical analysis
func Lexer(prompt string) {
	// len of prompt
  var promptLenght int = len(prompt);
  // index steps
	var steps int = 0;
  var indexList []Token;
	// here we iterate through each characters of the string
	for index, _ := range prompt {
		// checking if the actual char and the next char are special
    special := IsSpecialChar(prompt, index)
    //fmt.Println(special)
    // we won't accept three special characters
    if special.actual == true && special.last == true && special.next == true {
      fmt.Println("PARSING ERROR!")
      break;
    }
    // if we are at index zero, the last one don't exist
    if index == 0 && special.last == true {
      fmt.Println("PARSING ERROR!")
      break;
    }
    // we are at index zero and the next one is also a special one, we check wether the two of them are a special multi character
    if index == 0 && special.last != true && special.actual == true && special.next == true {
      if IsSpecialMultiChar(string(prompt[0])+string(prompt[1])) == true {
        indexList = append(indexList, Token{ lexic:Lexic{ word:false, special:true }, address:Tuple{ x:0, y:1 }})
        steps = index+1
      } else {
        indexList = append(indexList, Token{ lexic:Lexic{ word:false, special:true }, address:Tuple{ x:0, y:0 }})
        steps = index+1
      }
    }
    // we are in the sentence, the actual char is special, we check wether the next char is special or not, the two of them can also be a special multi character
    if index != 0 && special.last != true && special.actual == true {
      if special.next == true && IsSpecialMultiChar(string(prompt[index])+string(prompt[index+1])) == true {
        indexList = append(indexList, Token{ lexic:Lexic{ word:true, special:false }, address:Tuple{ x:steps, y:index }})
        indexList = append(indexList, Token{ lexic:Lexic{ word:false, special:true }, address:Tuple{ x:index, y:index+1 }})
        steps = index+1
      }
      if special.next == true && IsSpecialMultiChar(string(prompt[index])+string(prompt[index+1])) != true {
        fmt.Println("PARSING ERROR!")
        break;
      }
      if special.next != true {
        indexList = append(indexList, Token{ lexic:Lexic{ word:true, special:false }, address:Tuple{ x:steps, y:index }})
        indexList = append(indexList, Token{ lexic:Lexic{ word:false, special:true }, address:Tuple{ x:index, y:index+1 }})
        steps = index+1
      }
    }
    // we are in the sentence, the actual char is special, we check wether the last one was special too and increment steps
    if index != 0 && special.last == true && special.actual == true {
      steps = index+1
    }
    // we are at the end of the sentence, we check wether the last char and the actual ones are special, if not just append the end of the list as word
    if !(promptLenght > index+1) {
      if special.last != true && special.actual != true {
        indexList = append(indexList, Token{ lexic:Lexic{ word:true, special:false }, address:Tuple{ x:steps, y:index+1 }})
      }
    }
	}
  // just printouts to understand what the program does...
  fmt.Printf("sentence:{%s}, lenght:{%d}\n", prompt, promptLenght)
  for _, t := range indexList {
    fmt.Printf("%+v, %+v, {%+v}\n", t.lexic, t.address, string(prompt[t.address.x:t.address.y]))
  }
}
