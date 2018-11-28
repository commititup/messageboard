package messages

import (
	 "strconv"
	 "fmt"
	 "regexp"
	 "net/http"
	 "../output"
	 "../db"
	 "log"
	 "errors"
	 "encoding/json"
	 "strings"
	 "github.com/gorilla/mux"
)

//FUNCTION FOR GETTING ALL THE MESSAGES
func AllMessages(w http.ResponseWriter, r *http.Request) {

	query := "select id,title,author from messages"
	rows, err := db.DbConn.Query(query)
	if err != nil {
		log.Printf(fmt.Sprintf("Error to execute db Query, err: %s: %s", err.Error(), query))
		output.Send(w, nil, errors.New("Unable to fetch messages"))
		return
	}

	log.Printf(fmt.Sprintf(" executing Query: %s", query))

	var data  []map[string]interface{} 
	for rows.Next() {
		record := make(map[string]interface{})
		var title, username string
		var id int

		err = rows.Scan(&id, &title, &username)

		if err != nil {
			log.Printf(fmt.Sprintf("Error to read values of db Query, err: %s", err.Error()))
			output.Send(w, nil, errors.New("Unable to fetch messages"))
			return
		}
		record["id"] = id
		record["title"] = title
		record["author"] = username
		data = append(data, record)
	}

	output.RawSend(w, data, nil)
}

//FUNCTION TO ADD A MESSAGE TO DB
func AddMessage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	type Form struct {
		Title   string
		Body string
		Author    string
	}

	var form Form

	err := decoder.Decode(&form)
	if err != nil {
		output.Send(w, nil, err)
	}

	// If form is incomplete, throw error
	if form.Title == "" || form.Body == "" ||form.Author == "" {
		log.Printf("submitted incomplete form" )
		output.Send(w, nil, errors.New("Incomplete form submitted"))
		return
	}

	form.Body = strings.Replace(form.Body, "'", "''", -1) // Replace single quote with two single quoutes to insert string in postgresql ('test
	//stri''ng')

	query := fmt.Sprintf("insert into messages (title,body,author) values('%s', '%s', '%s')", form.Title, form.Body, form.Author)

	_,err = db.DbConn.Exec(query)

	if err != nil {
		log.Printf("Error to query , err: %s", err.Error())
		output.Send(w, nil, errors.New("Unable to add message"))
		return
	}

	log.Printf(" added a new message in database")
	data := map[string]interface{}{"added": true}
	output.Send(w, data, nil)

}


//FUNCTION TO DELETE A MESSAGE FROM DATABASE
func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	type Form struct {
		Id int
	}

	var form Form

	err := decoder.Decode(&form)
	if err != nil {
		output.Send(w, nil, err)
	}

	// If form is incomplete, throw error
	if form.Id <= 0 {
		log.Printf(" submitted incomplete form ")
		output.Send(w, nil, errors.New("id is empty"))
		return
	}

	query := fmt.Sprintf("delete from messages where id=%d", form.Id)
	log.Printf("Executing delete for id=%d",form.Id)

	res,err := db.DbConn.Exec(query)

	if err != nil {
		log.Printf("Error to query messages, err: %s", err.Error())
		output.Send(w, nil, errors.New("Unable to delete message"))
		return
	}

	count,err := res.RowsAffected()
	if err !=nil {
		log.Printf("Error to query messages, err: %s", err.Error())
		output.Send(w, nil, errors.New("Unable to delete message"))
		return
	}

	if count < 1{
	output.Send(w, nil, errors.New(" message doesn't exist"))
	return
	}

	log.Printf(" deleted message id %d from messages", form.Id)
	data := map[string]interface{}{"deleted": true}
	output.Send(w, data, nil)
}


//GET DETAILS OF A PARTICULAR MESSAGE
func GetDetailMessage(w http.ResponseWriter, r *http.Request) {
        params := mux.Vars(r)
	i    := params["id"]
	id, err := strconv.Atoi(i)
	if err != nil {
		output.Send(w, nil, errors.New("id is of Invalid Type "))
                return

	}

	// If form is incomplete, throw error
	if id <= 0 {
		log.Printf("%d submitted incomplete form",id)
		output.Send(w, nil, errors.New("id is empty"))
		return
	}

	query := fmt.Sprintf("select * from messages where id=%d", id)	
	rows, err := db.DbConn.Query(query)
	if err != nil {
		log.Printf(fmt.Sprintf("Error to execute db Query, err: %s: %s", err.Error(), query))
		output.Send(w, nil, errors.New("Unable to fetch messages"))
		return
	}


	var data  []map[string]interface{} 
	for rows.Next() {
		record := make(map[string]interface{})
		var title, username,body string
		var ids int

		err = rows.Scan(&id, &title, &body,&username)

		if err != nil {
			log.Printf(fmt.Sprintf("Error to read values of db Query, err: %s", err.Error()))
			output.Send(w, nil, errors.New("Unable to fetch messages"))
			return
		}

		record["id"] = ids
		record["title"] = title
		record["description"] = body
		record["author"] = username
		record["pallindrome"] =IsPalindrome(body) 
		data = append(data, record)
	}

	output.RawSend(w, data, nil)
}


//CHECK FOR STRING TO BE A PALLINDROME OR NOT
func IsPalindrome(value string) bool {
    value = sanitize(value)
    for i := 0; i < len(value)/2; i++ {
        if value[i] != value[len(value)-i-1] {
            return false
        }
    }
    return true
}

func sanitize(value string) string {
    reg, _ := regexp.Compile("[^A-Za-z0-9]+")
    safe := reg.ReplaceAllString(value, "")
    return strings.ToLower(strings.Trim(safe, ""))
}
