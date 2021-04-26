/*
	Created by XiaoInk at 2021/04/26 11:14:24
	GitHub: https://github.com/XiaoInk
*/

package service

import (
	"net/http"
	"time"
)

func NewClient(timeout time.Duration) *http.Client {
	return &http.Client{Timeout: timeout}
}
