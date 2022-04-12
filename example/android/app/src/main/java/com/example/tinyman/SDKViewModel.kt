package com.example.tinyman

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import kotlinx.coroutines.*
import tinyman.AssetAmountIterator
import java.math.BigDecimal

class SDKViewModel(): ViewModel(){
    fun createAsset(cb: (String) -> Unit) {
        viewModelScope.launch {
            try {
                val msg = withContext(Dispatchers.IO) {
                    val account = getAccount()
                    val userAddress = account.address
                    // Create algorand and tinyman clients
                    val (algodClient, _) = createClients(userAddress)

                    // Create a new asset
                    val assetId = createTestAsset(account, algodClient)

                    "Created an asset with id=$assetId, please copy this id and use it in other examples!"
                }

                cb(msg)
            } catch (err: Throwable) {
                cb(err.stackTraceToString())
            }
        }
    }

    fun boostrap(asset1Id: String, cb: (String) -> Unit) {
        viewModelScope.launch {
            try {
                val msg = withContext(Dispatchers.IO) {
                    val account = getAccount()
                    val userAddress = account.address
                    val (algodClient, tinymanClient) = createClients(userAddress)

                    // Check whether the user already opted in the app or not, if not let the user opt in
                    optInAppIfNeeded(tinymanClient, account)

                    // Check whether the user already opted in for the asset or not, if not let the user opt in
                    optInAssetIfNeeded(tinymanClient, account, asset1Id)

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

                    // Some transactions that need the user signatures are signed here
                    txGroup.sign(account)

                    // Submit a group of transaction to the blockchain
                    val txId = tinymanClient.submit(txGroup, true)

                    "Liquidity pool was bootstrapped with txid $txId"
                }

                cb(msg)
            } catch (err: Throwable) {
                cb(err.stackTraceToString())
            }
        }
    }

    fun mint(asset1Id: String, cb: (String) -> Unit) {
        viewModelScope.launch {
            try {
                val msg = withContext(Dispatchers.IO) {
                    var msg = ""
                    val account = getAccount()
                    val userAddress = account.address
                    val (algodClient, tinymanClient) = createClients(userAddress)

                    // Check whether the user already opted in the app or not, if not let the user opt in
                    optInAppIfNeeded(tinymanClient, account)

                    // Check whether the user already opted in for the asset or not, if not let the user opt in
                    optInAssetIfNeeded(tinymanClient, account, asset1Id)

                    // Fetch the asset1 token
                    val token = tinymanClient.fetchAsset(asset1Id)

                    // Fetch the ALGO token
                    val algo = tinymanClient.fetchAsset("0")

                    // Fetch TOKEN-ALGO pool
                    val pool = tinyman.Pool(algodClient, token, algo, null, tinymanClient.validatorAppID, account.address, true)

                    // Check whether the user already opted in for the liquidity asset or not, if not let the user opt in
                    optInAssetIfNeeded(tinymanClient, account, pool.liquidityAsset.id)

                    val tokenAssetAmount = tinyman.AssetAmount()

                    // Set asset
                    tokenAssetAmount.setAsset(token)

                    // Note that 10000000 is equal to 10000000 / 10 ** decimals (if decimals is 6), which is 10 tokens
                    tokenAssetAmount.setAmount("10000000")

                    val algoAssetAmount = tinyman.AssetAmount()

                    // Set asset
                    algoAssetAmount.setAsset(algo)

                    // Note that 300000 is equal to 300000 / 10 ** decimals (if decimals is 6), which is 0.3ALGO
                    algoAssetAmount.setAmount("300000")

                    // Fetch mint quote used when submit minting transactions
                    val quote = pool.fetchMintQuote(tokenAssetAmount, algoAssetAmount, "0.05")

                    // Calculate the liquidity asset amount after applying the slippage
                    val liquidityAssetAmountWithSlippage = quote.liquidityAssetAmountWithSlippage

                    msg += "Liquidity asset amount: ID-${quote.liquidityAssetAmount.asset.id} = ${quote.liquidityAssetAmount.amount}\n"
                    msg += "Liquidity asset amount with slippage: ID-${liquidityAssetAmountWithSlippage.asset.id} = ${liquidityAssetAmountWithSlippage.amount}\n"

                    // Prepare a transaction group for minting
                    // Note that some transactions need to be signed with LogicSig account, and they were signed in the function.
                    val txGroup = pool.prepareMintTransactionsFromQuote(quote, userAddress)

                    // Some transactions that need the user signatures are signed here
                    txGroup.sign(account)

                    // Submit a group of transaction to the blockchain
                    val txId = tinymanClient.submit(txGroup, true)

                    msg += "Liquidity was added in txid $txId\n"

                    val info = pool.fetchPoolPosition(userAddress)
                    val share = info.share.toBigDecimal().multiply(BigDecimal("100"))

                    msg += "Pool tokens: ID-${info.liquidityAssetAmount.asset.id} = ${info.liquidityAssetAmount.amount}\n"
                    msg += "Asset1: ID-${info.assetAmount1.asset.id} = ${info.assetAmount1.amount}\n"
                    msg += "Asset2: ID-${info.assetAmount2.asset.id} = ${info.assetAmount2.amount}\n"
                    msg += "Share of pool: ${share}%"

                    msg
                }

                cb(msg)
            } catch (e: Throwable) {
                cb(e.stackTraceToString())
            }
        }
    }

