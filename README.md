# tinyman-mobile-sdk

# Overview
This is a mobile SDK providing access to the [Tinyman AMM](https://docs.tinyman.org/) on the Algorand blockchain. It currently supports V1.1 Tinyman. This SDK is a wrapper for https://github.com/synycboom/tinyman-go-sdk (In case you are looking for Golang SDK).

## Installation
You can download libraries for iOS and Android from [releases](https://github.com/synycboom/tinyman-mobile-sdk/releases) page.

- For iOS, please see https://www.simpleswiftguide.com/how-to-add-xcframework-to-xcode-project/
- For Android, please see https://developer.android.com/studio/projects/android-library#psd-add-aar-jar-dependency

# Folder overview
`example/ios` contains example usage of iOS

`example/android` contains example usage of Android

`tinyman` provides wrappers for tinyman-go-sdk and some wrapped algorand-go-sdk functions

`misc` contains a release script

# Usage
See [example](https://github.com/synycboom/tinyman-mobile-sdk//tree/master/example) for how to integrate this SDK with both iOS and Android.

## Usage Limitation
Since this SDK is compiled from Go to mobile native libs, there are some limitations that this SDK cannot do.
- go-mobile has type restriction, so other unsupported types (from exported functions/parameters/returned values) are not able to be bundled to the libs.
- we have to create wrappers for the [algorand-go-sdk](https://github.com/algorand/go-algorand-sdk) such as the client, and other structs/functions and use them with this SDK, but it is not fully ported yet. because it requires a lot of effort to do.
- features that require algorand-go-sdk structs/functions like signing with kmd client [kmd-client](https://github.com/algorand/go-algorand-sdk#kmd-client) is not available right now.
- Some types used in [tinyman-go-sdk](https://github.com/synycboom/tinyman-go-sdk) like uint64, float64, or a slice of structs that are not []byte are not supported. We will use a string for uint64/float64 and an iterator for a slice and convert them inside the SDK.

## Basic Usage
Normally, the steps for calling the SDK are separated into 3 parts.
- call a prepareXXXX function to prepare transactions that you want
- sign transactions returned by the prepareXXXX function, will usually be signed by user account (other transactions will be signed by logic signature account inside the prepareXXXX function)
- submit transactions to the blockchain and wait for confirmation

## Supported Features
- Bootstrapping
- Minting
- Burning
- Swapping
- Redeeming
- OptInApp
- OptInAsset

### Supposing that we want to bootstrap a liquidity pool
1. prepare bootstrap transactions

Android
```kotlin
// Fetch the asset1 token
val token = tinymanClient.fetchAsset(asset1Id)

// Fetch the ALGO token
val algo = tinymanClient.fetchAsset("0")

// Fetch the created pool
val pool = tinyman.Pool(
    algodClient,
    token,
    algo,
    null,
    tinymanClient.validatorAppID,
    account.address,
    true
)

// Prepare a transaction group for bootstrapping
// Note that some transactions need to be signed with LogicSig account, and they were signed in the function.
val txGroup = pool.prepareBootstrapTransactions(account.address)
```

iOS
```swift
var error: NSError?

// Fetch the asset1 token
let token = try tinymanClient.fetchAsset(asset1Id)

// Fetch the ALGO token
let algo = try tinymanClient.fetchAsset("0")

// Fetch the created pool
let pool = TinymanNewPool(
    algodClient,
    token,
    algo,
    nil,
    tinymanClient.getValidatorAppID(),
    account.getAddress(),
    true,
    &error
)!
if (error != nil) {
    throw error!
}

// Prepare a transaction group for bootstrapping
// Note that some transactions need to be signed with LogicSig account, and they were signed in the function.
let txGroup = try pool.prepareBootstrapTransactions(account.getAddress())
```

2. sign transactions

Android
```kotlin
txGroup.sign(account)
```

iOS
```swift
try txGroup.sign(account)
```

3. submit transactions

Android
```kotlin
val txId = tinymanClient.submit(txGroup, true)
```

iOS
```swift
let txId = tinymanClient.submit(txGroup, wait: true, error: &error)
if (error != nil) {
    throw error!
}
```

## Running example
Open the example project with Android Studio or Xcode.

## For contributors
You can prepare environment for android development by running the following command.
It will install an SDK and NDK in the current directory.
```command
# in case you use linux
$ make development.linux

# in case you use mac
$ make development.osx

$ make init
```
You can run ```make build.ios``` or ```make build.android``` to build the SDK. It will automatically save the output to example directory. 

You can run ```make release``` to build a release version of the SDK and zip all the files. 

# Disclaimer
Nothing in the repo constitutes professional and/or financial advice. Use this SDK at your own risk.
