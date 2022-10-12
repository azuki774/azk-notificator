package sender

import (
	"azk-notificator/internal/model"
	"azk-notificator/internal/telemetry"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"

	"go.uber.org/zap"
)

type Sender struct {
	Logger     *zap.Logger
	SendClient SendClient
	ServerHost string
	ServerPort string
}

func (s *Sender) Run() (err error) {
	s.Logger.Debug("sender start")

	for {
		ctx := telemetry.NewCtxWithSpanID()
		l := telemetry.LoggerWithSpanID(ctx, s.Logger)

		// Fetch sending data
		url := "http://" + net.JoinHostPort(s.ServerHost, s.ServerPort) + "/dequeue"
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			l.Error("failed to create request", zap.Error(err))
		}

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			l.Error("failed to send the request", zap.Error(err))
		}

		if res.StatusCode != 200 {
			if res.StatusCode == 204 {
				l.Debug("not found notification sending")
				return nil
			} else {
				l.Error("dequeue unexpected status code")
				return errors.New("unexpected status code")
			}
		}

		b, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			l.Error("failed to read response", zap.Error(err))
		}

		var q model.Queue
		err = json.Unmarshal(b, &q)
		if err != nil {
			l.Error("failed to parse response", zap.Error(err))
		}

		q.Body = b

		err = s.SendClient.Send(ctx, q)
		if err != nil {
			l.Error("failed to send the notification", zap.Error(err))
			return err
		}

		l.Info("send the notification")
	}

}
