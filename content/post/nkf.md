+++
Categories = ["Shell"]
Description = "nkf で直す文字化け"
Tags = ["Mojibake"]
date = "2015-06-04T23:09:38+09:00"
title = "nkf で直す文字化け"

+++

Windows 時代の遺産で、文字コードが shift\_jis のファイルがたくさんありました。開いてびっくり文字化け。全部 utf-8 に修正してしまいます。

# 利用するもの

**nkf** ただこれだけ

# インストール

パッケージ管理でいれるのが楽です。無い場合は、[nkf Network Kanji Filter](http://sourceforge.jp/projects/nkf/) からダウンロードして、make してください。

# 利用例

-   文字コードを判定する場合
    
        % nkf -g shiftjis.txt
        output: Shift_JIS

-   文字コードを utf-8 + 改行コードを UNIX に変換してターミナル出力する場合
    
        % nkf -w -Lu shiftjis.txt

-   文字コードを utf-8 + 改行コードを UNIX に変換して元のファイルを上書きする場合
    
        % nkf -w --overwrite -Lu shiftjis.txt

# 参考

-   [Linux上で文字コードを変換できるコマンドnkfのオプション一覧](http://blog.layer8.sh/ja/2012/03/31/nkf_command_option/)
-   [文字コード変換コマンドnkfの使い方まとめ Linux](http://blog.layer8.sh/ja/2011/12/23/%E6%96%87%E5%AD%97%E3%82%B3%E3%83%BC%E3%83%89%E5%A4%89%E6%8F%9B%E3%82%B3%E3%83%9E%E3%83%B3%E3%83%89nkf%E3%81%AE%E4%BD%BF%E3%81%84%E6%96%B9%E3%81%BE%E3%81%A8%E3%82%81-linux/)
-   [6.9. 文字コードと改行コードの変換 【nkf】](http://www.turbolinux.co.jp/products/server/10s/manual/command_guide/command_guide/nkf.html)
