package controllers

import (
	"net/http"

	"github.com/FadyGamilM/photosharing/views"
)

// Genral method to wrap above all templates handlers
// INPUT => template struct
// OUTPUT => http.HandlerFunc
func StaticHandler(template views.Template) http.HandlerFunc {
   return func(w http.ResponseWriter, r *http.Request){
      template.Render(w, nil)
   }
}
