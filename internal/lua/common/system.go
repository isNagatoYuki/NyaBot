package common

import (
	"github.com/Elyart-Network/NyaBot/pkg/fastlib/system"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
	"log"
	"os"
	"time"
)

type SystemFunc struct{}

func (s *SystemFunc) Info() system.Info {
	return system.AllInfo()
}

func (s *SystemFunc) Exit(code int) {
	os.Exit(code)
}

func (s *SystemFunc) Sleep(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func (s *SystemFunc) GetEnv(key string) string {
	return os.Getenv(key)
}

func (s *SystemFunc) SetEnv(key, value string) {
	err := os.Setenv(key, value)
	if err != nil {
		log.Println("SetEnv error:", err)
	}
}

func (s *SystemFunc) UnsetEnv(key string) {
	err := os.Unsetenv(key)
	if err != nil {
		log.Println("UnsetEnv error:", err)
	}
}

func System(L *lua.LState) {
	var SystemFunc = &SystemFunc{}
	L.SetGlobal("Sys", luar.New(L, SystemFunc))
}