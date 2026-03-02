.intel_syntax noprefix

.global _start

.section .text

_start:
    # Leggiamo il primo numero
    call ReadNumber
    mov r12, rax

    # leggiamo il secondo numero
    call ReadNumber
    mov r13, rax

    # Sommiamo
    mov rdi, r12
    add rdi, r13

    # Chiamiamo la funzione per la stampa
    # il numero da stampare è in RDI
    call PrintNumber

    # Usciamo dal nostro programma
    call Exit;

# VOID ReadNumber()
ReadNumber:
    push r12
    push r13

    lea r12, pInputBuffer
    mov r13, 30

    mov rdi, r12
    mov rsi, r13

    call ReadLine

    mov rdi, r12
    mov rsi, rax
    dec rsi
    call ParseNumber

    pop r13
    pop r12
    ret

# /************************************/
# Converte una stringa in numero intero positivo
# /************************************/
# INT ParseNumber(CHAR* buffer, INT length)
ParseNumber:
    # Azzeriamo RAX per salvare il risultato
    # construito man mano
    xor rax, rax

_parseLoop:
    test rsi, rsi
    jz _parsePrologue

    imul rax, 10
    mov bl, byte ptr [rdi]
    sub rbx, '0'
    add rax, rbx
    inc rdi
    dec rsi
    jmp _parseLoop
_parsePrologue:
    ret

# /************************************/
# Stampa un numero passato come input a schermo
# /************************************/
# VOID PrintNumber(INT number)

PrintNumber:
    # Prologo
    push rbx

    # R9 = Contatore per il numero di caratteri scritti nel buffer
    xor r9, r9

    # Carichiamo in un registro il puntatore a un buffer
    # dove scrivere la nostra stringa e la sua dimensione massima
    lea rsi, pPrintBuffer
    mov ecx, dword ptr [iPrintBufferSize]

    # ci spostiamo all' ultimo elemento del buffer per scrivere
    # la prima cifra a partire da destra
    dec rcx
    add rsi, rcx

    # Carichiamo in RAX il numero da stampare
    # e il RBX il divisore
    mov rax, rdi
    mov rbx, 10

    # Gestiamo lo zero come caso particolare
    test rax, rax
    jz _handleZero

    # Scriviamo nel buffer il carattere "a capo"
    mov byte ptr [rsi], '\n'
    dec rsi
    inc r9

_printLoop:
    # In RDX sono contenuti i 64 bit più alti del dividendo (0) e
    # il resto della divisione dopo la chiamata a IDIV
    xor rdx, rdx
    idiv rbx

    # Il resto deve diventare una cifra: ci aggiungo zero
    # e lo salvo nel buffer, modificando poi puntatore e conteggio
    add rdx, '0'
    mov byte ptr [rsi], dl
    inc r9
    dec rsi

    # Se il risultato della divisione NON è zero, ciclo
    test rax, rax
    jnz _printLoop

    # Ho finito: sposto il puntatore sulla prima cifra da sinistra e
    # stampo
    inc rsi
    mov rdi, rsi
    mov rsi, r9
    call Print

    jmp _PrintNumberEpilogue

_handleZero:
    lea rdi, zeroCharacter
    mov rsi, 2
    call Print
_PrintNumberEpilogue:
    pop rbx
    ret

# VOID Print(CHAR* Buffer, INT Length)
Print:
    mov r8, rdi
    mov r9, rsi

    mov rax, 1                  # sys_write(
    mov rdi, 1                  # STDOUT_FILENO,
    mov rsi, r8		        # Buffer,
    mov rdx, r9                 # Length);
    syscall
    ret

# VOID ReadLine(CHAR* Buffer, INT Length)
ReadLine:
    mov r8, rdi
    mov r9, rsi

    xor rax, rax                # sys_read(
    mov rdi, 0                  # STDIN_FILENO,
    mov rsi, r8		        # Buffer,
    mov rdx, r9                 # Length);
    syscall
    ret

# VOID Exit()
Exit:
    mov rax, 60                 # sys_exit(
    xor rdi, rdi                # EXIT_SUCCESS
    syscall                     # );

.section .data

pInputBuffer:
    .space 30

iPrintBufferSize:
    .int 30
pPrintBuffer:
    .space 30

zeroCharacter:
    .ascii "0\n"
