package spider

import (
	"os"
)

func (s *App) readDefinitionFile(path string) ([]byte, error) {
	// return s.loader.ReadWithEnv(path)
	return os.ReadFile(path)
}
