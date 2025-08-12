package main

import "blogger/modules"

func main() {
	blogger := modules.NewBloggerServer()
	blogger.RunServer()
}