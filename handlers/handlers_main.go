package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/swanwish/go-helper/config"
	"github.com/swanwish/go-helper/logs"
	"github.com/swanwish/go-helper/web"
	. "github.com/swanwish/sb-admin-go/config"
	"github.com/swanwish/sb-admin-go/views"
)

type MainHandlers struct {
}

func (h MainHandlers) GetPathPrefix() string {
	return ""
}

func (h MainHandlers) InitRouter(r *mux.Router) {
	r.HandleFunc("/", web.MakeLogEnabledHandler(rootHandler)).Methods("GET")
	r.HandleFunc(ViewPathPrefix+"/{viewId}", web.MakeLogEnabledHandler(viewHandler)).Methods("GET")
}

func rootHandler(rw http.ResponseWriter, r *http.Request) {
	showView(views.DefaultViewId, rw, r, nil)
}

func viewHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	viewId := vars["viewId"]
	if viewId == "" {
		viewId = views.DefaultViewId
		logs.Debugf("The view id is %s", viewId)
	}
	showView(viewId, rw, r, nil)
}

func showView(viewId string, rw http.ResponseWriter, r *http.Request, model interface{}) {
	logs.Debugf("Show view %s", viewId)
	if viewId != "" {
		view, err := views.GetView(viewId)
		if err != nil {
			logs.Errorf("Failed to get view with name %s, the error is %v", viewId, err)
			replyOK(rw)
			return
		}
		tpl, err := views.GetTemplate(viewId)
		if err != nil {
			logs.Errorf("Failed to get view template with name %s, the error is %v", viewId, err)
			replyOK(rw)
			return
		}
		data := make(map[string]interface{}, 0)
		commonTemplateData := getCommonTemplateData(view)
		data["Common"] = commonTemplateData

		if model != nil {
			data["params"] = model
		}

		err = tpl.ExecuteTemplate(rw, view.View, data)
		if err != nil {
			logs.Errorf("Failed to execute template, the error is %v", err)
			replyOK(rw)
		}
		return
	}
	replyOK(rw)
}

func replyOK(rw http.ResponseWriter) {
	rw.Header().Set("Content-Type", "text/plain;charset=UTF-8")
	rw.Write([]byte("OK"))
}

func getCommonTemplateData(view views.ViewDefinition) map[string]interface{} {
	data := make(map[string]interface{}, 0)

	data["PageTitle"] = view.PageTitle
	data["PageHeader"] = view.PageHeader
	data["Scripts"] = view.Scripts
	data["Styles"] = view.Styles

	appBrand, err := config.Get("app_brand")
	if err == nil {
		data["AppBrand"] = appBrand
	}

	return data
}
