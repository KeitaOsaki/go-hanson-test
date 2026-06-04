package main

import (
	"sync"
	"time"
)

// Cache は TTL 付きのキャッシュ。Set した値は ttl 経過後、
// time.AfterFunc が起動する「別ゴルーチン」によって自動削除される。
//
//	削除は非同期 → 呼び出し元は「いつ消えたか」を直接は知れない
//
// この “時間経過 × 非同期の副作用” をテストでどう待つかが STEP 03 のテーマ。
type Cache struct {
	mu sync.Mutex
	m  map[string]string
}

func NewCache() *Cache {
	return &Cache{m: make(map[string]string)}
}

// Set は key=val を保存し、ttl 経過後の自動削除をスケジュールする。
// 削除は time.AfterFunc により別ゴルーチンで非同期に実行される。
func (c *Cache) Set(key, val string, ttl time.Duration) {
	c.mu.Lock()
	c.m[key] = val
	c.mu.Unlock()

	time.AfterFunc(ttl, func() {
		c.mu.Lock()
		delete(c.m, key)
		c.mu.Unlock()
	})
}

// Get は値と存在有無を返す。
func (c *Cache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.m[key]
	return v, ok
}

func main() {}
