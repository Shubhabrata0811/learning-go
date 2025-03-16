package main

// import "fmt";
// import "path";

import ("fmt"; "path");

func main(){
	var dir, file string;
	dir, file = path.Split("path-separator/main.go");
	fmt.Println("Dir:", dir);
	fmt.Println("ile:", file);
}