package model

// StarredRepositories is the struct we will use to unmarshal the github response
type StarredRepositories struct {
	ID          int      `json:"id" bson:"id"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	URL         string   `json:"html_url" bson:"url"`
	Language    string   `json:"language" bson:"language"`
	Tags        []string `json:"tags" bson:"tags"`
}

// User is the struct that will define our users in the DB
type User struct {
	ID           string       `json:"_id" bson:"_id"`
	Repositories []Repository `json:"repositories" bson:"repositories"`
}

// Repository is the struct that will define our repositories in the DB
type Repository struct {
	ID          int      `json:"id" bson:"id"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	URL         string   `json:"html_url" bson:"url"`
	Language    string   `json:"language" bson:"language"`
	Tags        []string `json:"tags" bson:"tags"`
}
