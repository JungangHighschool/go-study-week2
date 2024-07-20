package post

import "strings"

type ConsumedPost struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ProducedPost struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

func NewProducedPost(id int, status string) *ProducedPost {
	p := ProducedPost{Id: id, Status: status}

	return &p
}

var ForbiddenWords = []string{"defense", "communication", "command",
                              "tomato", "banana", "haesung", "mississippi",
                              "ishtar", "iris", "skku"}
// var piMap = map[string][]int{}

func isExist(s string, tgt string, out chan<-bool) {
  out<-strings.Contains(s, tgt)

//   pi := piMap[tgt]
//   if pi == nil {
//     pi = make([]int, len(tgt))
//     pi[0] = 0
//     j:=0
//     for i:=1; i<len(tgt); i++ {
//       for j > 0 && tgt[j] != tgt[i] {
//         j = pi[j-1]
//       }

//       if tgt[j] == tgt[i] {
//         pi[i] = j+1
//         j++
//       } else {
//         pi[i] = 0
//       }
//     }
//     piMap[tgt] = pi
//   }

//   for i, j := 0, 0; i<len(s); i++ {
//     if(s[i] == tgt[j]){
//       j++
//       if(j == len(tgt)){
//         out <- true
//         return
//       }
//     } else if j>0{
//       j = pi[j-1]
//     }
//   }

//   out <- false
}


func checkForbiddenWords(s string) string {
  resultCh := make(chan bool)

  for _, tgt := range ForbiddenWords {
    go isExist(s, tgt, resultCh)
  }

  for i := 0; i<len(ForbiddenWords); i++ {
    if <-resultCh{
      return "Failed"
    }
  }

  close(resultCh)

  return "Success"
}

func ProcessPost(msg ConsumedPost) ProducedPost {
  id := msg.Id
  status := checkForbiddenWords(msg.Title)

  return *NewProducedPost(id, status)
}
