/*

package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3, 4, 7, 9, 11, -5, 15, 17, 21, 23}

    // Write your implementation here.
}

Constraints
- Skip numbers smaller than 2.
- Use nested for loops.
- Use labelled continue to skip composite numbers.
- Do not use an isPrime boolean variable.
- Do not create additional functions.
- Optimize the divisor-checking loop.

*/

package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3, 4, 7, 9, 11, -5, 15, 17, 21, 23}

    // Write your implementation here.
    count := 0
    sum := 0

    OuterLoop:
    for i := range numbers {
    	num := numbers[i]

    	if num < 2 {
    		continue
    	}

    	for j := range num/2 + 1 {
    		if j <= 1 {
    			continue
    		}

    		if num % j == 0 {
    			continue OuterLoop
    			//Instead of break inner loop
    		}
    	}

    	count++
    	sum += num

    	fmt.Println("Prime : ", num)
    }

    fmt.Println("Prime count : ", count)
    fmt.Println("Prime sum : ", sum)


}