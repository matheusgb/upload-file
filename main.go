package main

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/firestore"
	cloud "cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/matheusgb/pocPostFile/handlers"
	routing "github.com/rmnoff/fasthttp-routing/v3"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/option"
)

type App struct {
	ctx     context.Context
	client  *firestore.Client
	storage *cloud.Client
}

func main() {
	application := App{}
	application.Init()

	router := routing.New()
	router.Get("/", handlers.Get)
	router.Post("/", application.Post)

	fasthttp.ListenAndServe(":3000", router.HandleRequest)
}

func (application *App) Init() {
	application.ctx = context.Background()
	serviceAccount := option.WithCredentialsFile("./serviceAccountKey.json")

	app, _ := firebase.NewApp(application.ctx, nil, serviceAccount)
	application.client, _ = app.Firestore(application.ctx)
	application.storage, _ = cloud.NewClient(application.ctx, serviceAccount)
}

func (application *App) Post(c *routing.Context) error {
	file, _ := c.FormFile("image")
	fileName := file.Filename
	bucket := "pocpostgolang.appspot.com"

	newWriter := application.storage.Bucket(bucket).Object(fileName).NewWriter(c)
	newWriter.ContentType = "image/jpeg"

}

func respondWithJSON(c *routing.Context, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	c.SetContentType("application/json")
	c.SetStatusCode(code)
	c.Write(response)
}
