# Chapter 3

## Arrays -- Too Rigid to Use Directly

1. There are a few declaration styles. In the first, we specify the size of the array and the type of the elements in
   the array:
    ```go
    // This creates an array of three ints with zero values and it is zero
    var x [3]int
    // If we have initial values for the array, we specify them with an array literal
    var x [12]int{1,5:4,6,10:100,15} 
    //This creats an array of 12 ints with [1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15]
    ```
2. When using an array literal to initialize an array, we can replace the number that specifies the number of elements
   in the array with ...:
   ```go
   var x = [...]int{10, 20, 30}
   ```
3. We can use == and != to compare two arrays. Arrays are equal if they are the same length and contain equal values:
   ```go
    var x = [...]int{1, 2, 3}
    var y = [3]int{1, 2, 3}
    fmt.Println(x == y) // prints true
   ```
4. Go has only one-dimensional arrays, but you can simulate multidimensional arrays:
   ```go
   var x [2][3]int
   // This declares x to be an array of length 2 whose type is an array of ints of length 3.
   // This sounds pedantic, but some languages have true matrix support, like Fortran or
   // Julia; Go isn’t one of them.
   ```
5. Like most languages, arrays in Go are read and written using bracket syntax:
   ```go
   x[0] = 10
   fmt.Println(x[2])
   ```
6. We cannot read or write past the end of an array or use a negative index. If we do this with a constant or literal
   index, it is a compile-time error. An out-of-bounds read or write with a variable index compiles but fails at runtime
   with a panic
7. The built-in function len takes in an array and returns its length:
   ```go 
   fmt.Println(len(x))
   ```
8. The reason why arrays are rarely used is that they come with an unusual limitation: Go considers the size of the
   array to be part of the type of the array. This makes an array that’s declared to be `[3]int` a different type from
   an array that’s declared to be `[4]int`. This also means that we cannot use a variable to specify the size of an
   array, because types must be resolved at compile time, not at runtime.
9. We cannot use a type conversion to directly convert arrays of different sizes to identical types. Because we cannot
   convert arrays of different sizes into each other, we cannot write a function that works with arrays of any size and
   we can’t assign arrays of different sizes to the same variable.
10. The main reason arrays exist in Go is to provide the backing store for slices, which are one of the most useful
    features of Go.

## Slices

1. Working with slices looks a lot like working with arrays, but subtle differences exist. The first thing to notice is
   that you don’t specify the size of the slice when you declare it:
   ```go
   var x = []int{10, 20, 30}
   ```
2. Using [...] makes an array. Using [] makes a slice.
3. This creates a slice of three ints using a slice literal. Just as with arrays, we can also specify only the indices
   with nonzero values in the slice literal:
   ```go
   var x = []int{1, 5: 4, 6, 10: 100, 15}
   // This creates a slice of 12 ints with the following values: [1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100,15].
   ```
4. We can simulate multidimensional slices and make a slice of slices:
   ```go
   var x [][]int
   ```
5. We read and write slices using bracket syntax, and, just as with arrays, we can’t read
   or write past the end or use a negative index:
   ```go
   x[0] = 10
   fmt.Println(x[2])
   ```
6. We can see the differences between arrays and slices when you look at declaring slices without using a literal:
   ```go
   var x []int // zero value is nil
   // This creates a slice of ints. Since no value is assigned, x is assigned the zero value for
   // a slice, which is something you haven’t seen before: nil.
   ```
7. In Go, *nil* is an identifier that represents the lack of a value for some types. Like the untyped numeric constants,
   *nil* has no type, so it can be assigned or compared against values of different types. A nil slice contains nothing.
8. A slice isn’t comparable. It is a compile-time error to use == to see if two slices are identical or != to see if
   they are different. The only thing WE can compare a slice with using == is nil:
   ```go
   fmt.Println(x == nil) // prints true
   ```
9. Since Go 1.21, the slices package in the standard library includes two functions to compare slices. The
   *slices.Equal*
   function takes in two slices and returns true if the slices are the same length, and all of the elements are equal.
   It requires the elements of the slice to be comparable. The other function, slices.EqualFunc, lets you pass in a
   function to determine equality and does not require the slice elements to be comparable. You’ll learn about passing
   functions into functions in “Passing Functions as Parameters” on page 107. The other functions in the slices package
   are covered in “Adding Generics to the Standard Library” on page 201.
   ```go
   x := []int{1, 2, 3, 4, 5}
   y := []int{1, 2, 3, 4, 5}
   z := []int{1, 2, 3, 4, 5, 6}
   s := []string{"a", "b", "c"}
   fmt.Println(slices.Equal(x, y)) // prints true
   fmt.Println(slices.Equal(x, z)) // prints false
   fmt.Println(slices.Equal(x, s)) // does not compile
   ```

## len

1. Passing a nil slice to len returns 0.
2. Functions like *len* are built into Go, because they can do things that can’t be done by the functions that you can
   write. You’ve already seen that len’s parameter can be any type of array or any type of slice, it also works for
   strings and maps, channels. Trying to pass a variable of any other type to len is a compile-time error. Go doesn’t
   let developers write a function that accepts any string, array, slice, channel, or map, but rejects other types.

## append

