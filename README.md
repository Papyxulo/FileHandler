# FileHandler

Simple File handler for golang

## Usage

``` Golang
...
// Create the File object
File := new(File)

// Write file, it will overrite if the file already exists
// Write(path string, text string)
err := File.Write("path and file name.txt", "some text")
if err != nil {
    println(err.error())
}

// Append file, it will append to the file if it exists, if the file does not exist it will create it
// The auto_newline will add a \n at the end of the string
//Append(path string, text string, auto_newline bool)
err := File.Append("path and file name.txt", "some text", true) 
if err != nil {
    println(err.error())
}

// Read a file and returns the content as a string
//Read(path string)
text, err := File.Read("path and file name.txt")
if err != nil {
    println(err.error())
}
println(text)

// Read a file and returns the content as a array os strings (split by line)
//ReadList(path string)
lines, err := File.ReadList("path and file name.txt")
if err != nil {
    println(err.error())
}
for _,line := range lines{
    println(line)
}

// Read a file and returns the content as a array os strings (split by line and by separator) 
//ReadCSV(path string, separator string)
lines, err := File.ReadCSV("path and file name.txt",";")
if err != nil {
    println(err.error())
}
for _,line := range lines{
    for _,field := range line{
        print(field + " ")
    }
    println("")
}

```

---

## V1.0.0

- Inicial commit
