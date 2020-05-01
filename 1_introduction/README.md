## Go introduction and essential API example

**Global const** def, may be a block of multiple const values

**iota**: block scoped, starts to 0 and increments by 1 at every call (inside a block is incremented for every const without value)
```
const (
    first = iota + 3
    second
)
```

It's possible to declare multiple values in one line

```stringDemo, number := "John", 3```

**Array** is of a defined length, **Slice** is a more dynamic, we can add elements with append and get chunks.
It's also possible to create a slice from an array (like a pointer to array's values).
```
arr := [3]int{1,2,3}
slice := arr[:]

slice := []int{1,2,4}
slice = append(slice, 5, 2, 4)

// Chunks with limits definition like [init:end)
chunk := slice[0:1]
```
A Map is as a JSON but with all keys and values of the defined type
```
m := map[string]int{"foo": 2}
```
A Struct is as a complete JSON (NOTE: in the values definitions formatted by row we need a last comma)
```
type user struct {
    ID int
    FirstName string
    LastName string
}
var u user

u.ID = 1
u.FirstName = "Arthur"
u.LastName = "Wane"

u2 := user {
    ID: 1,
    FirstName: "John",
    LastName: "Doe",
}
```

#### Controlling programming flow

Default FOR LOOP:
```
for i := 0; i < 5; i++ {
    println(i)
}
```
Infinite LOOP (empty for):
```
for {
    print("infinite")
}
```
List loop:
```
slice := []int{1,2,3}
for i, v := range slice {
    println(i, v)
}
```

### Go Packages

**Packages**: a package is constructed from one or more source files that together declare constants, types, 
variables and functions belonging to the package and which are accessible in all files of the same package.

#### Library Package
* consumed by another package
* name must match directory name
    * naming: short and clear, lowercase, no underscores, prefer nouns
* should provide a focused set of related features

#### Main Package
* application entry point
* contains a main() function
* can be in any directory
* focus on app setup and initialization

#### Member visibility

* Public scope:
    * capitalize member
    * available to all consumers
* Package scope:
    * lowercase member
    * only available within package
    * minimum scope
* Internal package scope:
    * can use public and package level members
    * scoped to parent package and its descendants
    * defined inside a folder named "internal"