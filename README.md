# deepequal2cmp
testingパッケージで自動生成されるDeeqequalをgo-cmpに自動変換する。

## why
社内プロジェクトで、go-cmpを利用することがルールづけられているが、手動でいちいち変換しなければならないため。
これを利用することで

- ルールを強制させる
- 手間を省く

ことを実現する。

## how to use
```
de2cmp ./...
```

を実行すると全ての*_test.goファイルが書き換えられる。

※ 現状は/cmd/deepequal2cmp/main.goをrunしてデバックしています
