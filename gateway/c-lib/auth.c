#include <string.h>
#include <jwt.h>

int base64_decode(const char *input, unsigned char **output){
	size_t len = strlen(input);
	size_t padding = 0;
	if (input[len-1] == '=') padding++;
	if (input[len-2] == '=') padding++;
	
	size_t out_len = (len * 3) / 4 - padding;
	*output = malloc(out_len + 1);

	BIO *bio, b64;
	b64 = BIO_NEW(BIO_f_base64());
	bio = BIO_new_mem_buf(input, -1);
	bio = BIO_push(b64, bio);
	BIO_set_flags(bio, BIO_FLAGS_BASE64_NO_NL);
	BIO_read(bio, *output, len);
	BIO_free_all(bio);

	return out_len;
}

void base64url_decode_to_str(const char *input, char *output){
	char temp[10240];
}

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
