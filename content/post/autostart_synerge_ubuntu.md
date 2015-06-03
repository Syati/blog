+++
date = "2015-03-01T21:37:06+09:00"
draft = true
title = "synerge ubuntu"

+++

Windows 7 と Ubuntu 12.04 のマシンを利用している Syati ですが、毎回 Ubuntu を立ち上げてログイン、synergy を立ち上げるのは面倒くさいので Ubuntu ログイン画面で自動で起動させるようにした。

# 方法

/etc/lightdm/lightdm に一行加える

    1  [SeatDefaults]
    2  greeter-session=unity-greeter
    3  greeter-setup-script=/usr/bin/synergyc シナジーサーバーが起動しているIP /* この行を加えた */

# 参考サイト

-   [Ubuntu11.10でSynergyを自動起動](http://yasu1973fc2.blog99.fc2.com/blog-entry-72.html)
