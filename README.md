# Backtracking Sudoku Solver

## Preparation

1. `git clone` this repository

2. Download dataset from [1 million Sudoku games](https://www.kaggle.com/datasets/bryanpark/sudoku) and `unzip sudoku.csv.zip`

```
$ ls -1
README.md
main.go
sudoku.csv
sudoku.csv.zip
```

3. `go run main.go`

## Performance

```
$ time go run main.go

real    0m37.780s
user    0m39.206s
sys     0m8.238s
```