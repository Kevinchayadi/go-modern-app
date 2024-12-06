package handlers

import (
	"net/http"

	"github.com/Kevinchayadi/go-modern-app/pkg/models"
	"github.com/Kevinchayadi/go-modern-app/pkg/config"
	"github.com/Kevinchayadi/go-modern-app/pkg/render"
)

//template for holds data had to send


var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepository(app *config.AppConfig) *Repository {
    return &Repository{App: app}
}

func NewHandler(r *Repository) {
	Repo = r;
}

func (m *Repository) Home(w http.ResponseWriter,r *http.Request){

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{
		
		
		} )
		// fmt.Fprintf(w, "this is the home page")
	}
	
	func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
		//logic
		StringMap := make(map[string]string)
		StringMap["test"] = "hello, again"

	//send data
		render.RenderTemplate(w, "about.page.html", &models.TemplateData{
			StringMap: StringMap,
		})

	// sum:= AddValue(2,2)
	// _, _= fmt.Fprintf(w, fmt.Sprintf("this is the about page, two plus two is %d",sum))
}

