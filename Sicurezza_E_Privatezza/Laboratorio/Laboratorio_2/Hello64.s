.intel_syntax noprefix

.global _start

.section .text

_start:
  mov rax, 1            # write(
  mov rdi, 1            #   STDOUT_FILENO,
  lea rsi, message     	#   "Hello, world!\n",
  mov rdx, 13           #   sizeof("Hello, world!\n")
  syscall               # );

  mov rax, 60       	  # exit(
  mov rdi, 0        	  #   EXIT_SUCCESS
  syscall           	  # );
message:
  .ascii  "Hello, world\n"


# To compile:
# gcc -nostdlib -static -o Hello64 Hello64.s

# Syscall doc:
# https://syscalls64.paolostivanin.com/
# https://x64.syscall.sh/
