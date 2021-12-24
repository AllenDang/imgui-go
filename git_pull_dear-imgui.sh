#!/bin/sh
set -ex

git subtree pull --prefix=dear-imgui https://github.com/ocornut/imgui.git "$1" --squash
