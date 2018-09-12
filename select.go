package main

import (
  "net/http"
  "log"
  "fmt"
  "os"
   "time"
  "html/template"
)


type RadioButton struct {
	Name       string
	Value      string
	IsDisabled bool
	IsChecked  bool
	Text       string
}

type PageVariables struct {
  PageTitle        string
  PageRadioButtons []RadioButton
  Answer           string
}

type HomeVariables struct {
	Date         string
	Time         string
}

func main() {
  http.HandleFunc("/", DisplayRadioButtons)
  http.HandleFunc("/selected", UserSelected)
  http.HandleFunc("/home", HomePage)
  http.HandleFunc("/test", Testfn)
  http.HandleFunc("/test2", Test02)
  http.HandleFunc("/test3", Test03)

//  log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("listening...")
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
        }
}

func DisplayRadioButtons(w http.ResponseWriter, r *http.Request){

   Title := "Que prefieres?"
   MyRadioButtons := []RadioButton{
     RadioButton{"animalselect", "cats", false, false, "Cats"},
     RadioButton{"animalselect", "dogs", false, false, "Dogs"},
   }

  MyPageVariables := PageVariables{
    PageTitle: Title,
    PageRadioButtons : MyRadioButtons,
    }

   t, err := template.ParseFiles("select.html")
   if err != nil {
     log.Print("template parsing error: ", err)
   }

   err = t.Execute(w, MyPageVariables)
   if err != nil {
     log.Print("template executing error: ", err) //log it
   }
}

func UserSelected(w http.ResponseWriter, r *http.Request){
  r.ParseForm()
  // r.Form is now either
  // map[animalselect:[cats]] OR
  // map[animalselect:[dogs]]
 // so get the animal which has been selected
  youranimal := r.Form.Get("animalselect")

  Title := "Tu animal preferido"
  MyPageVariables := PageVariables{
    PageTitle: Title,
    Answer : youranimal,
    }

    t, err := template.ParseFiles("select.html") //
    if err != nil { // 
      log.Print("template parsing error: ", err) //
    }

    err = t.Execute(w, MyPageVariables) //
    if err != nil { //
      log.Print("template executing error: ", err) //
    }
}


func HomePage(w http.ResponseWriter, r *http.Request){

    now := time.Now() //
    HomePageVars := HomeVariables{ //
      Date: now.Format("02-01-2006"),
      Time: now.Format("15:04:05"),
    }

    t, err := template.ParseFiles("homepage.html") //
    if err != nil { //
	  log.Print("template parsing error: ", err) //
	}
    err = t.Execute(w, HomePageVars) //
    if err != nil { //
	  log.Print("template executing error: ", err) //
	}
}

func Testfn(w http.ResponseWriter, r *http.Request){

    now := time.Now() //
    HomePageVars := HomeVariables{ //
      Date: now.Format("02-01-2006"),
      Time: now.Format("15:04:05"),
    }

    t, err := template.ParseFiles("fn01.html") //
    if err != nil { //
	  log.Print("template parsing error: ", err) //
	}
    err = t.Execute(w, HomePageVars) //
    if err != nil { //
	  log.Print("template executing error: ", err) //
	}
}

func Test02(w http.ResponseWriter, r *http.Request){

    now := time.Now() //
    HomePageVars := HomeVariables{ //
      Date: now.Format("02-01-2006"),
      Time: now.Format("15:04:05"),
    }

    t, err := template.ParseFiles("fn02.html") //
    if err != nil { //
	  log.Print("template parsing error: ", err) //
	}
    err = t.Execute(w, HomePageVars) //
    if err != nil { //
	  log.Print("template executing error: ", err) //
	}
}

func Test03(w http.ResponseWriter, r *http.Request){

    now := time.Now() //
    HomePageVars := HomeVariables{ //
      Date: now.Format("02-01-2006"),
      Time: now.Format("15:04:05"),
    }

    t, err := template.ParseFiles("fn03.html") //
    if err != nil { //
	  log.Print("template parsing error: ", err) //
	}
    err = t.Execute(w, HomePageVars) //
    if err != nil { //
	  log.Print("template executing error: ", err) //
	}
}


 func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
 }

