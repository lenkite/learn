#!/usr/bin/env zsh
sudo httpd -X -d . -e trace8 -f cgi.conf
