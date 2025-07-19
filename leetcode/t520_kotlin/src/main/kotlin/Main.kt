package org.example

import java.util.PriorityQueue

fun main() {
    println("Hello World!")
    val s = Solution()


}

class Solution {
    data class Prj (val profit: Int, val capital:  Int, var used: Boolean)
    fun findMaximizedCapital(k: Int, w: Int, profits: IntArray, capital: IntArray): Int {
        val prjs = profits.mapIndexed({i, v -> Prj(v, capital[i], false)})
            .sortedWith({a, b -> Integer.compare(a.capital, b.capital)})
        var profitQueue = PriorityQueue<Prj>({a,b -> Integer.compare(b.profit,a.profit)})
        var capital = w
        var prjI = 0
        for (i in 1..k) {
            while (prjI < prjs.size && prjs[prjI].capital <= capital) {
                profitQueue.offer(prjs[prjI])
                prjI++
            }
            //if (profitQueue.isEmpty()) break
            val polledPrj = profitQueue.poll()
            if (polledPrj == null) {
                break
            }
            capital += polledPrj.profit
        }
        return capital
    }

    fun findMaximizedCapitalBruteforce(k: Int, w: Int, profits: IntArray, capital: IntArray): Int {
        val prjs = profits.mapIndexed({i, v -> Prj(v, capital[i], false)})
            .sortedWith({a, b -> if (a.profit == b.profit) Integer.compare(a.capital, b.capital) else Integer.compare(b.profit, a.profit)})
        var capital = w
        for (i in 1..k) {
            var m:Prj? = null
            for (pr in prjs) {
                if (pr.used || pr.capital > capital) {
                    continue
                }
                if (m == null) {
                    m = pr
                    break
                } /*else if (m.profit <= pr.profit) {
                    m = pr
                }*/
            }
            if (m == null) break
            m.used = true
            capital += m.profit
        }
        return capital
    }
}