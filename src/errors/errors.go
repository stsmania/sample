package errors

import "fmt"

func HandleError(err error) error {
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
	}
	return err
}
