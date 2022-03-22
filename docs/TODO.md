- [X] Stack
- [X] Lexer && Tokens
- [X] Parser
- [X] AST
- [X] Interpreter
- [X] Colorize output
- [X] Repl
- [X] Src file interpreter
- [X] Src Files as repl functions
- [X] Project creator from the cli
- [X] UI
- [ ] Write Documentation
- [ ] Use files as functions inside the main
- [ ] Run Main file
- [ ] Example Function Programs
- [ ] Function Documentation syntax
- [ ] Syntax highlight VsCode, Vim




```es
( push 0 to the stack as a "default acomulator" and swap it with k
  |k| -> |k|0| -> |0|k| )
0 swap 

( repeat k times )
for {
  ( add 1 to the acomulator and duplicate it
    |0| -> |0|1| -> |1| -> |1|1| ...
    |1|2|2| -> |1|2|2|1| -> |1|2|3| -> |1|2|3|3| ... )
  1 + dup
}

( at the end we will get a duplicate number, so we drop it
  |1|2|...|k-1|k|k| -> |1|2|...|k-1|k| )
drop
```

```es
( |a|b| -> |a|b|a|b| -> |a|b|b|a| )
dup2
swap

( |a|b|b|a| -> |a|b|a>b| -> |a|b|1| -> |b|a|
                         -> |a|b|0| -> |a|b| )
> if { 
  swap 
}{}
```

```es
( |a|b| -> |a|b|a|b| -> |a|b|b|a| )
dup2
swap

( |a|b|b|a| -> |a|b|a<b| -> |a|b|1| -> |b|a|
                         -> |a|b|0| -> |a|b| )
< if { 
  swap 
}{}
```
