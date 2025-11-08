// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract IntegerToRoman {

    function intToRoman(uint256 num) external pure returns (string memory) {
        // 定义所有可能的数值及其对应的罗马数字
        uint256[] memory values = new uint256[](13);
        string[] memory symbols = new string[](13);

        values = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];
        symbols = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"];

        bytes memory result = new bytes(0);

        for (uint256 i = 0; i < 13; i++) {
            while (num >= values[i]) {
                num -= values[i];
                result = abi.encodePacked(result, symbols[i]);
            }
            if (num == 0) break;
        }

        return string(result);
    }
}