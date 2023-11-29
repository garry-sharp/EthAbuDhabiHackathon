// SPDX-License-Identifier: MIT
pragma solidity >=0.8.19;

import "./ERC20.sol";
import "./IUniswapV2Pair.sol";

contract Aggregator {
    struct PriceInfo {
        address quote;
        address base;
        uint256 quoteVolume;
        uint256 baseVolume;
        address dexAddress;
    }

    mapping(bytes32 => PriceInfo) public keyToBestBuy;
    mapping(bytes32 => PriceInfo) public keyToBestSell;

    function getPairKey(
        address base,
        address quote
    ) public pure returns (bytes32) {
        if (base < quote) {
            return keccak256(abi.encodePacked(base, quote));
        } else {
            return keccak256(abi.encodePacked(quote, base));
        }
    }

    function updateBestBuy(
        address base,
        address quote,
        uint256 baseVolume,
        uint256 quoteVolume,
        address tradeContract
    ) public {
        keyToBestBuy[getPairKey(base, quote)] = PriceInfo(
            quote,
            base,
            quoteVolume,
            baseVolume,
            tradeContract
        );
    }

    function updateBestSell(
        address base,
        address quote,
        uint256 baseVolume,
        uint256 quoteVolume,
        address tradeContract
    ) public {
        keyToBestSell[getPairKey(base, quote)] = PriceInfo(
            quote,
            base,
            quoteVolume,
            baseVolume,
            tradeContract
        );
    }

    function updateBestAll(
        address base,
        address quote,
        uint256 baseVolume,
        uint256 quoteVolume,
        address tradeContract
    ) public {
        PriceInfo memory pi = PriceInfo(
            quote,
            base,
            quoteVolume,
            baseVolume,
            tradeContract
        );
        keyToBestBuy[getPairKey(base, quote)] = pi;
        keyToBestSell[getPairKey(base, quote)] = pi;
    }

    function getERC20(address tokenAddress) private pure returns (ERC20) {
        return ERC20(tokenAddress);
    }

    function swap(address token0, address token1, uint256 amt) public {
        //TODO event
        // PriceInfo memory k = keyToPriceInfo[getPairKey(token0, token1)];
        // if (k.dexAddress != address(0)){
        //     IUniswapV2Pair pairContract = IUniswapV2Pair(k.dexAddress);
        //     pairContract.swap()
        // }else{
        //     //TODO can't do swap
        // }
    }
}
