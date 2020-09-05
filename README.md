# The Jumping Horse

This is a humble way to solve this problem (https://en.wikipedia.org/wiki/Knight%27s_tour).

## Consideration

- The representation of the chess board is like a matrix schema, the first item is a row and second is a column.
- The Horse start in 0 0
- The result is displayed by the console from the end to the beginning.
- No horses were harmed in the running of this program.


## How to run it

If you are a layz person how don't have golang in your PC, you can run it with docker. Follow the next step:

```bash
git clone https://github.com/quinont/jumpinghorse.git
```

And then, run docker

```bash
docker run --rm -v "$PWD/jumpinghorse/:/go/src/jumpinghorse/" -w "/go/src/jumpinghorse/" golang:1.15.1-alpine3.12 go run main.go
```

Finally,

```
Enjoy it
```

If you wanna build the code and execute the binary (to improve the execution time), you will need to do this:

```bash
docker run --rm -v "$PWD/jumpinghorse/:/go/src/jumpinghorse/" -w "/go/src/jumpinghorse/" golang:1.15.1-alpine3.12 go build main.go
```

and then...
```bash
./main
```


## TODO

- Documentation documentation documentation
- Improve the heuristic that deals with selecting the next step
- Improve the output
