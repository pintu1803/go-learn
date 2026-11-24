/*

Question 4/5 — Cumulative Membership Benefits
Difficulty: Medium

Concepts: Expression switch, fallthrough, loops.

An e-commerce platform provides cumulative benefits based on customer membership tiers.
Membership				Benefits
Platinum				Priority support, free shipping, reward points
Gold					Free shipping, reward points
Silver					Reward points
Basic					No benefits
Anything else			Invalid membership

Starter code
---------------------

package main

import "fmt"

func main() {
    memberships := []string{
        "Platinum",
        "Gold",
        "Silver",
        "Basic",
        "Unknown",
    }

    // Write your implementation here.
}


Expected output
-----------------
Platinum:
Priority support
Free shipping
Reward points

Gold:
Free shipping
Reward points

Silver:
Reward points

Basic:
No benefits

Unknown:
Invalid membership

Constraints
- Use a for loop.
- Use an expression switch.
- Use fallthrough to implement cumulative benefits.
- Each benefit's print statement must appear exactly once in the source code.
- Do not use additional functions.

*/

package main

import "fmt"

func main() {
    memberships := []string{
        "Platinum",
        "Gold",
        "Silver",
        "Basic",
        "Unknown",
    }

    // Write your implementation here.

    for _, membership := range memberships {

    	fmt.Println(membership, ":")

    	switch membership {
    		case "Platinum":
    			fmt.Println("Priority support")
    			fallthrough

    		case "Gold":
    			fmt.Println("Free shipping")
    			fallthrough

    		case "Silver":
    			fmt.Println("Reward points")

    		case "Basic":
    			fmt.Println("No benefits")

    		default:
    			fmt.Println("Invalid membership")
    	}

    	fmt.Println()
    }
}
