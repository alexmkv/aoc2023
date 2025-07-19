import java.io.File

class Day09(val fname: String, val method: Int = 1) {
    fun calc(): Long {
        val line = File(fname).readText()
        val r = mutableListOf<Long>()
        line.forEachIndexed{ i, c ->  addN(r, if (i % 2 == 0) (i/2).toLong() else -1, c.toString().toInt()) }

        if (method == 2) {
            return process2(line)
        } else {
            var i = 0
            var j = r.size - 1
            while (i < j) {
                if (r[j] == -1L) {
                    j--
                } else if (r[i] != -1L) {
                    i++
                } else {
                    r[i] = r[j]
                    r[j] = -1L
                    i++
                    j--
                }
            }
        }
        return r.foldIndexed(0L) {
            i, acc, v ->
            acc + (if (v == -1L) 0L else i*v)
        }
    }
    private fun process2(l: String): Long {
        var r = mutableListOf<Pair<Int, Long>>()
        l.forEachIndexed{ i, c -> r.add(Pair(c.toString().toInt(), if (i % 2 == 0) (i/2).toLong() else -1L)) }
        // we have array (cnt, idx or -1 - if empty)
        // starting from rightmost not empty try to fit in free position adding it to right and reducing size of existing free space
        var i = r.size
        while (i > 0) {
            i--
            if (r[i].second != -1L) {
                for (j in 0..<i) {
                    if (r[j].second == -1L && r[j].first >= r[i].first) {
                        if (r[j].first == r[i].first) {
                            r[j] = r[i]
                            r[i] = Pair(r[j].first, -1L)
                        } else {
                            r.add(j, r[i])
                            i++
                            r[j+1] = Pair(r[j+1].first - r[i].first, -1L)
                            r[i] = Pair(r[i].first, -1L)
                        }
                        break
                    }
                }
            }
        }
        var idx = 0
        var rr = 0L
        for (i in 0..<r.size) {
            if (r[i].second == -1L) {
                idx += r[i].first
            } else {
                for (j in 0..<r[i].first) {
                    rr += idx*r[i].second
                    idx++
                }
            }
        }
        return rr
    }
    private fun addN(l: MutableList<Long>, v: Long, c: Int) {
        for (i in 0..<c) {
            l.add(v)
        }
    }
}