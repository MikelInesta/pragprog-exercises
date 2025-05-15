package parse

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T){
	tests, err := os.ReadDir("./testdata")
	if err != nil {
		t.Fatal("unable to load test data: ", err)
	}

	for _, test := range tests {
		t.Run(test.Name(), func(t *testing.T){
			_, err := ParseFile(filepath.Join(".", "testdata", test.Name())) 
			if err != nil{
				t.Error("error parsing: err: ", err)
			}
		})	
	}
}