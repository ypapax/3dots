package main

import "log"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var ii = [...]int{1, 2, 3} //array // https://programming.guide/go/three-dots-ellipsis.html
	var jj = []int{1, 2, 3} // slice https://appliedgo.net/slices/#:~:text=A%20slice%20is%20just%20a,are%20in%20fact%20different%20types.&text=A%20slice%20is%20nothing%20but,some%20part%20of%20the%20array.
	log.Println(ii)
	log.Println(jj)
	log.Println(len(ii))
	log.Println(cap(jj))
	log.Println(len(ii))
	log.Println(cap(jj))

	log.Println(ii[:])
	log.Println(jj[:])
}