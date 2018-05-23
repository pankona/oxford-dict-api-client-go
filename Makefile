
SWAGGER_CODEGEN_CLI_VER  = 2.3.1
SWAGGER_CODEGEN_JAR      = swagger-codegen-cli-$(SWAGGER_CODEGEN_CLI_VER).jar
SWAGGER_CODEGEN_DEST_DIR = $(CURDIR)/client
SWAGGER_YAML             = public_doc_guest.json

all: $(CURDIR)/$(SWAGGER_CODEGEN_JAR) $(CURDIR)/public_doc_guest.json
	mkdir -p $(SWAGGER_CODEGEN_DEST_DIR)
	java -jar $(CURDIR)/$(SWAGGER_CODEGEN_JAR) generate \
		--lang go \
		-o $(SWAGGER_CODEGEN_DEST_DIR) \
		-i ./public_doc_guest.json

$(CURDIR)/public_doc_guest.json:
		cd $(CURDIR) && \
		wget https://developer.oxforddictionaries.com/swagger/spec/public_doc_guest.json

$(CURDIR)/$(SWAGGER_CODEGEN_JAR):
	wget http://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/$(SWAGGER_CODEGEN_CLI_VER)/$(SWAGGER_CODEGEN_JAR)
	@java -jar $(CURDIR)/$(SWAGGER_CODEGEN_JAR) help
	@java -jar $(CURDIR)/$(SWAGGER_CODEGEN_JAR) langs

clean:
	rm -f $(CURDIR)/public_doc_guest.json
	rm -f $(CURDIR)/*.jar
	rm -rf $(SWAGGER_CODEGEN_DEST_DIR)
