go-ipscan
=========

ネットワーク調査ツール
ネットワーク内にてどのホストがいるのか確認用に作った。pingに対応してないのが残念。
22, 80, 23 ポートが開いているか調査。開いているポートが一つでも見つかったら終了。

goのマルチスレッドを使ってみたかったんです。

複数調査
---------
* ./ipscan -i 192.168.11.
  * 192.168.11.1 - 192.168.11.254まで調査します。
* ./ipscan -i 192.168.11. -p 8080
  * 22, 80, 23 ポートのチェック後に8080ポートのチェックを追加します。

単体調査
---------
* ./ipscan -h mydomain.local
* ./ipscan -h 192.168.11.2

* ./ipscan -h mydomain.local -p 8080
* ./ipscan -h 192.168.11.2 -p 8080
  * 22, 80, 23 ポートのチェック後に8080ポートのチェックを追加します。
