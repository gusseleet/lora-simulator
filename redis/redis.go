package redis
import (
	"github.com/gomodule/redigo/redis"
	"fmt"
	"encoding/gob"
	"bytes"
	"github.com/brocaar/lorawan"
	"time"
)

type DeviceSession struct {
	// profile ids
	DeviceProfileID  string
	ServiceProfileID string
	RoutingProfileID string

	// session data
	DevAddr  lorawan.DevAddr
	DevEUI   lorawan.EUI64
	JoinEUI  lorawan.EUI64 // TODO: remove?
	NwkSKey  lorawan.AES128Key
	FCntUp   uint32
	FCntDown uint32

}

const (
	deviceSessionKeyTempl = "lora:ns:device:%s"  // contains the session of a DevEUI
)


const (
	redisMaxIdle        = 3
	redisIdleTimeoutSec = 240
)

func NewRedisPool(redisURL string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     redisMaxIdle,
		IdleTimeout: redisIdleTimeoutSec * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(redisURL)
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
			}
			return nil
		},
	}
}


func GetDeviceSession(p *redis.Pool, devEUI lorawan.EUI64) (DeviceSession, error) {
	var s DeviceSession

	c := p.Get()
	defer c.Close()

	val, err := redis.Bytes(c.Do("GET", fmt.Sprintf(deviceSessionKeyTempl, devEUI)))
	if err != nil {
		if err == redis.ErrNil {
			return s, redis.ErrNil
		}
		return s, err
	}

	err = gob.NewDecoder(bytes.NewReader(val)).Decode(&s)
	if err != nil {
		return s, err
	}


	return s, nil
}