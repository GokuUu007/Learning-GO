// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"errors"
// )


// var reader = bufio.NewReader(os.Stdin)


// func isPrime(n int) bool {
// 	if n <= 1 {
// 		return false
// 	}

// 	for i := 2; i * i <= n; i++ {
// 		if n % i == 0 {
// 			return false
// 		}
// 	}

// 	return true 
// }

// func min_max(a int, b int, c int) (int, int) { 
// 	return min(a, b, c), max(a, b, c)
// }

// func celsius_to_fahrenheit(c float64) (float64, error) {
// 	if c < -273.15 {
// 		return 0, errors.New("Temperature below absolute zero is not possible") 
// 	}
// 	return (c * 9 / 5) + 32, nil 
// }

// func main() {
// 	// num, err := reader.ReadString('\n')
// 	// if err != nil {
// 	// 	fmt.Print(err)
// 	// 	return 
// 	// }
// 	// fmt.Print(num)
// 	// num = strings.TrimSpace(num)
// 	// num2, err2 := strconv.ParseInt(num, 10, 64)
// 	// if err2 != nil {
// 	// 	fmt.Print(err2)
// 	// 	return
// 	// }
// 	// fmt.Println(num2);
// 	// fmt.Printf("%T\n", num2) 
// 	// pointers 	
// 	// var one int = 2 
// 	// var ptr *int
// 	// fmt.Println(ptr)
// 	// var ptr = &one 
// 	// fmt.Println(ptr)
// 	// fmt.Println("value =", *ptr) // similar to c++ mostly
// 	// slices
// 	// var a_list = []string{"abc", "def", "ghi", "jkl", "mno", "pqr"} 

// 	// var index int = 2
// 	// a_list = append(a_list[:index],  a_list[index + 2:]...) 

// 	// fmt.Println(a_list)

// 	langs := make(map[string]string) 
	
// 	langs["js"] = "javascript"
// 	langs["py"] = "python"
// 	langs["html"] = "hypertext markup language"
// 	langs["rb"] = "ruby"

// 	// fmt.Println(langs)
// 	// delete(langs, "py")
// 	// fmt.Println(langs)

// 	// for key, value := range langs {
// 	// 	fmt.Println(key, ":", value)
// 	// }

// 	// for i := 1; i <= 20; i++ {
// 	// 	if i % 3 == 0 && i % 5 == 0 {
// 	// 		fmt.Println("FizzBuzz")
// 	// 	} else if i % 3 == 0 {
// 	// 		fmt.Println("Fizz")
// 	// 	} else if i % 5 == 0 {
// 	// 		fmt.Println("Buzz")
// 	// 	} else {
// 	// 		fmt.Println(i);
// 	// 	}
// 	// }

// 	for i := 1; i <= 50; i++ {
// 		if isPrime(i) {
// 			fmt.Println(i, "is prime")
// 		}
// 	}

// }

// number guessing game 
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"strings"
// )
// func main() {
//     reader := bufio.NewReader(os.Stdin)

//     isPlaying := true

//     for isPlaying {
//         num := rand.Intn(100) + 1
//         attempts := 0
//         guessed := false

//         fmt.Println("I am thinking of a number in range 1-100. Guess it in at most 10 attempts.")

//         for attempts < 10 {
//             input, err := reader.ReadString('\n')
//             if err != nil {
//                 fmt.Println(err)
//                 break
//             }

//             guess, err := strconv.Atoi(strings.TrimSpace(input))
//             if err != nil {
//                 fmt.Println("Invalid input")
//                 continue
//             }

//             attempts++

//             if guess < num {
//                 fmt.Println("Your guess is too low")
//             } else if guess > num {
//                 fmt.Println("Your guess is too high")
//             } else {
//                 fmt.Println("Crazyyyyy!! You guessed it RIGHT")
//                 guessed = true
//                 break
//             }
//         }

//         if !guessed {
//             fmt.Println("Game over! The number was:", num)
//         }

//         fmt.Println("Do you wanna play again? (Y/N)")

//         resp, err := reader.ReadString('\n')
//         if err != nil {
//             fmt.Println(err)
//             break
//         }

//         resp = strings.TrimSpace(resp)

//         if strings.ToUpper(resp) != "Y" {
//             isPlaying = false
//         }
//     }
// }


package main 

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in go lang")
 
	rand.Seed(time.Now().UnixNano())
	diceNumber:= rand.Intn(6) + 1 
	  
	fmt.Println(diceNumber)
	  
	switch diceNumber {
	case 1:
		fmt.Println("Dice Value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 steps")
	case 3:
		fmt.Println("You can move 3 steps")
	case 4:
		fmt.Println("You can move 4 steps")
	case 5:
		fmt.Println("You can move 5 steps")
	case 6:
		fmt.Println("You can move 6 steps")
	default:
		fmt.Println("What was that ?")
	}

}