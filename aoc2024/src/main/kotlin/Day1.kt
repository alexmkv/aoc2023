package com.oubloom

import java.io.File
import kotlin.math.abs

class Day1 {
    fun part1_0(): Int {
        val r = parse("data/data1_0.txt")
        return r.first.foldIndexed(0, { index, acc, value ->
            //println("$value, ${r.second[index]}, ${abs(value - r.second[index])}, $acc");
            acc + abs(value - r.second[index])})
    }
    fun part1_1(): Int {
        val r = parse("data/data1_1.txt")
        return r.first.foldIndexed(0, { index, acc, value ->
            acc + abs(value - r.second[index])})
    }
    fun part2_0(): Int {
        val r = parse2("data/data1_0.txt")
        val freq = r.second.fold(mutableMapOf<Int, Int>(), { acc, value ->
            acc[value] = acc.getOrDefault(value, 0) + 1
            acc
        })
        return r.first.fold(0, { acc, value ->
            freq.getOrDefault(value, 0) * value + acc
        })
    }
    fun part2_1(): Int {
        val r = parse2("data/data1_1.txt")
        val freq = r.second.fold(mutableMapOf<Int, Int>(), { acc, value ->
            acc[value] = acc.getOrDefault(value, 0) + 1
            acc
        })
        return r.first.fold(0, { acc, value ->
            freq.getOrDefault(value, 0) * value + acc
        })
    }

    fun parse(fname: String): Pair<List<Int>, List<Int>> {
        val file = File(fname)
        val list1 = mutableListOf<Int>()
        val list2 = mutableListOf<Int>()
        file.forEachLine { line ->
            val (num1, num2) = line.split("\\s+".toRegex()).map { it.toInt() }
            list1.add(num1)
            list2.add(num2)
        }
        list1.sort()


        list2.sort()
        return Pair(list1, list2)
    }
    fun parse2(fname: String): Pair<List<Int>, List<Int>> {
        val file = File(fname)
        val list1 = mutableListOf<Int>()
        val list2 = mutableListOf<Int>()
        file.forEachLine { line ->
            val (num1, num2) = line.split("\\s+".toRegex()).map { it.toInt() }
            list1.add(num1)
            list2.add(num2)
        }
        return Pair(list1, list2)
    }
}

