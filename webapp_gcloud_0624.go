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
  gmap1 := r.FormValue("gmap1")
  gmap2 := r.FormValue("gmap2")

  //業務マップに関する処理
  type gmap struct {
    x, y string
  }

  var gmapStruct gmap
  gmapStruct.x = gmap1
  gmapStruct.y = gmap2

  var gmapPattern1 gmap
  gmapPattern1.x = ""
  gmapPattern1.y = ""

  var gmapPattern2 gmap
  gmapPattern2.x = "on"
  gmapPattern2.y = ""

  var gmapPattern3 gmap
  gmapPattern3.x = ""
  gmapPattern3.y = "on"

  var gmapPattern4 gmap
  gmapPattern4.x = "on"
  gmapPattern4.y = "on"

  var gmapResult string

  switch gmapStruct {
  case gmapPattern1:
    gmapResult = "Both are none."
  case gmapPattern2:
    gmapResult = "Above is on."
  case gmapPattern3:
    gmapResult = "Bottom is on."
  case gmapPattern4:
    gmapResult = "Both are on."
  }

  //


  //HTMLの文字列
  h := `
  <html><head><title>Governance Cloud</title></head><body>
    Welcome to "GovernanceCloud", produced by Amelys Co., Ltd.
    <br><br>
    <form action="/" method="post">
      <dev>
      <h3>業務の全体像の可視化</h3>
        <input type="checkbox" name="gmap1">
          業務の全体像を一覧できる資料がある<br>
        <input type="checkbox" name="gmap2">
          重要な業務の処理プロセスが書面などで可視化されている<br>
      </dev>

      チェックボックス：` + html.EscapeString(gmapResult) + `

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
