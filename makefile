# 用于跨平台构建的Makefile

# Go编译器设置
GO := go
GOARCH_386 := 386
GOARCH_amd64 := amd64
GOARCH_arm64 := arm64
GOARCH_armv6 := arm
GOARCH_armv7 := arm

# 输出目录
OUTPUT_DIR := build
LINUX_OUTPUT := $(OUTPUT_DIR)/linux
WINDOWS_OUTPUT := $(OUTPUT_DIR)/windows

# 源文件
SOURCE := main.go

.PHONY: all linux windows clean

all: linux windows

linux: linux_386 linux_amd64 linux_arm64 linux_armv6 linux_armv7

linux_386:
	@echo "Compiling for Linux 386..."
	@mkdir -p $(LINUX_OUTPUT)/linux_386
	GOOS=linux GOARCH=$(GOARCH_386) $(GO) build -o $(LINUX_OUTPUT)/linux_386/nbp -tags "netgo" -ldflags "-w -s" -buildmode=pie $(SOURCE)
	@echo "Done"

linux_amd64:
	@echo "Compiling for Linux amd64..."
	@mkdir -p $(LINUX_OUTPUT)/linux_amd64
	GOOS=linux GOARCH=$(GOARCH_amd64) $(GO) build -o $(LINUX_OUTPUT)/linux_amd64/nbp -tags "netgo" -ldflags "-w -s" -buildmode=pie $(SOURCE)
	@echo "Done"

linux_arm64:
	@echo "Compiling for Linux arm64..."
	@mkdir -p $(LINUX_OUTPUT)/linux_arm64
	GOOS=linux GOARCH=$(GOARCH_arm64) $(GO) build -o $(LINUX_OUTPUT)/linux_arm64/nbp -tags "netgo" -ldflags "-w -s" -buildmode=pie $(SOURCE)
	@echo "Done"

linux_armv6:
	@echo "Compiling for Linux armv6..."
	@mkdir -p $(LINUX_OUTPUT)/linux_armv6
	GOOS=linux GOARCH=$(GOARCH_armv6) $(GO) build -o $(LINUX_OUTPUT)/linux_armv6/nbp -tags "netgo" -ldflags "-w -s" -buildmode=pie $(SOURCE)
	@echo "Done"

linux_armv7:
	@echo "Compiling for Linux armv7..."
	@mkdir -p $(LINUX_OUTPUT)/linux_armv7
	GOOS=linux GOARCH=$(GOARCH_armv7) $(GO) build -o $(LINUX_OUTPUT)/linux_armv7/nbp -tags "netgo" -ldflags "-w -s" -buildmode=pie $(SOURCE)
	@echo "Done"

windows: windows_386 windows_amd64

windows_386:
	@echo "Compiling for Windows 386..."
	@mkdir -p $(WINDOWS_OUTPUT)/windows_386
	GOOS=windows GOARCH=$(GOARCH_386) $(GO) build -o $(WINDOWS_OUTPUT)/windows_386/nbp.exe -tags "netgo" -ldflags "-w -s" $(SOURCE)
	@echo "Done"

windows_amd64:
	@echo "Compiling for Windows amd64..."
	@mkdir -p $(WINDOWS_OUTPUT)/windows_amd64
	GOOS=windows GOARCH=$(GOARCH_amd64) $(GO) build -o $(WINDOWS_OUTPUT)/windows_amd64/nbp.exe -tags "netgo" -ldflags "-w -s" $(SOURCE)
	@echo "Done"

clean:
	@echo "Cleaning..."
	@rm -rf $(OUTPUT_DIR)
	@echo "Done"
