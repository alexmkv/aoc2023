package com.oubloom.Day08

import expectT
import java.io.File

private data class Pt(val x: Int, val y: Int);

class Day08 {
    fun go() {
        expectT("8.1.0", {process("data/data08_1_0.txt", 1)}, 14)
        expectT("8.1.1", {process("data/data08_1_1.txt", 1)}, 351)
        expectT("8.2.0", {process("data/data08_1_0.txt", 2)}, 34)
        expectT("8.2.0", {process("data/data08_2_0.txt", 2)}, 9)
        expectT("8.2.1", {process("data/data08_1_1.txt", 2)}, 1259)
    }

    private fun process(fname: String, method: Int): Long {
        val lines = File(fname).readLines()
        val a = lines.map { it.toCharArray()  }
        val uniqs = a.fold(mutableSetOf<Char>()) {
            acc, line -> line.fold(acc) { acc2, c -> if (c != '.') acc2.add(c); acc2 }
        }
        val res = mutableSetOf<Pt>()
        uniqs.forEach {
            val pos = mutableListOf<Pt>()
            a.forEachIndexed { y, line ->
                line.forEachIndexed { x, c ->
                    if (c == it) {
                        pos.add(Pt(x, y))
                    }
                }
            }
            for (i in 0..<pos.size) {
                for (j in i+1..<pos.size) {
                    if (method == 2) {
                        res.add(pos[i])
                        res.add(pos[j])
                    }
                    collectAll(pos[i], pos[j], a, res, method)
                    collectAll(pos[j], pos[i], a, res, method)
                }
            }
        }
        return res.size.toLong()
    }
    private fun collectAll(pt1: Pt, pt2: Pt, a: List<CharArray>, res: MutableSet<Pt>, method: Int) {
        var m = 1;
        while (true) {
            val newPt = Pt((m+1) * pt2.x - m*pt1.x, (m+1) * pt2.y - m*pt1.y)
            if (validPt(newPt, a)) res.add(newPt) else break
            if (method == 1) break
            m++
        }
    }
    private fun validPt(pt: Pt, a: List<CharArray>): Boolean {
        return pt.y >= 0 && pt.y < a.size && pt.x >= 0 && pt.x < a[0].size
    }

}

