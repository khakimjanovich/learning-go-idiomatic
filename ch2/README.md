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
14. Go doesn't allow *automatic type promotion* between variables. Type conversion must be used when variable types do
    not match. This makes it clear exactly what type you want without having to memorize any type conversion rules
15. In Go, no other type can be converted to a bool, implicitly or explicitly. You can do it by comparison operators
16. Type conversions are one of the places where Go chooses to add a little verbosity in exchange for a great deal of 
    simplicity and clarity. 
17. A literal string cannot be assigned to a variable with a numeric type or a literal number to a string variable, nor 
    can you assign a float literal to an int. These are all flagged by the compiler as errors. Size limitations also 
    exist; while you can write numeric literals that are larger than any integer can hold, it is a compile-time error to
    try to assign a literal whose value overflows the specified variable, such as trying to assign the literal 1000 to a
    variable of type byte.

## var Versus :=

1. The most verbose way to declare a variable in Go 
    ```go 
        var x int = 10
   ```
2. Default type of an integer literal is **int**.  
    ```go
        var x = 10
    ```
3. We can first declare a var and assign it the zero value
    ```go
        var x int
    ```
4. We can declare multiple variables at once with var, and they can be of the same type
    ```go
        var x, y int = 10, 20
    ```
5. We can declare all zero values of the same type, or of different types
    ```go
        var x, y int
        var z, w = 10, "hello"
    ```
6. We can wrap multiple variables in a declaration list
    ```go
        var (
            x    int
            y           = 20
            z    int
            d, e        = 40, "hello"
            f, g string
        )
    ```
7. Within a function we can use `:=` operator to replace var declaration that use type inference
    ```go
        var x = 10
        x := 10 
        var x, y = 10, "hello"
        x, y := 10, "hello"
    ```
8. The `:=` operator can do one trick that you cannot do with `var`: it allows you to assign values to existing 
    variables too. As long as at least one new variable is on the lefthand side of the `:=`, any of the other variables
    can already exist:
    ```go
         x := 10
         x, y := 30, "hello"
    ```
9. Using `:=` has one limitation. If you are declaring a variable at the package level, you must use `var` because `:=` 
   is not legal outside of functions.
10. The most common declaration style within functions is `:=`
11. Outside of a function, use declaration lists on the rare occasions when you are declaring multiple package-level 
    variables.
12. In some situations within functions, you should avoid `:=`:
    * When initializing a variable to its zero value, use `var x int`. This makes it clear
      that the zero value is intended.
    * When assigning an untyped constant or a literal to a variable and the default type
      for the constant or literal isn’t the type you want for the variable, use the long var
      form with the type specified. While it is *legal* to use a type conversion to specify
      the type of the value and use `:=` to write `x := byte(20)`, it is **idiomatic** to write
      `var x byte = 20`.
    * Because `:=` allows you to assign to both new and existing variables, it sometimes
      creates new variables when you think you are reusing existing ones. In those situations, 
      explicitly declare all your new variables with `var` to make it clear which variables are 
      **new**, and then use the assignment operator `(=)` to assign values to both **new** and **old variables**.
13. While `var` and `:=` allow you to declare multiple variables on the same line, use this style only when assigning 
    multiple values returned from a function or ``the comma ok idiom``
14. You should rarely declare variables outside of functions, in what’s called the *package block*
15. Package-level variables whose values change are a `bad idea`. When you have a variable outside of a function, it can 
    be difficult to track the changes made to it, which makes it hard to understand how data is flowing through your 
    program. This can lead to subtle bugs. As a general rule, you should only declare variables in the package block 
    that are effectively `immutable`.
16. Avoid declaring variables outside of functions because they complicate data flow analysis.

## Using Const
1. Go provide a way to declare a value as immutable by `const` keyword. 
    ```go
    package main
   
    import "fmt"
   
    const x int64 = 10
   
    const (
	    idKey = "id"
	    nameKey = "name"
    )
   
    const z = 20 * 10
   
    func main() {
	    const y = "hello"
	    fmt.Println(x)
	    fmt.Println(y)
	    x = x + 1 // this will not compile!
	    y = "bye" // this will not compile!
	    fmt.Println(x)
	    fmt.Println(y)
    }
    ```
   When you run this code, compilation fails with the following error messages:
   ```
   ./prog.go:20:2: cannot assign to x (constant 10 of type int64)
   ./prog.go:21:2: cannot assign to y (untyped string constant "hello")
   ```
2. Be aware that const in Go is very limited. Constants in Go are a way to give names to literals. They can only hold 
   values that the compiler can figure out at compile time. This means that they can be assigned:
   * Numeric literals 
   * true and false 
   * Strings 
   * Runes
   * The values returned by the built-in functions complex, real, imag, len, and cap
   * Expressions that consist of operators and the preceding values
3. Go doesn’t provide a way to specify that a value calculated at runtime is immutable.
    ```go
    x := 5
    y := 10
    const z = x + y // this won't compile!
    ```
4. There are no immutable arrays, slices, maps, or structs, and there’s no way to declare that a field in a struct is 
   immutable
5. Constants in Go are a way to give names to literals. There is no way in Go to declare that a variable is immutable.

## Typed and Untyped Constants

1. Constants can be typed and untyped. An untyped constant works exactly like a literal; it has no type of its own  but
   does have a default type when no other type can be inferred
2. A typed constant can be directly assigned only to a variable of that type
   ```go
    //untyped constant
    const x = 10  
    // All of the following assignments are legal:
    var y int = x
    var z float64 = x
    var d byte = x
    //typed constant
    const typedX int = 10 
    // This constant can be assigned directly only to an int. Assigning it to any other type produces a compile-time 
    // error like this:
    // cannot use typedX (type int) as type float64 in assignment
   ```
3. Another Go requirement is that *every declared local variable must be read*. It is a *compile-time error* to declare a 
   local variable and to not read its value.
4. The Go compiler won’t stop you from creating unread package-level variables. This is one more reason you should avoid
   creating package-level variables.
5. The Go compiler allows you to create unread constants with const. This is because constants in Go are calculated at 
   compile time and cannot have any side effects. This makes them easy to eliminate: if a constant isn’t used, it is 
   simply not included in the compiled binary.
6. 