#include <stdio.h>

int Add(int a, int b){
    printf("Adding inside C!\n");

    return a + b;
}

int Subtract(int a, int b){
    printf("Subtracting inside C!\n");
    
    if(a > b)
        return a - b;
    return b -a;
}
