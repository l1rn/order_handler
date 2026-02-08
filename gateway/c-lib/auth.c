#include <string.h>

int check_token(const char *t){
	if(strncmp(t, "Bearer ", 7) != 0){
		return 0;
	}

	const char *actual_t = t + 7;
	
	int dot_count = 0;
	for(int i = 0; actual_t[i]; i++){
		if(actual_t[i] == '.') dot_count++;
	}
	return (dot_count == 2) ? 1 : 0;
}

void get_user_id(const char *t, char *buf){
	if(check_token(t)){
		strcpy(buf, "user123");
	} else {
		strcpy(buf, "invalid token:(");
	}
}
