package db

import (
	"testing"

	"github.com/RomeroGabriel/go-api-standards/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db := createDb()
	userDB := NewUserDB(db)
	user, _ := entity.NewUser("Gabriel", "g@g.com", "123456")
	err := userDB.Create(user)
	assert.Nil(t, err)
	deleteDbTest()
}

func TestUserFindByEmail(t *testing.T) {
	db := createDb()

	userDB := NewUserDB(db)
	user, _ := entity.NewUser("Gabriel", "g@g.com", "123456")
	err := userDB.Create(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail("g@g.com")
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
	deleteDbTest()
}
