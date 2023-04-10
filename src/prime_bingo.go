// My first program will print the Numbers from 1 to a chosen parameter and replace all primes with the word PRIME.

package main

import "fmt"
import "math/big"


func prime_bingo(last_int int64) {
  var i int64 = 1
  for {
    if big.NewInt(i).ProbablyPrime(0) {
     fmt.Println("PRIME")
    } else if i==1 {
      fmt.Println("PRIME")
      } else {
        fmt.Println(i)
        }
    i = i+1
    if i>last_int {
      break 
      }
    }  
 }


func main() {    
    prime_bingo(100)
}
