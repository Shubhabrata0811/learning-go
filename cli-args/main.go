package main
import (
	"fmt"
	"os"
)

func main()  {
	fmt.Printf("%#v\n",os.Args)
	fmt.Println("Number of cli arguments:", len(os.Args))
	fmt.Printf("Path: %s\n", os.Args[0])
	fmt.Printf("1st argument: %s\n", os.Args[1])
	fmt.Printf("2nd argument: %s\n", os.Args[2])
}