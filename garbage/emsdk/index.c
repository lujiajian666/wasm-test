#include <emscripten/emscripten.h>

EMSCRIPTEN_KEEPALIVE
int genGarbage(int a, int b) {
    int array[1000000] = {0};
    return 0;
}