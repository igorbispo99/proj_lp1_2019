#include <stdio.h>
#include <stdlib.h>
#include <string.h>


int main() {

  int B[5] = {0,1,2,3,4};
  int size = sizeof(B)/sizeof(B[0]);

  for(int i = 0; i < size; i++){
    printf("%d %d\n", B[i], i);
  }

  // Não há arranjos em C. Mas a implementação ficaria
  // parecido com esse trecho
  int* sub[5];
  sub[0] = &(B[2]);
  sub[1] = &(B[3]);

  printf("[");
  printf("%d %d", *sub[0], *sub[1]);
  printf("]\n");


  // Copy array
  int* sub2[5];
  for(int i = 0; i < size; i++){
    sub2[i] = &B[i];
  }

  printf("[");
  printf("%d", *sub2[0]);

  for(int i = 1; i < size; i++){
    printf(" %d", *sub2[i]);
  }
  printf("]\n");

  int C[3] = {2,-3,0};
  int* sub3 = malloc(8*sizeof(int));

  memcpy(sub3, B, 5*sizeof(B[0]));
  memcpy(sub3+5, C, 3*sizeof(C[0]));

  printf("[");
  printf("%d", sub3[0]);
  for(int i = 1; i < 8; i++){
    printf(" %d", sub3[i]);
  }
  printf("]\n");

  
  return 0;

}
