package keeper

import (
	"example/x/example/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetTextStorage set a specific textStorage in the store from its index
func (k Keeper) SetTextStorage(ctx sdk.Context, textStorage types.TextStorage) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TextStorageKeyPrefix))
	b := k.cdc.MustMarshal(&textStorage)
	store.Set(types.TextStorageKey(
		textStorage.Index,
	), b)
}

// GetTextStorage returns a textStorage from its index
func (k Keeper) GetTextStorage(
	ctx sdk.Context,
	index string,

) (val types.TextStorage, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TextStorageKeyPrefix))

	b := store.Get(types.TextStorageKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTextStorage removes a textStorage from the store
func (k Keeper) RemoveTextStorage(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TextStorageKeyPrefix))
	store.Delete(types.TextStorageKey(
		index,
	))
}

// GetAllTextStorage returns all textStorage
func (k Keeper) GetAllTextStorage(ctx sdk.Context) (list []types.TextStorage) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TextStorageKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TextStorage
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
