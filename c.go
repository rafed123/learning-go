package main

// #include <stdio.h>
// void loop(){
// 	int i;
// 	for (i=0; i<10; i++){
// 		printf("%d\n", i);
// 	}
// }
import "C"

func main() {
	C.loop()
}
