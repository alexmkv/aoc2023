

fun sign(v: Int): Int {
    return if (v > 0) 1 else if (v < 0) -1 else 0
}

fun expect(header: String, v: Int, e: Int) {
    print(header)
    print(": $v")
    if (v == e)
        println(" - OK")
    else
        println(" <> $e - FAIL")
}

fun <T> expectT(header: String, act: ()->T,  e: T) {
    print(header)
    val start = System.currentTimeMillis()
    val v= act()
    val duration = System.currentTimeMillis() - start
    print(": $v")
    if (v == e)
        println(", ($duration) - OK")
    else
        println(" <> $e, ($duration) - FAIL")
}