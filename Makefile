SRC_PATH := $(shell pwd)
MODULE := github.com/synycboom/tinyman-mobile-sdk/tinyman
GO_MOBILE := golang.org/x/mobile/cmd/gomobile

ANDROID_BUILD_PATH := $(SRC_PATH)/build/android
ANDROID_BUILD_AAR := $(ANDROID_BUILD_PATH)/tinyman.aar

ANDROID_BUILD_TOOLS := $(SRC_PATH)/android-build-tools
ANDROID_SDK_MANAGER := $(ANDROID_BUILD_TOOLS)/cmdline-tools/latest/bin/sdkmanager

ANDROID_NDK_ROOT := $(ANDROID_BUILD_TOOLS)/ndk
ANDROID_SDK_ROOT := $(ANDROID_BUILD_TOOLS)/sdk
ANDROID_CMAKE_ROOT := $(ANDROID_BUILD_TOOLS)/cmake

ANDROID_SDK_VERSION := android-32
ANDROID_NDK_VERSION := 24.0.8215888
ANDROID_CMAKE_VERSION := 3.18.1

development.linux:
	# Setup command line tools for Android
	rm -r $(ANDROID_BUILD_TOOLS)
	wget https://dl.google.com/android/repository/commandlinetools-linux-8092744_latest.zip
	mkdir -p $(ANDROID_BUILD_TOOLS)/cmdline-tools
	unzip commandlinetools-linux-8092744_latest.zip -d $(ANDROID_BUILD_TOOLS)/cmdline-tools
	rm commandlinetools-linux-8092744_latest.zip
	mv $(ANDROID_BUILD_TOOLS)/cmdline-tools/cmdline-tools $(ANDROID_BUILD_TOOLS)/cmdline-tools/latest

	# Install sdk, ndk, and cmake
	$(ANDROID_SDK_MANAGER) --sdk_root=$(ANDROID_SDK_ROOT) --install "platforms;$(ANDROID_SDK_VERSION)"
	$(ANDROID_SDK_MANAGER) --sdk_root=$(ANDROID_NDK_ROOT) --install "ndk;$(ANDROID_NDK_VERSION)"
	$(ANDROID_SDK_MANAGER) --sdk_root=$(ANDROID_CMAKE_ROOT) --install "cmake;$(ANDROID_CMAKE_VERSION)"

	# NDK 24+ does not contain abi16 any more. see [https://github.com/golang/go/issues/35030#issuecomment-1026887111]
	ln -sf $(ANDROID_NDK_ROOT)/ndk/$(ANDROID_NDK_VERSION)/toolchains/llvm/prebuilt/linux-x86_64/bin/i686-linux-android32-clang $(ANDROID_NDK_ROOT)/ndk/$(ANDROID_NDK_VERSION)/toolchains/llvm/prebuilt/linux-x86_64/bin/i686-linux-android16-clang
	ln -sf $(ANDROID_NDK_ROOT)/ndk/$(ANDROID_NDK_VERSION)/toolchains/llvm/prebuilt/linux-x86_64/bin/i686-linux-android32-clang++ $(ANDROID_NDK_ROOT)/ndk/$(ANDROID_NDK_VERSION)/toolchains/llvm/prebuilt/linux-x86_64/bin/i686-linux-android16-clang++
	ln -sf $(ANDROID_NDK_ROOT)/ndk/$(ANDROID_NDK_VERSION)/toolchains/llvm/prebuilt/linux-x86_64/bin/armv7a-linux-androideabi32-clang $(ANDROID_NDK_ROOT)/ndk/$(ANDROID_NDK_VERSION)/toolchains/llvm/prebuilt/linux-x86_64/bin/armv7a-linux-androideabi16-clang
	ln -sf $(ANDROID_NDK_ROOT)/ndk/$(ANDROID_NDK_VERSION)/toolchains/llvm/prebuilt/linux-x86_64/bin/armv7a-linux-androideabi32-clang++ $(ANDROID_NDK_ROOT)/ndk/$(ANDROID_NDK_VERSION)/toolchains/llvm/prebuilt/linux-x86_64/bin/armv7a-linux-androideabi16-clang++
	
init:
	go run $(GO_MOBILE) init

build.ios:
	go run $(GO_MOBILE) bind -target=ios $(MODULE)

build.android:
	mkdir -p $(ANDROID_BUILD_PATH)
	ANDROID_HOME=$(ANDROID_SDK_ROOT) ANDROID_NDK_HOME=$(ANDROID_NDK_ROOT)/ndk/$(ANDROID_NDK_VERSION) go run $(GO_MOBILE) bind -target=android -o $(ANDROID_BUILD_AAR) $(MODULE)
