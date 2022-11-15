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
	s.mu.Lock()
	defer s.mu.Unlock()
	r := s.Client.Keys(s.getCookieID("*"))
	keys, err := r.Result()
	if err != nil {
		return err
	}
	r2 := s.Client.Keys(s.Prefix + ":request:*")
	keys2, err := r2.Result()
	if err != nil {
		return err
	}
	keys = append(keys, keys2...)
	keys = append(keys, s.getQueueID())
	return s.Client.Del(keys...).Err()
}

// QueueSize implements queue.Storage.QueueSize() function
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
