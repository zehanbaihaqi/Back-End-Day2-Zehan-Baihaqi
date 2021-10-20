
package main
 

import "fmt"

 

func playWithAsterik(n int) {

  i:= 0
  j:= 0

    for i = 1; i <= n; i++ {
        for j = 1; j <= n-i; j++ {
            fmt.Print(" ")
        }
        for k := 0; k != i; k++ {
            fmt.Printf(" *")
        }
        fmt.Println()
    }

}

 

func main() {

 playWithAsterik(5)

 /*

       *

      * *

     * * *

    * * * *

   * * * * *

 */

}