#! /bin/bash -e

SWAGGER_CODEGEN_CLI_VER=2.2.1
SWAGGER_CODEGEN_JAR_NAME=swagger-codegen-cli-$SWAGGER_CODEGEN_CLI_VER.jar
SWAGGER_CODEGEN_DEST_DIR=client

if [ ! -e $SWAGGER_CODEGEN_JAR_NAME ]; then
    wget http://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/$SWAGGER_CODEGEN_CLI_VER/$SWAGGER_CODEGEN_JAR_NAME
fi

java -jar $SWAGGER_CODEGEN_JAR_NAME help
java -jar $SWAGGER_CODEGEN_JAR_NAME langs

mkdir -p $SWAGGER_CODEGEN_DEST_DIR
java -jar $SWAGGER_CODEGEN_JAR_NAME generate \
   --lang go \
   -o $SWAGGER_CODEGEN_DEST_DIR \
   -i https://developer.oxforddictionaries.com/swagger/spec/public_doc_guest.json


