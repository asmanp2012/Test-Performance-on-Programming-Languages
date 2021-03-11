package main

import (
	/*"context"*/
	"fmt"
	"time"
)
var count int64 = 0 
func delta() int64 {
	return   10 * 15 / 20  * 40
}
func main() {
	
	fmt.Println("Hi")
	var StartTime = time.Now()
	var duration int64 = 0
	for ; duration < 1000; count+= 100000 {
		for i := 0; i < 100000; i++ {
			delta()
		}
		duration = time.Since(StartTime).Milliseconds()
	}
	/* 9223372036854775807 */
	/* 1257894000000       */
	/* 99999999999999999 */
	/* 1000000000        */
	fmt.Println(count)
	/*fmt.Println(time.Now().Sub(StartTime).Milliseconds() )*/
}