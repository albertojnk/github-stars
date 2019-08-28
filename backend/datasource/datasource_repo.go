package datasource

import (
	"golang-crud-spa/backend/model"
	"log"

	"github.com/globalsign/mgo/bson"
)

// CreateUserRepositories creates a new user with the repositories on DB
func CreateUserRepositories(username string, repositories []model.Repository) error {
	db := database
	user := model.User{}

	db.C("users").Find(bson.M{"_id": username}).One(&user)

	updateRepos := []model.Repository{}
	repoMap := make(map[int]model.Repository, 0)

	for _, git := range repositories {
		repoMap[git.ID] = model.Repository{
			ID:           git.ID,
			Name:         git.Name,
			Description:  git.Description,
			URL:          git.URL,
			Language:     git.Language,
			Tags:         git.Tags,
			TagSuggester: git.TagSuggester,
		}
	}

	for _, db := range user.Repositories {
		repoMap[db.ID] = db
	}
	for _, repo := range repoMap {
		updateRepos = append(updateRepos, repo)
	}

	_, err := db.C("users").Upsert(
		bson.M{"_id": username},
		bson.M{
			"$set": bson.M{"repositories": updateRepos},
		},
	)
	if err != nil {
		log.Printf("error creating repositories, err: %s", err)
		return err
	}

	log.Println("Repositories creation succeed")

	return nil
}

// ListUserRepositories gets the repositories of a user from DB
func ListUserRepositories(username string) (model.User, error) {
	db := database

	user := model.User{}

	err := db.C("users").Find(bson.M{"_id": username}).One(&user)
	if err != nil {
		log.Printf("something went wrong while accessing the DB, err: %s", err)
		return user, err
	}

	log.Println("List repositories succeed")

	return user, nil
}

// UpdateUserRepositoryTags upadte the tags of a specific repository and return the new values
func UpdateUserRepositoryTags(username string, repositoryID int, tags []string) (model.User, error) {
	db := database

	user := model.User{}

	err := db.C("users").Update(
		bson.M{"_id": username, "repositories.id": repositoryID},
		bson.M{
			"$set": bson.M{"repositories.$.tags": tags},
		},
	)
	if err != nil {
		log.Printf("error updating repository, err: %s", err)
		return user, err
	}

	err = db.C("users").Find(bson.M{"_id": username}).One(&user)
	if err != nil {
		log.Printf("error finding repository, err: %s", err)
		return user, err
	}

	return user, nil
}

// DeleteUserRepositoryTags delete all tags of a specific repository and return the new values
func DeleteUserRepositoryTags(username string, repoID int) (model.User, error) {
	db := database

	user := model.User{}

	_, err := db.C("users").Upsert(
		bson.M{"_id": username, "repositories.id": repoID},
		bson.M{
			"$set": bson.M{"repositories.$.tags": []string{}},
		},
	)
	if err != nil {
		log.Printf("error deleting tags, err: %s", err)
		return user, err
	}

	err = db.C("users").Find(bson.M{"_id": username}).One(&user)
	if err != nil {
		log.Printf("error finding repository, err: %s", err)
		return user, err
	}

	return user, nil
}
