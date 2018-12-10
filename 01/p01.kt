import java.io.File

val p01 =  fun(input: List<String>) {
    var numbers = input.map {it.toInt()}

    val sum = numbers.reduce{sum, it->sum+it}
    println(sum)

    var freq = 0
    var freqs = HashSet<Int>(freq)
    for (number in numbers.asSequence().infinite()){
        freq += number
        if (!freqs.add(freq)){
            println(freq)
            return
        }
    }
}