package com.example.tinyman

import java.util.UUID
import android.util.Log
import tinyman.Tinyman
import tinyman.AlgodClient
import tinyman.Client
import tinyman.Account

var base64PrivateKey: String = "oo/eWVfOrzp7a8TTGx8dIkNE6Q/S4WZryYtI8U6tj29CBv+Xdbp6Bi7i/S76E9v4+Jp+iX9Zx7FJS52tBRSDGA=="
var mnemonic: String = "tell that observe tag sick leg wish possible jungle wonder diamond among draw guard dice staff donor impact cancel chest world equal access abandon math"

fun createClients(userAddress: String): Pair<AlgodClient, Client> {
    val algoClient = Tinyman.makeAlgodClient(Tinyman.AlgodTestnetHost, "")
    val tinyManClient = Tinyman.newTestNetClient(algoClient, userAddress)

    return algoClient to tinyManClient
}

fun getAccount(): Account {
    if (mnemonic.isNotEmpty()) {
        return Account(Tinyman.AccountFromMnemonic, mnemonic)
    }

    return Account(Tinyman.AccountFromPrivateKey, base64PrivateKey)
}

fun optInAppIfNeeded(tinymanClient: Client, account: Account) {
    val userAddress = account.address()
    val isOptedIn = tinymanClient.isOptedIn(userAddress)
    if (isOptedIn) {
        return
    }

    Log.i("TINY_MAN_MOBILE_SDK", "Account is not opted into app, opting in now...")

    val txGroup = tinymanClient.prepareAppOptInTransaction("")

    txGroup.sign(account)

    val txID = tinymanClient.submit(txGroup, true)

    Log.i("TINY_MAN_MOBILE_SDK", "Submitted opt-in tx $txID")
}

fun optInAssetIfNeeded(tinymanClient: Client, account: Account, assetId: String) {
    val userAddress = account.address()
    val isOptedIn = tinymanClient.isAssetOptedIn(assetId, userAddress)
    if (isOptedIn) {
        return
    }

    Log.i("TINY_MAN_MOBILE_SDK", "$userAddress not opted in for asset $assetId, opting in...")

    val txGroup = tinymanClient.prepareAssetOptInTransactions(assetId, userAddress)

    txGroup.sign(account)

    val txID = tinymanClient.submit(txGroup, true)

    Log.i("TINY_MAN_MOBILE_SDK", "Opted in for asset $assetId with txid $txID")
}

fun createTestAsset(account: Account, ac: AlgodClient): String {
    val uuid = UUID.randomUUID().toString().substring(0, 5)
    return Tinyman.createAnAsset(
        "sdk-test-$uuid", "st", "6", "1000000000", account.address(), account, ac,
    )
}