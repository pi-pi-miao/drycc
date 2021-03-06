#!/bin/bash

case $1 in
  mariadb)
    chown -R mysql:mysql /data
    chmod 0700 /data
    shift
    exec gosu mysql \
      /bin/drycc-mariadb $*
    ;;
  api)
    shift
    exec /bin/drycc-mariadb-api $*
    ;;
  *)
    echo "Usage: $0 {mariadb|api}"
    exit 2
    ;;
esac
