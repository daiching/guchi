# guchi
## Download
```bash
# 実行バイナリのみ(only executable binary)
$ wget https://github.com/daiching/guchi/raw/master/guchi

$ ./guchi

# go get
$ go get github.com/daiching/guchi

$ guchi
```
## 日本語(English is below)
### 概要
愚痴を投稿するCLIクライアントです。コマンドラインから愚痴の投稿および取得ができます。   
なお、愚痴は以下のURLにただの文字列で表示されます。   
https://guchis.site/guchis

動作確認は WSL (Ubuntu) と CentOS7 で実施しましたので、恐らく多くの Linux では動くと思われます。   
サーバー作業しているときに負の感情が募ったら使ってみてください。。。

### コマンド
- 愚痴を取得する場合
``` bash
# 何も指定しないと最新の愚痴が10件取得されます。
$ guchi
Nanashi:
仕事辞めたい

daiching:
頭痛い、腹が痛い、心が痛い、家帰りたい
(...略)

# -n オプションを指定すると、その件数のみ取得されます
$ guchi -n 2
(...略)
```

- 愚痴を投稿する場合
``` bash
# オプションを指定しないと Nanashi という名前で投稿されます。
$ guchi キラキラネームだからNanashiでいいです
Nanashi:
キラキラネームだからNanashiでいいです
(...略)

# -u オプションを指定すると名前を決めて投稿できます。
$ guchi -u アラレちゃん んちゃ！
アラレちゃん:
んちゃ！
(...略)
```

- config ファイルの設定   
~ ディレクトリ以下に .guchi ファイルを置くことで、オプション指定ができます。
といっても、いまのところは投稿者名のみです。以下のように指定します(# でコメントアウト可能です。)
```bash
Name=daiching
```

## English
"guchi" is Japanese bitches.This is CLI vomit tool.   
They display the following URL   
https://guchis.site/guchis   
Operation Check: WSL(Ubuntu), CentOS7   

### Command
- Read "guchi"
``` bash
# If nothing is specified, the latest 10 bitches will be acquired.
$ guchi
Nanashi:
I want to quit my job.

daiching:
I have a headache, stomach ache, heart ache, I want to go home.
(...Less than)

# -n If you specify an option, only that number will be retrieved
$ guchi -n 2
(...Less than)
```

- Write "guchi"
``` bash
# If no option is specified, the post will be posted as "Nanashi".
$ guchi 
Nanashi:
Because it is a bizarre name, Nanashi is fine.
(...Less than)

# -u If you specify an option, you can decide the name and post.
$ guchi -u Arale Ncha!
Arale:
Ncha!
(...Less than)
```

- config file   
Options can be specified by placing a .guchi file under the ~ directory.   
But for now, only the name of the contributor is available. Specify as follows (you can comment out with #).
```bash
Name=daiching
```
