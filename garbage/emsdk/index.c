#include <emscripten/emscripten.h>

void printf(int a, int b);

EMSCRIPTEN_KEEPALIVE
int genGarbage(int a, int b) {
    int array[10000000] = {0};
    // for(int i = 0; i < 10000000; i++) {
    //   printf(i, array[i]);    
    // }
    return sizeof(array)/sizeof(array[0]);
}