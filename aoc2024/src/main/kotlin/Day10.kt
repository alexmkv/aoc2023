import java.io.File

class Day10(val fname: String, val method: Int=1) {
    var gp = mutableSetOf<Int>()
    var cp = 1;
    fun calc():Long {
        val map = File(fname).readLines()
        return map.foldIndexed(0L) { y, acc, line ->
            acc + line.foldIndexed(0L) { x, a2, c -> a2 + if (c=='0') calcScore(map, y, x) else 0L }
        }
    }
    fun calcScore(mp: List<String>, y: Int, x: Int ): Long {
        val m = mp.map { MutableList(it.length, { 0}) }
        if (method == 1)
            return tryP(mp, m, y, x, 0 )
        else {
            return tryP2(mp, m, y, x, 0 )

        }
    }
    fun tryP(mp: List<String>, m: List<MutableList<Int>>, y: Int, x: Int, exp: Int): Long {
        if (x < 0 || x >= mp[0].length) return 0
        if (y < 0 || y >= mp.size) return 0
        if (m[y][x] != 0 ) return 0
        if (mp[y][x] != (exp+'0'.toInt()).toChar()) return 0
        m[y][x] = 1
        if (mp[y][x] == '9') return 1
        //println("$y, $x, ${mp[y][x]}")
        return tryP(mp, m, y-1, x, exp+1)+
                tryP(mp, m, y+1, x, exp+1)+
                tryP(mp, m, y, x-1, exp+1)+
                tryP(mp, m, y, x+1, exp+1)
    }

    fun tryP2(mp: List<String>, m: List<MutableList<Int>>, y: Int, x: Int, exp: Int): Long {
        if (x < 0 || x >= mp[0].length) return 0
        if (y < 0 || y >= mp.size) return 0
        if (mp[y][x] != (exp+'0'.toInt()).toChar()) return 0
        println("$y, $x, ${mp[y][x]}")
        if (m[y][x] != 0 ) {
            println("good")
            return 1
        }
        m[y][x] = 1
        if (mp[y][x] == '9') {
            println("good")
            return 1
        }

        val res = tryP2(mp, m, y-1, x, exp+1)+
                tryP2(mp, m, y+1, x, exp+1)+
                tryP2(mp, m, y, x-1, exp+1)+
                tryP2(mp, m, y, x+1, exp+1)
        if (res == 0L) {
            m[y][x] = 0
        }
        return res
    }

}