# cyoa

Ever wanted to build your own storyline? `cyoa` is primarily a webtool which reads a JSON file and provides each chapter of the story on a webpage with a specific route. At the end of each chapter you will have a few options of how you would like to continue the story. What if you have no web browser ? No problem the tool runs also in your command line with the same functionality. Reference to [gophercise - Exercise 3](https://github.com/gophercises/cyoa).

## Installation

Create a new directory in your `$GOPATH` and make sure that in your `$GOPATH` exists a `bin` directory.

```bash
# Project structure
$GOPATH
|
--- bin
--- src
     |
     --- <YOUR-PROJECT>
```

## Build

Install the `make` tool and run :

```bash
make build
```

## Run the program

```bash
$GOPATH/bin/cyoa --story=<path-to-json> --cli <boolean>
```

Examples of the accepted story JSON format can be found in the `/assets` folder. Feel free to change the `html` template which also can be found in the `/assets` folder.
