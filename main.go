package main

import "fmt"

func main() {
	var (
		i      int     = 21
		j      bool    = true
		russia int     = '\u042F'
		k      float64 = 123.456
	)

	fmt.Printf("Nilai i = %v \n", i)
	fmt.Printf("Type data i = %T \n", i)
	fmt.Printf("Symbol = %% \n")
	fmt.Printf("Nilai boolean j = %t \n", j)
	fmt.Printf("Nilai boolean j = %t \n", j)
	fmt.Printf("Unicode russia = %c \n", russia)
	fmt.Printf("Nilai base 10 = %d \n", i)
	fmt.Printf("Nilai base 8 = %o \n", i)
	fmt.Printf("Nilai base 16 = %x \n", 15)
	fmt.Printf("Nilai base 16 = %X \n", 15)
	fmt.Printf("Unicode char Я = %U \n", russia)
	fmt.Printf("Float = %.6f \n", k)
	fmt.Printf("Float Scientific = %E \n", k)
}
