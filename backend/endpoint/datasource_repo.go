package endpoint

import (
	"log"
	"strings"

	"github.com/golang-crud-spa/backend/search"

	"github.com/globalsign/mgo/bson"
)

// CreateUserRepositories creates a new user with the repositories on DB
func CreateUserRepositories(username string, repositories []StarredRepositories) error {
	db := Connect()

	// Writing down the results on the database
	err := db.C("users").Insert(
		bson.M{"_id": username, "repositories": repositories},
	)
	if err != nil && !strings.Contains(err.Error(), "dup key") {
		log.Printf("error creating repositories, err: %s", err)
		return err
	}

	log.Println("Repositories creation succeed")

	index, err := search.NewBleveMapping(username)
	if err != nil {
		log.Printf("something wrong happend, err: %s", err)
		return err
	}

	err = search.NewBleveIndex(index, repositories)
	if err != nil {
		log.Printf("something wrong happend, err: %s", err)
		return err
	}

	return nil
}

// ListUserRepositories gets the repositories of a user from DB
func ListUserRepositories(username string) ([]User, error) {
	db := Connect()

	users := []User{}

	err := db.C("users").Find(bson.M{"_id": username}).All(&users)
	if err != nil {
		log.Printf("something went wrong while accessing the DB, err: ", err)
		return nil, err
	}

	log.Println("List repositories succeed")

	return users, nil
}

// UpdateUserRepositoryTags upadte the tags of a specific repository and return the new values
func UpdateUserRepositoryTags(username string, repositoryID int, tags []string) (User, error) {
	db := Connect()

	err := db.C("users").Update(
		bson.M{"_id": username, "repositories.id": repositoryID},
		bson.M{
			"$set": bson.M{"repositories.$.tags": tags},
		},
	)
	if err != nil {
		log.Printf("error updating repository, err: %s", err)
		return User{}, err
	}

	user := User{}

	err = db.C("users").Find(bson.M{"_id": username}).One(&user)
	if err != nil {
		log.Printf("error finding repository, err: %s", err)
		return User{}, err
	}

	return user, nil
}

// DeleteUserRepositoryTags delete all tags of a specific repository and return the new values
func DeleteUserRepositoryTags(username string, repoID int) (User, error) {
	db := Connect()

	_, err := db.C("users").Upsert(
		bson.M{"_id": username, "repositories.id": repoID},
		bson.M{
			"$set": bson.M{"repositories.$.tags": []string{}},
		},
	)
	if err != nil {
		log.Printf("error deleting tags, err: %s", err)
		return User{}, err
	}

	user := User{}

	err = db.C("users").Find(bson.M{"_id": username}).One(&user)
	if err != nil {
		log.Printf("error finding repository, err: %s", err)
		return User{}, err
	}

	return user, nil
}
