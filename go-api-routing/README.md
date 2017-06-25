# Go API routing

In this project I am expand even further on my previous hello-complex example.
Here I take it one step further and start breaking everything into separate files.

The starting file is `hello.go`.
It does very little besides run the server.
The only other task it has is to extract the first portion of the URL and forward the request to the correct API.

Each API is in its own package.
My starting point was to create a v1 API and a dev API.
They each have a `Root()` that as their starting point.
From here requests will be incrementally forwarded further down the API as appropriate.

The rules that I am using to direct the flow are simply to shift off the first portion of the path.
Two other basic conditions include checking the query string and the HTTP method.
