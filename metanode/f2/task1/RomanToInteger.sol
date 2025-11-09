// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract RomanToInteger {
    
    function romanToInt(string memory s) external pure returns (uint) {
        // 罗马数字到数值的映射
        mapping(bytes1 => uint256) private romanMap;
        romanMap["I"] = 1;
        romanMap["V"] = 5;
        romanMap["X"] = 10;
        romanMap["L"] = 50;
        romanMap["C"] = 100;
        romanMap["D"] = 500;
        romanMap["M"] = 1000;

        bytes memory sBytes = bytes(s);
        uint256 total = 0;
        uint256 n = sBytes.length;

        for (uint i = 0; i < n; i++) {
            // 前一个字符值小于当前，说明是减法组合（如 IV=4）
            if (i > 0 && romanMap[sBytes[i]] > romanMap[sBytes[i-1]]) {
                total += romanMap[sBytes[i]] - 2 * romanMap[sBytes[i-1]];
            } else {
                total += romanMap[sBytes[i]];
            }
        }

        return total;
    }

}