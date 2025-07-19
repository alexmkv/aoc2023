package com.oubloom.Day04

import expect
import java.io.File

class Day04 {
    fun go() {
        expect("4.0", process(parse("data/data04_0.txt")), 18)
        expect("4.1", process(parse("data/data04_1.txt")), 2557)
        expect("4.0.2", process2(parse("data/data04_0.txt")), 9)
        expect("4.0.2", process2(parse("data/data04_1.txt")), 1854)
    }

    private fun process2(fld: Field):Int {
        var res: Int = 0
        for (y in 1..<fld.sizeY()-1) {
            for (x in 1 ..< fld.sizeX()-1) {
                val pt = Pt(x, y)
                if (fld.valueAt(pt) == 'A') {
                    if ((fld.at(x-1, y-1) == 'M' && fld.valueAtDir(pt, 7) == 'S')
                        || (fld.valueAtDir(pt, 0) == 'S' && fld.valueAtDir(pt, 7) == 'M')
                    ) {
                        if ((fld.valueAtDir(pt, 2) == 'M' && fld.valueAtDir(pt, 5) == 'S')
                            || (fld.valueAtDir(pt, 2) == 'S' && fld.valueAtDir(pt, 5) == 'M')
                        ) {
                            res++
                        }
                    }
                }
            }
        }
        return res
    }

    private fun process(fld: Field):Int {
        var pt = Pt(0, 0)
        var res: Int = 0
        while (fld.validPt(pt)) {
            while (fld.validPt(pt)) {
                if (fld.valueAt(pt) == 'X') {
                    //println("Found tree at $pt")
                    for (dir in 0..7) {
                        val inc = processFrom(fld, pt, dir, 3)
                        if (inc > 0) {
                            //println("Found at $pt, dir $dir, inc $inc")
                        }
                        res += inc
                    }
                }
                pt = Pt(pt.x+1, pt.y)
            }
            pt = Pt(0, pt.y+1)
        }
        //println(res)
        return res
    }
    private fun processFrom(fld: Field, pt: Pt, dir: Int, left: Int): Int {
        if (left == 0) {
           // println("good")
            return 1
        }
        val pt0 = pt.move(dir)
        if (!fld.validPt(pt0)) {
            return 0
        }
        val nextVal = if (fld.valueAt(pt0) == validLetter(left)) {
            //println("$pt - $left")
            processFrom(fld, pt0, dir, left-1)
        } else 0
        return /*processFrom(fld, pt0, dir, left) +*/ nextVal
    }

    fun validLetter(left:Int):Char {
        //if (left == 4) return 'X'
        if (left == 3) return 'M'
        if (left == 2) return 'A'
        if (left == 1) return 'S'
        throw IllegalArgumentException("Invalid left $left")
    }

    private fun parse(fileName: String): Field {
        val f = File(fileName)
        val lines = f.readLines().map {
                line -> line.toCharArray()
        }
        return Field(lines)
    }
}

private data class Pt(val x: Int, val y: Int) {
    fun move(dir: Int): Pt {
        val cdir = if (dir >= 4) dir+1 else dir
        val dx = cdir % 3 - 1
        val dy = cdir / 3 - 1
        return Pt(x + dx, y + dy)
    }
}

private data class Field(val field: List<CharArray>) {
    fun validPt(pt: Pt): Boolean {
        return pt.x >= 0 && pt.x < sizeX() && pt.y >= 0 && pt.y < sizeY()
    }
    fun at(x: Int, y: Int): Char {
        return field[y][x]
    }

    fun valueAt(pt: Pt): Char {
        return field[pt.y][pt.x]
    }
    fun valueAtDir(pt: Pt, dir: Int): Char {
        val pt0 = pt.move(dir)
        return valueAt(pt0)
    }
    fun sizeX(): Int {
        return field[0].size
    }
    fun sizeY(): Int {
        return field.size
    }
}

