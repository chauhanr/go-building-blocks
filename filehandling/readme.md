## File Handling

## Reading and writing from user

In order to capture strings from the command line we need to use a bufio package in go. The method that is used declaring a reader to read the standard output os NewReader() -  func NewReader(rd io.Reader) *Reader.
os.Stdin implements the io.Reader interface and therefore is a good choice to put into the method

```
     inputReader := bufio(os.Stdin)
     input, err := inputReader.ReadString('\n')
     // ReadString() takes a delimiter which signifies that the method will keep reading till it encounters a delimiter.
     if err != nil{
       fmt.Printf("The input from the command line is : %s\n", input)
     }
```
The code can be used with switch statements to capture inputs from the command line and processing the input. 


2. Reading from a file and compressed files
3. Writing to a file
4. Copying files
5. Reading arguments from the command line
6. Reading files from buffer
7. Reading and writing file with slices
8. Using interfaces with read and write
9. Json and XML data format
10. Data transport using gob
11. Cryptography with Go
