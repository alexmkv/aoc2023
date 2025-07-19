package org.example


fun main() {
    println("Run tests")
}

fun subarraysDivByK(nums: IntArray, k: Int):Int {
    val sameKModeCount = IntArray(k)
    sameKModeCount[0] = 1
    var sum = 0
    var result = 0
    for (v in nums) {
        sum = (sum + v)%k
        if (sum < 0) {
            sum += k
        }
        result += sameKModeCount[sum]
        sameKModeCount[sum]++
    }
    return result
}