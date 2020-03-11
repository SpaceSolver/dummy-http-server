# AndroidエミュレータからホストPCにアクセスする用のダミーサーバ

* アドレス := http://10.0.2.2:8088
* Android 9以降はデフォルトでhttps通信になるので↓の対応を入れる
   * https://qiita.com/b_a_a_d_o/items/afa0d83bbffdb5d4f6be
   * 1) `res/xml/network_security_config.xml` を作る
   * 2) `AndroidManifest.xml`に登録する
* permission.INTERNET も忘れずに
* 応答は `{num: "123"}` とかの適当なJSONデータ

* 起動は、管理者権限で実行すること
