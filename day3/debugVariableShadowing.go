/*
Question 5/5 — Variable shadowing and scope
Difficulty: Medium–Hard

Concepts: Variable shadowing, :=, assignment, if initialization statements and loops.

An order service calculates the final payable amount for each order.

The following code compiles, but produces incorrect results.


package main

import "fmt"

func main() {
    prices := []int{400, 1200, 800, 2000}

    totalPayable := 0

    for _, price := range prices {
        payable := price

        if discount := price * 10 / 100; price >= 1000 {
            payable := price - discount
            fmt.Println("Discounted price:", payable)
        }

        totalPayable := totalPayable + payable

        fmt.Println("Running total:", totalPayable)
    }

    fmt.Println("Final payable:", totalPayable)
}


Business requirements
Orders worth ₹1,000 or more receive a 10% discount. All other orders receive no discount.

Expected output
---------------------
Running total: 400
Discounted price: 1080
Running total: 1480
Running total: 2280
Discounted price: 1800
Running total: 4080
Final payable: 4080

Your tasks
1. Predict the actual output of the original code.
2. Identify the bugs and explain why they occur.
3. Correct the code with minimal changes.

*/

package main

import "fmt"

func main() {
    prices := []int{400, 1200, 800, 2000}

    totalPayable := 0

    for _, price := range prices {
        payable := price

        if discount := price * 10 / 100; price >= 1000 {
            payable = price - discount
            fmt.Println("Discounted price:", payable)
        }

        totalPayable = totalPayable + payable

        fmt.Println("Running total:", totalPayable)
    }

    fmt.Println("Final payable:", totalPayable)
}

//Changes made:
// 1. Removed the := operator in the if statement to avoid shadowing the payable variable.
// 2. Removed the := operator when updating totalPayable to ensure it updates the existing variable instead of creating a new one.