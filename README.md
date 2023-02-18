# Go Shell (that is) Hackable.

## I wanted to do a shell in go. Simple and made to be hacked by design.
## This project is a way of learning computers, programming and so on, you can help me but don't give away the answers, i want to learn as much as i can.

You can help me to keep this program simple and stupid, i'm learning so don't also correct everything directly and let me dig the problem a bit with my stupid mind.

Thank you.

### I'm not responsible of any damage taken or due with this program, i'm just beginning programming in Go so DONT'T EVER use it in production. Or at YOUR OWN risks.

#### ALREADY IMPLEMENTED:

##### I/O/err like the unix file descriptors:
    * 0: Standard Input
    * 1: Standard Output
    * 2: Standard Error

##### A parser to make words from sentences:
    e.g:    ps -ef | grep sh   # is a sentence
            ps -ef             # will be launched first and the output pushed to grep
            grep sh            # will receive the ps -ef's result as parameter

    WIP special characters:
        |   pipe
        &   launch second command in background
        ;   launch second command after the first one, not affected by error handling
        &&  launch second command after the first one but won't launch if error is raise
        <   redirect from
        >   redirect to
        \   used to read multiple lines as one command but need to be placed at end of the line
        ""  consider multiple words as one
        ''  same as quotes
        ``  same as quotes but multilines
        #   comment/ignore the line after the # character

##### Commands to use that are not usable at the moment since we don't have any shell to work with:
implemented:
        exit
        clear
        help
        pipe
