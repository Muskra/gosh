package std

// map functions to strings
type fnCmd func([]string) Descriptor;
var cmdMap map[string] fnCmd;
var specialArgMap map[string] int;//fnCmd;

// init
func init() {
	// this init function is defined in the config.go as a config file flavour

	// the list of all commands you add to your program can be added here
	// you can also remove with a simple comment the standard library functions you don't want in your final program
	// don't use the init() function in another way than "command declaration purposes" to make it clear and readable
	cmdMap = map[string] fnCmd {

	// standard library
	"clear":  clear,
  "exit":   exit,
  "help":   help,
  //"ps":     ps,
	// put custom commands below
	// if you print something, add \n to the end of your string to make it prettier
	// and don't forget the last ","

	/*
	exemple:
	// we also put the final comma because it's required
	"gosay": hellosay,
	*/

	}
  // this map is actually not working since i didn't implemented the associated functions for each special chars
	specialArgMap = map[string] int {//fnCmd {
		/*
		special characters:
		        |   the output of the first command is sent to the second one
		        &   launch command in background, the second one is foregrounded
		        ;   launch second command after the first one, not affected by error handling
		        &&  launch second command after the first one but won't launch if error is raise
		        <   redirect from
		        >   redirect to
		        \   used to cancel the next character
		        ""  consider multiple words as one
		        ''  same as quotes
		        ``  same as quotes but multilines
		        #   comment/ignore the line after the # character
		
		here we need to define a queue for the commands to launch properly in adequacy to the pipe parameter
		for instance, a pipe is no more than:
		  - launching the first command and keep the result
		  - launching the second command and pass it the result of the previous one
		in the program we will prevent the result to be parsed if it's an error.
		*/
    // WIP state
    "|": 1,
    // TODO state
    // i only added the basics here
    "||": 2,
    "&": 3,
    "&&": 4,
    ";": 4,
	}
}
