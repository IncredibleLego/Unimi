.intel_syntax noprefix
.global _start
.section .text

_start:
    # Open File
    mov rax, 2
    lea rdi, [rip + file]
    xor rsi, rsi
    xor rdx, rdx
    syscall

    # Output file content
    mov rsi, rax
    mov rax, 28
    mov rdi, 1
    mov rdx, 0
    mov r10, 100
    syscall

    # Exit (opzionale)
    mov rax, 60
    mov rdi, 0
    syscall

file:
    .string "flag"
