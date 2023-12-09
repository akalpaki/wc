WC.go
---
Golang implementation of the unix WC program

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