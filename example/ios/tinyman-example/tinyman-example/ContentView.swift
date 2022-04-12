//
//  ContentView.swift
//  tinyman-example
//
//  Created by Wutichai on 12/4/2565 BE.
//

import SwiftUI
import Tinyman

var base64PrivateKey: String = "oo/eWVfOrzp7a8TTGx8dIkNE6Q/S4WZryYtI8U6tj29CBv+Xdbp6Bi7i/S76E9v4+Jp+iX9Zx7FJS52tBRSDGA=="
var mnemonic: String = "tell that observe tag sick leg wish possible jungle wonder diamond among draw guard dice staff donor impact cancel chest world equal access abandon math"

func createClients(userAddress: String) throws -> (TinymanAlgodClient, TinymanClient) {
    var error: NSError?
    let algodClient = TinymanMakeAlgodClient(TinymanAlgodTestnetHost, "", &error)
    if (error != nil) {
        throw error!;
    }
    let tinyManClient = TinymanNewTestNetClient(algodClient, userAddress, &error)
    if (error != nil) {
        throw error!;
    }
    
    return (algodClient!, tinyManClient!);
}

func getAccount() throws -> TinymanAccount {
    if (!mnemonic.isEmpty) {
        return TinymanAccount(TinymanAccountFromMnemonic, value: mnemonic)!
    }

    return TinymanAccount(TinymanAccountFromPrivateKey, value: base64PrivateKey)!
}

func optInAppIfNeeded(tinymanClient: TinymanClient, account: TinymanAccount) throws {
    let userAddress = account.getAddress()
    var isOptedIn: ObjCBool = false
    
    try tinymanClient.isOpted(in: userAddress, ret0_: &isOptedIn)
    
    if (isOptedIn.boolValue) {
        return
    }
    
    print("Account is not opted into app, opting in now...")
    
    let txGroup = try tinymanClient.prepareAppOpt(inTransaction: "")
    
    try txGroup.sign(account)
    
    var error: NSError?
    let txId = tinymanClient.submit(txGroup, wait: true, error: &error)
    if (error != nil) {
        throw error!;
    }
    
    print("Submitted opt-in tx: \(txId)")
}

func optInAssetIfNeeded(tinymanClient: TinymanClient, account: TinymanAccount, assetId: String) throws {
    let userAddress = account.getAddress()
    var isOptedIn: ObjCBool = false
    
    try tinymanClient.isAssetOpted(in: assetId, userAddress: userAddress, ret0_: &isOptedIn)
    
    if (isOptedIn.boolValue) {
        return
    }

    print("\(userAddress) not opted in for asset \(assetId), opting in...")

    let txGroup = try tinymanClient.prepareAssetOpt(inTransactions: assetId, userAddress: userAddress)
    
    try txGroup.sign(account)
    
    var error: NSError?
    let txId = tinymanClient.submit(txGroup, wait: true, error: &error)
    if (error != nil) {
        throw error!;
    }

    print("Opted in for asset \(assetId) with txid \(txId)")
}

func createTestAsset(account: TinymanAccount, ac: TinymanAlgodClient) throws -> String {
    let uuid = UUID().uuidString.prefix(5)
    var error: NSError?
    
    let assetId = TinymanCreateAnAsset(
        "sdk-test-\(uuid)",
        "st",
        "6",
        "1000000000",
        account.getAddress(),
        account,
        ac,
        &error
    )
    if (error != nil) {
        throw error!;
    }
    
    return assetId
}

func displayAmountOut(iter: TinymanAssetAmountIterator) -> String {
    var msg: String = ""
    while (iter.hasNext()) {
        let asset = iter.next()!
        msg += "\t - ID:\(asset.getAsset()!.getId()) = \(asset.getAmount())\n"
    }
    
    return msg
}

