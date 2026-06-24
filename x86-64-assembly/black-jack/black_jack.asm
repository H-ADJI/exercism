	C2  equ 2
	C3  equ 3
	C4  equ 4
	C5  equ 5
	C6  equ 6
	C7  equ 7
	C8  equ 8
	C9  equ 9
	C10 equ 10
	CJ  equ 11
	CQ  equ 12
	CK  equ 13
	CA  equ 14

	TRUE  equ 1
	FALSE equ 0

	section .text

	global value_of_card

value_of_card:
	cmp rdi, C10
	jbe .c10
	cmp rdi, CA
	je  .ca
	;   face cards J/Q/K = 10
	mov rax, 10
	ret

.ca:
	mov rax, 1
	ret

.c10:
	mov rax, rdi
	ret

global higher_card

higher_card:
	;    Get values first
	push rdi
	push rsi
	call value_of_card
	mov  rcx, rax; rcx = value of card_one
	pop  rsi
	pop  rdi
	push rdi
	push rsi
	mov  rdi, rsi
	call value_of_card
	mov  rdx, rax; rdx = value of card_two
	pop  rsi
	pop  rdi

	cmp rcx, rdx
	je  .eq
	ja  .g
	;   .l: card_two higher
	mov rax, rsi
	mov rdx, 0
	ret

.g:
	mov rax, rdi
	mov rdx, 0
	ret

.eq:
	mov rax, rdi
	mov rdx, rsi
	ret

global value_of_ace

value_of_ace:
	push rdi
	push rsi
	call value_of_card
	mov  rcx, rax; value of card_one
	pop  rsi
	pop  rdi
	push rdi
	push rsi
	mov  rdi, rsi
	call value_of_card
	mov  rdx, rax; value of card_two
	pop  rsi
	pop  rdi

	;   Check if either card is an ace
	cmp rdi, CA
	je  .ace_in_hand
	cmp rsi, CA
	je  .ace_in_hand

	;   Sum current hand
	add rcx, rdx
	;   If sum + 11 <= 21, return 11, else 1
	add rcx, 11
	cmp rcx, 21
	ja  .return_one
	mov rax, 11
	ret

.ace_in_hand:
	mov rax, 1
	ret

.return_one:
	mov rax, 1
	ret

global is_blackjack

is_blackjack:
	;   Check if one card is ace and other is ten-value
	mov rax, FALSE

	;   Check if card_one is ace
	cmp rdi, CA
	je  .check_ten_two

	;   Check if card_two is ace
	cmp rsi, CA
	je  .check_ten_one

	ret

.check_ten_two:
	;   card_one is ace, check if card_two is 10, J, Q, K
	cmp rsi, C10
	je  .blackjack
	cmp rsi, CJ
	je  .blackjack
	cmp rsi, CQ
	je  .blackjack
	cmp rsi, CK
	je  .blackjack
	ret

.check_ten_one:
	;   card_two is ace, check if card_one is 10, J, Q, K
	cmp rdi, C10
	je  .blackjack
	cmp rdi, CJ
	je  .blackjack
	cmp rdi, CQ
	je  .blackjack
	cmp rdi, CK
	je  .blackjack
	ret

.blackjack:
	mov rax, TRUE
	ret

global can_split_pairs

can_split_pairs:
	;    Get values of both cards
	push rdi
	push rsi
	call value_of_card
	mov  rcx, rax
	pop  rsi
	pop  rdi
	push rdi
	push rsi
	mov  rdi, rsi
	call value_of_card
	mov  rdx, rax
	pop  rsi
	pop  rdi

	cmp rcx, rdx
	je  .can_split
	mov rax, FALSE
	ret

.can_split:
	mov rax, TRUE
	ret

global can_double_down

can_double_down:
	;    Get values of both cards
	push rdi
	push rsi
	call value_of_card
	mov  rcx, rax
	pop  rsi
	pop  rdi
	push rdi
	push rsi
	mov  rdi, rsi
	call value_of_card
	mov  rdx, rax
	pop  rsi
	pop  rdi

	add rcx, rdx
	cmp rcx, 9
	jb  .no
	cmp rcx, 11
	ja  .no
	mov rax, TRUE
	ret

.no:
	mov rax, FALSE
	ret

%ifidn  __OUTPUT_FORMAT__, elf64
section .note.GNU-stack noalloc noexec nowrite progbits
%endif
