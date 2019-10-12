package ws

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	cache "github.com/patrickmn/go-cache"
)

/*
队列进行SSDB/REDIS持久化存储
**/
type GlobalProcess struct {
	Process    map[string]JobProcess
	ProceMutex sync.RWMutex
}

type GlobalUserMessage struct {
	Message  map[string]UserMessage
	MsgMutex sync.Mutex
}

type JobProcess struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Status  string  `json:"status"`
	Process float64 `json:"process"`
}

type UserMessage struct {
	Id        string        `json:"id"`
	Broadcast bool          `json:"broadcast"`
	Header    MessageHeader `json:"header"`
	Body      MessageBody   `json:"body"`
}

type MessageHeader struct {
	Source      string    `json:"source"`
	Target      string    `json:"target"`
	SourceIp    string    `json:"sourceip"`
	TargetIp    string    `json:"targetip"`
	Time        time.Time `json:"time"`
	ContentType string    `json:"contenttype"`
}

type MessageBody struct {
	Body []byte `json:"body"`
}

var UploadProcess GlobalProcess
var LocalCache *cache.Cache

func init() {
	log.Println("init message service")
	UploadProcess.Process = make(map[string]JobProcess, 1024)
	LocalCache = cache.New(5*time.Minute, 10*time.Minute)
}

var Wsupgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Wshandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err    error
		conn   *Connection
		projob JobProcess
	)

	if wsConn, err = Wsupgrader.Upgrade(w, r, nil); err != nil {
		return
	}
	if conn, err = InitConnection(wsConn); err != nil {
		conn.Close()
	}

	uid := strings.Replace(r.URL.Query().Get("uid"), "-", "", -1)

	PushJob(uid, JobProcess{
		Id:      uid,
		Name:    "open connection",
		Status:  "ready",
		Process: 0,
	})

	go func() {
		var err error
		for {
			if err = conn.WriteMessage([]byte("--heartbeat--")); err != nil {
				return
			}
			//SetRandomProcess(uid)
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	for {
		if job, ok := GetJobProcess(uid); ok {
			if job.Id == uid && projob != job {
				b, _ := json.Marshal(job)
				if err_w := conn.WriteMessage(b); err_w != nil {
					goto ERR
				}
				projob = job
			}
		}
	}

ERR:
	{
		CleanJob(uid)
		conn.Close()
	}
}

func SetRandomProcess(uid string) {
	var (
		job JobProcess
		ok  bool
	)
	UploadProcess.ProceMutex.RLock()
	defer UploadProcess.ProceMutex.RUnlock()
	if job, ok = UploadProcess.Process[uid]; ok {
		job.Process = float64(rand.Intn(100))
		UploadProcess.Process[uid] = job
	}
}

func GetJobProcess(uid string) (job JobProcess, ok bool) {
	UploadProcess.ProceMutex.Lock()
	defer UploadProcess.ProceMutex.Unlock()
	job, ok = UploadProcess.Process[uid]
	return job, ok
}

func PushJob(uid string, job JobProcess) {
	UploadProcess.ProceMutex.Lock()
	UploadProcess.Process[uid] = job
	log.Println(UploadProcess.Process[uid])
	UploadProcess.ProceMutex.Unlock()
}

func CleanJob(uid string) {
	UploadProcess.ProceMutex.Lock()
	defer UploadProcess.ProceMutex.Unlock()
	delete(UploadProcess.Process, uid)
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
