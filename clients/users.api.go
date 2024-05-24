package clients

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"fmt"
    "io/ioutil"
	"net/http"
    "net/url"
    "os"
    "strconv"
)

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Name 	  string `json:"first_name"`
}

type Support struct {
	URL		  string `json:"url"`
	Text	  string `json:"text"`
}

type Response struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	Support Support `json:"support"`
	Data    []User `json:"data"`
}

var baseUri string

func init() {

	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf("failed to load .env file: %v", err))
	}

	baseUri = os.Getenv("ACCOUNTS_IMPOTER_BASE_URI")
}

func GetUserAccounts(startDate string, endDate string, offset int, limit int) Response {
	
	u, _ := url.Parse(baseUri + "/users")
	q := u.Query()
	q.Set("start_date", startDate)
	q.Set("end_date", endDate)
	q.Set("page", strconv.Itoa(offset))
	q.Set("per_page", strconv.Itoa(limit))
	u.RawQuery = q.Encode()
	fmt.Println(u)
	
	// Get request
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {  // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	// fmt.Println(PrettyPrint(result))

	// Loop through the data node for the FirstName
	for _, rec := range result.Data {
		fmt.Println(rec.Name)
	}
	
	return result
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
