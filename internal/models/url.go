package models

type Url struct {
	TinyUrl string `json:"tinyUrl" bson:"tinyUrl"`
	LongUrl string `json:"longUrl" bson:"longUrl"`
}
