#!/bin/bash

CURRENT_DIR=`pwd`
DIST_DIR=${CURRENT_DIR}/dist
PACKAGE=github.com/argoproj/argo-cd/common
VERSION='123'
BUILD_DATE=`date -u +'%Y-%m-%dT%H:%M:%SZ'`
GIT_COMMIT='abcdefg'
GIT_TAG='cxl-test'
GIT_TREE_STATE='clean'

GOPATH='/go'

docker build -t argocd-base --target argocd-base .
docker build -t argocd-ui --target argocd-ui .
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 dist/packr build -v -i  -o ${DIST_DIR}/argocd-server ./cmd/argocd-server
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 dist/packr build -v -i  -o ${DIST_DIR}/argocd-application-controller ./cmd/argocd-application-controller
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 dist/packr build -v -i  -o ${DIST_DIR}/argocd-repo-server ./cmd/argocd-repo-server
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 dist/packr build -v -i  -o ${DIST_DIR}/argocd-util ./cmd/argocd-util
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 dist/packr build -v -i  -o ${DIST_DIR}/argocd ./cmd/argocd
mv ${DIST_DIR}/argocd-server ${DIST_DIR}/appcenter-server
mv ${DIST_DIR}/argocd-application-controller ${DIST_DIR}/appcenter-application-controller
mv ${DIST_DIR}/argocd-repo-server ${DIST_DIR}/appcenter-repo-server
mv ${DIST_DIR}/argocd-util ${DIST_DIR}/appcenter-util
mv ${DIST_DIR}/argocd ${DIST_DIR}/appcenter
cp Dockerfile.dev dist
docker build -t argocd:cxl-test -f dist/Dockerfile.dev dist
docker tag argocd:cxl-test registry.cn-beijing.aliyuncs.com/xianlu/appcenter:cxl-test
docker push registry.cn-beijing.aliyuncs.com/xianlu/appcenter:cxl-test