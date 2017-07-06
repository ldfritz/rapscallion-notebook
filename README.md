# Rapscallion Notebook

This is a growing collection of Go examples.

I will attempt to get them setup as `godoc` examples.
There are a few I'll need to convert to get started.

## How to use

Get the repo.

```
go get github.com/ldfritz/rapscallion-notebook
```

Ignore it if it complains about nothing to build.
Maybe I'll toss a snippet in later to satisfy that.


Then run the Go doc server.

```
godoc -http:8080
```

Visit the package in your browser.

```
http://localhost:8080/pkg/github.com/ldfritz/rapscallion-notebook/
```

## Testing the examples

If you are concerned about the accuracy of an example, you can test them.

Browse to the folder with the appropriate tests and run the test suite.

```
go test -v
```

## License

> MIT License

> Copyright (c) 2017 Luke Fritz

> Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

> The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
