#!/usr/bin/env bash

# 使用方法
# ./genModel.sh 数据库名称 表名称

# table
tables=$2
# genmodel dir
genmodel_dir=./genmodel

# database config
host="127.0.0.1"
port="3306"
dbname=looklook_$1
username="root"
passwd="123"

# 从数据库中导出model
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${genmodel_dir}" -cache=true --style=goZero