struct ContentView: View {
    var body: some View {
        NavigationView {
            VStack {
                Text("Please choose an example. I suggest that you go to create an asset first")
                    .offset(y: -60)
                
                NavigationLink(destination: CreateAssetView(), label: {
                    Text("Go to create an asset")
                        .bold()
                        .frame(width: 280, height: 50)
                        .background(Color.blue)
                        .foregroundColor(.white)
                        .cornerRadius(10)
                })
                
                NavigationLink(destination: BootstrapLiquidityView(), label: {
                    Text("Go to bootstrap liquidity")
                        .bold()
                        .frame(width: 280, height: 50)
                        .background(Color.blue)
                        .foregroundColor(.white)
                        .cornerRadius(10)
                })
                
                NavigationLink(destination: AddLiquidityView(), label: {
                    Text("Go to add liquidity")
                        .bold()
                        .frame(width: 280, height: 50)
                        .background(Color.blue)
                        .foregroundColor(.white)
                        .cornerRadius(10)
                })
                
                NavigationLink(destination: RemoveLiquidityView(), label: {
                    Text("Go to remove liquidity")
                        .bold()
                        .frame(width: 280, height: 50)
                        .background(Color.blue)
                        .foregroundColor(.white)
                        .cornerRadius(10)
                })
                
                NavigationLink(destination: SwapView(), label: {
                    Text("Go to swap")
                        .bold()
                        .frame(width: 280, height: 50)
                        .background(Color.blue)
                        .foregroundColor(.white)
                        .cornerRadius(10)
                })
                
                NavigationLink(destination: RedeemView(), label: {
                    Text("Go to redeem")
                        .bold()
                        .frame(width: 280, height: 50)
                        .background(Color.blue)
                        .foregroundColor(.white)
                        .cornerRadius(10)
                })
            }
        }
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}

struct CreateAssetView: View {
    @State var output: String = "Click create to create an asset, it might take time to create"
    
    func createAsset() {
        DispatchQueue.global(qos: .userInitiated).async {
            output = "loading..."
            
            do {
                let account = try getAccount()
                let (algodClient, tinyManClient) = try createClients(userAddress: account.getAddress())
                
                try optInAppIfNeeded(tinymanClient: tinyManClient, account: account)
                let assetId = try createTestAsset(account: account, ac: algodClient)
                
                output = "Created an asset with id=\(assetId), please copy this id and use it in other examples!"
                print(output)
            } catch let error {
                print(error)
                output = error.localizedDescription
            }
        }
    }
    
    var body: some View {
        NavigationView {
            VStack(spacing: 20) {
                Text(output)
                    .offset(y: -60)
                
                Text("Create")
                    .bold()
                    .frame(width: 280, height: 50)
                    .background(Color.blue)
                    .foregroundColor(.white)
                    .cornerRadius(10)
                    .onTapGesture {
                        self.createAsset()
                    }
            }
        }
    }
}

struct BootstrapLiquidityView: View {
    @State var output: String = "This example bootstraps liquidity with a given asset id and ALGO token"
    @State var asset1Id: String = ""
    
    func bootstrap() {
        DispatchQueue.global(qos: .userInitiated).async {
            output = "loading..."
            
            do {
                if (asset1Id.isEmpty) {
                    output = "Asset 1 ID is empty"
                    
                    return
                }
                
                var error: NSError?
                let account = try getAccount()
                let userAddress = account.getAddress()
                let (algodClient, tinymanClient) = try createClients(userAddress: userAddress)

                // Check whether the user already opted in the app or not, if not let the user opt in
                try optInAppIfNeeded(tinymanClient: tinymanClient, account: account)

                // Check whether the user already opted in for the asset or not, if not let the user opt in
                try optInAssetIfNeeded(tinymanClient: tinymanClient, account: account, assetId: asset1Id)

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
                    throw error!;
                }

                // Prepare a transaction group for bootstrapping
                // Note that some transactions need to be signed with LogicSig account, and they were signed in the function.
                let txGroup = try pool.prepareBootstrapTransactions(account.getAddress())

                // Some transactions that need the user signatures are signed here
                try txGroup.sign(account)

                // Submit a group of transaction to the blockchain
                let txId = tinymanClient.submit(txGroup, wait: true, error: &error)
                if (error != nil) {
                    throw error!;
                }
            
                output = "Liquidity pool was bootstrapped with txid \(txId)"
                print(output)
            } catch let error {
                print(error)
                output = error.localizedDescription
            }
        }
    }
    
