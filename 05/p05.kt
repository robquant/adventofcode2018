import java.io.File

fun react(input: ArrayList<Char>) : ArrayList<Char> {
    var i = 0
    var result = ArrayList<Char>()
    val reversed = input.map { if (it.isUpperCase()) it.toLowerCase() else it.toUpperCase() }
    while (i < input.size - 1){
        val next_changed = reversed[i+1]
        if (input[i] == next_changed){
            i += 2
            continue
        }
        result.add(input[i])
        i += 1
    }
    if (i == input.size - 1){
        result.add(input[i])
    }
    return result
}


fun fullyReact(input: Collection<Char>, count:Boolean = false) : ArrayList<Char> {
    var reacted = ArrayList(input)
    var old_size: Int
    var counter:Int = 0
    do {
        old_size = reacted.size
        reacted = react(reacted)
        counter += 1
    } while (reacted.size < old_size)
    if (count) {
        println(counter)
    }
    return reacted
}

val p05 = fun(input: List<String>){
    val input = ArrayList(input[0].toList())
    val reacted = fullyReact(input, count=true)
    println(reacted.size)

    val min_length = ('a'..'z').map { c -> fullyReact(reacted.filter { it -> it.toLowerCase() != c }).size}.min()
    println(min_length)
}
