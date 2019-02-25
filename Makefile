include version.mk
GO = go
PROG1 = postCode
VERSION = $(VERSION_MAJOR).$(VERSION_MINOR).$(VERSION_BUILD)
VERSION_INT = $(VERSION_MAJOR)$(VERSION_MINOR)$(VERSION_BUILD)

all::
	$(GO) build -ldflags "-X main.Version=${VERSION} -X main.VersionInt=${VERSION_INT}" -o ${PROG1}
clean::
	rm payment

