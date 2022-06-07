#include <stdio.h>
#include <errno.h>
#include <string.h>
#include <stdlib.h>
#include "number.h"

int number_div(int a, int b)
{
	if (b == 0)
	{
		errno = EINVAL;
		fprintf(stderr, "Division by zero! Exiting...\n");
		return 0;
		// exit(1);
	}

	// FILE *fp;

	// fp = fopen("123.txt", "r");
	// printf("The error message is: %s\n", strerror(errno));
	// perror("Message from perror");
	return a / b;
}

char *foo(char *input)
{
	char *m = malloc(16 * sizeof(char));
	sprintf(m, "%s", input);
	// go_debug_log("Read done...............");

	char str[255] = {0};
	snprintf(str, sizeof(str), "foo char m is: %s", m);  //记录文件，行号以及日志信息
	printf("222222222222222 %s\n", str);
	go_debug_log(str);

	go_debug_log_char("Read done...............", m);
	return m;
}