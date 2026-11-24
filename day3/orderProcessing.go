/*

You are implementing an order-processing system. You receive a sequence of item prices in rupees.
Requirements
Process the prices sequentially according to these rules:
Condition					Action
Price is negative			Skip the item
Price is zero				Stop processing
Price is below ₹500			No discount
Price is ₹500–999			Apply 10% discount
Price is ₹1,000 or above	Apply 20% discount

After processing, print the number of valid items processed, total original price, total discount and final payable amount.

Sample input

prices := []int{200, 800, -100, 1500, 400, 0, 2000}

Expected output

Items processed: 4
Subtotal: 2900
Discount: 380
Payable amount: 2520

Starter code

package main
import "fmt"
func main() {
    prices := []int{200, 800, -100, 1500, 400, 0, 2000}

    // Write your implementation here.
}

Constraints:
- Use a for loop to process the prices.
- Use continue and break appropriately.
- Use an expressionless switch for discount calculation.
- Use integer arithmetic.
- No additional functions, structs or pointers required.

*/

package main

import "fmt"

func main() {
    prices := []int{200, 800, -100, 1500, 400, 0, 2000}

    // Write your implementation here.
    count := 0
    discount := 0
    totaldiscount := 0
    subtotal := 0
    payable := 0
    temp := 0

    for _, price := range prices {
    	if price < 0 {
    		continue
    	}

    	if price == 0 {
    		break
    	}

    	switch {
    		case price < 500:
    			discount = 0
    		case price <= 999: 
    			discount = 10
    		default:
    			discount = 20
    	}

    	temp = price * discount / 100
    	totaldiscount += temp

    	subtotal += price
    	count++
    }

    payable = subtotal - totaldiscount

    fmt.Println("Valid items processed count : ", count)
    fmt.Println("Total original price : ", subtotal)
    fmt.Println("Total discount : ", totaldiscount)
    fmt.Println("Final payable amount : ", payable)



}