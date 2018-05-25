
SWAGGER_CODEGEN_CLI_VER   := 2.3.1
SWAGGER_CODEGEN_JAR       := swagger-codegen-cli-$(SWAGGER_CODEGEN_CLI_VER).jar
SWAGGER_CODEGEN_DEST_DIR  := client
SWAGGER_YAML              := public_doc_guest.json
SWAGGER_YAML_GET_FROM_URL := https://developer.oxforddictionaries.com/swagger/spec/$(SWAGGER_YAML)
TMPDIR                    := $(shell mktemp -d)

all:
	mkdir -p $(TMPDIR)
	@wget $(SWAGGER_YAML_GET_FROM_URL) -O $(TMPDIR)/tmp_swagger.yaml
	@rsync -c $(TMPDIR)/tmp_swagger.yaml ./$(SWAGGER_YAML) 
	rm -rf $(TMPDIR)
	make $(SWAGGER_CODEGEN_DEST_DIR)/README.md

$(SWAGGER_CODEGEN_DEST_DIR)/README.md: $(CURDIR)/$(SWAGGER_CODEGEN_JAR) $(CURDIR)/$(SWAGGER_YAML)
	mkdir -p $(SWAGGER_CODEGEN_DEST_DIR)
	java -jar $(CURDIR)/$(SWAGGER_CODEGEN_JAR) generate \
		--lang go \
		-o $(SWAGGER_CODEGEN_DEST_DIR) \
		-i ./$(SWAGGER_YAML)
	goimports -w $(SWAGGER_CODEGEN_DEST_DIR)

$(CURDIR)/$(SWAGGER_YAML):
		cd $(CURDIR) && \
		@wget $(SWAGGER_YAML_GET_FROM_URL)

$(CURDIR)/$(SWAGGER_CODEGEN_JAR):
	wget http://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/$(SWAGGER_CODEGEN_CLI_VER)/$(SWAGGER_CODEGEN_JAR)
	@java -jar $(CURDIR)/$(SWAGGER_CODEGEN_JAR) help
	@java -jar $(CURDIR)/$(SWAGGER_CODEGEN_JAR) langs

clean:
	rm -f $(CURDIR)/$(SWAGGER_YAML)
	rm -f $(CURDIR)/*.jar
	rm -rf $(SWAGGER_CODEGEN_DEST_DIR)
