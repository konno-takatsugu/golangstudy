package main

import (
  "fmt"
  "html"  //htmlパッケージ
  "math/big"
  "net/http"
)

//ServeHTTPメソッド用の構造体型
type Server struct{}

//httpリクエストを受け取るメソッド
func (Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  //フォームの入力値を取得
  left := r.FormValue("left")
  right := r.FormValue("right")
  op := r.FormValue("op")

  gmap1 := r.FormValue("gmap1")
  gmap2 := r.FormValue("gmap2")

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

  //文字列を整数に変換
  leftInt := &big.Int{}
  rightInt := &big.Int{}
  _, leftOK := leftInt.SetString(left, 10)
  _, rightOK := rightInt.SetString(right, 10)
  //四則演算の処理
  //変換エラーがなければ、演算子に従って計算
  var result string
  if leftOK && rightOK {
    resultInt := &big.Int{}
    //演算子ごとに分岐
    switch op {
    case "add": //加算
      resultInt.Add(leftInt, rightInt)
    case "sub": //減算
      resultInt.Sub(leftInt, rightInt)
    case "multi": //乗算
      resultInt.Mul(leftInt, rightInt)
    case "div": //除算
      resultInt.Div(leftInt, rightInt)
  }
  result = resultInt.String()
}

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
      左項目：<input type="text" name="left"><br>
      右項目：<input type="text" name="right"><br>
      演算子：
      <input type="radio" name="op" value="add" checked> ＋
      <input type="radio" name="op" value="sub"> ー
      <input type="radio" name="op" value="multi"> ×
      <input type="radio" name="op" value="div"> ÷
      <br><input type="submit" name="送信"><hr>

      [フォームの入力値]<br>
      左項目：` + html.EscapeString(left) + `<br>
      右項目：` + html.EscapeString(right) + `<br>
      演算子：` + html.EscapeString(op) + `<hr>

      演算結果：` + html.EscapeString(result) + `<hr>

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
