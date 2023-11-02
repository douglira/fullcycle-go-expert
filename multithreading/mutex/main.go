package main

import (
	"fmt"
	"net/http"

	// "sync"
	"sync/atomic"
	"time"
)

var number uint64 = 0

/*
Como a função da HandleFunc pode ser executada simultaneamente
em cenários de milhares de acessos simultâneas, haverá uma certa
concorrência ao acessar o valor da variável "number". A chamada
"m.Lock()" garante que a variável esteja sendo incrementada para
cada acesso corretamente
*/
func main() {
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		// number++
		// m.Unlock()

		/*
			O package atomic possue funções que auxiliam na manipulação de
			operações atômicas. O código abaixo executa da mesma maneira
			que o código comentado acima.
		*/
		atomic.AddUint64(&number, 1)
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o visitante número: %d", number)))
	})
	http.ListenAndServe(":3000", nil)
}
