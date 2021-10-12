package pkg

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Wget(filepath string, url string) (err error) {
	out, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
