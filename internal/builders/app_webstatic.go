package builders

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fengdotdev/coipossr/internal/render"
)


type AppWebStatic struct{
	
}


// model
type StaticBuilder struct {
	RootPath string
	pages map[string]render.StaticWebPage
}
// constructors
func NewStaticBuilder(outputPath ...string) *StaticBuilder {
	rootPath := "output"
	if len(outputPath) > 0 {
		rootPath = outputPath[0]
	}
	return &StaticBuilder{
		RootPath: rootPath,
		pages: make(map[string]render.StaticWebPage),
	}
}

// methods

func (s *StaticBuilder) Build() error {
	wd, err := os.Getwd()
		if err != nil {
			return err
		}

	unixwd := filepath.ToSlash(wd)
	index := "index.html"
	for path, page := range s.pages {

		fullPathDir := filepath.Join(unixwd, s.RootPath, path)
		fullPath:= filepath.Join(fullPathDir, index)
		err := os.MkdirAll(fullPathDir, os.ModePerm)
		if err != nil {
			return err
		}
		file, err := os.Create(fullPath)
		if err != nil {
			return err
		}
		page.Render(file)
		fmt.Println(file.Name() + " created")
		file.Close()
	}
	return nil
}

func (s *StaticBuilder) HandleFunc(pattern string, page render.StaticWebPage) {
	s.pages[pattern] = page
}

func (s *StaticBuilder)AddPage(pattern string, page render.StaticWebPage){
	s.pages[pattern] = page
}