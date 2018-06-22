package funcs

import (
	"fmt"
	"log"

	drive "google.golang.org/api/drive/v2"
)

type Functions struct {
	Service *drive.Service
}

var service *drive.Service

func (f *Functions) GetFuncs(serv *drive.Service) map[string]interface{} {
	service = serv
	return map[string]interface{}{"foo": foo, "bar": bar, "ls": ls}
}

func foo() {
	print("foo")
}
func bar(a, b, c string) {
	print("bar")
	print(a)
	print(b)
	print(c)
}

func ls() {
	r, err := service.Files.List().Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	fmt.Println("Files:")
	if len(r.Items) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, i := range r.Items {
			fmt.Printf("%s (%s)\n", i.Title, i.Id)
		}
	}
}
