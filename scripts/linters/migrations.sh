#!/usr/bin/env bash
# linter/formatter for migrations

# Why: Used by the script that calls us
# shellcheck disable=SC2034
extensions=(sql)

version="v1.131.0"

migrations_linter() {
  "$DIR/gobin.sh" "github.com/getoutreach/smartstore/cmd/smartstore@${version}" migrations lint
  return $?
}

migrations_formatter() {
  "$DIR/gobin.sh" "github.com/getoutreach/smartstore/cmd/smartstore@${version}" migrations revises --save
}

linter() {
  run_command "migrations" migrations_linter
}

formatter() {
  run_command "migrations" migrations_formatter
}
