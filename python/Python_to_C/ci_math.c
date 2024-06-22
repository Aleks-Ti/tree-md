#include <stdio.h>

int add(int a, int b)
{
    return a + b;
}

/*
gcc -shared -o ci_math.so ci_math.c
скомпилировать под убунту
*/