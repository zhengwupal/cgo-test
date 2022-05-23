#include <stdio.h>
#include <errno.h>
#include <string.h>
#include <stdlib.h>
#include "number.h"

int number_div(int a, int b) {
	if(b == 0) {
		errno = EINVAL;
		fprintf(stderr, "Division by zero! Exiting...\n");
		return 0;
		// exit(EXIT_FAILURE);
	}

	// FILE *fp;

	// fp = fopen("123.txt", "r");
	// printf("The error message is: %s\n", strerror(errno));
	// perror("Message from perror");
	return a/b;
}

char* foo(char *input) {
	char* m = malloc(16*sizeof(char));
	sprintf(m, "%s", input);
	go_debug_log("Read done...............");
	return m;
}