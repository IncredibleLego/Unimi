// Programma vulterabile a code injection
// Per compilare
// gcc -g -z execstack -o first_injection first_injection.c 2> /dev/null

#include <stdio.h>
#include <time.h>

void bye1() { puts("Goodbye!"); }
void bye2() { puts("Farewell!"); }

void hello(char *name, void (*bye_func)())
{
    printf("Hello, %s!\n", name);
    bye_func();
}
int main(int argc, char **argv)
{
    char name[1024];
    gets(name);

    srand(time(0));
    if (rand() % 2)
        hello(bye1, name);
    else
        hello(name, bye2);
}