    var body: some View {
        NavigationView {
            VStack(spacing: 20) {
                Text(output)
                    .offset(y: -60)

                TextField("Enter asset 1 id...", text: $asset1Id)
                    .multilineTextAlignment(.center)
                
                Text("Bootstrap")
                    .bold()
                    .frame(width: 280, height: 50)
                    .background(Color.blue)
                    .foregroundColor(.white)
                    .cornerRadius(10)
                    .onTapGesture {
                        self.bootstrap()
                    }
            }
        }
    }
}


struct AddLiquidityView: View {
    @State var output: String = "This example add liquidity with a given asset id and ALGO token"
    @State var asset1Id: String = ""
    
    func addLiquidity() {
        DispatchQueue.global(qos: .userInitiated).async {
            output = "loading..."
            
            do {
                if (asset1Id.isEmpty) {
                    output = "Asset 1 ID is empty"
                    
                    return
                }
                
                var error: NSError?
                let account = try getAccount()
                let userAddress = account.getAddress()
                let (algodClient, tinymanClient) = try createClients(userAddress: userAddress)

                // Check whether the user already opted in the app or not, if not let the user opt in
                try optInAppIfNeeded(tinymanClient: tinymanClient, account: account)

                // Check whether the user already opted in for the asset or not, if not let the user opt in
                try optInAssetIfNeeded(tinymanClient: tinymanClient, account: account, assetId: asset1Id)

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
                    throw error!;
                }
                
                // Check whether the user already opted in for the liquidity asset or not, if not let the user opt in
                try optInAssetIfNeeded(tinymanClient: tinymanClient, account: account, assetId: pool.getLiquidityAsset()!.getId())
                
                let tokenAssetAmount = TinymanNewAssetAmount()!
                
                // Set asset
                try tokenAssetAmount.setAsset(token)

                // Note that 10000000 is equal to 10000000 / 10 ** decimals (if decimals is 6), which is 10 tokens
                try tokenAssetAmount.setAmount("10000000")
                
                let algoAssetAmount = TinymanNewAssetAmount()!

                // Set asset
                try algoAssetAmount.setAsset(algo)

                // Note that 300000 is equal to 300000 / 10 ** decimals (if decimals is 6), which is 0.3ALGO
                try algoAssetAmount.setAmount("300000")

                // Fetch mint quote used when submit minting transactions
                let quote = try pool.fetchMintQuote(tokenAssetAmount, amountB: algoAssetAmount, slippage: "0.05")

                // Calculate the liquidity asset amount after applying the slippage
                let liquidityAssetAmountWithSlippage = try quote.getLiquidityAssetAmountWithSlippage()

                output = "Liquidity asset amount: ID-\(quote.getLiquidityAssetAmount()!.getAsset()!.getId()) = \(quote.getLiquidityAssetAmount()!.getAmount())\n"
                output += "Liquidity asset amount with slippage: ID-\(liquidityAssetAmountWithSlippage.getAsset()!.getId()) = \(liquidityAssetAmountWithSlippage.getAmount())\n"

                // Prepare a transaction group for minting
                // Note that some transactions need to be signed with LogicSig account, and they were signed in the function.
                let txGroup = try pool.prepareMintTransactions(from: quote, minterAddress: userAddress)

                // Some transactions that need the user signatures are signed here
                try txGroup.sign(account)

                // Submit a group of transaction to the blockchain
                let txId = tinymanClient.submit(txGroup, wait: true, error: &error)
                if (error != nil) {
                    throw error!;
                }
                
                output += "Liquidity was added in txid \(txId)\n"

                let info = try pool.fetchPosition(userAddress)
                let share = Decimal(string: info.getShare())! * Decimal(string: "100")!
                
                output += "Pool tokens: ID-\(info.getLiquidityAssetAmount()!.getAsset()!.getId()) = \(info.getLiquidityAssetAmount()!.getAmount())\n"
                output += "Asset1: ID-\(info.getAssetAmount1()!.getAsset()!.getId()) = \(info.getAssetAmount1()!.getAmount())\n"
                output += "Asset2: ID-\(info.getAssetAmount2()!.getAsset()!.getId()) = \(info.getAssetAmount2()!.getAmount())\n"
                output += "Share of pool: \(share)%"
                
                print(output)
            } catch let error {
                print(error)
                output = error.localizedDescription
            }
        }
    }
    
    var body: some View {
        NavigationView {
            VStack(spacing: 20) {
                Text(output)
                    .offset(y: -60)

                TextField("Enter asset 1 id...", text: $asset1Id)
                    .multilineTextAlignment(.center)
                
                Text("Add liquidity")
                    .bold()
                    .frame(width: 280, height: 50)
                    .background(Color.blue)
                    .foregroundColor(.white)
                    .cornerRadius(10)
                    .onTapGesture {
                        self.addLiquidity()
                    }
            }
        }
    }
}

