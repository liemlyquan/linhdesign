package hello

import (
  "html/template"
  "net/http"
  "fmt"

  "google.golang.org/appengine"
  "google.golang.org/appengine/mail"
  "google.golang.org/appengine/log"

)

func init() {
  http.HandleFunc("/", mainHandler)
  http.HandleFunc("/buk", bukHandler)
  http.HandleFunc("/caterall", caterallHandler)
  http.HandleFunc("/fireflies", firefliesHandler)
  http.HandleFunc("/helvetica", helveticaHandler)
  http.HandleFunc("/jibo", jiboHandler)
  http.HandleFunc("/lamont", lamontHandler)
  http.HandleFunc("/nhan", nhanHandler)
  http.HandleFunc("/owl", owlHandler)
  http.HandleFunc("/sport", sportHandler)
  http.HandleFunc("/thePalette", thePaletteHandler)
  http.HandleFunc("/ustair", ustairHandler)
}

var tpl = template.Must(template.ParseGlob("templates/*.html"))

func mainHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "GET" {
    err := template.Must(template.ParseGlob("*.html")).ExecuteTemplate(w, "index.html", nil)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
  } else {
    ctx := appengine.NewContext(r)
    name := r.FormValue("name")
    email := r.FormValue("email")
    phone := r.FormValue("phone")
    message := r.FormValue("message")
    const confirmMessage = `
    You have received a new message with following information
    Name: %s
    Email: %s
    Phone: %s
    Message: %s
    `
    msg := &mail.Message{
      Sender:  "linhdoboi@gmail.com",
      To:      []string{"linhdoboi@gmail.com"},
      Subject: "You have received a new message from linhdo.me",
      Body:    fmt.Sprintf(confirmMessage, name, email,phone, message),
    }
    if err := mail.Send(ctx, msg); err != nil {
        log.Errorf(ctx, "Couldn't send email: %v", err)
    }

    http.Redirect(w, r, "/", http.StatusSeeOther)

  }
}

func bukHandler(w http.ResponseWriter, r *http.Request) {
  err := tpl.ExecuteTemplate(w, "buk.html", nil)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func caterallHandler(w http.ResponseWriter, r *http.Request) {
  err := tpl.ExecuteTemplate(w, "caterall.html", nil)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func firefliesHandler(w http.ResponseWriter, r *http.Request) {
  err := tpl.ExecuteTemplate(w, "fireflies.html", nil)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func helveticaHandler(w http.ResponseWriter, r *http.Request) {
  err := tpl.ExecuteTemplate(w, "helvetica.html", nil)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func jiboHandler(w http.ResponseWriter, r *http.Request) {
  err := tpl.ExecuteTemplate(w, "jibo.html", nil)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func lamontHandler(w http.ResponseWriter, r *http.Request) {
  err := tpl.ExecuteTemplate(w, "lamont.html", nil)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func nhanHandler(w http.ResponseWriter, r *http.Request) {
  err := tpl.ExecuteTemplate(w, "nhan.html", nil)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func owlHandler(w http.ResponseWriter, r *http.Request) {
  err := tpl.ExecuteTemplate(w, "owl.html", nil)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func sportHandler(w http.ResponseWriter, r *http.Request) {
  err := tpl.ExecuteTemplate(w, "sport.html", nil)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func thePaletteHandler(w http.ResponseWriter, r *http.Request) {
  err := tpl.ExecuteTemplate(w, "thePalette.html", nil)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func ustairHandler(w http.ResponseWriter, r *http.Request) {
  err := tpl.ExecuteTemplate(w, "ustair.html", nil)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}
