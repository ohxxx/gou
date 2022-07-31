package packet_book

import (
	"errors"
	"strconv"
)

func ShowBookInfo(bookName string, bookPrice float64) (string, error) {
	if bookName == "" {
		return "", errors.New("bookName is empty")
	}
	if bookPrice < 0 {
		return "", errors.New("bookPrice is negative")
	}
	return "书名：" + bookName + "，价格：" + strconv.FormatFloat(bookPrice, 'f', 2, 64), nil
}
