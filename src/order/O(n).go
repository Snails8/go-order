package order

import "fmt"

func On(num int32)  {
	
	for i := 0; i < 100000; i++ {
		if i == int(num) {
			fmt.Println(i)
			break;
		}
	}
	
	return 
}