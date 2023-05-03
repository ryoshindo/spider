package spider

import "io/ioutil"

func (s *App) readDefinitionFile(path string) ([]byte, error) {
	// return s.loader.ReadWithEnv(path)
	return ioutil.ReadFile(path)
}
