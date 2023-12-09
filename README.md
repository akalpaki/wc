WC.go
---
Golang implementation of the unix WC program

Uses:
1. You can pass a file name as an argument.
2. If you choose not to pass an argument, you will instead provide input from stdin. This allows you to use UNIX pipes to pipe input into the product.

Setup
---
```go
git clone https://github.com/akalpaki/wc.git

go build .
chmod +x ./wc
./wc [flags] [input]
```

Flags
---
- c : prints the number of bytes in the input
- l : prints the number of lines in the input
- w : prints the number of words in the input
- m : prints the number of characters in the input (UTF-8 chars)


You can test the correctness of the program by running it with the included
test.txt file and comparing it to `wc`.