package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/suatacikel/golang-rest-api/entity"
	"github.com/suatacikel/golang-rest-api/repository"
	"github.com/suatacikel/golang-rest-api/service"

	"github.com/stretchr/testify/assert"
)

const (
	id    int    = 1
	title string = "Title 1"
	text  string = "Text 1"
)

var (
	postRepo       repository.PostRepository = repository.NewSQLiteRepository()
	postSrv        service.PostService       = service.NewPostService(postRepo)
	postController PostController            = NewPostController(postSrv)
)

func TestAddPost(t *testing.T) {
	//create a new http post request
	jsonStr := []byte(`{"title": ` + title + `, "text": ` + text + `}`)
	req, _ := http.NewRequest("POST", "./posts", bytes.NewBuffer(jsonStr))

	//http handler function
	handler := http.HandlerFunc(postController.AddPost)

	//record http response
	response := httptest.NewRecorder()

	//dispacth the http request
	handler.ServeHTTP(response, req)

	//assertions
	status := response.Code
	if status != http.StatusOK {
		t.Error(fmt.Printf("handler returned a wrong status code: got %d want %d", status, http.StatusOK))
	}

	//decode http response
	var post entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&post)

	//assert http response
	assert.NotNil(t, post.Id)
	assert.Equal(t, title, post.Title)
	assert.Equal(t, text, post.Text)

	// Cleanup database
	tearDown(post.Id)

}

func TestGetPosts(t *testing.T) {

	// Insert new post
	setup()

	// Create new HTTP request
	req, _ := http.NewRequest("GET", "/posts", nil)

	// Assing HTTP Request handler Function (controller function)
	handler := http.HandlerFunc(postController.GetPosts)
	// Record the HTTP Response
	response := httptest.NewRecorder()
	// Dispatch the HTTP Request
	handler.ServeHTTP(response, req)

	// Assert HTTP status
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode HTTP response
	var posts []entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&posts)

	// Assert HTTP response
	assert.Equal(t, id, posts[0].Id)
	assert.Equal(t, title, posts[0].Title)
	assert.Equal(t, text, posts[0].Text)

	// Cleanup database
	tearDown(id)
}

func tearDown(postID int) {
	var post entity.Post = entity.Post{
		Id: postID,
	}
	postRepo.Delete(&post)
}

func setup() {
	var post entity.Post = entity.Post{
		Id:    id,
		Title: title,
		Text:  text,
	}
	postRepo.Save(&post)
}
