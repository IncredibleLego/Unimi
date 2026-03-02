.intel_syntax noprefix

.global _start
.section .text

_start:
    xor rax, rax		    # sys_read(
    mov rdi, 0 			    #   STDIN_FILENO,
    lea rsi, pBuffer		#   char* buf,
    mov edx, [iBufferSize]	#   buflen)
    syscall

    # Move the number of read characters from rax
    # to rbx (4th syscall parameter)
    mov rbx, rax

    mov rax, 1			    # sys_write(
    mov rdi, 1			    #   STDOUT_FILENO
    lea rsi, pBuffer		#   pBuffer
    mov rdx, rbx		    #   readLength);
    syscall

    mov rax, 60			# sys_exit(
    xor rdi, rdi		#   EXIT_SUCCESS
    syscall			    # );

# We need a writeable section to hold our input data
.section .data
iBufferSize:
    .int 100
pBuffer:
    .space 100
    .ascii "\n"
