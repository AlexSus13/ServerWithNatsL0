package app

import (
	"ServerWithNatsL0/cache"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"html/template"
	"net/http"
)

func (app *App) HomePage(w http.ResponseWriter, r *http.Request) {
	app.MyLogger.Info("HomePage")
	//Check the cache if is empty...
	if len(app.CacheMap) == 0 {
		//calling the cache recovery function
		err := cache.CacheRecovery(app.Db, app.CacheMap)
		if err != nil {
			app.MyLogger.WithFields(logrus.Fields{
				"func":    "cache.CacheRecovery",
				"package": "app",
			}).Info(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
	}
	//Template is parsed using the ParseFile() function
	tmpl, err := template.ParseFiles(app.Config.PathToHTMLhome)
	if err != nil {
		app.MyLogger.WithFields(logrus.Fields{
			"func":    "template.ParseFiles home_page.html",
			"package": "app",
		}).Info(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	//Execute() method applies a parsed template to the specified data object, and writes the output to an output writer.
	tmpl.Execute(w, app.CacheMap)

}

func (app *App) PublishOrdersData(w http.ResponseWriter, r *http.Request) {
	//We get the parameters of the request path, the mux.Vars function.
	vars := mux.Vars(r)
	//We extract the name of the parameter we need from the vars object
	OrderUid := vars["orderuid"]
	//Template is parsed using the ParseFile() function
	tmpl, err := template.ParseFiles(app.Config.PathToHTMLorders)
	if err != nil {
		app.MyLogger.WithFields(logrus.Fields{
			"func":    "template.ParseFiles orders_page.html",
			"package": "app",
		}).Info(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	//Check the cache if is empty...
	if len(app.CacheMap) == 0 {
		//calling the cache recovery function
		err := cache.CacheRecovery(app.Db, app.CacheMap)
		if err != nil {
			app.MyLogger.WithFields(logrus.Fields{
				"func":    "cache.CacheRecovery",
				"package": "app",
			}).Info(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
	}
	//Getting the data structure from the cache map
	OrdersData := app.CacheMap[OrderUid]
	//Execute() method applies a parsed template to the specified data object, and writes the output to an output writer.
	tmpl.Execute(w, OrdersData)
}
