
import java.io.File

class Day06 {
    fun go() {
        expectT("6.1.0", {process("data/data06_0.txt", 1)}, 41)
        expectT("6.1.1", {process("data/data06_1.txt", 1)}, 5269)
        expectT("6.2.0", {process("data/data06_0.txt", 2)}, 6)
        expectT("6.2.1", {process("data/data06_1.txt", 2)}, 1957)
    }
    fun process(fname: String, method: Int):Int {
        if (method == 2) {
            return process2(fname)
        }
        var lines = File(fname).readLines()
        val pr = MutableList(lines.size) { MutableList(lines[0].length) { 0} }
        var x = 0;
        var y = 0;
        var dir = 0
        for (i in 0..<lines.size) {
            val p = findAny(lines[i])
            if (p != -1) {
                y = i
                x = p
                dir = dirNum(lines[i][p])
                break
            }
        }
        while (true) {
            if (dir == 1) {
                x += 1
                if (x >= lines[0].length) break
                if (lines[y][x] == '#') {
                    dir = 2
                    x -= 1
                }
            } else if (dir == 2) {
                y += 1
                if (y >= lines.size) break
                if (lines[y][x] == '#') {
                    dir = 4
                    y -= 1
                }
            } else if (dir == 4) {
                x -= 1
                if (x < 0) break
                if (lines[y][x] == '#') {
                    dir = 8
                    x += 1
                }
            } else if (dir == 8) {
                y -= 1
                if (y < 0) break
                if (lines[y][x] == '#') {
                    dir = 1
                    y += 1
                }
            }
            if (pr[y][x] and dir == 0) {
                pr[y][x] = pr[y][x] or dir
            } else {
                break
            }
        }
        var r = 0
        pr.forEach { v -> v.forEach { if (it != 0) r++ } }
        return r
    }

    fun process2(fname: String):Int {
        var lines = File(fname).readLines()
        var l2 = lines.map { it.toCharArray() }
        var r = 0
        for (i in  0..<l2.size) {
            for (j in 0..<l2[i].size) {
                if (l2[i][j] != '.') continue
                l2[i][j] = '#'
                if (process3(l2)) {
                    r++
                }
                l2[i][j] = '.'
            }
        }

        return r
    }
    fun process3(lines: List<CharArray>):Boolean {
        val pr = MutableList(lines.size) { MutableList(lines[0].size) { 0} }
        var x = 0;
        var y = 0;
        var dir = 0
        for (i in 0..<lines.size) {
            val p = findAny(String(lines[i]))
            if (p != -1) {
                y = i
                x = p
                dir = dirNum(lines[i][p])
                break
            }
        }
        while (true) {
            if (dir == 1) {
                x += 1
                if (x >= lines[0].size) break
                if (lines[y][x] == '#') {
                    dir = 2
                    x -= 1
                }
            } else if (dir == 2) {
                y += 1
                if (y >= lines.size) break
                if (lines[y][x] == '#') {
                    dir = 4
                    y -= 1
                }
            } else if (dir == 4) {
                x -= 1
                if (x < 0) break
                if (lines[y][x] == '#') {
                    dir = 8
                    x += 1
                }
            } else if (dir == 8) {
                y -= 1
                if (y < 0) break
                if (lines[y][x] == '#') {
                    dir = 1
                    y += 1
                }
            }
            if (pr[y][x] and dir == 0) {
                pr[y][x] = pr[y][x] or dir
            } else {
                return true
            }
        }
        return false
    }




    fun dirNum(c: Char) : Int {
        return when (c) {
            '>' -> 1
            'v' -> 2
            '<' -> 4
            '^' -> 8
            else -> 0
        }
    }

    fun findAny(l: String): Int {
        val r = Regex("(\\<|\\>|\\^|\\v)").find(l)
        if (r != null) {
            return r.range.start
        }
        return -1
    }

}