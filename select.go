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

//  log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("listening...")
 	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
 		log.Fatal("ListenAndServe: ", err)
        }
}

func DisplayRadioButtons(w http.ResponseWriter, r *http.Request){
 // Display some radio buttons to the user

   Title := "Which do you prefer?"
   MyRadioButtons := []RadioButton{
     RadioButton{"animalselect", "cats", false, false, "Cats"},
     RadioButton{"animalselect", "dogs", false, false, "Dogs"},
   }

  MyPageVariables := PageVariables{
    PageTitle: Title,
    PageRadioButtons : MyRadioButtons,
    }

   t, err := template.ParseFiles("select.html") //parse the html file homepage.html
   if err != nil { // if there is an error
     log.Print("template parsing error: ", err) // log it
   }

   err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
   if err != nil { // if there is an error
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

  Title := "Your preferred animal"
  MyPageVariables := PageVariables{
    PageTitle: Title,
    Answer : youranimal,
    }

 // generate page by passing page variables into template
    t, err := template.ParseFiles("select.html") //parse the html file homepage.html
    if err != nil { // if there is an error
      log.Print("template parsing error: ", err) // log it
    }

    err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
    if err != nil { // if there is an error
      log.Print("template executing error: ", err) //log it
    }
}


func HomePage(w http.ResponseWriter, r *http.Request){

    now := time.Now() // find the time right now
    HomePageVars := HomeVariables{ //store the date and time in a struct
      Date: now.Format("02-01-2006"),
      Time: now.Format("15:04:05"),
    }

    t, err := template.ParseFiles("homepage.html") //parse the html file homepage.html
    if err != nil { // if there is an error
	  log.Print("template parsing error: ", err) // log it
	}
    err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
    if err != nil { // if there is an error
	  log.Print("template executing error: ", err) //log it
	}
}

func Testfn(w http.ResponseWriter, r *http.Request){

    now := time.Now() // find the time right now
    HomePageVars := HomeVariables{ //store the date and time in a struct
      Date: now.Format("02-01-2006"),
      Time: now.Format("15:04:05"),
    }

    t, err := template.ParseFiles("fn01.html") //parse the html file homepage.html
    if err != nil { // if there is an error
	  log.Print("template parsing error: ", err) // log it
	}
    err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
    if err != nil { // if there is an error
	  log.Print("template executing error: ", err) //log it
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

