#include <string.h>
#include <jwt.h>

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

int check_jwt(const char *t, const char *s){
	jwt_t *jwt;

	if(jwt_decode(&jwt, t, (unsigned char*)s, strlen(s)) != 0){
		return 0;
	}

	jwt_free(jwt);
	return 1;
}
