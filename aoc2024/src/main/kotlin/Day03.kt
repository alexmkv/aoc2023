package com.oubloom

import expect
import java.io.File

class Day03 {
    private val MUL_REGEXP = "mul\\(([0-9]{1,3}),([0-9]{1,3})\\)"

    fun go() {
        expect("Part0.0", processFile("data/data03_0.txt", 1), 161)
        expect("Part0.1", processFile("data/data03_1.txt", 1), 170068701)
        expect("Part1.0", processFile("data/data03_0_1.txt", 2), 48)
        expect("Part1.0", processFile("data/data03_1.txt", 2), 78683433)
    }

    fun processFile(fileName: String, calcMethod: Int):Int {
        val r = Regex(if (calcMethod == 2) "($MUL_REGEXP|do\\(\\)|don\\'t\\(\\))" else MUL_REGEXP)
        val startGroupIdx = if (calcMethod == 2) 2 else 1
        return r.findAll(File(fileName).readText()).fold(Pair(true,0)) { acc, v ->
            if (v.groupValues[0].startsWith("mul")) {
                if (acc.first) {
                    Pair(true, acc.second + calcMul(v, startGroupIdx))
                } else
                    Pair(false, acc.second)
            } else {
                Pair(!v.groupValues[0].startsWith("don't"), acc.second)
            }
        }.second
    }
    fun calcMul(mr: MatchResult, startGroupIdx: Int): Int {
        return mr.groupValues[startGroupIdx].toInt() * mr.groupValues[startGroupIdx+1].toInt()
    }

}