    fun burn(asset1Id: String, cb: (String) -> Unit) {
        viewModelScope.launch {
            try {
                val msg = withContext(Dispatchers.IO) {
                    var msg = ""
                    val account = getAccount()
                    val userAddress = account.address
                    val (algodClient, tinymanClient) = createClients(userAddress)

                    // Check whether the user already opted in the app or not, if not let the user opt in
                    optInAppIfNeeded(tinymanClient, account)

                    // Check whether the user already opted in for the asset or not, if not let the user opt in
                    optInAssetIfNeeded(tinymanClient, account, asset1Id)

                    // Fetch the asset1 token
                    val token = tinymanClient.fetchAsset(asset1Id)

                    // Fetch the ALGO token
                    val algo = tinymanClient.fetchAsset("0")

                    // Fetch TOKEN-ALGO pool
                    val pool = tinyman.Pool(algodClient, token, algo, null, tinymanClient.validatorAppID, account.address, true)

                    // Check whether the user already opted in for the liquidity asset or not, if not let the user opt in
                    optInAssetIfNeeded(tinymanClient, account, pool.info.liquidityAssetID)

                    // Fetch total balance of liquidtiy asset that the user has
                    val balance = tinymanClient.fetchBalance(pool.liquidityAsset, userAddress)

                    msg += "Current balance of liquidity \n\t - ID:${pool.liquidityAsset.id} = ${balance.amount}\n"

                    if (balance.amount.toBigInteger().equals(0)) {
                        msg += "User does not have liquidity balance\n"

                        return@withContext msg
                    }

                    // Fetch burn quote used when burning liquidity asset with slippage equals 0.05
                    val quote = pool.fetchBurnQuote(balance, "0.05")

                    msg += "Liquidity asset amount \n\t - ID-${quote.liquidityAssetAmount.asset.id} = ${quote.liquidityAssetAmount.amount}\n"
                    msg += "Output amounts\n"
                    msg += displayAmountOut(quote.assetAmountsOutIterator)
                    msg += "Output amounts with slippage\n"
                    msg += displayAmountOut(quote.assetAmountsOutWithSlippageIterator)

                    // Prepare a transaction group for burning
                    // Note that some transactions need to be signed with LogicSig account, and they were signed in the function.
                    val txGroup = pool.prepareBurnTransactionsFromQuote(quote, userAddress)

                    // Some transactions that need the user signatures are signed here
                    txGroup.sign(account)

                    // Submit a group of transaction to the blockchain
                    val txId = tinymanClient.submit(txGroup, true)

                    msg += "Liquidity was removed in txid $txId\n"

                    val info = pool.fetchPoolPosition(userAddress)
                    val share = info.share.toBigDecimal().multiply(BigDecimal("100"))

                    msg += "Pool tokens: ID-${info.liquidityAssetAmount.asset.id} = ${info.liquidityAssetAmount.amount}\n"
                    msg += "Asset1: ID-${info.assetAmount1.asset.id} = ${info.assetAmount1.amount}\n"
                    msg += "Asset2: ID-${info.assetAmount2.asset.id} = ${info.assetAmount2.amount}\n"
                    msg += "Share of pool: ${share}%"

                    msg
                }

                cb(msg)
            } catch (err: Throwable) {
                cb(err.stackTraceToString())
            }
        }
    }

