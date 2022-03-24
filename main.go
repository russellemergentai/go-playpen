package main

import (
	greet "myPOC/myPackage"
	webstuff "myPOC/myWebPackage"
)

func main() {
	greet.Hello()
	webstuff.IDWebcall()
}