struct RemoveLiquidityView: View {
    @State var output: String = "This example remove all liquidity with a given asset id and ALGO token"
    @State var asset1Id: String = ""
    
    func removeLiquidity() {
        DispatchQueue.global(qos: .userInitiated).async {
            output = "loading..."
            
            do {
                if (asset1Id.isEmpty) {
                    output = "Asset 1 ID is empty"
                    
                    return
                }
                
                var error: NSError?
                let account = try getAccount()
                let userAddress = account.getAddress()
                let (algodClient, tinymanClient) = try createClients(userAddress: userAddress)

                // Check whether the user already opted in the app or not, if not let the user opt in
                try optInAppIfNeeded(tinymanClient: tinymanClient, account: account)

                // Check whether the user already opted in for the asset or not, if not let the user opt in
                try optInAssetIfNeeded(tinymanClient: tinymanClient, account: account, assetId: asset1Id)

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
                    throw error!;
                }
                
                // Check whether the user already opted in for the liquidity asset or not, if not let the user opt in
                try optInAssetIfNeeded(tinymanClient: tinymanClient, account: account, assetId: pool.getLiquidityAsset()!.getId())
                
                // Fetch total balance of liquidtiy asset that the user has
                let balance = try tinymanClient.fetchBalance(pool.getLiquidityAsset(), userAddress: userAddress)

                output = "Current balance of liquidity \n\t - ID:\(pool.getLiquidityAsset()!.getId()) = \(balance.getAmount())\n"

                if (Decimal(string: balance.getAmount())! == Decimal(string: "0")!) {
                    output += "User does not have liquidity balance\n"
                    print(output)

                    return
                }

                // Fetch burn quote used when burning liquidity asset with slippage equals 0.05
                let quote = try pool.fetchBurnQuote(balance, slippage: "0.05")

                output += "Liquidity asset amount \n\t - ID-\(quote.getLiquidityAssetAmount()!.getAsset()!.getId()) = \(quote.getLiquidityAssetAmount()!.getAmount())\n"
                output += "Output amounts\n"
                output += displayAmountOut(iter: quote.getAssetAmountsOutIterator()!)
                output += "Output amounts with slippage\n"
                output += displayAmountOut(iter: try quote.getAssetAmountsOutWithSlippageIterator())

                // Prepare a transaction group for burning
                // Note that some transactions need to be signed with LogicSig account, and they were signed in the function.
                let txGroup = try pool.prepareBurnTransactions(from: quote, burnerAddress: userAddress)

                // Some transactions that need the user signatures are signed here
                try txGroup.sign(account)

                // Submit a group of transaction to the blockchain
                let txId = tinymanClient.submit(txGroup, wait: true, error: &error)
                if (error != nil) {
                    throw error!;
                }

                output += "Liquidity was removed in txid \(txId)\n"
                
                let info = try pool.fetchPosition(userAddress)
                let share = Decimal(string: info.getShare())! * Decimal(string: "100")!

                output += "Pool tokens: ID-\(info.getLiquidityAssetAmount()!.getAsset()!.getId()) = \(info.getLiquidityAssetAmount()!.getAmount())\n"
                output += "Asset1: ID-\(info.getAssetAmount1()!.getAsset()!.getId()) = \(info.getAssetAmount1()!.getAmount())\n"
                output += "Asset2: ID-\(info.getAssetAmount2()!.getAsset()!.getId()) = \(info.getAssetAmount2()!.getAmount())\n"
                output += "Share of pool: \(share)%"
                print(output)
            } catch let error {
                print(error)
                output = error.localizedDescription
            }
        }
    }
    
    var body: some View {
        NavigationView {
            VStack(spacing: 20) {
                Text(output)
                    .offset(y: -60)

                TextField("Enter asset 1 id...", text: $asset1Id)
                    .multilineTextAlignment(.center)
                
                Text("Remove liquidity")
                    .bold()
                    .frame(width: 280, height: 50)
                    .background(Color.blue)
                    .foregroundColor(.white)
                    .cornerRadius(10)
                    .onTapGesture {
                        self.removeLiquidity()
                    }
            }
        }
    }
}

