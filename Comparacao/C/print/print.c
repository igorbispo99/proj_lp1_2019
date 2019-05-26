#define BYTE_TO_BINARY_PATTERN "%c%c%c%c%c%c%c%c"
#define BYTE_TO_BINARY(byte)  \
  (byte & 0x80 ? '1' : '0'), \
  (byte & 0x40 ? '1' : '0'), \
  (byte & 0x20 ? '1' : '0'), \
  (byte & 0x10 ? '1' : '0'), \
  (byte & 0x08 ? '1' : '0'), \
  (byte & 0x04 ? '1' : '0'), \
  (byte & 0x02 ? '1' : '0'), \
  (byte & 0x01 ? '1' : '0')

#include <stdio.h>
#include <string.h>
#include <stdbool.h>

int main() {

  // Problemas: tamanho limitado na string final
  char a[50] = "Some text";
  char *b = " more text"; 

  bool win = true;

  float pi = 3.141592;

  strcat(a,b);

  printf("%d \n", 3);

  printf("%s\n", a);

  printf(win? "true" : "false");
  printf("\n");

  printf(BYTE_TO_BINARY_PATTERN, BYTE_TO_BINARY(12));
  printf("\n");

  printf("%c \n", 34);
	printf("%X \n", 12);
	printf("%e \n", pi);

  return 0;
}