package response

import (
	"context"
	"mainmod/api/grpc/proto"
	"mainmod/internal/request"
	"mainmod/internal/storage"
	"net/http"
	"strings"
	"time"
)

var stateandcode = make(map[string]string)

func Handler(res http.ResponseWriter, req *http.Request) {

	URL := req.RequestURI

	if strings.Contains(URL, "?code=") {
		ParseUrl(URL)
	}

}

func ParseUrl(URL string) {

	URLstateandcode := strings.Split(URL, "&")

	forFindCode := strings.Split(URLstateandcode[0], "=")
	forFindState := strings.Split(URLstateandcode[1], "=")

	state := forFindState[1]
	code := forFindCode[1]

	stateandcode[state] = code
}

type GRPServer struct {
	proto.UnimplementedRegistrServer
}

func (GRPServer) mustEmbedUnimplementedRegistrServer() {
	panic("unimplemented")
}

func (GRPServer) Reg(ctx context.Context, req *proto.Request) (*proto.Response, error) {

	time.Sleep(30 * time.Second)

	var result *storage.Response

	for i := 1; i <= 2; i++ {
		for chekState, Code := range stateandcode {
			if chekState == req.State {
				ch := make(chan *storage.Response)
				go RegisterUser(Code, req.Nickname, req.Region, ch)

				result = <-ch
				defer close(ch)
				i = 2
			}
		}
		if i == 1 {
			result = &storage.Response{AccessToken: "101", RefreshToken: "101", Err: "101"}
			//101 - Превышено время ожидания пользователя
		}
		time.Sleep(30 * time.Second)
	}
	return &proto.Response{AccessToken: result.AccessToken, RefreshToken: result.RefreshToken, Err: result.Err}, nil
}

func RegisterUser(code, nickname, region string, ch chan *storage.Response) {

	token, reftoken := request.Request(code)
	if token == "404" {
		ch <- &storage.Response{
			AccessToken:  token, //404 - ошибка запроса
			RefreshToken: reftoken,
			Err:          "404",
		}
	}

	result := request.RequestCheckNickname(token, nickname, region)

	if result == "303" || result == "404" { // 303 - ошибка, ник не совпал или не тот регион
		ch <- &storage.Response{
			AccessToken:  result,
			RefreshToken: result,
			Err:          result,
		}
	} else {
		ch <- &storage.Response{
			AccessToken:  token,
			RefreshToken: reftoken,
			Err:          "200", //200 - ok
		}
	}

}
