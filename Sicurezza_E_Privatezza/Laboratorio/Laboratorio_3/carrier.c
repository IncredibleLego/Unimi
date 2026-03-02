// Per compilare
// gcc -o carrier carrier.c

#include <sys/mman.h>
#include <unistd.h>
#include <sys/stat.h>
#include <stdio.h>
#include <stddef.h>

int main(void)
{
    // Per testare il carrier come programma setuid, rimuovere la riga sequente commentata e ricompilare
    setreuid(0, 0);
    // Una volta ricompilato, modficare l' owner tramite il comando sudo chown root:root carrier
    // e aggiungere lo sticky bit tramite sudo chmod +s carrier
    // Testare lo shellcode che apre /bin/sh e verificare whoami di essere diventati root a partire dal vostro 
    // normale utente
    
    void *page = mmap(0x1337000, 0x1000,
                      PROT_READ | PROT_WRITE | PROT_EXEC, MAP_PRIVATE | MAP_ANON, 0, 0);
    read(0, page, 0x1000);
    ((void (*)())page)();
    return 0;
}
