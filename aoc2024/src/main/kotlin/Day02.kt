package com.oubloom

import expect
import sign
import java.io.File
import kotlin.math.abs
import kotlin.math.max

const val IMPOSSIBLE_VALUE = Int.MAX_VALUE

class Day02 {
    fun go() {
        expect("Part0.0", processFile("data/data02_0.txt", false), 2)
        expect("Part0.1", processFile("data/data02_1.txt", false), 314)
        expect("Part1.0", processFile("data/data02_0.txt", true), 4)
        expect("Part1.1", processFile("data/data02_1.txt", true), 373)
    }
    fun processFile(fname: String, retry: Boolean):Int {
        return parse(fname).map{ if (isSafeLine(it, retry)) 1 else 0 }.sum()
    }
    fun isSafeLine(r: List<Int>, retry: Boolean, startI: Int = 0, skipI: Int = IMPOSSIBLE_VALUE): Boolean {
        var previousValue = IMPOSSIBLE_VALUE
        var previousDiff = IMPOSSIBLE_VALUE
        for (i in startI..<r.size) {
            if (i == skipI) continue
            if (previousValue != IMPOSSIBLE_VALUE) {
                val diff = r[i] - previousValue
                if (!checkDiff(diff, previousDiff)) {
                    if (retry) {
                        if (isSafeLine(r, false, max(i-2, 0), i)) return true
                        if (isSafeLine(r, false, max(i-3, 0), i-1)) return true
                        if (isSafeLine(r, false, max(i-3, 0), i-2)) return true
                    }
                    return false
                }
                previousDiff = diff
            }
            previousValue = r[i]
        }
        return true
    }

    fun checkDiff(diff: Int, previousDiff: Int): Boolean {
        if (abs(diff) > 3 || diff == 0) {
            return false
        }
        if (previousDiff != IMPOSSIBLE_VALUE && sign(diff) != sign(previousDiff)) {
            return false
        }
        return true
    }

    fun parse(fileName: String): List<List<Int>> {
        return File(fileName).readLines().map {
            line -> line.split("\\s+".toRegex()).map { it.toInt() }
        }
    }

}

