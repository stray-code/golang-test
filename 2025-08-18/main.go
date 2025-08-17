package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// ルートパス "/" へのリクエストを処理するハンドラ関数を登録します。
	// http.HandleFunc は、特定のパスへのリクエストを処理する関数を登録するGoの標準ライブラリの機能です。
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// http.ResponseWriter はクライアントへのレスポンスを書き込むために使用します。
		// *http.Request はクライアントからのリクエストに関する情報を含みます。

		// fmt.Fprintf を使用して、レスポンスとして「Hello, World!」を書き込みます。
		// 第二引数の w は、レスポンスを書き込む対象である http.ResponseWriter です。
		queryParams := r.URL.Query()
		name := queryParams.Get("name")

		if name == "" {
			name = "Guest"
		}

		fmt.Fprintf(w, "Hello, %s", name)
	})

	// サーバーをポート8080で起動します。
	// http.ListenAndServe は、指定されたポートでTCPアドレスをリッスンし、着信リクエストを処理するサーバーを起動します。
	log.Println("サーバーがポート8080で起動しました...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
