.intel_syntax noprefix

.global _start
.section .text

_start:
    sub rsp, 8
    mov rax, 0x8899aabbccddeeff
    mov [rsp], rax
    add rsp, 8

    mov rax, 60
    xor rdi, rdi
    syscall
