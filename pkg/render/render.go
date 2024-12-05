package render

import (
	// "fmt"
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate is a function to render HTML templates
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	//create a template cache
	tc , err := createTemplateCache()
	if err!= nil {
        log.Fatal(err)
    }


	//get Request template from cache
	t, ok := tc[tmpl]
	if !ok{
		log.Fatal(err)
	}

	buf:= new(bytes.Buffer);

	err = t.Execute(buf,nil)

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

func createTemplateCache()(map[string]*template.Template, error){
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