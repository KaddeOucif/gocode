package main
import ("github.com/ant0ine/go-json-rest/rest"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"os")
func main() {
    // try making database connection and query
    db, err := sql.Open("mysql", "Captain:welcome1@tcp(http://129.157.179.180:3000)/deathstar")
    defer db.Close()

    if err != nil {
        fmt.Println("Failed to connect", err)
        return
    }
	
	var xCoordinate int
	var yCoordinate int
	err = db.QueryRow("SELECT * from SecretTable").Scan(&xCoordinate, &yCoordinate)
	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		fmt.Println("Failed to get results" ,err)
	default:
        fmt.Println(xCoordinate, yCoordinate)
	}
	
	// try making http request against sample site
	for i := 0; i < 10; i++ {
		response, err := http.Get("http://129.157.179.180:3000/fighters/shield/45/" + i + "/blue/KaddeOucif")
	if err != nil {
		fmt.Println("Error making GET call")
	} else {
	    defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
		  fmt.Println("Error parsing GET call")
		} else {
		  fmt.Printf("%s\n", string(contents))
		}
		
	}
	}
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"Body": "Figher is running!"})
	}))

	fmt.Println(http.ListenAndServe(":"+os.Getenv("PORT"), api.MakeHandler()))
}
