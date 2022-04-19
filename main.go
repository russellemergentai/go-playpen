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

	// az aks get-credentials --resource-group K8SRG --name aksbetatest
	// kubectl apply -f deploy.yml
	// curl http://20.108.239.131:30475/users/bob
	//  kubectl logs web-65f9f75fdf-dccct
	//
	greet.Hello()
	webstuff.IDWebcall()
}
