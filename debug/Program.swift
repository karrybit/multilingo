func main() {
let n = Int(readLine()!)!

var years: [Int] = []
(1...n).forEach { _ in
    years.append(Int(readLine()!)!)
}
years.forEach {
    let isLeap = "is a leap year"
    let isNotLeap = "is not a leap year"

    if $0 % 4 == 0, $0 % 100 != 0 {
        print("\($0) \(isLeap)")
    } else if $0 % 4 == 0, $0 % 400 == 0 {
        print("\($0) \(isLeap)")
    } else {
        print("\($0) \(isNotLeap)")
    }
}
}

main()
