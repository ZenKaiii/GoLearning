package filelisting

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return errors.New("path must start with "+ prefix)
	}
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil{
		//panic(err)
		//http.Error(writer,
		//	err.Error(),
		//	http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}

	writer.Write(all)
	return nil
}
