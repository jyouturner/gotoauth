package awssolution

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/jyouturner/gotoauth"
	"golang.org/x/oauth2"
)

//redisTokenStorage stores oauth tokens in Redis
type redisTokenStorage struct {
	ctx    context.Context
	client *redis.Client
	key    string
}

//NewRedisTokenStorage create a redis client to load or store oauth tokens, return TokenStorage
func NewRedisTokenStorage(ctx context.Context, c *redis.Client, k string) gotoauth.TokenStorage {
	return redisTokenStorage{
		ctx:    ctx,
		client: c,
		key:    k,
	}
}

//LoadToken load oauth token from Redis, implement the LoadToken of TokenStorage
func (p redisTokenStorage) LoadToken() (*oauth2.Token, error) {
	val, err := p.client.Get(p.ctx, p.key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get from redis %v", err)
	}
	tok := &oauth2.Token{}

	err = json.Unmarshal([]byte(val), &tok)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON from S3, %v", err)
	}
	return tok, nil

}

//SaveNewToken stores oauth token in Redis, implement the SaveNewToken of TokenStorage
func (p redisTokenStorage) SaveNewToken(token *oauth2.Token) error {
	body, err := json.Marshal(token)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data, %v", err)
	}

	err = p.client.Set(p.ctx, p.key, string(body), 0).Err()
	if err != nil {
		return fmt.Errorf("failed to save the token to redis %v", err)
	}
	return nil

}
