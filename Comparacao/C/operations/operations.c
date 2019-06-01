#include <stdio.h>
#include <stdbool.h>

int main() {
	int x = 5;
	int y = 3;

	printf("x + y = %d\n", x+y);
	printf("x - y = %d\n", x-y);
	printf("x * y = %d\n", x*y);
	printf("x / y = %d\n", x/y);
	printf("x mod y = %d\n", x%y);

	bool varbool = true;
	bool otherbool = false;

	// C é incapaz de imprimir tipos booleanos
	printf(varbool && otherbool? "true" : "false");
	printf("\n");
	printf(varbool || otherbool? "true" : "false");
	printf("\n");
	printf(!(varbool) == otherbool?"true" : "false");
	printf("\n");

	return 0;
}
