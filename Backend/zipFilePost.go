package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func zipHandler(w http.ResponseWriter, r *http.Request) {
	filename := "sample2.txt"
	buf := new(bytes.Buffer)
	writer := zip.NewWriter(buf)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	f, err := writer.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write([]byte(data))
	if err != nil {
		log.Fatal(err)
	}
	err = writer.Close()
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.zip\"", filename))
	//io.Copy(w, buf)
	w.Write(buf.Bytes())
}

// func main() {
// 	http.HandleFunc("/zip", zipHandler)
// 	http.ListenAndServe(":8080", nil)
// }

/*
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
"os"
"io"
	"github.com/gorilla/mux"
)

// Contact struct (Model)
type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

// Init contacts var as a slice Contact struct
var contacts []Contact

// Add new contact
// func createContact(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var contact Contact
// 	_ = json.NewDecoder(r.Body).Decode(&contact)
// 	contacts = append(contacts, contact)
// 	json.NewEncoder(w).Encode(contact)
// }
*/
// Main function
func main() {
	// Init router

	fmt.Println("Hello")

	r := mux.NewRouter()

	// Route handles & endpoints

	r.HandleFunc("/zip", zipHandler).Methods("POST")
	//http.HandleFunc("/zip", zipHandler)
	//	http.ListenAndServe(":8080", nil)

	// Start server
	log.Fatal(http.ListenAndServe(":3000", r))
}

/*
func Upload(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Upload files\n")

    file, handler, err := r.FormFile("file")
    if err != nil {
        panic(err) //dont do this
    }
    defer file.Close()

    // copy example
    f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        panic(err) //please dont
    }
    defer f.Close()
    io.Copy(f, file)

}

*/
