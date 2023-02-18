package std

import(
  "fmt"
	"bufio"
	"os"
	"strings"
)

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
      OutMsg: "I/O Error\n" + desc,
    }
	  return newDesc
  }
}

