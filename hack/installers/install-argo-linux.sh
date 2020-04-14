#!/bin/bash
set -eux -o pipefail

curl -sLf --retry 3 -o $BIN/appcenter-application-controller https://hk-public-bucket.oss-cn-hongkong.aliyuncs.com/argocd/appcenter-application-controller
chmod +x $BIN/appcenter-application-controller

curl -sLf --retry 3 -o $BIN/appcenter-repo-server https://hk-public-bucket.oss-cn-hongkong.aliyuncs.com/argocd/appcenter-repo-server
chmod +x $BIN/appcenter-repo-server

curl -sLf --retry 3 -o $BIN/appcenter-util https://hk-public-bucket.oss-cn-hongkong.aliyuncs.com/argocd/appcenter-util
chmod +x $BIN/appcenter-util

curl -sLf --retry 3 -o $BIN/appcenter https://hk-public-bucket.oss-cn-hongkong.aliyuncs.com/argocd/appcenter
chmod +x $BIN/appcenter

curl -sLf --retry 3 -o $BIN/appcenter-server https://hk-public-bucket.oss-cn-hongkong.aliyuncs.com/argocd/appcenter-server
chmod +x $BIN/appcenter-server
