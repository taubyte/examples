package lib

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"image/png"
	"strings"
	"time"

	"github.com/mailru/easyjson"
	"github.com/o1egl/govatar"
	utils "github.com/tafseer-khan/tb_utils/err"
	srand "github.com/taubyte/go-sdk/crypto/rand"
	"github.com/taubyte/go-sdk/database"
	"github.com/taubyte/go-sdk/ethereum/client/logs"
	"github.com/taubyte/go-sdk/event"
	ipfs "github.com/taubyte/go-sdk/ipfs/client"
	pubsub "github.com/taubyte/go-sdk/pubsub/node"
)

//export _sleep
func sleep(dur time.Duration)

//export function
func function(e event.Event) uint32 {
	p, err := e.PubSub()
	if err != nil {
		panic(err)
	}

	channel, err := p.Channel()
	if err != nil {
		panic(err)
	}

	data, err := p.Data()
	if err != nil {
		return utils.Publish(channel, err)
	}

	_logs := logs.EventLog{}
	if err = _logs.UnmarshalJSON(data); err != nil {
		return utils.Publish(channel, err)
	}

	computeEvent := &computeEvent{}
	if err = computeEvent.UnmarshalJSON(_logs.Log.Data); err != nil {
		return utils.Publish(channel, err)
	}

	notifChannel, err := pubsub.Channel("notifications/" + strings.ToLower(computeEvent.Sender))
	if err != nil {
		return utils.Publish(channel, err)
	}

	kvdb, err := database.New("cidStore")
	if err != nil {
		return utils.Publish(notifChannel, err)
	}

	jobKey := fmt.Sprintf("jobs/%d", computeEvent.JobId.Int64())
	jobData, err := kvdb.Get(jobKey)
	if err != nil || len(jobData) == 0 {
		rand.Reader = srand.NewReader()
		jobIndexData := make([]byte, 16)
		if _, err = rand.Read(jobIndexData); err != nil {
			return utils.Publish(notifChannel, err)
		}

		if err = kvdb.Put(jobKey, jobIndexData); err != nil {
			return utils.Publish(notifChannel, err)
		}

		sleep(500 * time.Millisecond)

		jobData, err = kvdb.Get(jobKey)
		if err != nil {
			return utils.Publish(notifChannel, err)
		}

		if bytes.Equal(jobData, jobIndexData) {
			var genderParsed govatar.Gender
			switch computeEvent.Gender {
			case "male":
				genderParsed = govatar.MALE
			case "female":
				genderParsed = govatar.FEMALE
			default:
				if time.Now().UnixNano()%2 == 0 {
					genderParsed = govatar.FEMALE
				} else {
					genderParsed = govatar.MALE
				}
			}

			_avatar, err := govatar.Generate(genderParsed)
			if err != nil {
				return utils.Publish(notifChannel, err)
			}

			var buffer bytes.Buffer
			if err = png.Encode(&buffer, _avatar); err != nil {
				return utils.Publish(notifChannel, err)
			}

			ipfsClient, err := ipfs.New()
			if err != nil {
				return utils.Publish(notifChannel, err)
			}

			rw, err := ipfsClient.Create()
			if err != nil {
				return utils.Publish(notifChannel, err)
			}

			if _, err = rw.Write(buffer.Bytes()); err != nil {
				return utils.Publish(notifChannel, err)
			}

			cid, err := rw.Push()
			if err != nil {
				return utils.Publish(notifChannel, err)
			}

			if err := kvdb.Put(computeEvent.Sender+"/"+cid.String(), cid.Bytes()); err != nil {
				return utils.Publish(notifChannel, err)
			}

			json, err := easyjson.Marshal(&avatar{
				Cid: cid.String(),
			})
			if err != nil {
				return utils.Publish(notifChannel, err)
			}

			notifChannel.Publish(json)
		}
	}

	return 0
}
