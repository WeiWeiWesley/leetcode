# Decode String

## 題目連結
https://leetcode.com/problems/decode-string/

## 題目描述
Given an encoded string, return its decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; No extra white spaces, square brackets are well-formed, etc.

## 範例

### Example 1:
```
Input: s = "3[a]2[bc]"
Output: "aaabcbc"
```

### Example 2:
```
Input: s = "3[a2[c]]"
Output: "accaccacc"
```

### Example 3:
```
Input: s = "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"
```

## 難度
Medium

## 標籤
String, Stack, Recursion
