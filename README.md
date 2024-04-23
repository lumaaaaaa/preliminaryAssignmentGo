# Preliminary Assignment (Go)
Author: Devin Patel
## Introduction
This is a preliminary assignment completed as part of the application process for an Android Developer Co-op position. 
The resources in this repository is my submission for the assignment.

## Structure
The repository contains the following relevant resources:
1. **README.md**: This file contains the introduction and the structure of the repository.
2. **NIO Assignment Solution.pdf**: This file contains my write-up of my submission.
3. **main.go**: This file contains a simple demo program that calls the algorithm I implemented in the assignment.
4. **parse/parse.go**: This file contains the implementation of the algorithm.
5. **parse/parse_test.go**: This file contains the test cases for the algorithm.

## Prerequisites
To run the program, you need to have Go installed on your machine. You can download and install Go from the official 
site [here](https://go.dev/dl/), or if you are using a Unix-based system, I would recommend installing Go using the 
package manager of your choice. On Arch Linux, you can install Go using the following command:
```bash
sudo pacman -S go
```

## Running the Program
To run the program, you can use the following command inside of the root directory of the repository:
```bash
go run .
```

## Testing the Algorithm
To run the test cases for the algorithm, you can use the following command inside of the `parse` directory:
```bash
go test
```

To see the test coverage of the algorithm, you can use the following command inside of the `parse` directory:
```bash
go test -cover
```

If you would like to see the individual test results, append the `-v` flag to your go test commands.

## Conclusion
Thank you for taking the time to review my submission. 

If you have any questions regarding my algorithm, implementation, or tests, please do not hesitate to reach out to me.

I look forward to hearing from you soon!