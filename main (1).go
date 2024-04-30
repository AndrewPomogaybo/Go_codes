package main

import (
	"fmt"
	"math"
)

func main() {

  var(
    min int
    max int
    num int
    number int
    elements_size int
  )
  
  fmt.Println("Задание 1")
  fmt.Println("Введите число")
  fmt.Scan(&number)
	fmt.Println(PluralizeComputer(number))
  
  fmt.Println("Задание 2")

  fmt.Println("Введите размер массива")
  fmt.Scan(&elements_size)

  
  numbers := make([]int,elements_size)

  fmt.Println("Введите элементы массива:")
  for i := 0; i < elements_size; i++ {
      fmt.Printf("Введите элемент %d: ", i+1)
      fmt.Scanln(&numbers[i])
  }

  commonDivisors := FindCommonDivisors(numbers)

  fmt.Println("Общие делители", commonDivisors)

  fmt.Println("Задание 3")
  
  fmt.Println("Введите начало диапазона")
  fmt.Scan(&min)
  fmt.Println("Введите конец диапазона")
  fmt.Scan(&max)

  if min > max{
    fmt.Println("Начальное значение должно быть меньше")
    return
  }
  
  primes := findPrimesInRange(min,max)
  fmt.Println("Простые числа в диапазоне:",primes)

  fmt.Println("Задание 4")
  

  fmt.Println("Введите число")
  fmt.Scan(&num)
  printMultiplicationTable(num)
}
//Задание 4
func printMultiplicationTable(n int) {
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            fmt.Printf("%d\t", i*j)
        }
        fmt.Println()
    }
}

//Задание 1
func PluralizeComputer(num int)string{
  last_digit := num % 10
  last_two_digit := num % 100
  
  switch{
    case last_digit == 1 && last_two_digit != 11:
      return fmt.Sprintf("%d компьютер",num)
    case last_digit >= 2 && last_digit <= 4 && !(last_two_digit >= 12 && last_two_digit <= 14):
      return fmt.Sprintf("%d компьютера",num)
    default:
      return fmt.Sprintf("%d компьютеров", num)
 
  }
  
}

//задание 2
func Maximum(x, y int) int {
 for y != 0 {
  x, y = y, x%y
 }
 return x
}

// gcdArray находит НОД для массива чисел задание 2
func MaximumArray(numbers []int) int {
 result := numbers[0]
 for _, num := range numbers[1:] {
  result = Maximum(result, num)
  if result == 1 {
   return 1
  }
 }
 return result
}

// findDivisors находит все делители числа задание 2
func FindDivisors(n int) []int {
 divisors := []int{}
 for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
  if n%i == 0 {
   divisors = append(divisors, i)
   if i != n/i {
    divisors = append(divisors, n/i)
   }
  }
 }
 return divisors
}

// findCommonDivisors находит общие делители для всех чисел в массиве, задание 2
func FindCommonDivisors(numbers []int) []int {
 gcdResult := MaximumArray(numbers)
 divisors := FindDivisors(gcdResult)
 return divisors
}

//задание 3
func isPrime(n int) bool {
 if n <= 1 {
  return false
 }
 if n <= 3 {
  return true
 }
 if n%2 == 0 || n%3 == 0 {
  return false
 }
 for i := 5; i*i <= n; i += 6 {
  if n%i == 0 || n%(i+2) == 0 {
   return false
  }
 }
 return true
}

// findPrimesInRange находит все простые числа в заданном диапазоне.задание 3
func findPrimesInRange(min, max int) []int {
 var primes []int
 for i := min; i <= max; i++ {
  if isPrime(i) {
   primes = append(primes, i)
  }
 }
 return primes
}

