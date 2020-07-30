// Создаем конфиг

package main

import (
	"flag"
	"fmt"
	"time"
)

// Config represents some external client config.
type Config struct {
	Timeout time.Duration
	Addr    string
}

func NewConfig(prefix string) func() *Config {
	var (
		addr    = flag.String(prefix+".addr", "", "")
		timeout = flag.Duration(prefix+"timeout", time.Second, "")
	)
	return func() *Config {
		return &Config{
			Addr:    *addr,
			Timeout: *timeout,
		}
	}
}

var (
	mailRuAddr        = flag.String("mailru.addr", "http://mail.ru", "")
	mailRuTimeout     = flag.Duration("mailru.timeout", time.Second, "")
	geekBrainsAddr    = flag.String("geekbrains.addr", "http://geekbrains.ru", "")
	geekBrainsTimeout = flag.Duration("geekbrains.timeout", 2*time.Second, "")
)

func main() {
	citymobilConfig := NewConfig("citymobil")
	flag.Parse()
	mailRuConfig := Config{
		Addr:    *mailRuAddr,
		Timeout: *mailRuTimeout,
	}
	geekBrainsConfig := Config{
		Addr:    *geekBrainsAddr,
		Timeout: *geekBrainsTimeout,
	}
	fmt.Println(mailRuConfig)
	fmt.Println(geekBrainsConfig)
	fmt.Println(*citymobilConfig())
	// citymobilConfig("http://citimobil.ru")
	fmt.Println(citymobilConfig())
}
