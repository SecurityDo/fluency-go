package fluency

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(x interface{}) {
	b, _ := json.Marshal(x)
	fmt.Println(string(b))
}

func PrettyPrintJSON(x interface{}) {
	pretty, _ := json.MarshalIndent(x, "", "   ")
	fmt.Printf("%s\n", pretty)
}
