# fuzzqueens

How to fuzz wrong

Or: solving 8-queen problem with electricity.

```bash
$ go install golang.org/dl/go1.18beta1@latest

$ go1.18beta1 download

# will be installed into your GOPATH

$ go1.18beta1 test -v -fuzz=FuzzQueens
fuzz: elapsed: 0s, gathering baseline coverage: 0/7 completed
fuzz: elapsed: 0s, gathering baseline coverage: 7/7 completed, now fuzzing with 4 workers
fuzz: elapsed: 3s, execs: 213620 (71152/sec), new interesting: 0 (total: 6)
fuzz: elapsed: 6s, execs: 438033 (74778/sec), new interesting: 0 (total: 6)
fuzz: elapsed: 9s, execs: 666796 (76291/sec), new interesting: 0 (total: 6)
fuzz: elapsed: 12s, execs: 858814 (64042/sec), new interesting: 0 (total: 6)
fuzz: elapsed: 15s, execs: 1053941 (65033/sec), new interesting: 0 (total: 6)
fuzz: elapsed: 18s, execs: 1250438 (65507/sec), new interesting: 0 (total: 6)
fuzz: elapsed: 21s, execs: 1482186 (77250/sec), new interesting: 0 (total: 6)
fuzz: minimizing 57-byte failing input file
fuzz: elapsed: 21s, minimizing
--- FAIL: FuzzQueens (21.39s)
    --- FAIL: FuzzQueens (0.00s)
        testing.go:1320: panic: (0,3), (1, 1), (2, 4), (3, 7), (4, 5), (5, 0), (6, 2), (7, 6)
```

Note: process termiantes after a first solution. 

From what I know currently it's not possible to ignore failing input and continue fuzzing.
