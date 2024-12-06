package render

import (
	// "fmt"
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Kevinchayadi/go-modern-app/pkg/models"
	"github.com/Kevinchayadi/go-modern-app/pkg/config"

)

var function = template.FuncMap{

}
var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td;
}

// RenderTemplate is a function to render HTML templates
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache{
		tc =app.TemplateCache
	}else{
		tc,_ = CreateTemplateCache()
	}

	//create a template cache



	//get Request template from cache
	t, ok := tc[tmpl]
	if !ok{
		log.Fatal("could not get template from cache template")
	}


	buf:= new(bytes.Buffer);

	td = AddDefaultData(td)

	_ = t.Execute(buf,td)

	_, err := buf.WriteTo(w)
	if(err != nil){
		log.Println(err)
	}

	//render the template
	_,err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
    }

	// parseTemplate, _ := template.ParseFiles("./template/" + tmpl, "./template/base.template.html")
	// err := parseTemplate.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println("error parsing template: ", err)
	// 	return
	// }
}

func CreateTemplateCache() (map[string]*template.Template, error){
	// mycache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./template/*.page.html")
	if err != nil {
		return myCache ,err
	}
	
	for _, page := range pages {
		name:=filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache ,err
		}
		matches, err := filepath.Glob("./template/*.template.html")
		if err!= nil {
            return myCache, err
        }
		if len(matches) > 0 {
			ts,err = ts.ParseGlob("./template/*.template.html")
			if err!= nil {
                return myCache, err
            }
		}

		myCache[name] = ts
	}
	return myCache, nil

}








// RenderTemplate is a function to render
// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter,t string){
// 	var tmpl *template.Template
// 	var err error

// 	//check to see if we allready have the template in out cache
// 	_, inMap := tc[t];
// 	if !inMap {
// 		fmt.Println("adding cache and create the template")
// 		err = createTemplateCache(t)
// 		if err!= nil {
//             fmt.Println("error creating template cache: ", err)
//         }
// 	} else {
// 		fmt.Println("using cache template!")
// 	}
// 	tmpl = tc[t]

// 	err = tmpl.Execute(w, nil)
// 	if err!= nil {
//             fmt.Println("error creating template cache: ", err)
//         }
// }

// func createTemplateCache(t string) error{
// 	templates:= []string{
// 		fmt.Sprintf("./template/%s",t), "./template/base.template.html",

// 	}

// 	tmpl, err := template.ParseFiles(templates...)
// 	if err!= nil {
//         return err
//     }
// 	tc[t] = tmpl
// 	return nil
// }