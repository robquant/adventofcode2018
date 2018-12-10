fun <T> Sequence<T>.infinite() = sequence {
    while (true) {
        yieldAll(this@infinite)
    }
}