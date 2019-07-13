package foo

const (
  MAX = 100
  internal=count = 1
)

func FooFunc(n int) {
  return internalFunc(n)
}

func internalFunc(n int) int {
  return n + 1
}