    fun swap(asset1Id: String, cb: (String) -> Unit) {
        viewModelScope.launch {
            try {
                val msg = withContext(Dispatchers.IO) {
                    var msg = ""
                    val account = getAccount()
                    val userAddress = account.address
                    val (algodClient, tinymanClient) = createClients(userAddress)

                    // Check whether the user already opted in the app or not, if not let the user opt in
                    optInAppIfNeeded(tinymanClient, account)

                    // Check whether the user already opted in for the asset or not, if not let the user opt in
                    optInAssetIfNeeded(tinymanClient, account, asset1Id)

                    // Fetch the asset1 token
                    val token = tinymanClient.fetchAsset(asset1Id)

                    // Fetch the ALGO token
                    val algo = tinymanClient.fetchAsset("0")

                    // Fetch TOKEN-ALGO pool
                    val pool = tinyman.Pool(algodClient, token, algo, null, tinymanClient.validatorAppID, account.address, true)

                    // Check whether the user already opted in for the liquidity asset or not, if not let the user opt in
                    optInAssetIfNeeded(tinymanClient, account, pool.liquidityAsset.id)

                    val tokenAssetAmount = tinyman.AssetAmount()

                    // Set asset
                    tokenAssetAmount.setAsset(token)

                    // Note that 500000 is equal to 500000 / 10 ** decimals (if decimals is 6), which is 0.5 tokens
                    tokenAssetAmount.setAmount("500000")

                    // Fetch mint quote used when submit minting transactions
                    val quote = pool.fetchFixedInputSwapQuote(tokenAssetAmount, "0.05")

                    // Calculate price after applying the slippage
                    val priceWithSlippage = quote.priceWithSlippage

                    // Calculate output amount after applying the slippage
                    val amountOutWithSlippage = quote.assetAmountOutWithSlippage

                    msg += "ID-${token.id} per ALGO: ${quote.price}\n"
                    msg += "ID-${token.id} per ALGO (worst case): ${priceWithSlippage}\n"
                    msg += "Swapping ID-${quote.assetAmountIn.asset.id} = ${quote.assetAmountIn.amount} to ID-${amountOutWithSlippage.asset.id} = ${amountOutWithSlippage.amount}\n"

                    // Prepare a transaction group for swappingg
                    // Note that some transactions need to be signed with LogicSig account, and they were signed in the function.
                    val txGroup = pool.prepareSwapTransactionsFromQuote(quote, userAddress)

                    // Some transactions that need the user signatures are signed here
                    txGroup.sign(account)

                    // Submit a group of transaction to the blockchain
                    val txId = tinymanClient.submit(txGroup, true)

                   msg += "Swapped with txid $txId"

                    msg
                }

                cb(msg)
            } catch (err: Throwable) {
                cb(err.stackTraceToString())
            }
        }
    }

    fun redeem(asset1Id: String, cb: (String) -> Unit) {
        viewModelScope.launch {
            try {
                val msg = withContext(Dispatchers.IO) {
                    var msg = ""
                    val account = getAccount()
                    val userAddress = account.address
                    val (algodClient, tinymanClient) = createClients(userAddress)

                    // Check whether the user already opted in the app or not, if not let the user opt in
                    optInAppIfNeeded(tinymanClient, account)

                    // Check whether the user already opted in for the asset or not, if not let the user opt in
                    optInAssetIfNeeded(tinymanClient, account, asset1Id)

                    // Fetch the asset1 token
                    val token = tinymanClient.fetchAsset(asset1Id)

                    // Fetch the ALGO token
                    val algo = tinymanClient.fetchAsset("0")

                    // Fetch TOKEN-ALGO pool
                    val pool = tinyman.Pool(algodClient, token, algo, null, tinymanClient.validatorAppID, account.address, true)

                    // Fetch excess amount resulting from the swap
                    val redeemQuoteIterator = tinymanClient.fetchExcessAmount(userAddress)

                    // Filter redeem quotes matching with TOKEN
                    val quote = pool.getRedeemQuoteMatchesAssetID(asset1Id, redeemQuoteIterator)
                    if (quote == null) {
                        msg += "No excess amount to be redeemed\n"

                        return@withContext msg
                    }

                    msg += "There is ID-${quote.assetAmount.asset.id} = ${quote.assetAmount.amount} to be redeemed\n"

                    // Prepare a transaction group for redeeming
                    // Note that some transactions need to be signed with LogicSig account, and they were signed in the function.
                    val txGroup = pool.prepareRedeemTransactionsFromQuote(quote, userAddress)

                    // Some transactions that need the user signatures are signed here
                    txGroup.sign(account)

                    // Submit a group of transaction to the blockchain
                    val txId = tinymanClient.submit(txGroup, true)

                    msg += "Redeemed excess amount with txid $txId\n"

                    msg
                }

                cb(msg)
            } catch (err: Throwable) {
                cb(err.stackTraceToString())
            }
        }
    }
}

fun displayAmountOut(outputAmountsWithSlippageIter: AssetAmountIterator): String {
    var msg = ""
    while (outputAmountsWithSlippageIter.hasNext()) {
        val asset = outputAmountsWithSlippageIter.next()
        msg += "\t - ID:${asset.asset.id} = ${asset.amount}\n"
    }

    return msg
}
