package api

import (
	"fmt"

	"github.com/stonelike/CleanGo/src/domain/entity"
)

func main() {
	u, err := entity.NewUser("aa")

	fmt.Println(u, err)
}
