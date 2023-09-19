#! /usr/bin/env bash

src=../../sql/usercenter.sql
dir=./genmodel

# generate model code
goctl model mysql ddl --src ${src} --dir ${dir} .