#ifndef __CUPSMOD_H__
#define __CUPSMOD_H__
#endif

typedef struct{
	int num_dests;
	cups_dest_t *dests;
} user_data_t;

extern int cups_enum_dests(cups_ptype_t type, cups_ptype_t mask, cups_dest_t **dests);
extern int cups_enum_dests_cb(user_data_t *user_data, unsigned flags, cups_dest_t *dest);
extern void test_get_info();