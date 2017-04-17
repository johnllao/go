package main

import "fmt"
import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, indexHTML)
}

var indexHTML = `<html>
	<head>
		<title>File Upload</title>
		<link rel="stylesheet" type="text/css" href="/css" />
	</head>
	<body>
		<div id="container">
			<h2>File Upload</h2>
			<form method="POST" action="/upload" enctype="multipart/form-data">
				<div>
					<input type="file" name="file" />
				</div>
				<div>
					<input type="submit" value="Upload" />
				</div>
			</form>
		</div>
	</body>
</html>`