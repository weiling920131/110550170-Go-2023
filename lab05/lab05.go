package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"os"
	"html/template"
	"strings"
)

func changetimeformat(inpu interface {}) interface{}{
	input := inpu.(string)
	input = strings.Split(input, "T")[0]
	inputs := strings.Split(input, "-")
	var output = inputs[0]+"年"+inputs[1]+"月"+inputs[2]+"日"
	return output
}

func changeformat(inpu interface {}) interface{}{
	input := inpu.(string)
	var l = len(input)
	var cnt = l / 3
	var i int
	var output = ""
	for i=0;i<cnt;i++{
		var tmp = input[l-3:l]
		output = ","+tmp+output
		l = l-3
	}
	output = input[0:l]+output
	if len(input)%3==0{
		output = output[1:]
	}
	return output
}
// TODO: Please create a struct to include the information of a video

type data struct{
	Title	interface {}
	Id		interface {}
	ChannelTitle	interface {}
	LikeCount		interface {}
	ViewCount		interface {}
	PublishedAt		interface {}
	CommentCount	interface {}

}


func YouTubePage(w http.ResponseWriter, r *http.Request) {
	// TODO: Get API token from .env file
	err := godotenv.Load()
	if err != nil{
		http.ServeFile(w, r, "error.html")
		return
	}
	password := os.Getenv("YOUTUBE_API_KEY") 

	// TODO: Get video ID from URL query `v`
	v := r.URL.Query().Get("v")
	if v == ""{
		http.ServeFile(w, r, "error.html")
		return
	}
	// TODO: Get video information from YouTube API
	urll := "https://www.googleapis.com/youtube/v3/videos?key="+password+"&id="+v+"&part=statistics,snippet"
	resp, err := http.Get(urll)
	if err != nil{
		http.ServeFile(w, r, "error.html")
		return
	}
	defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		http.ServeFile(w, r, "error.html")
		return
	}

	// // TODO: Parse the JSON response and store the information into a struct
	var m map[string]interface{}
	err = json.Unmarshal(body, &m)
	if err != nil{
		http.ServeFile(w, r, "error.html")
		return
	}
	if len(m["items"].([]interface{})) ==0{
		http.ServeFile(w, r, "error.html")
		return;
	}
	test := m["items"].([]interface{})[0].(map[string]interface{})
	snippet :=test["snippet"].(map[string]interface{})
	statistics := test["statistics"].(map[string]interface{})

	if snippet["title"] ==nil || snippet["channelTitle"] ==nil|| statistics["likeCount"] ==nil|| statistics["viewCount"] ==nil|| snippet["publishedAt"] ==nil|| statistics["commentCount"]==nil{
		http.ServeFile(w, r, "error.html")
		return
	}

	// TODO: Display the information in an HTML page through `template`

	var stru = data{
		Title	: snippet["title"],
	 	Id:	v,
		ChannelTitle	: snippet["channelTilte"],
		LikeCount		: changeformat(statistics["likeCount"]),	
		ViewCount		: changeformat(statistics["viewCount"]),
		PublishedAt		: changetimeformat(snippet["publishedAt"]),
		CommentCount	: changeformat(statistics["commentCount"]),
	}
	err = template.Must(template.ParseFiles("index.html")).Execute(w, stru)
}

func main() {
	http.HandleFunc("/", YouTubePage)
	log.Fatal(http.ListenAndServe(":8085", nil))
}
