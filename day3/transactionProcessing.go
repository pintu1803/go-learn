/*

Question 3/5 — Transaction Processing
Difficulty: Medium–Hard

Concepts: Labelled break, nested loops, expressionless switch, continue.

A payment service receives transactions in batches.
batches := [][]int{
    {100, 200, -50, 300},
    {400, 0, 500},
    {600, 700, -100},
    {800, 9999, 900},
    {1000, 1100},
}

Business rules
Transaction value					Action
Negative							Skip the transaction
Zero								Stop processing the current batch only
9999								Stop processing all batches immediately
Any other positive value			Process the transaction

Calculate the total processed amount, number of successful transactions and number of skipped negative transactions.

Expected output
Successful transactions: 7
Skipped transactions: 2
Total processed amount: 3100


Constraints
- Use nested for loops.
- Use an expressionless switch.
- Use labelled break where necessary.
- Do not use additional functions or boolean flags.

*/


package main

import "fmt"

func main() {

	batches := [][]int{
	    {100, 200, -50, 300},
	    {400, 0, 500},
	    {600, 700, -100},
	    {800, 9999, 900},
	    {1000, 1100},
	}

	skipCount := 0
	processCount := 0
	totalAmount := 0

	OuterLoop:
	for i := range batches {

		InnerLoop:
		for _, txn := range batches[i] {

			switch {

				case txn < 0: {
					skipCount++
					continue
				}

				case txn == 0:
					break InnerLoop
					//otherwise, it just break switch

				case txn == 9999:
					break OuterLoop
			}

			processCount++
			totalAmount += txn

		}
	}

	fmt.Println("Successful transactions: ", processCount)
	fmt.Println("Skipped transactions: ", skipCount)
	fmt.Println("Total processed amount: ", totalAmount)

}