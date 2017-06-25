# For the fun of it - Quadratic equation solver

## Limitations
* Uses spaces to split the equation into terms
* + and - signs must be attached to each term in the equation (+5x etc)
* x must be used to denote the variable
* Can only solve up to x^2

## Build

```sh
$ go build -o main
```

## Execute

```sh
$ ./main -equation="x^2 +5x +18 = 7363094" -constraint_lower=0 -constraint_upper =10000
```
