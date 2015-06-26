#ifdef __cplusplus
extern "C" {
#endif
typedef	void* Kvstore;
Kvstore Copen(char *, int, int);
int Cget_key(Kvstore, char *key, int value);
int Cset_key(Kvstore, char *key, int value);
#ifdef __cplusplus
}
#endif
