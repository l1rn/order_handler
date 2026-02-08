#include <string.h>

int check_access(const char *t){
	if (t == NULL) return 0;

	if(strcmp(t, "open") == 0){
		return 1;
	}

	return 0;
}

