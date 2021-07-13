package main

//Go Service Pattern template to get you started!

import "net/http"

func main() {

	/* Configure Port */
	//TODO: make Configurable
	PORT := "3000"

	http.ListenAndServe(":"+PORT, NewChiRouter().InitRouter())

}
