#include <stdio.h>
#include "cups/cups.h"
#include "cupsmod.h"

int cups_enum_dests_cb(user_data_t *user_data, unsigned flags, cups_dest_t *dest){

  if (flags & CUPS_DEST_FLAGS_REMOVED){
    //Remove destination from array...
    user_data->num_dests = cupsRemoveDest(dest->name, dest->instance, user_data->num_dests, &(user_data->dests));
  }else{
	// add destination to array...
    user_data->num_dests = cupsCopyDest(dest, user_data->num_dests, &(user_data->dests));
  }

  return (1);
}

int cups_enum_dests(cups_ptype_t type, cups_ptype_t mask, cups_dest_t **dests){
  user_data_t user_data = { 0, NULL };

  if (!cupsEnumDests(CUPS_DEST_FLAGS_NONE, 1000, NULL, type,
                     mask, (cups_dest_cb_t)cups_enum_dests_cb,
                     &user_data)){
   /*
    * An error occurred, free all of the destinations and
    * return...
    */

    cupsFreeDests(user_data.num_dests, user_data.dests);
    *dests = NULL;
    return (0);
  }

 /*
  * Return the destination array...
  */

  *dests = user_data.dests;
  return (user_data.num_dests);
}
