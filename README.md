# Coreutils-Basem_Elgalfy

This repository is a Go implementation of some of the core utilities commonly found in Unix-like operating systems. The project aims to provide a learning platform for Go programming and an understanding of how these fundamental utilities work.

## Utilities Included

- `cat`: Concatenate and print files, -n flag : by default false , used to show the line number in the output
- `echo`: Display a line of text
- `env`: Display the environment variables
- `false`: Do nothing, unsuccessfully
- `head`: Output the first part of files, -n flag : by default 10 , used to specify how many lines to be printed 
- `tail`: Output the last part of files, -n flag : by default 10 , used to specify how many lines to be printed
- `tree`: List contents of directories in a tree-like format, supports the -L flag
- `true`: Do nothing, successfully
- `wc`: Print newline, word, and byte counts for each file, -w,-c,-l flags : by default true , used to display the word , bytes and lines counts.
- `yes`: Output a string repeatedly until killed

## Getting Started

To get started with these utilities, clone the repository and build the utilities using Go.

### Prerequisites

- Go version 1.22.4 or higher

### Installation

1. Clone the repository:

```sh
git clone https://github.com/codescalersinternships/Coreutils-Basem_Elgalfy.git