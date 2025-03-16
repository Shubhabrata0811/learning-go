package main

import "fmt"

func main() {
	/*int8	8-bit	-128 to 127
	int16	16-bit	-32,768 to 32,767
	int32	32-bit	-2,147,483,648 to 2,147,483,647
	int64	64-bit	-9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
	int	System-dependent	32-bit or 64-bit*/

	var num1 int = 10; //default value 0
	fmt.Println("num1 int:", num1);
	
	var num2 int16 = 16;
	fmt.Println("num2 int:", num2);

 var str string = "shubha";
 fmt.Println("Hi, I'm", str);

 var complex1 complex64 = 2 + 3i;
 var complex2 complex64 = 5 + 4i;
 var complex3 complex64 = complex1 + complex2;
 fmt.Println("Sum:", complex3);

}