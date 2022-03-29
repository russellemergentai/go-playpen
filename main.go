package main

import (
	greet "myPOC/myPackage"
	webstuff "myPOC/myWebPackage"
)

func main() {
	// docker build --tag alpha .
	// docker image ls
	// docker run --publish 8080:8080 beta
	// http://localhost:8080/users/bob
	greet.Hello()
	webstuff.IDWebcall()
}
