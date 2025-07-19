package org.example

fun main() {
    println("Hello World!")
}

class Solution {
    fun minPatches(nums: IntArray, n: Int): Int {
        // On each step we will add up
        var maxReached: Long = 0
        var patches = 0
        for (v in nums) {
            while (v > maxReached+1) {
                patches++
                maxReached += maxReached + 1
                if (maxReached >= n) {
                    break
                }
            }
            maxReached += v
            if (maxReached >= n) {
                break
            }
        }
        while (maxReached < n) {
            patches++
            maxReached += maxReached + 1
        }
        return patches
    }
}