struct SwapView: View {
    @State var output: String = "This example swap an asset id for ALGO token"
    @State var asset1Id: String = ""
    
    func swap() {
        DispatchQueue.global(qos: .userInitiated).async {
            output = "loading..."
            
            do {
                if (asset1Id.isEmpty) {
                    output = "Asset 1 ID is empty"
                    
                    return
                }
                
                var error: NSError?
                let account = try getAccount()
                let userAddress = account.getAddress()
                let (algodClient, tinymanClient) = try createClients(userAddress: userAddress)

                // Check whether the user already opted in the app or not, if not let the user opt in
                try optInAppIfNeeded(tinymanClient: tinymanClient, account: account)

                // Check whether the user already opted in for the asset or not, if not let the user opt in
                try optInAssetIfNeeded(tinymanClient: tinymanClient, account: account, assetId: asset1Id)

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
                    throw error!;
                }
                
                // Check whether the user already opted in for the liquidity asset or not, if not let the user opt in
                try optInAssetIfNeeded(tinymanClient: tinymanClient, account: account, assetId: pool.getLiquidityAsset()!.getId())
                
                let tokenAssetAmount = TinymanNewAssetAmount()!

                // Set asset
                try tokenAssetAmount.setAsset(token)

                // Note that 500000 is equal to 500000 / 10 ** decimals (if decimals is 6), which is 0.5 tokens
                try tokenAssetAmount.setAmount("500000")

                // Fetch mint quote used when submit minting transactions
                let quote = try pool.fetchFixedInputSwapQuote(tokenAssetAmount, slippage: "0.05")

                // Calculate price after applying the slippage
                let priceWithSlippage = quote.getPriceWithSlippage(&error)
                if (error != nil) {
                    throw error!;
                }

                // Calculate output amount after applying the slippage
                let amountOutWithSlippage = try quote.getAssetAmountOutWithSlippage()

                output = "ID-\(token.getId()) per ALGO: \(quote.getPrice())\n"
                output += "ID-\(token.getId()) per ALGO (worst case): \(priceWithSlippage)\n"
                output += "Swapping ID-\(quote.getAssetAmountIn()!.getAsset()!.getId()) = \(quote.getAssetAmountIn()!.getAmount()) to ID-\(amountOutWithSlippage.getAsset()!.getId()) = \(amountOutWithSlippage.getAmount())\n"
                
                // Prepare a transaction group for swappingg
                // Note that some transactions need to be signed with LogicSig account, and they were signed in the function.
                let txGroup = try pool.prepareSwapTransactions(from: quote, swapperAddress: userAddress)

                // Some transactions that need the user signatures are signed here
                try txGroup.sign(account)

                // Submit a group of transaction to the blockchain
                let txId = tinymanClient.submit(txGroup, wait: true, error: &error)
                if (error != nil) {
                    throw error!;
                }

                output += "Swapped with txid \(txId)"
                print(output)
            } catch let error {
                print(error)
                output = error.localizedDescription
            }
        }
    }
    
    var body: some View {
        NavigationView {
            VStack(spacing: 20) {
                Text(output)
                    .offset(y: -60)

                TextField("Enter asset 1 id...", text: $asset1Id)
                    .multilineTextAlignment(.center)
                
                Text("Swap")
                    .bold()
                    .frame(width: 280, height: 50)
                    .background(Color.blue)
                    .foregroundColor(.white)
                    .cornerRadius(10)
                    .onTapGesture {
                        self.swap()
                    }
            }
        }
    }
}

