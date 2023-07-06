package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/cblokkeel/hotel-reservation/db"
	"github.com/cblokkeel/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27030"

type testdb struct {
	db.UserStore
}

func (tdb *testdb) teardown(ctx context.Context) {
	tdb.UserStore.Drop(ctx)
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		t.Error(err)
	}

	return &testdb{
		UserStore: db.NewMongoUserStore(client, db.TestDbName),
	}
}

func TestInsertUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(context.Background())

	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)
	app.Post("/", userHandler.HandleInsertUser)

	params := types.CreateUserParams{
		Email:     "someemail@foo.com",
		FirstName: "James",
		LastName:  "Foo",
		Password:  "foobar1234",
	}

	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	var user types.User

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		t.Error(err)
	}
	if len(user.ID) == 0 {
		t.Error("excepted user id to be set")
	}
	if len(user.EncryptedPassword) > 0 {
		t.Error("expected the encrypted password to be null")
	}
	if user.Email != params.Email {
		t.Errorf("expected email: %s, got: %s\n", user.Email, params.Email)
	}
	if user.FirstName != params.FirstName {
		t.Errorf("expected first name: %s, got: %s\n", user.FirstName, params.FirstName)
	}
	if user.LastName != params.LastName {
		t.Errorf("expected last name: %s, got: %s\n", user.LastName, params.LastName)
	}
}
