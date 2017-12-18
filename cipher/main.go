package main

import (
	"fmt"
	"github.com/skilstak/go-input"
	"os"
	"time"
)

var encodeDatabase = map[string]string{
	"a": "1",
	"b": "2",
	"c": "3",
	"d": "4",
	"e": "5",
	"f": "6",
	"g": "7",
	"h": "8",
	"i": "9",
	"j": "a",
	"k": "b",
	"l": "c",
	"m": "d",
	"n": "e",
	"o": "f",
	"p": "g",
	"q": "h",
	"r": "i",
	"s": "j",
	"t": "k",
	"u": "l",
	"v": "m",
	"w": "n",
	"x": "o",
	"y": "p",
	"z": "q",
	" ": "r",
}
var decodeDatabase = map[string]string{
	"1": "a",
	"2": "b",
	"3": "c",
	"4": "d",
	"5": "e",
	"6": "f",
	"7": "g",
	"8": "h",
	"9": "i",
	"a": "j",
	"b": "k",
	"c": "l",
	"d": "m",
	"e": "n",
	"f": "o",
	"g": "p",
	"h": "q",
	"i": "r",
	"j": "s",
	"k": "t",
	"l": "u",
	"m": "v",
	"n": "w",
	"o": "x",
	"p": "y",
	"q": "z",
	"r": " ",
}

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "-e" {
			a := input.Ask("What is your lowcase message> ")
			encode(a)
		} else if os.Args[1] == "-d" {
			a := input.Ask("What is your encoded message> ")
			decode(a)
		} else if os.Args[1] == "-g" {
			fmt.Println("Key generating in file: pubkey.json")
			generate()
		} else {

		}

	} else {
		menu := input.Ask("Do you want to encode (e) or decode (d) a message> ")
		if menu == "encode" {
			a := input.Ask("What is your lowcase message> ")
			encode(a)
		} else if menu == "e" {
			a := input.Ask("What is your lowcase message> ")
			encode(a)
		} else if menu == "decode" {
			a := input.Ask("What is your encoded message> ")
			decode(a)
		} else if menu == "d" {
			a := input.Ask("What is your encoded message> ")
			decode(a)
		} else if menu == "g" {
			fmt.Println("Key generating in file: pubkey.json")
			generate()

		} else {
			os.Exit(0)
		}

	}
}

func encode(a string) {
	n := ""
	for _, r := range a {
		c := string(r)
		n += encodeDatabase[c]
		fmt.Print(n)
	}
	fmt.Print("\n")
	after()
}

func decode(a string) {
	n := ""
	for _, r := range a {
		c := string(r)
		n += decodeDatabase[c]
	}
	for _, t := range n {
		d := string(t)
		time.Sleep(500 * time.Millisecond)
		fmt.Print(d)
	}
	fmt.Print("\n")
	after()

}
func generate() {
	fmt.Println("hi")
}

func after() {
	menu := input.Ask("Do you want to encode (e) or decode (d) a message or stop (s)> ")
	if menu == "encode" {
		a := input.Ask("What is your lowcase message> ")
		encode(a)
	} else if menu == "e" {
		a := input.Ask("What is your lowcase message> ")
		encode(a)
	} else if menu == "decode" {
		a := input.Ask("What is your encoded message> ")
		decode(a)
	} else if menu == "d" {
		a := input.Ask("What is your encoded message> ")
		decode(a)
	} else if menu == "stop" {
		os.Exit(0)
	} else if menu == "s" {
		os.Exit(0)
	} else {
		os.Exit(0)
	}

}
