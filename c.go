package main

/*
#include "c.c"
void loop(){
	int i;
	for(i=0; i<10; i++)
		printf("%d ", i);
	printf("%d\n");
}
*/
import "C"
import "fmt"

func main() {
	x := C.Add(4, 5)
	y := C.Subtract(100, 50)

	fmt.Println(x, y)

	C.loop()
}
