package main

import (
  "fmt"
  "html"  //htmlパッケージ
  "net/http"
)

//ServeHTTPメソッド用の構造体型
type Server struct{}

//httpリクエストを受け取るメソッド
func (Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  //フォームの入力値を取得
  v := r.FormValue("input_value")

  //HTMLの文字列
  h := `
  <html><head><title>HTMLのフォーム</title></head><body>
    <form action="/" method="post">
      <input type="text" name="input_value">
      <input type="submit" name="送信"><br>
      入力値: ` + html.EscapeString(v) + `
    </form>
  </body></html>
  `

  //クライアント（ブラウザ）にHTMLを送信
  fmt.Fprint(w, h)
}

func main() {
  //webサーバを起動
  http.ListenAndServe(":4000", Server{})

}
