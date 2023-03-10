package response

import (
	"context"
	"log"
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
				i = 2
			}
		}
		if i == 1 {
			result = &storage.Response{AccessToken: "101", RefreshToken: "101"}
			//101 - Превышено время ожидания пользователя
		}
		time.Sleep(30 * time.Second)
	}
	return &proto.Response{AccessToken: result.AccessToken, RefreshToken: result.RefreshToken}, nil
}

func RegisterUser(code, nickname, region string, ch chan *storage.Response) {

	token, reftoken, err := request.Request(code)
	if err != nil {
		log.Print(err)
		ch <- &storage.Response{
			AccessToken:  "404", //404 - ошибка запроса
			RefreshToken: "404",
		}
	}

	result := request.RequestCheckNickname(token, nickname, region)

	if result == "303" || result == "404" { // 303 - ошибка, ник не совпал
		ch <- &storage.Response{
			AccessToken:  result,
			RefreshToken: result,
		}
	} else {
		ch <- &storage.Response{
			AccessToken:  token,
			RefreshToken: reftoken,
		}
	}

}