struct RedeemView: View {
    @State var output: String = "This example redeem excess tokens from swapping"
    @State var asset1Id: String = ""

    func redeem() {
        DispatchQueue.global(qos: .userInitiated).async {
            output = "loading..."
            
            do {
                if (asset1Id.isEmpty) {
                    output = "Asset 1 ID is empty"
                    
                    return
                }
                
                var error: NSError?
                let account = try getAccount()
                let userAddress = account.getAddress()
                let (algodClient, tinymanClient) = try createClients(userAddress: userAddress)

                // Check whether the user already opted in the app or not, if not let the user opt in
                try optInAppIfNeeded(tinymanClient: tinymanClient, account: account)

                // Check whether the user already opted in for the asset or not, if not let the user opt in
                try optInAssetIfNeeded(tinymanClient: tinymanClient, account: account, assetId: asset1Id)

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
                    throw error!;
                }
                
                // Check whether the user already opted in for the liquidity asset or not, if not let the user opt in
                try optInAssetIfNeeded(tinymanClient: tinymanClient, account: account, assetId: pool.getLiquidityAsset()!.getId())
                
                // Fetch excess amount resulting from the swap
                let redeemQuoteIterator = try tinymanClient.fetchExcessAmount(userAddress)
                
                // Filter redeem quotes matching with TOKEN
                let quote = try pool.getRedeemQuoteMatchesAssetID(asset1Id, iter: redeemQuoteIterator)
                
                if (TinymanIsNullRedeemQuote(quote)) {
                    output = "No excess amount to be redeemed\n"
                    print(output)

                    return
                }

                output = "There is ID-\(quote.getAssetAmount()!.getAsset()!.getId()) = \(quote.getAssetAmount()!.getAmount()) to be redeemed\n"

                // Prepare a transaction group for redeeming
                // Note that some transactions need to be signed with LogicSig account, and they were signed in the function.
                let txGroup = try pool.prepareRedeemTransactions(from: quote, redeemerAddress: userAddress)

                // Some transactions that need the user signatures are signed here
                try txGroup.sign(account)

                // Submit a group of transaction to the blockchain
                let txId = tinymanClient.submit(txGroup, wait: true, error: &error)
                if (error != nil) {
                    throw error!;
                }

                output += "Swapped with txid \(txId)"
                print(output)
            } catch let error {
                print(error)
                output = error.localizedDescription
            }
        }
    }
    
    var body: some View {
        NavigationView {
            VStack(spacing: 20) {
                Text(output)
                    .offset(y: -60)

                TextField("Enter asset 1 id...", text: $asset1Id)
                    .multilineTextAlignment(.center)
                
                Text("Redeem")
                    .bold()
                    .frame(width: 280, height: 50)
                    .background(Color.blue)
                    .foregroundColor(.white)
                    .cornerRadius(10)
                    .onTapGesture {
                        self.redeem()
                    }
            }
        }
    }
}
