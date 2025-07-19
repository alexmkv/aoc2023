import java.io.File

class Day05 {
    fun go() {
        expect("5.0", process("data/data05_0.txt", 1), 143)
        expect("5.1", process("data/data05_1.txt", 1), 5964)
        expect("5.2.0", process("data/data05_0.txt", 2), 123)
        expect("5.2.1", process("data/data05_1.txt", 2), 4719)
    }

    fun process(fileName: String, method: Int): Int {
        var before = mutableMapOf<String, MutableSet<String>>()
        return File(fileName).readLines().fold(-1) { acc, line ->
            if (acc == -1) {
                if (line == "") {
                    0
                } else {
                    val r = line.split("|")
                    addBefore(before, r[0], r[1])
                    acc
                }
            } else {
                if (method == 1) {
                    acc + calculateOne(before, line)
                } else {
                    acc + calculateTwo(before, line)
                }
            }
        }
    }


    fun calculateOne(before: MutableMap<String, MutableSet<String>>, line: String): Int {
        val r = line.split(",")
        var thisValid = true
        for (i in 1..<r.size) {
            for (j in 0..<i) {
                thisValid = thisValid && valid(before, r[j], r[i])
                if (!thisValid) break
            }
            if (!thisValid) break
        }
        if (thisValid) {
            return r[r.size/2].toInt()
        }
        return 0
    }
    fun calculateTwo(before: MutableMap<String, MutableSet<String>>, line: String): Int {
        val r = ArrayList(line.split(","))
        var allValid = true
        for (i in 1..<r.size) {
            for (j in 0..<i) {
                val thisValid = valid(before, r[j], r[i])
                if (!thisValid) {
                    val t = r[j]
                    r[j] = r[i]
                    r[i] = t
                    allValid = false
                }
            }
        }
        if (!allValid) {
            return r[r.size/2].toInt()
        }
        return 0
    }


    fun addBefore(before: MutableMap<String, MutableSet<String>>, beforeN: String, afterN: String) {
        if (before.containsKey(afterN)) {
            before[afterN]!!.add(beforeN)
        } else {
            before[afterN] = mutableSetOf(beforeN)
        }
    }

    fun valid(before: MutableMap<String, MutableSet<String>>, beforeN: String, afterN: String): Boolean {
        val b = before[beforeN]
        if (b == null) {
            return true
        }
        if (b.contains(afterN)) {
            return false
        }
        return true
    }
}