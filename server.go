package main

import (
	"os"

	"github.com/suatacikel/golang-rest-api/controller"
	router "github.com/suatacikel/golang-rest-api/http"
	"github.com/suatacikel/golang-rest-api/repository"
	"github.com/suatacikel/golang-rest-api/service"
)

var (
	PostRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(PostRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	os.Setenv("PORT", "8000")

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(os.Getenv("PORT"))
}
