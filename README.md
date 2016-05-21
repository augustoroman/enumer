#Enumer

Enumer is a tool to generate Go code that adds useful methods to Go enums
(constants with a specific type). It started as a fork of
[alvaroloes's Enumer tool](https://godoc.org/github.com/alvaroloes/enumer)
which started as a fork of
[Rob Pike’s Stringer tool](https://godoc.org/golang.org/x/tools/cmd/stringer).

##Generated functions and methods
When Enumer is applied to a type, it will generate three methods and one function:
* A method `String()` that returns the string representation of the enum value.
This makes the enum conform the `Stringer` interface, so whenever you print an
enum value, you'll get the display name instead of a number.
* A function `<Type>_Parse(s string)` to get the enum value from its display
representation. This is useful when you need to read enum values from the
command line arguments, from a configuration file, from a REST API request... In
short, from those places where using the real enum value (an integer) would be
almost meaningless or hard to trace or use by a human.
* And two more methods, `MarshalText()` and `UnmarshalText()`, that makes the
enum conform the `encoding.Marshaler` and `encoding.Unmarshaler` interfaces.
Very useful to use it in JSON or XML APIs.

Generally, the display representation is just the identifier of the constant,
but enumer also looks at the comment following the declaration.  If the comment
looks like a string (starts and ends with a quote), it will use that as the
string representation of the enum.

For example, if we have an enum type called `Pill`,
```go
type Pill int

const (
	PLACEBO Pill = iota         // "a Placebo, yo!"
	ASPIRIN                     // "Aspirin"
	IBUPROFEN                   // "Ibuprofen"
	PARACETAMOL                 // "Paracetamol"
	ACETAMINOPHEN = Paracetamol // "Acetaminophen"
)
```
executing `enumer -type=Pill` will generate a new file with four methods:
```go
func (i Pill) String() string {
    //...
}

func PillString(s string) (Pill, error) {
    //...
}

func (i Pill) MarshalText() ([]byte, error) {
	//...
}

func (i *Pill) UnmarshalText(data []byte) error {
	//...
}
```
From now on, we can:
```go
// Convert any Pill value to string
var aspirinString string = ASPIRIN.String()
// (or use it in any place where a Stringer is accepted)
fmt.Println("I need ", PLACEBO) // Will print "I need a Placebo, yo!"

// Convert a string with the enum name to the corresponding enum value
pill, err := PillString("Ibuprofen")
if err != nil {
    fmt.Println("Unrecognized pill: ", err)
    return
}
// Now pill == IBUPROFEN

// Marshal/unmarshal to/from json strings, either directly or automatically when
// the enum is a field of a struct
pilltext, err := ASPIRIN.MarshalText()
// Now pilltext == []byte(`Aspirin`)
```

The generated code is exactly the same as the Stringer tool plus the mentioned
additions, so you can use **Enumer** where you are already using **Stringer**
without any code change.

## How to use
The usage of Enumer is the same as Stringer, so you can refer to the [Stringer
docs](https://godoc.org/golang.org/x/tools/cmd/stringer) for more information.

There is only one flag added: `noMarshal`. If this flag is set to true (i.e.
`enumer -type=Pill -noMarshal`), the JSON related methods won't be generated.

## Inspiring projects
* [Enumer](https://godoc.org/github.com/alvaroloes/enumer)
* [Stringer](https://godoc.org/golang.org/x/tools/cmd/stringer)
* [jsonenums](https://github.com/campoy/jsonenums)

