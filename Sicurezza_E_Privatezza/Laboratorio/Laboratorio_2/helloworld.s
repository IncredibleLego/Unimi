.intel_syntax noprefix

.global _start
.section .text

_start:
    mov rax, 1
    mov rdi, 1
    lea rsi, message
    mov rdx, 13
    syscall

    mov rax, 60
    mov rdi, 0
    syscall

message:
    .ascii "Hello, World\n"
