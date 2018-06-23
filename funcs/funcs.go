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

var functions = map[string]interface{}{
	"foo": foo,
	"bar": bar,
	"ls":  ls,
}

func (f *Functions) GetFuncs(serv *drive.Service) map[string]interface{} {
	service = serv
	return functions
}

// sample to show how reflection works, try passing only two params
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

// test function for working with APIs
func foo() {
	// this is currently hardcoded to my own "My Drive", still trying to figure out how to get this programatically
	r, err := service.Files.List().MaxResults(2000).Q("'0AExqA7NOpVhmUk9PVA' in parents").Do()

	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	fmt.Println("Files:")
	if len(r.Items) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, i := range r.Items {
			fmt.Printf("%s (%s) %s %s %s\n", i.Title, i.Id, i.Kind, i.TeamDriveId, i.MimeType)
			/*
				for _, i := range i.Parents {
					fmt.Printf("%t (%s) (%s) (%s) (%s) (%s)\n", i.IsRoot, i.Id, i.Header, i.ParentLink, i.SelfLink, i.Kind)
				}
			*/
		}
	}
}
