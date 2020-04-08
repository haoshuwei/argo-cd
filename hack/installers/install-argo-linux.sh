#!/bin/bash
set -eux -o pipefail

curl -sLf --retry 3 -o $BIN/argocd-application-controller https://hk-public-bucket.oss-cn-hongkong.aliyuncs.com/argocd/argocd-application-controller
chmod +x $BIN/argocd-application-controller

curl -sLf --retry 3 -o $BIN/argocd-repo-server https://hk-public-bucket.oss-cn-hongkong.aliyuncs.com/argocd/argocd-repo-server
chmod +x $BIN/argocd-repo-server

curl -sLf --retry 3 -o $BIN/argocd-util https://hk-public-bucket.oss-cn-hongkong.aliyuncs.com/argocd/argocd-util
chmod +x $BIN/argocd-util

curl -sLf --retry 3 -o $BIN/argocd https://hk-public-bucket.oss-cn-hongkong.aliyuncs.com/argocd/argocd
chmod +x $BIN/argocd

curl -sLf --retry 3 -o $BIN/argocd-server https://hk-public-bucket.oss-cn-hongkong.aliyuncs.com/argocd/argocd-server
chmod +x $BIN/argocd-server
