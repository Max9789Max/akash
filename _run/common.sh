AKASH_ROOT=../..

DATA_ROOT=data
AKASH_DIR=$DATA_ROOT/client
PROVIDER_DIR=$DATA_ROOT/provider
AKASHD_DIR=$DATA_ROOT/node

_akash() {
  "$AKASH_ROOT/akash" -d "$AKASH_DIR" "$@"
}

_akashd() {
  "$AKASH_ROOT/akashd" -d "$AKASHD_DIR" "$@"
}

stripkey() {
  sed -e 's/.*: //'
}
