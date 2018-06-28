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

  mtg1 := r.FormValue("mtg1")
  mtg2 := r.FormValue("mtg2")
  mtg3 := r.FormValue("mtg3")

  sls1 := r.FormValue("sls1")
  sls2 := r.FormValue("sls2")

  ifm1 := r.FormValue("ifm1")
  ifm2 := r.FormValue("ifm2")
  ifm3 := r.FormValue("ifm3")

  dcm1 := r.FormValue("dcm1")
  dcm2 := r.FormValue("dcm2")

  ctm1 := r.FormValue("ctm1")
  ctm2 := r.FormValue("ctm2")

  stp1 := r.FormValue("stp1")
  stp2 := r.FormValue("stp2")

  act1 := r.FormValue("act1")
  act2 := r.FormValue("act2")

  shi1 := r.FormValue("shi1")

  its1 := r.FormValue("its1")
  its2 := r.FormValue("its2")



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
    gmapResult = "業務を改革・改善するには、業務の全体像が社内で共有されていることが必須です。アメリスとともに、まずは「業務マップ」を整理しましょう。"
  case gmapPattern2:
    gmapResult = "業務の全体像から、改善すべき「キー業務」を絞り込みましょう。まずはその業務について、アメリスとともに、業務プロセスの可視化をしましょう。"
  case gmapPattern3:
    gmapResult = "重要業務のプロセスを改善するために、それに関連する業務の可視化を、アメリスとともに始めましょう。"
  case gmapPattern4:
    gmapResult = "業務の全体像のうち、キーとなる業務について、アメリスとともに業務の「あるべき姿」を構築しましょう。"
  }

  //会議体に関する処理
  type mtg struct {
    x, y, z string
  }

  var mtgStruct mtg
  mtgStruct.x = mtg1
  mtgStruct.y = mtg2
  mtgStruct.z = mtg3

  var (
    mtgPattern1 mtg
    mtgPattern2 mtg
    mtgPattern3 mtg
    mtgPattern4 mtg
    mtgPattern5 mtg
    mtgPattern6 mtg
    mtgPattern7 mtg
    mtgPattern8 mtg
  )

  mtgPattern1.x = ""
  mtgPattern1.y = ""
  mtgPattern1.z = ""

  mtgPattern2.x = "on"
  mtgPattern2.y = ""
  mtgPattern2.z = ""

  mtgPattern3.x = "on"
  mtgPattern3.y = "on"
  mtgPattern3.z = ""

  mtgPattern4.x = "on"
  mtgPattern4.y = ""
  mtgPattern4.z = "on"

  mtgPattern5.x = "on"
  mtgPattern5.y = "on"
  mtgPattern5.z = "on"

  mtgPattern6.x = ""
  mtgPattern6.y = "on"
  mtgPattern6.z = ""

  mtgPattern7.x = ""
  mtgPattern7.y = "on"
  mtgPattern7.z = "on"

  mtgPattern8.x = ""
  mtgPattern8.y = ""
  mtgPattern8.z = "on"

  var mtgResult string

  switch mtgStruct {
  case mtgPattern1:
    mtgResult = "定例会議を経営のリズムの要として活用することで、組織は良くなります。 まずはチーム単位での週次の会議から始めましょう。 "
  case mtgPattern2:
    mtgResult = "定例会議を経営のリズムの要として活用することで、組織は良くなります。 週次会議の上位の定例会議として、月次の定例会議を設定しましょう。 "
  case mtgPattern3:
    mtgResult = "会議を有効活用するために「運営ルール」と「開催プロセス」を整備しましょう。 アメリスにはそれを実現・浸透させるノウハウがあります。 "
  case mtgPattern4:
    mtgResult = "一か月に1度程度、理念・ビジョンを再確認し、これから1か月の経営方針を伝える会議を行いましょう。 "
  case mtgPattern5:
    mtgResult = "定例会議をうまく運営されているようですね。 定例会議で集約される経営情報を、多方面の業務に活用できていますか。 "
  case mtgPattern6:
    mtgResult = "定例会議を経営のリズムの要として活用することで、組織は良くなります。 チーム単位の週次会議は全ての業務の基点になりますので、そこから始めましょう。 "
  case mtgPattern7:
    mtgResult = "定例会議は現場レベルでも実施することで、組織全体がまとまります。 チーム単位の週次会議は全ての業務の基点になりますので、そこから始めましょう。 "
  case mtgPattern8:
    mtgResult = "定例会議は現場レベルでも実施することで、組織全体がまとまります。 チーム単位の週次会議は全ての業務の基点になりますので、そこから始めましょう。 "
}

  //営業に関する処理
  type sls struct {
    x, y string
  }

  var slsStruct sls
  slsStruct.x = sls1
  slsStruct.y = sls2

  var slsPattern1 sls
  slsPattern1.x = ""
  slsPattern1.y = ""

  var slsPattern2 sls
  slsPattern2.x = "on"
  slsPattern2.y = ""

  var slsPattern3 sls
  slsPattern3.x = ""
  slsPattern3.y = "on"

  var slsPattern4 sls
  slsPattern4.x = "on"
  slsPattern4.y = "on"

  var slsResult string

  switch slsStruct {
  case slsPattern1:
    slsResult = "営業マンは忙しく、社内への進捗報告は後回しにされがちですが、組織で戦っていく上では営業進捗の社内共有は必須です。 定例会議等を仕組みとして取り入れつつ改革していきましょう。 "
  case slsPattern2:
    slsResult = "せっかく営業から共有される情報を、組織の資産として集約しましょう。 集約したうえで、必要な人がいつでもそこにアクセスできることが必要です。 "
  case slsPattern3:
    slsResult = "営業マンは忙しく、社内への進捗報告は後回しにされがちですが、組織で戦っていく上では営業進捗の社内共有は必須です。 営業マンに過度な負担をかけずに、仕組みを作り上げていきましょう。 "
  case slsPattern4:
    slsResult = "営業マンからの情報の収集、集約ができているのは素晴らしいことです。 営業の成果を月単位でレビューし、翌月の改善に活かしていきましょう。 "
  }

  //情報管理に関する処理
  type ifm struct {
    x, y, z string
  }

  var ifmStruct ifm
  ifmStruct.x = ifm1
  ifmStruct.y = ifm2
  ifmStruct.z = ifm3

  var (
    ifmPattern1 ifm
    ifmPattern2 ifm
    ifmPattern3 ifm
    ifmPattern4 ifm
    ifmPattern5 ifm
    ifmPattern6 ifm
    ifmPattern7 ifm
    ifmPattern8 ifm
  )

  ifmPattern1.x = ""
  ifmPattern1.y = ""
  ifmPattern1.z = ""

  ifmPattern2.x = "on"
  ifmPattern2.y = ""
  ifmPattern2.z = ""

  ifmPattern3.x = "on"
  ifmPattern3.y = "on"
  ifmPattern3.z = ""

  ifmPattern4.x = "on"
  ifmPattern4.y = ""
  ifmPattern4.z = "on"

  ifmPattern5.x = "on"
  ifmPattern5.y = "on"
  ifmPattern5.z = "on"

  ifmPattern6.x = ""
  ifmPattern6.y = "on"
  ifmPattern6.z = ""

  ifmPattern7.x = ""
  ifmPattern7.y = "on"
  ifmPattern7.z = "on"

  ifmPattern8.x = ""
  ifmPattern8.y = ""
  ifmPattern8.z = "on"

  var ifmResult string

  switch ifmStruct {
  case ifmPattern1:
    ifmResult = "文書の管理体系を定め、定期的な棚卸で規律を維持しましょう。 アメリスには、それを徹底するノウハウがあります。 "
  case ifmPattern2:
    ifmResult = "ファイルの命名ルールを統一することは、必要な情報の検索性を上げるために必須です。「コード」の活用も含めて、簡明な命名ルールを作っていきましょう。 "
  case ifmPattern3:
    ifmResult = "フォルダ管理体系、命名ルールが定まったら、次は資料の棚卸サイクルを作りましょう。このサイクルによって、ルール運用が徹底されます。 "
  case ifmPattern4:
    ifmResult = "ファイルの命名ルールを統一することは、必要な情報の検索性を上げるために必須です。「コード」の活用も含めて、簡明な命名ルールを作っていきましょう。 "
  case ifmPattern5:
    ifmResult = "情報管理の体制がしっかりと整っているとお見受けしました。整理された情報を各業務で活用していけるよう、各業務の「業務プロセス」を定めていきましょう。 "
  case ifmPattern6:
    ifmResult = "社内にある紙ファイル、電子ファイルを管理しきれているでしょうか。簡明なフォルダ管理体系を導入して、資料へのアクセス性を高めましょう。 "
  case ifmPattern7:
    ifmResult = "社内にある紙ファイル、電子ファイルを管理しきれているでしょうか。簡明なフォルダ管理体系を導入して、資料へのアクセス性を高めましょう。 "
  case ifmPattern8:
    ifmResult = "社内にある紙ファイル、電子ファイルを管理しきれているでしょうか。簡明なフォルダ管理体系を導入して、資料へのアクセス性を高めましょう。 "
}

  //決定に関する処理
  type dcm struct {
    x, y string
  }

  var dcmStruct dcm
  dcmStruct.x = dcm1
  dcmStruct.y = dcm2

  var dcmPattern1 dcm
  dcmPattern1.x = ""
  dcmPattern1.y = ""

  var dcmPattern2 dcm
  dcmPattern2.x = "on"
  dcmPattern2.y = ""

  var dcmPattern3 dcm
  dcmPattern3.x = ""
  dcmPattern3.y = "on"

  var dcmPattern4 dcm
  dcmPattern4.x = "on"
  dcmPattern4.y = "on"

  var dcmResult string

  switch dcmStruct {
  case dcmPattern1:
    dcmResult = "社内意思決定のルールを個別業務レベルで設定し、その履歴を残しておく仕組みを取り入れ、フェアな職場を作りましょう。アメリスなら、簡単にそれができるノウハウをご提供できます。 "
  case dcmPattern2:
    dcmResult = "社内意思決定ルールは具体的な業務シーンに沿ったものでしょうか。意思決定の履歴を記録していくことで、組織の意思決定の積み上げができます。 "
  case dcmPattern3:
    dcmResult = "社内意思決定のルールを個別業務レベルで設定し、その履歴を残しておく仕組みを取り入れ、フェアな職場を作りましょう。アメリスなら、簡単にそれができるノウハウをご提供できます。 "
  case dcmPattern4:
    dcmResult = "社内意思決定ルールは具体的な業務シーンに沿ったものでしょうか。重要な意思決定のルールは、アメリスのノウハウが詰まった「業務プロセス文書」の様式を用いて社内の誰もが閲覧できるようにしておきましょう。 "
  }

  //契約管理に関する処理
  type ctm struct {
    x, y string
  }

  var ctmStruct ctm
  ctmStruct.x = ctm1
  ctmStruct.y = ctm2

  var ctmPattern1 ctm
  ctmPattern1.x = ""
  ctmPattern1.y = ""

  var ctmPattern2 ctm
  ctmPattern2.x = "on"
  ctmPattern2.y = ""

  var ctmPattern3 ctm
  ctmPattern3.x = ""
  ctmPattern3.y = "on"

  var ctmPattern4 ctm
  ctmPattern4.x = "on"
  ctmPattern4.y = "on"

  var ctmResult string

  switch ctmStruct {
  case ctmPattern1:
    ctmResult = "契約書の管理は経営管理の基本といえるでしょう。契約書をきちんと管理したうえで契約条項の着実な履行ができる体制を作りましょう。 "
  case ctmPattern2:
    ctmResult = "締結した契約のリストを作成し、必要な人がすぐに閲覧できる体制、契約条項の着実な履行ができる体制を整備しましょう。 "
  case ctmPattern3:
    ctmResult = "契約書の管理は経営管理の基本といえるでしょう。契約書をきちんと管理したうえで契約条項の着実な履行ができる体制を作りましょう。 "
  case ctmPattern4:
    ctmResult = "契約書の管理体制が整っているようです。契約交渉から契約書の締結、契約書の保管まで、契約にまつわる一連の業務プロセスは整備されていますか。 "
  }

  //社判（社印）管理に関する処理
  type stp struct {
    x, y string
  }

  var stpStruct stp
  stpStruct.x = stp1
  stpStruct.y = stp2

  var stpPattern1 stp
  stpPattern1.x = ""
  stpPattern1.y = ""

  var stpPattern2 stp
  stpPattern2.x = "on"
  stpPattern2.y = ""

  var stpPattern3 stp
  stpPattern3.x = ""
  stpPattern3.y = "on"

  var stpPattern4 stp
  stpPattern4.x = "on"
  stpPattern4.y = "on"

  var stpResult string

  switch stpStruct {
  case stpPattern1:
    stpResult = "社印の管理は不正利用の防止上必須ですが、厳重にしすぎると業務効率が落ちてしまいます。アメリスにノウハウがあります。お問い合わせください。 "
  case stpPattern2:
    stpResult = "これまで社印を押捺した履歴を残しておくことで、外部とのやり取りの記録をすぐに閲覧することができます。 "
  case stpPattern3:
    stpResult = "社印の管理は不正利用の防止上必須ですが、厳重にしすぎると業務効率が落ちてしまいます。 アメリスにノウハウがあります。お問い合わせください。 "
  case stpPattern4:
    stpResult = "社印の管理体制が整っているようです。社印を押す前後の業務のプロセスは整備されていますか。 "
  }

  //請求・支払に関する処理
  type act struct {
    x, y string
  }

  var actStruct act
  actStruct.x = act1
  actStruct.y = act2

  var actPattern1 act
  actPattern1.x = ""
  actPattern1.y = ""

  var actPattern2 act
  actPattern2.x = "on"
  actPattern2.y = ""

  var actPattern3 act
  actPattern3.x = ""
  actPattern3.y = "on"

  var actPattern4 act
  actPattern4.x = "on"
  actPattern4.y = "on"

  var actResult string

  switch actStruct {
  case actPattern1:
    actResult = "顧客との請求・支払のやり取りは、経営管理上最も注意を払う必要があります。業務のルールおよびプロセスを整備して、安定した業務の進め方を整備しましょう。 "
  case actPattern2:
    actResult = "請求・支払情報を集約して管理することで、請求漏れ、支払漏れを防ぎ、予測可能性が生まれます。請求・支払案件に関わる全ての社員が情報集約に参加できる体制を作りましょう。 "
  case actPattern3:
    actResult = "顧客との請求・支払のやり取りは、経営管理上最も注意を払う必要があります。業務のルールおよびプロセスを整備して、安定した業務の進め方を整備しましょう。 "
  case actPattern4:
    actResult = "請求・支払管理の体制は整っているようです。新しく組織に参加する社員でもすぐに請求・支払管理ができるよう、業務プロセスを整備しましょう。 "
  }

  //浸透に関する処理
  type shi struct {
    x string
  }

  var shiStruct shi
  shiStruct.x = shi1

  var shiPattern1 shi
  shiPattern1.x = ""

  var shiPattern2 shi
  shiPattern2.x = "on"

  var shiResult string

  switch shiStruct {
  case shiPattern1:
    shiResult = "業務のルール、プロセスを整備するだけでは不十分です。ツールを整え、社員がルール、プロセスに従って行動できる組織に変えていきましょう。 "
  case shiPattern2:
    shiResult = "業務プロセスの浸透がなされていることは素晴らしいことです。定められた業務プロセスに沿って行動すれば、自然と業務の改善が行われるように、アメリスとともに仕組みを作り上げていきませんか。 "
  }

  //ITシステムに関する処理
  type its struct {
    x, y string
  }

  var itsStruct its
  itsStruct.x = its1
  itsStruct.y = its2

  var itsPattern1 its
  itsPattern1.x = ""
  itsPattern1.y = ""

  var itsPattern2 its
  itsPattern2.x = "on"
  itsPattern2.y = ""

  var itsPattern3 its
  itsPattern3.x = ""
  itsPattern3.y = "on"

  var itsPattern4 its
  itsPattern4.x = "on"
  itsPattern4.y = "on"

  var itsResult string

  switch itsStruct {
  case itsPattern1:
    itsResult = "ITシステムは業務プロセスに沿って構築されることで初めて役に立つものです。まずは既存ITシステムの社員の利用状況をモニタリングしましょう。 "
  case itsPattern2:
    itsResult = "ITシステムは業務プロセスに沿って構築されることで初めて役に立つものです。利用状況を元に、ITシステムの改善課題を洗い出しましょう。 "
  case itsPattern3:
    itsResult = "ITシステムは業務プロセスに沿って構築されることで初めて役に立つものです。まずは既存ITシステムの社員の利用状況をモニタリングしましょう。 "
  case itsPattern4:
    itsResult = "ITシステムは業務プロセスに沿って構築されることで初めて役に立つものです。ITシステムと業務の進め方の連携を常に最適化し続けていきましょう。"
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
      <dev>
      <h3>定例会議を結節点とした経営</h3>
        <input type="checkbox" name="mtg1">
          チーム単位での週次の定例会議がある<br>
        <input type="checkbox" name="mtg2">
          会社の経営方針を社内に伝える月次の定例会議がある<br>
        <input type="checkbox" name="mtg3">
          定例の会議について、運営ルールおよび開催プロセスが存在する<br>
      </dev>
      <dev>
      <h3>営業情報の有効活用</h3>
        <input type="checkbox" name="sls1">
          営業の進捗を社内に共有する仕組みがある<br>
        <input type="checkbox" name="sls2">
          営業情報を集約できている<br>
      </dev>
      <dev>
      <h3>情報管理</h3>
        <input type="checkbox" name="ifm1">
        紙ファイル、電子ファイルのフォルダ管理体系が定まっている<br>
        <input type="checkbox" name="ifm2">
          ファイル命名ルールが定まっている<br>
        <input type="checkbox" name="ifm3">
        資料の棚卸・再整理を定期的に行っている<br>
      </dev>
      <dev>
      <h3>意思決定</h3>
        <input type="checkbox" name="dcm1">
        社内意思決定ルールがある<br>
        <input type="checkbox" name="dcm2">
        過去の社内意思決定が保存されている <br>
      </dev>
      <dev>
      <h3>契約管理</h3>
        <input type="checkbox" name="ctm1">
        契約書が整理され管理されている<br>
        <input type="checkbox" name="ctm2">
        締結した契約のリストがすぐに閲覧できるようになっている<br>
      </dev>
      <dev>
      <h3>社判（社印）管理</h3>
        <input type="checkbox" name="stp1">
        社印の管理ルールが定まっている<br>
        <input type="checkbox" name="stp2">
        これまで社印を押した書類の履歴が残っている<br>
      </dev>
      <dev>
      <h3>請求・支払管理</h3>
        <input type="checkbox" name="act1">
        顧客への請求書発出、顧客からの請求書対応に関する業務のルールおよびプロセスが定まっている<br>
        <input type="checkbox" name="act2">
        今後の請求案件、支払情報が集約管理できている<br>
      </dev>
      <dev>
      <h3>業務プロセスの浸透</h3>
        <input type="checkbox" name="shi1">
        業務のルールおよびプロセスに沿って社員が行動できている<br>
      </dev>
      <dev>
      <h3>ITシステムの管理</h3>
        <input type="checkbox" name="its1">
        ITシステムの社員の利用状況を定期的にレビューしている<br>
        <input type="checkbox" name="its2">
        ITシステムに関する課題を洗い出せている<br>
      </dev>

      <br><input type="submit" name="送信"><hr>

      <h3>診断結果</h3>
      業務の全体像の可視化：<br>
      ` + html.EscapeString(gmapResult) + `<br><br>
      定例会議を結節点とした経営：<br>
      ` + html.EscapeString(mtgResult) + `<br><br>
      営業情報の有効活用：<br>
      ` + html.EscapeString(slsResult) + `<br><br>
      情報管理：<br>
      ` + html.EscapeString(ifmResult) + `<br><br>
      意思決定：<br>
      ` + html.EscapeString(dcmResult) + `<br><br>
      契約管理：<br>
      ` + html.EscapeString(ctmResult) + `<br><br>
      社判（社印）管理：<br>
      ` + html.EscapeString(stpResult) + `<br><br>
      請求・支払管理：<br>
      ` + html.EscapeString(actResult) + `<br><br>
      業務プロセスの浸透：<br>
      ` + html.EscapeString(shiResult) + `<br><br>
      ITシステムの管理：<br>
      ` + html.EscapeString(itsResult) + `<br><br>

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
