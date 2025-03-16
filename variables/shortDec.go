package main

import "fmt"

func main()  {
	name, age := "Shubha", 22;
	fmt.Printf("Hi, I am %s, my age is %d.\n", name, age);
	name, age = "Broto", 23;
	fmt.Printf("Now, I am %s, my age is %d.\n", name, age);
	fmt.Printf("My age as float is %.2f\n", float32(age));

}