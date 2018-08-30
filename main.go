package main

import (
	"fmt"

	"github.com/mauleyzaola/tak-proto/pb/services"
)

func main() {
	a := &pb.AreaService{}
	fmt.Println(a.Id)
}
