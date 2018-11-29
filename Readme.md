 #  App
![Alt text](https://sonarcloud.io/api/project_badges/measure?project=commititup_messageboard&metric=alert_status)

  # see [Demo](http://rohit.one)

This is A Demo App which support REST API. It is build using

    - GOLANG (Backend)
    - Vue js  (Frontend)
    - Postgres (Database)

# Features!
  - Lists received messages
  - Allows users to submit/post messages
  - Drag and drop images (requires your Dropbox account be linked)
  - Retrieves a specific message on demand, and determines if it is a
       palindrome.
  - Allows users to delete specific messages

# Architecture


we have nginx running in front of the backend API acting as a reverse proxy routing request to the backend.


Each Parts of API will run on separate docker containers and talk with each other through internal inbuilt docker networks

![Alt text](doc/Arch.png?raw=true "Architecture")

### Installation
     - Method 1:
        Run initial-node-bootstrap.sh with sudo privileges and it will take care of all the things
     - Method 2
         To run the app as a whole you should have docker and docker-compose installed
         See [How to install docker and docker-compose](https://medium.com/@praaveen/part-2-docker-ce-and-docker-compose-installation-with-ubuntu-ef7b16bd3531)


         After install is finished just run from root folder
	 ```sh
	 $ docker-compose up -d 
	 Starting postgresql_db_1
	 Starting postgresql_web_1
	 Starting postgresql_proxy_1
	 Attaching to postgresql_db_1, postgresql_web_1, postgresql_proxy_1
	 ```
	 This will will and start the containers and the app

	 Next create the db schema for the app

	 ```sh
	 $ cd conf/postgres/
	 $ ./sql.sh
	 Password for user messageuser: 
	 ```
## Development

Want to contribute? Great!

 All the App config is available in the conf directory

 For Backend changes you should have go installed.
 All the source code are here.

 ```sh
 $ cd backend/src/
 $ ls
 config  db  github.com  messages  output  server.go
 ```

 After code changes, build it and place it to backend/bin directory
 ```sh
 $ go build
 ```

 For Frontend Changes you should have node installed.
 All the backend codes are here
 ```sh
 $ cd frontend/src/
 $ ls App.vue  assets  components  functions  main.js
 ```
 After code changes , just build it and all done

 ```sh
  $ npm run build
  ```

# API Documentation
**Show  a single message**
----
  Returns json data about a single message.

* **URL**

  /message/:id

* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**
 
   `id=[integer]`

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200
    **Content:** `{
    "data": [
        {
            "author": "rohit",
            "description": "Testing",
            "id": 1,
            "pallindrome": false,
            "title": "testing"
        }
    ],
    "error": null,
    "success": true
}`
 
* **Error Response:**

  * **Code:** 404 NOT FOUND 
    **Content:** `{ error : "message doesn't exist" }`

 

**Delete a message**
----
  Returns json data about a message being deleted.

* **URL**

  /message/:id

* **Method:**

  `DELETE`
  
*  **URL Params**

   **Required:**
	`id=[integer]` 

* **Data Params**

  {"id":1}

* **Success Response:**

  * **Code:** 200
    **Content:** `{"data":{"deleted":true},"error":null,"success":true}`
 
* **Error Response:**

  * **Code:** 200 OK
    **Content:** `{"data":null,"error":" message doesn't exist","success":false}`
or
* **Code:** 200 OK
    **Content:** `{"data":null,"error":"EOF","success":false}{"data":null,"error":"id is empty","success":false}`



**List all messages**
----
  Returns json data about all messages stored.

* **URL**

  /messages

* **Method:**

  `GET`
  
*  **URL Params**
    None
   **Required:**
    None

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200
    **Content:** `{
    "data": [
        {
            "author": "Rohit",
            "id": 10,
            "title": "Test Title"
        },
        {
            "author": "Rohit",
            "id": 11,
            "title": "Test Title"
        }
    ],
    "error": null,
    "success": true
}`
 
* **Empty message list Response:**

  * **Code:** 200 OK
    **Content:** `{
    "data": null,
    "error": null,
    "success": true
}`



**Create a message**
----
  Returns json data about message being created.

* **URL**

  /message

* **Method:**

  `POST`
  
*  **URL Params**
    None
   
* **Data Params**
`{
	"title":"Test Title",
	"author":"Rohit",
	"body": "Test message desc"
}`
* **Success Response:**

  * **Code:** 200
    **Content:** `{"data":{"added":true},"error":null,"success":true}`
 
* **Error Response:**

  * **Code:** 200 OK
    **Content:** `{"data":null,"error":"Incomplete form submitted","success":false}`



