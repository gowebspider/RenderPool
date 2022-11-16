package redisstorage

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

// Storage implements the redis storage backend
type Storage struct {
	// Address is the redis server address
	Address string
	// Password is the password for the redis server
	Password string
	// DB is the redis database. Default is 0
	DB int
	// Prefix is an optional string in the keys. It can be used
	// to use one redis database for independent scraping tasks.
	Prefix string
	// Client is the redis connection
	Client *redis.Client

	// Expiration time for Visited keys. After expiration pages
	// are to be visited again.
	Expires time.Duration

	mu sync.RWMutex // Only used for cookie methods.
}

// Init initializes the redis storage
func (s *Storage) Init() error {
	if s.Client == nil {
		s.Client = redis.NewClient(&redis.Options{
			Addr:     s.Address,
			Password: s.Password,
			DB:       s.DB,
		})
	}
	if _, err := s.Client.Ping().Result(); err != nil {
		return fmt.Errorf("redisstorage: Redis connection error %s", err.Error())
	}
	return nil
}

// Clear removes all entries from the storage
func (s *Storage) Clear() error {
	return s.Client.Del(s.Prefix).Err()
}

// Visited implements colly/storage.Visited()
func (s *Storage) Visited(requestID uint64) error {
	return s.Client.Set(s.getIDStr(requestID), "1", s.Expires).Err()
}

// IsVisited implements colly/storage.IsVisited()
func (s *Storage) IsVisited(requestID uint64) (bool, error) {
	_, err := s.Client.Get(s.getIDStr(requestID)).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Storage) AddRequest(r []byte) error {
	return s.Client.RPush(s.getQueueID(), r).Err()
}

func (s *Storage) GetRequest() ([]byte, error) {
	r, err := s.Client.LPop(s.getQueueID()).Bytes()
	if err != nil {
		return nil, err
	}
	return r, err
}

func (s *Storage) QueueSize() (int, error) {
	i, err := s.Client.LLen(s.getQueueID()).Result()
	return int(i), err
}

func (s *Storage) getIDStr(ID uint64) string {
	return fmt.Sprintf("%s:request:%d", s.Prefix, ID)
}

func (s *Storage) getCookieID(c string) string {
	return fmt.Sprintf("%s:cookie:%s", s.Prefix, c)
}

func (s *Storage) getQueueID() string {
	return fmt.Sprintf("%s:queue", s.Prefix)
}
