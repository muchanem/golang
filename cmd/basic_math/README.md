## Basic Math

> Math Operators   

  * Go has also the same basic operators as C and others.
  * ```+``` addition
  * ```-``` subtraction
  * ```*``` multiplication
  * ```/``` division
  * ```%``` Reminder
  * ```&``` Bitwise AND
  * ```|``` Bitwise OR
  * ```^``` Bitwise XOR
  * ```&^``` Bit clear
  * ```<<``` Left shift
  * ```>>``` Right shift

> Rules

  * Math requires same types
  * You cant add int to float
  * Convert int to float by:
  ```go
  var int1 = 5
	var float1 = 3.14
	sum := float64(int1) + float1
	fmt.Println(sum)
  ```
## Math package

> Basics

  * Implements common mathematical functions
  * Provides constants and functions  

> Verbs  

* Basic verbs     
```go
%.10g
```
* %10.g compacts the output. Replace 10 with how many numbers you want it to compact to
```go
%.10f
```
* %.f sets how many characters after the decimal you want. Replace 2 with however many characters after the decimal you want
