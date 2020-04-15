package env

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Load func loads all env variable from .env file
// placed at the project root.
// Each variable need to be putted in new line.
// Key and value need to be separated by `=`
//	KEY=value:
// 	DB_USER=test123
//
// allowed comments:
//
//	# some comment
//	DB_USER=test123
//
// inline comments are prohibited, will cause a bug:
//
//	DB_USER=test123 # wrong comment
// will set env as:
//	key= DB_USER
//	value= test123 # wrong comment
func Load() error {
	ef, err := os.Open(".env")
	if err != nil && err.Error() == "open .env: no such file or directory" {
		return nil
	}
	if err != nil {
		return err
	}
	defer ef.Close()
	s := bufio.NewScanner(ef)
	for s.Scan() {
		ev := strings.TrimSpace(s.Text())
		if ev == "" || strings.HasPrefix(ev, "#") {
			continue
		}
		if !strings.Contains(ev, "=") {
			return fmt.Errorf("env %v, do not contain separator `=`", ev)
		}
		kv := strings.SplitN(ev, "=", 2)
		os.Setenv(kv[0], kv[1])
	}
	return s.Err()
}
