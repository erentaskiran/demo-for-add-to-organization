package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type application struct {
	Db  *sql.DB
	Cfg aws.Config
}

const (
	Sender    = "erentskrn7@gmail.com"
	Recipient = "eren.tas3535@gmail.com"
	Subject   = "asdasdasdasd"
	HtmlBody  = "<h1 style='color:blue'>Amazon SES Test Email (AWS SDK for Go)</h1><p>This email was sent with " +
		"<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
		"<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"
	TextBody = "This email was sent with Amazon SES using the AWS SDK for Go."
	CharSet  = "UTF-8"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db, err := newDb()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app := &application{
		Db:  db,
		Cfg: cfg,
	}

	mux := app.NewRouter()

	fmt.Println("server started at 3000 port.")

	http.ListenAndServe(":3000", mux)
}
