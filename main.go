package main

import (
    "encoding/json"
    "strconv"
    "os"
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
	"math/rand"
	"time"
)

// our main function
func main() {	
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	} else {
		fmt.Sprintf("<%s>", port)
		log.Println("port - "+port)
	}

    router := mux.NewRouter()
    router.HandleFunc("/health",statusHandler).Methods("GET")
    router.HandleFunc("/name",contactHandler).Methods("POST")
    router.HandleFunc("/move",moveHandler).Methods("POST")	

    log.Fatal(http.ListenAndServe(":"+port, router))
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - /health")
	data, _ := json.Marshal(HealthCheckResponse{Status: "UP"})
    writeJsonResponse(w, http.StatusOK, data)
}

func contactHandler(w http.ResponseWriter, r *http.Request){
	log.Println("POST - /name")

	data, _ := json.Marshal(Contact{Name:"DareDave", Email:"david.rodriguez@privalia.com"})
	writeJsonResponse(w, http.StatusOK, data)
}

func moveHandler(w http.ResponseWriter, r *http.Request){
	log.Println("POST - /move")

	rand.Seed(time.Now().UnixNano())
	MOVEMENTS := [] string{ "up",	 "down", "left", "right", "fire-up", "fire-down", "fire-right", "fire-left"}
	var mIndex int = rand.Intn(len(MOVEMENTS))
//	params := mux.Vars(r)
    // var game Game
    // _ = json.NewDecoder(r.Body).Decode(&game)

    //  var player Player
    // _ = json.NewDecoder(r.Body).Decode(&player)

    //  var board Board
    // _ = json.NewDecoder(r.Body).Decode(&board)

 // var moverequest MoveRequest
  //  _ = json.NewDecoder(r.Body).Decode(&moverequest)

 // var game = moverequest.Game


	data, _ := json.Marshal(MoveResponse{Move: MOVEMENTS[mIndex]})
	log.Println("bodyRequest - /move" + string(data))
	//log.Println("bodyRequest - /move" + string(game.id))
	writeJsonResponse(w, http.StatusOK, data)
}



func writeJsonResponse(w http.ResponseWriter, status int, data []byte) {
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Content-Length", strconv.Itoa(len(data)))
        w.WriteHeader(status)
        w.Write(data)
}


type HealthCheckResponse struct {
        Status string `json:"status"`
}

type MoveResponse struct {
	Move string `json:"move"`
}

type Contact struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

type MoveRequest struct {
	Game *Game `json:"game"`
	Player *Player `json:"player"`
	Board *Board `json:"board"`
	Players []*Position `json:"players"`
	Invaders []*Invader `json:"invaders"`
}

type Game struct{
	id string 
}

type Board struct{
	size *Size `json:"size"`
	walls []*Position `json:"walls"`
}

type Player struct{
	id string `json:"id"`
	name string `json:"name"`
	position *Position `json:"position"`
	area *Area `json:"area"`
	fire bool `json:"fire"`
}

type Area struct{
	x1 int 
	x2 int
	y1 int
	y2 int
}

type Position struct {
	x int `json:"x"`
	y int `json:"y"`
}

type Invader struct {
	x int `json:"x"`
	y int `json:"y"`
	neutral bool `json:"neutral`
}

type Size struct{
	width int `json:"width"`
	height int `json:"height"`
}