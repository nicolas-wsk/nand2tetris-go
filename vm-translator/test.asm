






@17
D=A
@SP
A=M
M=D
@SP
M=M+1

@17
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
AM=M-1
D=M
A=A-1
D=M-D
@EQ.true.1
D;JEQ
@SP
A=M-1
M=0
@EQ.after.1
0;JMP
(EQ.true.1)
@SP
A=M-1
M=-1
(EQ.after.1)

@17
D=A
@SP
A=M
M=D
@SP
M=M+1

@16
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
AM=M-1
D=M
A=A-1
D=M-D
@EQ.true.2
D;JEQ
@SP
A=M-1
M=0
@EQ.after.2
0;JMP
(EQ.true.2)
@SP
A=M-1
M=-1
(EQ.after.2)

@16
D=A
@SP
A=M
M=D
@SP
M=M+1

@17
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
AM=M-1
D=M
A=A-1
D=M-D
@EQ.true.3
D;JEQ
@SP
A=M-1
M=0
@EQ.after.3
0;JMP
(EQ.true.3)
@SP
A=M-1
M=-1
(EQ.after.3)

@892
D=A
@SP
A=M
M=D
@SP
M=M+1

@891
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
AM=M-1
D=M
A=A-1
D=M-D
@LT.true.4
D;JLT
@SP
A=M-1
M=0
@LT.after.4
0;JMP
(LT.true.4)
@SP
A=M-1
M=-1
(LT.after.4)

@891
D=A
@SP
A=M
M=D
@SP
M=M+1

@892
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
AM=M-1
D=M
A=A-1
D=M-D
@LT.true.5
D;JLT
@SP
A=M-1
M=0
@LT.after.5
0;JMP
(LT.true.5)
@SP
A=M-1
M=-1
(LT.after.5)

@891
D=A
@SP
A=M
M=D
@SP
M=M+1

@891
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
AM=M-1
D=M
A=A-1
D=M-D
@LT.true.6
D;JLT
@SP
A=M-1
M=0
@LT.after.6
0;JMP
(LT.true.6)
@SP
A=M-1
M=-1
(LT.after.6)

@32767
D=A
@SP
A=M
M=D
@SP
M=M+1

@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
AM=M-1
D=M
A=A-1
D=M-D
@GT.true.7
D;JGT
@SP
A=M-1
M=0
@GT.after.7
0;JMP
(GT.true.7)
@SP
A=M-1
M=-1
(GT.after.7)

@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

@32767
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
AM=M-1
D=M
A=A-1
D=M-D
@GT.true.8
D;JGT
@SP
A=M-1
M=0
@GT.after.8
0;JMP
(GT.true.8)
@SP
A=M-1
M=-1
(GT.after.8)

@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
AM=M-1
D=M
A=A-1
D=M-D
@GT.true.9
D;JGT
@SP
A=M-1
M=0
@GT.after.9
0;JMP
(GT.true.9)
@SP
A=M-1
M=-1
(GT.after.9)

@57
D=A
@SP
A=M
M=D
@SP
M=M+1

@31
D=A
@SP
A=M
M=D
@SP
M=M+1

@53
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
AM=M-1
D=M
A=A-1
M=D+M

@112
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
AM=M-1
D=M
A=A-1
M=M-D

@SP
A=M-1
M=-M
@SP
AM=M-1
D=M
A=A-1
M=D&M

@82
D=A
@SP
A=M
M=D
@SP
M=M+1

@SP
AM=M-1
D=M
A=A-1
M=D|M

@SP
A=M-1
M=!M

