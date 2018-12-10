import java.io.File

val p02 =  fun(input: List<String>) {

    var count_2 = 0
    var count_3 = 0
    for (line in input){
        val counts = line.groupingBy { it }.eachCount()
        count_2 += if (counts.containsValue(2)) 1 else 0
        count_3 += if (counts.containsValue(3)) 1 else 0
    }
    println(count_2 * count_3)

    for ((i, line1) in input.withIndex()){
        for (line2 in input.drop(i)) {
            val sameCount = line1.zip(line2).count {(c1,c2) -> c1 == c2 }
            if (sameCount == line1.length - 1) {
                val sameLetters = line1.zip(line2).filter { (c1, c2) -> (c1 == c2) }
                .map { (c1, _) -> c1 }
                println(sameLetters.joinToString(""))
            }
        }
    }
}