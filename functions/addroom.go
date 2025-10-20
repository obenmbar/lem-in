package lemino

import (
	"fmt"
	"strconv"
	"strings"
)

func (R *Room) AddRoom(Dataroom string) error {
	name := strings.Fields(Dataroom)
	if len(name) != 3 {
		return fmt.Errorf("error en Add room ,les cordonner de room pas exist")
	}
	if strings.HasPrefix(name[0], "L") || strings.HasPrefix(name[0], "#")  {
		return fmt.Errorf("error en addroom , le nom de room n'est pas exist il contien espace ou L ou #")
	}
	R.Name = name[0]
	
	R.X, err = strconv.Atoi(name[1])
	if err != nil {
		return fmt.Errorf("error en addroom , X n' est pas valide", err)
	}
	R.Y, err = strconv.Atoi(name[2])

	if err != nil {
		return fmt.Errorf("error en addroom , Y n' est pas valide", err)
	}
	
	return nil
}
