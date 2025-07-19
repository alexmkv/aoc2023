import java.io.File

class Day07 {
    fun go() {
        expectT("7.1.0", {process("data/data07_1_0.txt", 1)}, 3749)
        expectT("7.1.1", {process("data/data07_1_1.txt", 1)}, 3312271365652)
        expectT("7.1.0", {process("data/data07_1_0.txt", 2)}, 11387)
        expectT("7.1.1", {process("data/data07_1_1.txt", 2)}, 509463489296712)
    }

    fun process(fname: String, method: Int): Long {
        val lines = File(fname).readLines()
        return lines.fold(0) { acc, line ->
            val s0 = line.split(':')
            val r = s0[0].toLong()
            val s1 = s0[1].split(" ").filter{ it != "" }.map{ it.toLong()}
            acc + (if (tryThis(r, s1, s1.size - 1, method)) r else 0)
        }
    }
    fun tryThis(r: Long, s1: List<Long>, pos: Int, method: Int): Boolean {
        if (pos == 0) return s1[0] == r
        if (r - s1[pos] > 0 && tryThis(r - s1[pos], s1, pos - 1, method)) return true
        if (r % s1[pos] == 0L && tryThis(r / s1[pos], s1, pos - 1, method)) return true
        if (method == 2) {
            val rs = r.toString()
            val rs1 = rs.removeSuffix(s1[pos].toString())
            if (rs != rs1) {
                if (rs1 != "") {
                    return tryThis(rs1.toLong(), s1, pos - 1, method)
                } else {
                    return tryThis(0L, s1, pos - 1, method)
                }
            }
            if (r.toString().endsWith(s1[pos].toString())) {
                return tryThis(r.toString().removeSuffix(s1[pos].toString()).toLong(), s1, pos - 1, method)
            }
        }
        return false
    }

}