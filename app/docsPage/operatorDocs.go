package docs

//////////////////////////////////////////////////////////////
// MATH OPERATORS                                           //
//////////////////////////////////////////////////////////////

var (
	add = operator{"Add", "+", 2, "a + b"}
	sub = operator{"Subtract", "-", 2, "a - b"}
	mul = operator{"Multiply", "*", 2, "a * b"}
	div = operator{"Divide", "/", 2, "a / b"}
	mod = operator{"Modulo", "%", 2, "a mod b"}
	pow = operator{"Power", "^", 2, "a ^ b"}
)

//////////////////////////////////////////////////////////////
// COMPARISSON OPERATORS                                    //
//////////////////////////////////////////////////////////////

var (
	less    = operator{"Lesser", "<", 2, "a < b"}
	lesseq  = operator{"Lesser or Equal", "<=", 2, "a <= b"}
	great   = operator{"Greater", ">", 2, "a > b"}
	greateq = operator{"Greater or Equal", ">=", 2, "a >= b"}
	equal   = operator{"Equal", "==", 2, "a == b"}
	diff    = operator{"Different", "!=", 2, "a != b"}
)

//////////////////////////////////////////////////////////////
// LOGICAL OPERATORS                                        //
//////////////////////////////////////////////////////////////

var (
	and = operator{"Logical And", "&&, and", 2, "a and b"}
	or  = operator{"Logical Or", "||, or", 2, "a or b"}
	xor = operator{"Logical Xor", "XX, xor", 2, "a xor b"}
	not = operator{"Logical Not", "!!, not", 1, "not a"}
)

//////////////////////////////////////////////////////////////
// STACK MANIPULATION OPERATORS                             //
//////////////////////////////////////////////////////////////

var (
	swap = operator{"Swap", "swap", 2, "|c|b|a| -> |c|a|b|"}
	drop = operator{"Drop", "drop", 1, "|c|b|a| -> |c|b|"}
	dup  = operator{"Duplicate", "dup", 1, "|c|b|a| -> |c|b|a|a|"}
	dup2 = operator{"Duplicate 2", "dup2", 2, "|c|b|a| -> |c|b|a|b|a|"}
	rot  = operator{"Rotate", "rot", 3, "|c|b|a| -> |b|a|c|"}
	over = operator{"Over", "over", 2, "|c|b|a| -> |c|b|a|b|"}
)

//////////////////////////////////////////////////////////////
// MATH FUNCTION OPERATORS                                  //
//////////////////////////////////////////////////////////////

var (
	sqrt  = operator{"Square Root", "sqrt, √x", 1, "√a"}
	fact  = operator{"Factorial", "fact, x!", 1, "a!"}
	sin   = operator{"Sine", "sin", 1, "sin(a)"}
	cos   = operator{"Cosine", "cos", 1, "cos(a)"}
	tan   = operator{"Tangent", "tan", 1, "tan(a)"}
	exp   = operator{"Exponential", "exp", 1, "exp(a)"}
	log   = operator{"Logarithm", "log", 1, "log(1)"}
	sim   = operator{"Simetric", "sim, ±", 1, "-a"}
	abs   = operator{"Absolute", "abs, |x|", 1, "|a|"}
	ceil  = operator{"Ceil", "ceil, ⌈x⌉", 1, "⌈a⌉"}
	floor = operator{"Floor", "floor, ⌊x⌋", 1, "⌊a⌋"}
	max   = operator{"Maximum", "max, ↑", 2, "max{a, b}"}
	min   = operator{"Minimum", "min, ↓", 2, "min{a, b}"}
)

//////////////////////////////////////////////////////////////
// ALL OPERATORS                                            //
//////////////////////////////////////////////////////////////

var allOps = []operator{
	add, sub, mul, div, mod, pow,
	less, lesseq, great, greateq, equal, diff,
	and, or, xor, not,
	swap, drop, dup, dup2, rot, over,
	sqrt, fact, sin, cos, tan, exp, log, sim, abs, ceil, floor, max, min,
}