1. The built-in append function is used to grow slices
2. One slice is appended onto another by using the ... operator to expand the source slice into individual values
   ```go
   var x []int
   x = append(x, 10) // assign result to the variable that's passed in
   var x = []int{1, 2, 3}
   x = append(x, 4)
   // You can append more than one value at a time:
   x = append(x, 5, 6, 7)
   y := []int{20, 30, 40}
   x = append(x, y...)
   ```
3. It is a compile-time error if you forget to assign the value returned from append.
   Capacity
   A slice is a sequence of values. Each element in a slice is assigned to consecutive memory locations, which makes it
   quick to read or write these values. The length of a slice is the number of consecutive memory locations that have
   been assigned a value. Every slice also has a capacity, which is the number of consecutive memory locations reserved.
   This can be larger than the length. Each time you append to a slice, one or more values is added to the end of the
   slice. Each value added increases the length by one. When the length reaches the capacity, there’s no more room to
   put values. If you try to add additional values when the length equals the capacity, the append function uses the Go
   runtime to allocate a new backing array for the slice with a larger capacity. The values in the original backing
   array are copied to the new one, the new values are added to the end of the new backing array, and the slice is
   updated to refer to the new backing array. Finally, the updated slice is returned.

   The Go Runtime
   Every high-level language relies on a set of libraries to enable programs written in that language to run, and Go is
   no exception. The Go runtime provides services like memory allocation and garbage collection, concurrency support,
   networking, and implementations of built-in types and functions. The Go runtime is compiled into every Go binary.
   This is different from languages that use a virtual machine, which must be installed separately to allow programs
   written in those languages to function. Including the runtime in the binary makes it easier to distribute Go programs
   and avoids worries about compatibility issues between the runtime and the program. The drawback of including the
   runtime in the binary is that even the simplest Go program produces a binary that’s about 2 MB.


4. When a slice grows via append, it takes time for the Go runtime to allocate new memory and copy the existing data
   from the old memory to the new.
5. Make function allows you to specify the type, the length, and, optionally the capacity of the slice.
   ```go
   x := make([]int, 5) //This creates an int slice with a length of 5 and a capacity of 5.
   ```
6. One common beginner mistake is to try to populate those initial elements using append:
   ```go
   x := make([]int, 5) 
   x = append(x, 10) 
   // The 10 is placed at the end of the slice, after the zero values in elements 0–4 because
   // append always increases the length of a slice. The value of x is now [0 0 0 0 0 10],
   // with a length of 6 and a capacity of 10 (the capacity was doubled as soon as the sixth
   // element was appended).
   ```
7. You can also specify an initial capacity with make:
   ```go
   x: = make([]int, 5, 10)
   // This creates an int slice with a length of 5 and a capacity of 10.
   ```

8. You can also create a slice with zero length but a capacity that’s greater than zero:
   ```go
   x := make([]int, 0, 10)
   // In this case, you have a non-nil slice with a length of 0 but a capacity of 10. Since the
   // length is 0, you can’t directly index into it, but you can append values to it:
   // x := make([]int, 0, 10) x = append(x, 5,6,7,8)
   // The value of x is now [5 6 7 8], with a length of 4 and a capacity of 10
   ```
9. Emptying a slice: Go 1.21 added a clear function that takes in a slice and sets all the slice’s elements
   to their zero value. The length of the slice remains unchanged. The following code:
   ```go
   s := []string{"first", "second", "third"}
   fmt.Println(s, len(s))
   clear(s)
   fmt.Println(s, len(s))
   ```
10. Declaring a slice: The primary goal is to minimize the number of times the slice needs to grow. 
    1. If it’s possible that the slice won’t need to grow at all, use a var declaration with no assigned value to create
    a nil slice 
    ```go
       var data []int
    ```
    **NOTE!** You can create a slice using an empty slice literal:
      ```go
      var x = []int{}
      ```
      This creates a slice with zero length and zero capacity. It is confusingly different from a nil slice. Because of 
      implementation reasons, comparing a zero-length slice to nil returns false, while comparing a nil slice to nil 
      returns true. For simplicity, favor nil slices. A zero-length slice is useful only when converting a slice to JSON 
    2. If you have some starting values, or if a slice’s values aren’t going to change, then a slice literal is a good choice
    ```go
     data := []int{2, 4, 6, 8} // numbers we appreciate
    ```
    3. If you have a good idea of how large your slice needs to be, but don’t know what those values will be when you 
       are writing the program, use make. The question then becomes whether you should specify a nonzero length in the 
       call to make or specify a zero length and a nonzero capacity. There are three possibilities:
       1. If you are using a slice as a buffer, then specify a nonzero length.
       2. If you are sure you know the exact size you want, you can specify the length and index into the slice to set 
          the values. This is often done when transforming values in one slice and storing them in a second. The 
          downside to this approach is that if you have the size wrong, you’ll end up with either zero values at the end
          of the slice or a panic from trying to access elements that don’t exist.
       3. In other situations, use make with a zero length and a specified capacity. This allows you to use append to 
          add items to the slice. If the number of items turns out to be smaller, you won’t have an extraneous zero 
          value at the end. If the number of items is larger, your code will not panic
11. Slicing slices: A slice expression creates a slice from a slice. It's written inside brackets and consists of a 
    starting of a starting offset and an ending offset, separated by a colon (:). The starting offset is the first 
    position in the slice that is included in the new slice, and the ending offset is one past the last position to 
    include. If you leave off the starting offset, 0 is assumed. Likewise, if you leave off the ending offset, the end 
    of the slice is substituted.
12. 