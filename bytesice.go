package main

//We want to say get the value or variable greeting and convert it to byteslice and then print it.

import (
	"fmt"
	"strings"
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



