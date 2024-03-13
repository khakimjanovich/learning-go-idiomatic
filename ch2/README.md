# Chapter 2

1. Go has many types built into the language. Those are called *predeclared types*
2. Go assigns a default *zero value* to any variable that is declared but not assigned a value. Having an explicit zero
   value makes code clearer and removes a source of bugs fund in C and C++ programs.

## Literals

1. When storage is allocated for a variable, either through a declaration or a call of new, or when a new value is 
   created, either through a composite literal or a call of make, and no explicit initialization is provided, the 
   variable or value is given a default value. Each element of such a variable or value is set to the zero value for its
   type: false for booleans, 0 for numeric types, "" for strings, and nil for pointers, functions, interfaces, slices, 
   channels, and maps. This initialization is done recursively, so for instance each element of an array of structs will
   have its fields zeroed if no value is specified.
2. a Go *literal* is an explicitly specified number, character, or string. Go programs have four common and one rare
   kind of literals. 
3. An *integer literal* is a sequence of numbers. Integer literals as base 10 by default, but different prefixes are
   used to indicate other bases: 0b for binary (base 2), 0o for octal (base 8), or 0x for hexadecimal (base 16). A
   leading 0 with no letter after it is another way to represent an octal literal. Do not use it, as it is confusing!
4. Go allows you put underscores in the middle of your literal. This allows, group by thousands in base 10 (1_234).
   The only limitations on underscore are that they cannot be at the beginning or end of numbers, and you can't have 
   them next to each other.
5. A *floating-point literal* has a decimal point to indicate the fractional portion of the value. They can also have
   an exponent specified with the letter **e** and a positive and a negative number(such as 6.03e23). You also have the 
   option to write them in hexadecimal by using the 0x prefix and the letter p for indicating any exponent (0x12.34p5, 
   which is equal to 582.5 in base 10). As integer literals, you can use underscores to format your floating-point 
   literals.
6. A *rune literal* represents a character and is surrounded by single quotes. Unlike many other languages, in Go single 
   quotes and double quotes are *not* interchangeable. Rune literals can be written as single Unicode characters ('a'), 
   8-bit octal numbers ('\141'), 8-bit hexadecimal numbers ('\x61'), 16-bit hexadecimal numbers ('\u0061'), or 32-bit 
   Unicode numbers ('\U00000061'). There are also several backslash-escaped rune literals, with the most useful ones 
   being newline ('\n'), tab ('\t'), single quote ('\''), and backslash ('\\').
7. Use base 10 to represent your integer and floating-point literals.Octal representations are rare, mostly used to 
   represent POSIX permission flag values (such as 0o777 for rwxrwxrwx). Hexadecimal and binary are sometimes used for 
   bit filters or networking and infrastructure applications. Avoid using any of the numeric escapes for rune literals, 
   unless the context makes your code clearer.
8. There are two ways to indicate *string literals*. Most of the time, you should use double quotes to create an 
   *interpreted string literal* (e.g., type "Greetings and Salutations"). These contain zero or more rune literals. They 
   are called “interpreted” because they interpret rune literals (both numeric and backslash escaped) into single 
   characters.
   ```
   One rune literal backslash escape is not legal in a string literal: the
   single quote escape. It is replaced by a backslash escape for double
   quotes.
   ```
9. The only characters that cannot appear in an interpreted string literal are unescaped backslashes, unescaped 
   newlines, and unescaped double quotes.
10. Literals are considered untyped.
11. In case you were wondering what the fifth kind of primitive literal was, Go supports imaginary literals to represent
    the imaginary portion of a complex number. They look just like floating-point literals, but they have an i for a 
    suffix.

## Types

1. The bool type represents Boolean values. Variables of bool type can have one of two values: true or false. The zero
   value for boolean is false:
    ```
    var flag bool
    var isAwesome = true
    ```
2. Go has a large number of numeric types: 12 types (and a few special names) that are grouped into three categories.
3. Choosing which integer to use:
   * If you are working with a binary file format or network protocol that has an integer of a specific size or sign, 
     use the corresponding integer type.
   * If you are writing a library function that should work with any integer type, take advantage of Go’s generics 
     support and use a generic type parameter to represent any integer type.
   * In all other cases, just use **int**.
4. Go integers support the usual arithmetic operators: +, -, *, /, with % for modulus.
5. You can combine any of the arithmetic operators with = to modify a variable: +=, -=,
   *=, /=, and %=. For example, the following code results in x having the value 20:
   ```
   var x int = 10
   x *= 2 
   ```
6. You compare integers with ==, !=, >, >=, <, and <=.
7. Go also has bit-manipulation operators for integers. You can bit shift left and right with << and >>, or do bit masks
   with & (bitwise AND), | (bitwise OR), ^ (bitwise XOR), and &^ (bitwise AND NOT). As with the arithmetic operators, 
   you can also combine all the bitwise operators with = to modify a variable: &=, |=, ^=, &^=, <<=, and >>=.
8. Go has 2 floating types
9. You can use all the standard mathematical and comparison operators with floats, except %. Floating-point division has 
   a couple of interesting properties. Dividing a nonzero floating-point variable by 0 returns +Inf or -Inf (positive or
   negative infinity), depending on the sign of the number. Dividing a floating-point variable set to 0 by 0 returns NaN
   (Not a Number).
10. While Go lets you use == and != to compare floats, don’t do it. Because of the inexact nature of floats, two 
    floating-point values might not be equal when you think they should be. Instead, define a maximum allowed variance 
    and see if the difference between two floats is less than that. This value (sometimes called epsilon) depends on 
    your accuracy needs;
11. Go includes strings as a built-in type. The zero value for a string is the empty string. Like integers and floats, 
    strings are compared for equality using ==, difference with !=, or ordering with >, >=, <, or <=. They are 
    concatenated by using the + operator.
12. Strings in Go are immutable; you can reassign the value of a string variable, but you cannot change the value of the
    string that is assigned to it.
13. Go also has a type that represents a single code point. The rune type is an alias for the int32 type, just as byte 
    is an alias for uint8. As you could probably guess, a rune literal’s default type is a rune, and a string literal’s 
    default type is a string. If you are referring to a character, use the rune type, not the int32 type. They might
    be the same to the compiler, but you want to use the type that clarifies the intent of your code:
    ```
    var myFirstInitial rune = 'J' // good - the type name matches the usage
    var myLastInitial int32 = 'B' // bad - legal but confusing
    ```
    