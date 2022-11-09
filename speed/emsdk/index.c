#include <emscripten/emscripten.h>


EMSCRIPTEN_KEEPALIVE
int fib(int n) {
  if (n < 2) {
    return 1;
  } else {
    return fib(n - 1) + fib(n - 2);
  }
}

EMSCRIPTEN_KEEPALIVE
int cFib30Times(int n) {
  int sum = 0;
	for (int i = 0; i < 30; i++) {
		sum += fib(n);
	}
  return sum;
}