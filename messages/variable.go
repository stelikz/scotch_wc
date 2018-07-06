package messages

type Message struct {
    Username string `json:"username"`
    Message string `json:"message"`
    ID uint64 `json:"id"`

}

type Messages []Message