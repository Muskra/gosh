package std

//import(
  //"fmt"
	//"bufio"
	//"os"
	//"strings"
//)

// descriptor a structure file descriptor like scheme input/output/error
// to explain, it's like in UNIX:
//		0 = INPUT	|	READ
//		1 = OUTPUT	|	WRITE
//		2 = ERROR	|	READ
// here READ means that the I/O will read values from while WRITE means I/O will write to
// the write is specific as it will be reused after the command or program has ended
type Descriptor struct {

  // defined to be the I/O/ERR value wich a program returns
  Id      int
  // defined if the id value is set to 0 as INPUT to be READ before being sent
  InMsg   []string
  // defined if the id value is set to 1 as OUTPUT to be WRITTEN
  OutMsg  string
  // defined if the command or program throws error
  ErrMsg  string
}

// ioErr the input/output/error handler
func IOE(desc Descriptor) Descriptor {
	// with descriptor id of 0 we just pass the actual command throught
	if desc.Id == 0 {
		cmdMap[desc.InMsg[0]](desc.InMsg[1:])
  }
  // separately, we check if an error has been raised (value of 2)
  if desc.Id == 2 {
    newDesc := Descriptor{
      Id:     2,
      ErrMsg: "Error: \n%s\n" + desc.ErrMsg,
    }
    return newDesc
  }
  // as a normal return (value of 1) we just send back the new message
  if desc.Id == 1 {
    newDesc := Descriptor{
      Id:     1,
      OutMsg: desc.OutMsg,
    }
    return newDesc
  }
  // in case of a bad Id has been declared for the descriptor
  if desc.Id != 0 || desc.Id != 1 || desc.Id != 2 {
    newDesc := Descriptor{
      Id:     2,
      OutMsg: "I/O Error: can't process I/O!\n" + desc.ErrMsg,
    }
    return newDesc
  }
  return Descriptor{Id: 2, OutMsg: "I/O Error: can't process I/O!\n"}
}
// another tool that wee need is a function that verify if a command or a special argument exist, the function might be similar to IOE but the purpose isn't really the same as we are not checking for commands and trying to launch, nor communicate with it. So here we are trying to make a census of a command or an arguments that exists, and we simply return a bool true if it exists
// we also need to define what we want to check, so a simple int (from 0 to 2) will define the type as:
//  * 0: is all types
//  * 1: is a command
//  * 2: is an argument
func Census(xType int, xName string) bool {
  //defer fmt.Println("I/O Error: can't conduct census!\n")
  // if type is all types
  if xType == 0 {
    // declaring the check values
    _, cm := cmdMap[xName]
    _, sam := specialArgMap[xName]
    // comparison
    if cm == true || sam == true {
      return true;
    } else {
      return false
    }
  // if type is command
  }
  if xType == 1 {
    if _, ok := cmdMap[xName]; ok == true {
      return true;
    } else {
      return false;
    }
  // if type is argument
  }
  if xType == 2 {
    if _, ok := specialArgMap[xName]; ok == true {
      return true;
    } else {
      return false;
    }
  }
  return false;
}
