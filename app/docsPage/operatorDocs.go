package docs

//////////////////////////////////////////////////////////////
// MATH OPERATORS                                           //
//////////////////////////////////////////////////////////////

var mathOpDocs = operatorList([]operator{
	{"Add", "+", 2, "a + b"},
	{"Subtract", "-", 2, "a - b"},
	{"Multiply", "*", 2, "a * b"},
	{"Divide", "/", 2, "a / b"},
	{"Modulo", "%", 2, "a mod b"},
	{"Power", "^", 2, "a ^ b"},
})

//////////////////////////////////////////////////////////////
// COMPARISSON OPERATORS                                    //
//////////////////////////////////////////////////////////////

var compOpDocs = operatorList([]operator{
	{"Lesser", "<", 2, "a < b"},
	{"Lesser or Equal", "<=", 2, "a <= b"},
	{"Greater", ">", 2, "a > b"},
	{"Greater or Equal", ">=", 2, "a >= b"},
	{"Equal", "==", 2, "a == b"},
	{"Different", "!=", 2, "a != b"},
})

//////////////////////////////////////////////////////////////
// LOGICAL OPERATORS                                        //
//////////////////////////////////////////////////////////////

var logicOpDocs = operatorList([]operator{
	{"Logical And", "&&, and", 2, "a and b"},
	{"Logical Or", "||, or", 2, "a or b"},
	{"Logical Exclusive Or", "XX, xor", 2, "a xor b"},
	{"Logical Negation", "!!, not", 1, "not a"},
})

//////////////////////////////////////////////////////////////
// STACK MANIPULATION OPERATORS                             //
//////////////////////////////////////////////////////////////

var stackOpDocs = operatorList([]operator{
	{"Swap", "swap", 2, "|c|b|a| -> |c|a|b|"},
	{"Drop", "drop", 1, "|c|b|a| -> |c|b|"},
	{"Duplicate", "dup", 1, "|c|b|a| -> |c|b|a|a|"},
	{"Duplicate 2", "dup2", 2, "|c|b|a| -> |c|b|a|b|a|"},
	{"Rotate", "rot", 3, "|c|b|a| -> |b|a|c|"},
	{"Over", "over", 2, "|c|b|a| -> |c|b|a|b|"},
})

//////////////////////////////////////////////////////////////
// MATH FUNCTION OPERATORS                                  //
//////////////////////////////////////////////////////////////

var funcOpDocs = operatorList([]operator{
	{"Square Root", "sqrt, √x", 1, "√a"},
	{"Factorial", "fact, x!", 1, "a! = a*(a-1)*(a-2)*...*1"},
	{"Sine", "sin", 1, "sin(a)"},
	{"Cosine", "cos", 1, "cos(a)"},
	{"Tangent", "tan", 1, "tan(a)"},
	{"Exponential", "exp", 1, "exp(a) = e ^ a"},
	{"Logarithm (base e)", "log", 1, "log(1)"},
	{"Simetric", "sim, ±", 1, "-a"},
	{"Absolute Value", "abs, |x|", 1, "|a|"},
	{"Ceil", "ceil, ⌈x⌉", 1, "⌈a⌉"},
	{"Floor", "floor, ⌊x⌋", 1, "⌊a⌋"},
	{"Maximum", "max, ↑", 2, "max{a, b}"},
	{"Minimum", "min, ↓", 2, "min{a, b}"},
})
