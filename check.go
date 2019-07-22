package otac

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func getExec() (string, error) {
	ex, err := os.Executable(); 
	if err != nil {
		return "", err
	}

	fi, err := os.Lstat(ex)
	if err != nil {
		return "", err
	}
	
	if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
		link, _ := filepath.EvalSymlinks(ex)
		return link, nil
	} else {
		return ex, nil
	}
}

func OTACheck(AppName, AppVersion, OTAUrl string) error {
	req, err := http.NewRequest("GET", OTAUrl, nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-App-Name", AppName)
	req.Header.Add("X-App-Version", AppVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.Header.Get("Content-Type") == "application/octet-stream" {
		ex, err := getExec()
		if err != nil {
			return err
		}

		if err := os.Remove(ex); err != nil {
			return err
		}
		
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(ex, body, 0777); err != nil {
			return err
		}
		return nil
	}
	return nil
}
