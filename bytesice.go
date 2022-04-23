package main

//We want to say get the value or variable greeting and convert it to byteslice and then print it.

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func main() {

	greeting := "Hi there"

	fmt.Println([]byte(greeting))

}

//output: [72 105 32 116 104 101 114 101 33 ]

//Next we want to convert the deck to byteslice ،
 //the roadmap : First convert the deck to a string and then convert the string to a byteslice.

 // deck = A slice containing a string

 //This turns an array of strings into byteslice.

 //reminder :
 // save to file function: A function that stores the list of cards in a file on a local device.

 //We need to create a help function to take a deck and return a string.
 
 func tostring()


 func (d deck) tostring() string {
	 
	return strings.Join ([]string(d),",")

	[] string (d)



 }

 
 cards.tostring()

 // tostring : This function had a deck receiver, ie it returns a string in exchange for receiving a deck.
 // d = It is a receiver. It consists of a slice of strings.
 //deck changedd to d and d is a slice of string

 //The next step in converting a deck to a byteslice is to convert slice of string to a string.

 // steps = deck (we have) changed to []string (slice of string) and then changing to byteslice ([]byteslice)

 //We have to use comma any value in sslice of string so that the strings are separated.
 //for example :
//["red","yellow","green"]  In converting to string, change this way   [red,yellow,green]
//This is very time consuming if done manually.
//We need to find a package called strings in the documentation section of the GO site and enter the coding space.
//One of the functions of this package is called join
//func Join (a[]string , sep string )string
//According to the description, this function takes an slice of string and converts them all into a string.
// sep = short for seperation , Used for spacing between strings.
//We need to add the package for this function above.



//Use the function writefile :

func (d deck)savetofile(filename string){


}

cards.savetofile()

//This function requires a receiver ,Add a deck receiver.

func writefile(filename string, data[]byteslice,[perm os.filemode])error

//According to the standard package description and defaults, we need the function type to be a string.
//the word error means the value that the function returns is of the error type.


func (d deck) savetofile (filename string)error

//Inside the body we have to write a command function to save files and code on the local hard disk,
// to do this we use the ioutil package.

func ioutil.Write(filename ,[]byteslice(d.tostring()))

//The first argument should be the name of the function 
//and the second argument a byteslice of the data we want to store on the hard drive.
//We also convert the deck to byteslice at the same time.

//The third argument is for permissions.
// Permission is for who has access to the files or who edits them, etc.
// If we use the default permissions, it means that anyone can read and edit these files.

return ioutil.writefile(filename,[]byteslice,d.tostring(),066)

//066 is desired ,use for default

//We must also enter the ioutil package
//To do this, we need to import it from a subpackage called io.
//io/ioutil